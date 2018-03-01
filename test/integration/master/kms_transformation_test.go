// +build !windows

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package master

import (
	"bytes"
	"context"
	"crypto/aes"
	"encoding/binary"
	"fmt"
	"strings"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/storage/value"
	aestransformer "k8s.io/apiserver/pkg/storage/value/encrypt/aes"

	kmsapi "k8s.io/apiserver/pkg/storage/value/encrypt/envelope/v1beta1"
)

const (
	kmsPrefix     = "k8s:enc:kms:v1:grpc-kms-provider:"
	dekKeySizeLen = 2

	kmsConfigYAML = `
kind: EncryptionConfig
apiVersion: v1
resources:
  - resources:
    - secrets
    providers:
    - kms:
       name: grpc-kms-provider
       cachesize: 1000
       endpoint: unix:///tmp/kms-provider.sock
`
)

// rawDEKKEKSecret provides operations for working with secrets transformed with Data Encryption Key(DEK) Key Encryption Kye(KEK) envelop.
type rawDEKKEKSecret []byte

func (r rawDEKKEKSecret) getDEKLen() int {
	// DEK's length is stored in the two bytes that follow the prefix.
	return int(binary.BigEndian.Uint16(r[len(kmsPrefix) : len(kmsPrefix)+dekKeySizeLen]))
}

func (r rawDEKKEKSecret) getDEK() []byte {
	return r[len(kmsPrefix)+dekKeySizeLen : len(kmsPrefix)+dekKeySizeLen+r.getDEKLen()]
}

func (r rawDEKKEKSecret) getStartOfPayload() int {
	return len(kmsPrefix) + dekKeySizeLen + r.getDEKLen()
}

func (r rawDEKKEKSecret) getPayload() []byte {
	return r[r.getStartOfPayload():]
}

// TestKMSProvider is an integration test between KubAPI, ETCD and KMS Plugin
// Concretely, this test verifies the following integration contracts:
// 1. Raw records in ETCD that were processed by KMS Provider should be prefixed with k8s:enc:kms:v1:grpc-kms-provider-name:
// 2. Data Encryption Key (DEK) should be generated by envelopeTransformer and passed to KMS gRPC Plugin
// 3. KMS gRPC Plugin should encrypt the DEK with a Key Encryption Key (KEK) and pass it back to envelopeTransformer
// 4. The payload (ex. Secret) should be encrypted via AES CBC transform
// 5. Prefix-EncryptedDEK-EncryptedPayload structure should be deposited to ETCD
func TestKMSProvider(t *testing.T) {
	pluginMock, err := NewBase64Plugin()
	if err != nil {
		t.Fatalf("failed to create mock of KMS Plugin: %v", err)
	}
	defer pluginMock.cleanUp()
	go pluginMock.grpcServer.Serve(pluginMock.listener)

	test, err := newTransformTest(t, kmsConfigYAML)
	if err != nil {
		t.Fatalf("failed to start KUBE API Server with encryptionConfig\n %s", kmsConfigYAML)
	}
	defer test.cleanUp()

	secretETCDPath := test.getETCDPath()
	var rawSecretAsSeenByETCD rawDEKKEKSecret
	rawSecretAsSeenByETCD, err = test.getRawSecretFromETCD()
	if err != nil {
		t.Fatalf("failed to read %s from etcd: %v", secretETCDPath, err)
	}

	if !bytes.HasPrefix(rawSecretAsSeenByETCD, []byte(kmsPrefix)) {
		t.Fatalf("expected secret to be prefixed with %s, but got %s", kmsPrefix, rawSecretAsSeenByETCD)
	}

	// Since Data Encryption Key (DEK) is randomly generated (per encryption operation), we need to ask KMS Mock for it.
	dekPlainAsSeenByKMS, err := getDEKFromKMSPlugin(pluginMock)
	if err != nil {
		t.Fatalf("failed to get DEK from KMS: %v", err)
	}

	decryptResponse, err := pluginMock.Decrypt(context.Background(),
		&kmsapi.DecryptRequest{Version: kmsAPIVersion, Cipher: rawSecretAsSeenByETCD.getDEK()})
	if err != nil {
		t.Fatalf("failed to decrypt DEK, %v", err)
	}
	dekPlainAsWouldBeSeenByETCD := decryptResponse.Plain

	if !bytes.Equal(dekPlainAsSeenByKMS, dekPlainAsWouldBeSeenByETCD) {
		t.Fatalf("expected dekPlainAsSeenByKMS %v to be passed to KMS Plugin, but got %s",
			dekPlainAsSeenByKMS, dekPlainAsWouldBeSeenByETCD)
	}

	plainSecret, err := decryptPayload(dekPlainAsWouldBeSeenByETCD, rawSecretAsSeenByETCD, secretETCDPath)
	if err != nil {
		t.Fatalf("failed to transform from storage via AESCBC, err: %v", err)
	}

	if !strings.Contains(string(plainSecret), secretVal) {
		t.Fatalf("expected %q after decryption, but got %q", secretVal, string(plainSecret))
	}

	// Secrets should be un-enveloped on direct reads from Kube API Server.
	s, err := test.restClient.CoreV1().Secrets(testNamespace).Get(testSecret, metav1.GetOptions{})
	if secretVal != string(s.Data[secretKey]) {
		t.Fatalf("expected %s from KubeAPI, but got %s", secretVal, string(s.Data[secretKey]))
	}
}

func getDEKFromKMSPlugin(pluginMock *base64Plugin) ([]byte, error) {
	select {
	case e := <-pluginMock.encryptRequest:
		return e.Plain, nil
	case <-time.After(time.Second):
		return nil, fmt.Errorf("timed-out while getting encryption request from KMS Plugin Mock")
	}
}

func decryptPayload(key []byte, secret rawDEKKEKSecret, secretETCDPath string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize AES Cipher: %v", err)
	}
	// etcd path of the key is used as the authenticated context - need to pass it to decrypt
	ctx := value.DefaultContext([]byte(secretETCDPath))
	aescbcTransformer := aestransformer.NewCBCTransformer(block)
	plainSecret, _, err := aescbcTransformer.TransformFromStorage(secret.getPayload(), ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to transform from storage via AESCBC, err: %v", err)
	}

	return plainSecret, nil
}
