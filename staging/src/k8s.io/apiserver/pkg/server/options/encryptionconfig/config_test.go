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

package encryptionconfig

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apiserverconfig "k8s.io/apiserver/pkg/apis/config"
	"k8s.io/apiserver/pkg/features"
	"k8s.io/apiserver/pkg/storage/value"
	"k8s.io/apiserver/pkg/storage/value/encrypt/envelope"
	envelopekmsv2 "k8s.io/apiserver/pkg/storage/value/encrypt/envelope/kmsv2"
	"k8s.io/apiserver/pkg/storage/value/encrypt/envelope/metrics"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	featuregatetesting "k8s.io/component-base/featuregate/testing"
	"k8s.io/component-base/metrics/legacyregistry"
	"k8s.io/component-base/metrics/testutil"
	kmsservice "k8s.io/kms/pkg/service"
)

const (
	sampleText        = "abcdefghijklmnopqrstuvwxyz"
	sampleContextText = "0123456789"
)

var (
	sampleInvalidKeyID = string(make([]byte, envelopekmsv2.KeyIDMaxSize+1))
)

// testEnvelopeService is a mock envelope service which can be used to simulate remote Envelope services
// for testing of the envelope transformer with other transformers.
type testEnvelopeService struct {
	err error
}

func (t *testEnvelopeService) Decrypt(data []byte) ([]byte, error) {
	if t.err != nil {
		return nil, t.err
	}
	return base64.StdEncoding.DecodeString(string(data))
}

func (t *testEnvelopeService) Encrypt(data []byte) ([]byte, error) {
	if t.err != nil {
		return nil, t.err
	}
	return []byte(base64.StdEncoding.EncodeToString(data)), nil
}

// testKMSv2EnvelopeService is a mock kmsv2 envelope service which can be used to simulate remote Envelope v2 services
// for testing of the envelope transformer with other transformers.
type testKMSv2EnvelopeService struct {
	err   error
	keyID string
}

func (t *testKMSv2EnvelopeService) Decrypt(ctx context.Context, uid string, req *kmsservice.DecryptRequest) ([]byte, error) {
	if t.err != nil {
		return nil, t.err
	}
	return base64.StdEncoding.DecodeString(string(req.Ciphertext))
}

func (t *testKMSv2EnvelopeService) Encrypt(ctx context.Context, uid string, data []byte) (*kmsservice.EncryptResponse, error) {
	if t.err != nil {
		return nil, t.err
	}
	return &kmsservice.EncryptResponse{
		Ciphertext: []byte(base64.StdEncoding.EncodeToString(data)),
		KeyID:      t.keyID,
	}, nil
}

func (t *testKMSv2EnvelopeService) Status(ctx context.Context) (*kmsservice.StatusResponse, error) {
	if t.err != nil {
		return nil, t.err
	}
	return &kmsservice.StatusResponse{Healthz: "ok", KeyID: t.keyID, Version: "v2alpha1"}, nil
}

// The factory method to create mock envelope service.
func newMockEnvelopeService(ctx context.Context, endpoint string, timeout time.Duration) (envelope.Service, error) {
	return &testEnvelopeService{nil}, nil
}

// The factory method to create mock envelope service which always returns error.
func newMockErrorEnvelopeService(endpoint string, timeout time.Duration) (envelope.Service, error) {
	return &testEnvelopeService{errors.New("test")}, nil
}

// The factory method to create mock envelope kmsv2 service.
func newMockEnvelopeKMSv2Service(ctx context.Context, endpoint, providerName string, timeout time.Duration) (kmsservice.Service, error) {
	return &testKMSv2EnvelopeService{nil, "1"}, nil
}

// The factory method to create mock envelope kmsv2 service which always returns error.
func newMockErrorEnvelopeKMSv2Service(endpoint string, timeout time.Duration) (kmsservice.Service, error) {
	return &testKMSv2EnvelopeService{errors.New("test"), "1"}, nil
}

// The factory method to create mock envelope kmsv2 service that always returns invalid keyID.
func newMockInvalidKeyIDEnvelopeKMSv2Service(ctx context.Context, endpoint string, timeout time.Duration, keyID string) (kmsservice.Service, error) {
	return &testKMSv2EnvelopeService{nil, keyID}, nil
}

func TestLegacyConfig(t *testing.T) {
	legacyV1Config := "testdata/valid-configs/legacy.yaml"
	legacyConfigObject, _, err := loadConfig(legacyV1Config, false)
	cacheSize := int32(10)
	if err != nil {
		t.Fatalf("error while parsing configuration file: %s.\nThe file was:\n%s", err, legacyV1Config)
	}

	expected := &apiserverconfig.EncryptionConfiguration{
		Resources: []apiserverconfig.ResourceConfiguration{
			{
				Resources: []string{"secrets", "namespaces"},
				Providers: []apiserverconfig.ProviderConfiguration{
					{Identity: &apiserverconfig.IdentityConfiguration{}},
					{AESGCM: &apiserverconfig.AESConfiguration{
						Keys: []apiserverconfig.Key{
							{Name: "key1", Secret: "c2VjcmV0IGlzIHNlY3VyZQ=="},
							{Name: "key2", Secret: "dGhpcyBpcyBwYXNzd29yZA=="},
						},
					}},
					{KMS: &apiserverconfig.KMSConfiguration{
						APIVersion: "v1",
						Name:       "testprovider",
						Endpoint:   "unix:///tmp/testprovider.sock",
						CacheSize:  &cacheSize,
						Timeout:    &metav1.Duration{Duration: 3 * time.Second},
					}},
					{AESCBC: &apiserverconfig.AESConfiguration{
						Keys: []apiserverconfig.Key{
							{Name: "key1", Secret: "c2VjcmV0IGlzIHNlY3VyZQ=="},
							{Name: "key2", Secret: "dGhpcyBpcyBwYXNzd29yZA=="},
						},
					}},
					{Secretbox: &apiserverconfig.SecretboxConfiguration{
						Keys: []apiserverconfig.Key{
							{Name: "key1", Secret: "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY="},
						},
					}},
				},
			},
		},
	}
	if d := cmp.Diff(expected, legacyConfigObject); d != "" {
		t.Fatalf("EncryptionConfig mismatch (-want +got):\n%s", d)
	}
}

func TestEncryptionProviderConfigCorrect(t *testing.T) {
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.KMSv2, true)()
	// Set factory for mock envelope service
	factory := envelopeServiceFactory
	factoryKMSv2 := EnvelopeKMSv2ServiceFactory
	envelopeServiceFactory = newMockEnvelopeService
	EnvelopeKMSv2ServiceFactory = newMockEnvelopeKMSv2Service
	defer func() {
		envelopeServiceFactory = factory
		EnvelopeKMSv2ServiceFactory = factoryKMSv2
	}()

	ctx := testContext(t)

	// Creates compound/prefix transformers with different ordering of available transformers.
	// Transforms data using one of them, and tries to untransform using the others.
	// Repeats this for all possible combinations.
	// Math for GracePeriod is explained at - https://github.com/kubernetes/kubernetes/blob/c9ed04762f94a319d7b1fb718dc345491a32bea6/staging/src/k8s.io/apiserver/pkg/server/options/encryptionconfig/config.go#L159-L163
	expectedKMSCloseGracePeriod := 46 * time.Second
	correctConfigWithIdentityFirst := "testdata/valid-configs/identity-first.yaml"
	identityFirstEncryptionConfiguration, err := LoadEncryptionConfig(ctx, correctConfigWithIdentityFirst, false)
	if err != nil {
		t.Fatalf("error while parsing configuration file: %s.\nThe file was:\n%s", err, correctConfigWithIdentityFirst)
	}
	if identityFirstEncryptionConfiguration.KMSCloseGracePeriod != expectedKMSCloseGracePeriod {
		t.Fatalf("KMSCloseGracePeriod mismatch (-want +got):\n%s", cmp.Diff(expectedKMSCloseGracePeriod, identityFirstEncryptionConfiguration.KMSCloseGracePeriod))
	}

	// Math for GracePeriod is explained at - https://github.com/kubernetes/kubernetes/blob/c9ed04762f94a319d7b1fb718dc345491a32bea6/staging/src/k8s.io/apiserver/pkg/server/options/encryptionconfig/config.go#L159-L163
	expectedKMSCloseGracePeriod = 32 * time.Second
	correctConfigWithAesGcmFirst := "testdata/valid-configs/aes-gcm-first.yaml"
	aesGcmFirstEncryptionConfiguration, err := LoadEncryptionConfig(ctx, correctConfigWithAesGcmFirst, false)
	if err != nil {
		t.Fatalf("error while parsing configuration file: %s.\nThe file was:\n%s", err, correctConfigWithAesGcmFirst)
	}
	if aesGcmFirstEncryptionConfiguration.KMSCloseGracePeriod != expectedKMSCloseGracePeriod {
		t.Fatalf("KMSCloseGracePeriod mismatch (-want +got):\n%s", cmp.Diff(expectedKMSCloseGracePeriod, aesGcmFirstEncryptionConfiguration.KMSCloseGracePeriod))
	}

	invalidConfigWithAesGcm := "testdata/invalid-configs/invalid-aes-gcm.yaml"
	_, err = LoadEncryptionConfig(ctx, invalidConfigWithAesGcm, false)
	if !strings.Contains(errString(err), "error while parsing file") {
		t.Fatalf("should result in error while parsing configuration file: %s.\nThe file was:\n%s", err, invalidConfigWithAesGcm)
	}

	// Math for GracePeriod is explained at - https://github.com/kubernetes/kubernetes/blob/c9ed04762f94a319d7b1fb718dc345491a32bea6/staging/src/k8s.io/apiserver/pkg/server/options/encryptionconfig/config.go#L159-L163
	expectedKMSCloseGracePeriod = 26 * time.Second
	correctConfigWithAesCbcFirst := "testdata/valid-configs/aes-cbc-first.yaml"
	aesCbcFirstEncryptionConfiguration, err := LoadEncryptionConfig(ctx, correctConfigWithAesCbcFirst, false)
	if err != nil {
		t.Fatalf("error while parsing configuration file: %s.\nThe file was:\n%s", err, correctConfigWithAesCbcFirst)
	}
	if aesCbcFirstEncryptionConfiguration.KMSCloseGracePeriod != expectedKMSCloseGracePeriod {
		t.Fatalf("KMSCloseGracePeriod mismatch (-want +got):\n%s", cmp.Diff(expectedKMSCloseGracePeriod, aesCbcFirstEncryptionConfiguration.KMSCloseGracePeriod))
	}

	// Math for GracePeriod is explained at - https://github.com/kubernetes/kubernetes/blob/c9ed04762f94a319d7b1fb718dc345491a32bea6/staging/src/k8s.io/apiserver/pkg/server/options/encryptionconfig/config.go#L159-L163
	expectedKMSCloseGracePeriod = 14 * time.Second
	correctConfigWithSecretboxFirst := "testdata/valid-configs/secret-box-first.yaml"
	secretboxFirstEncryptionConfiguration, err := LoadEncryptionConfig(ctx, correctConfigWithSecretboxFirst, false)
	if err != nil {
		t.Fatalf("error while parsing configuration file: %s.\nThe file was:\n%s", err, correctConfigWithSecretboxFirst)
	}
	if secretboxFirstEncryptionConfiguration.KMSCloseGracePeriod != expectedKMSCloseGracePeriod {
		t.Fatalf("KMSCloseGracePeriod mismatch (-want +got):\n%s", cmp.Diff(expectedKMSCloseGracePeriod, secretboxFirstEncryptionConfiguration.KMSCloseGracePeriod))
	}

	// Math for GracePeriod is explained at - https://github.com/kubernetes/kubernetes/blob/c9ed04762f94a319d7b1fb718dc345491a32bea6/staging/src/k8s.io/apiserver/pkg/server/options/encryptionconfig/config.go#L159-L163
	expectedKMSCloseGracePeriod = 34 * time.Second
	correctConfigWithKMSFirst := "testdata/valid-configs/kms-first.yaml"
	kmsFirstEncryptionConfiguration, err := LoadEncryptionConfig(ctx, correctConfigWithKMSFirst, false)
	if err != nil {
		t.Fatalf("error while parsing configuration file: %s.\nThe file was:\n%s", err, correctConfigWithKMSFirst)
	}
	if kmsFirstEncryptionConfiguration.KMSCloseGracePeriod != expectedKMSCloseGracePeriod {
		t.Fatalf("KMSCloseGracePeriod mismatch (-want +got):\n%s", cmp.Diff(expectedKMSCloseGracePeriod, kmsFirstEncryptionConfiguration.KMSCloseGracePeriod))
	}

	// Math for GracePeriod is explained at - https://github.com/kubernetes/kubernetes/blob/c9ed04762f94a319d7b1fb718dc345491a32bea6/staging/src/k8s.io/apiserver/pkg/server/options/encryptionconfig/config.go#L159-L163
	expectedKMSCloseGracePeriod = 42 * time.Second
	correctConfigWithKMSv2First := "testdata/valid-configs/kmsv2-first.yaml"
	kmsv2FirstEncryptionConfiguration, err := LoadEncryptionConfig(ctx, correctConfigWithKMSv2First, false)
	if err != nil {
		t.Fatalf("error while parsing configuration file: %s.\nThe file was:\n%s", err, correctConfigWithKMSv2First)
	}
	if kmsv2FirstEncryptionConfiguration.KMSCloseGracePeriod != expectedKMSCloseGracePeriod {
		t.Fatalf("KMSCloseGracePeriod mismatch (-want +got):\n%s", cmp.Diff(expectedKMSCloseGracePeriod, kmsv2FirstEncryptionConfiguration.KMSCloseGracePeriod))
	}

	// Pick the transformer for any of the returned resources.
	identityFirstTransformer := identityFirstEncryptionConfiguration.Transformers[schema.ParseGroupResource("secrets")]
	aesGcmFirstTransformer := aesGcmFirstEncryptionConfiguration.Transformers[schema.ParseGroupResource("secrets")]
	aesCbcFirstTransformer := aesCbcFirstEncryptionConfiguration.Transformers[schema.ParseGroupResource("secrets")]
	secretboxFirstTransformer := secretboxFirstEncryptionConfiguration.Transformers[schema.ParseGroupResource("secrets")]
	kmsFirstTransformer := kmsFirstEncryptionConfiguration.Transformers[schema.ParseGroupResource("secrets")]
	kmsv2FirstTransformer := kmsv2FirstEncryptionConfiguration.Transformers[schema.ParseGroupResource("secrets")]

	dataCtx := value.DefaultContext([]byte(sampleContextText))
	originalText := []byte(sampleText)

	transformers := []struct {
		Transformer value.Transformer
		Name        string
	}{
		{aesGcmFirstTransformer, "aesGcmFirst"},
		{aesCbcFirstTransformer, "aesCbcFirst"},
		{secretboxFirstTransformer, "secretboxFirst"},
		{identityFirstTransformer, "identityFirst"},
		{kmsFirstTransformer, "kmsFirst"},
		{kmsv2FirstTransformer, "kmvs2First"},
	}

	for _, testCase := range transformers {
		transformedData, err := testCase.Transformer.TransformToStorage(ctx, originalText, dataCtx)
		if err != nil {
			t.Fatalf("%s: error while transforming data to storage: %s", testCase.Name, err)
		}

		for _, transformer := range transformers {
			untransformedData, stale, err := transformer.Transformer.TransformFromStorage(ctx, transformedData, dataCtx)
			if err != nil {
				t.Fatalf("%s: error while reading using %s transformer: %s", testCase.Name, transformer.Name, err)
			}
			if stale != (transformer.Name != testCase.Name) {
				t.Fatalf("%s: wrong stale information on reading using %s transformer, should be %v", testCase.Name, transformer.Name, testCase.Name == transformer.Name)
			}
			if !bytes.Equal(untransformedData, originalText) {
				t.Fatalf("%s: %s transformer transformed data incorrectly. Expected: %v, got %v", testCase.Name, transformer.Name, originalText, untransformedData)
			}
		}
	}
}

func TestKMSMaxTimeout(t *testing.T) {
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.KMSv2, true)()

	testCases := []struct {
		name            string
		expectedErr     string
		expectedTimeout time.Duration
		config          apiserverconfig.EncryptionConfiguration
	}{
		{
			name: "config with bad provider",
			config: apiserverconfig.EncryptionConfiguration{
				Resources: []apiserverconfig.ResourceConfiguration{
					{
						Resources: []string{"secrets"},
						Providers: []apiserverconfig.ProviderConfiguration{
							{
								KMS: nil,
							},
						},
					},
				},
			},
			expectedErr:     "provider does not contain any of the expected providers: KMS, AESGCM, AESCBC, Secretbox, Identity",
			expectedTimeout: 6 * time.Second,
		},
		{
			name: "default timeout",
			config: apiserverconfig.EncryptionConfiguration{
				Resources: []apiserverconfig.ResourceConfiguration{
					{
						Resources: []string{"secrets"},
						Providers: []apiserverconfig.ProviderConfiguration{
							{
								KMS: &apiserverconfig.KMSConfiguration{
									Name:       "kms",
									APIVersion: "v1",
									Timeout: &metav1.Duration{
										// default timeout is 3s
										// this will be set automatically if not provided in config file
										Duration: 3 * time.Second,
									},
									Endpoint: "unix:///tmp/testprovider.sock",
								},
							},
						},
					},
				},
			},
			expectedErr:     "",
			expectedTimeout: 6 * time.Second,
		},
		{
			name: "with v1 provider",
			config: apiserverconfig.EncryptionConfiguration{
				Resources: []apiserverconfig.ResourceConfiguration{
					{
						Resources: []string{"secrets"},
						Providers: []apiserverconfig.ProviderConfiguration{
							{
								KMS: &apiserverconfig.KMSConfiguration{
									Name:       "kms",
									APIVersion: "v1",
									Timeout: &metav1.Duration{
										// default timeout is 3s
										// this will be set automatically if not provided in config file
										Duration: 3 * time.Second,
									},
									Endpoint: "unix:///tmp/testprovider.sock",
								},
							},
						},
					},
					{
						Resources: []string{"configmaps"},
						Providers: []apiserverconfig.ProviderConfiguration{
							{
								KMS: &apiserverconfig.KMSConfiguration{
									Name:       "kms",
									APIVersion: "v1",
									Timeout: &metav1.Duration{
										// default timeout is 3s
										// this will be set automatically if not provided in config file
										Duration: 3 * time.Second,
									},
									Endpoint: "unix:///tmp/testprovider.sock",
								},
							},
						},
					},
				},
			},
			expectedErr:     "",
			expectedTimeout: 12 * time.Second,
		},
		{
			name: "with v2 provider",
			config: apiserverconfig.EncryptionConfiguration{
				Resources: []apiserverconfig.ResourceConfiguration{
					{
						Resources: []string{"secrets"},
						Providers: []apiserverconfig.ProviderConfiguration{
							{
								KMS: &apiserverconfig.KMSConfiguration{
									Name:       "kms",
									APIVersion: "v2",
									Timeout: &metav1.Duration{
										Duration: 15 * time.Second,
									},
									Endpoint: "unix:///tmp/testprovider.sock",
								},
							},
							{
								KMS: &apiserverconfig.KMSConfiguration{
									Name:       "new-kms",
									APIVersion: "v2",
									Timeout: &metav1.Duration{
										Duration: 5 * time.Second,
									},
									Endpoint: "unix:///tmp/anothertestprovider.sock",
								},
							},
						},
					},
					{
						Resources: []string{"configmaps"},
						Providers: []apiserverconfig.ProviderConfiguration{
							{
								KMS: &apiserverconfig.KMSConfiguration{
									Name:       "another-kms",
									APIVersion: "v2",
									Timeout: &metav1.Duration{
										Duration: 10 * time.Second,
									},
									Endpoint: "unix:///tmp/testprovider.sock",
								},
							},
							{
								KMS: &apiserverconfig.KMSConfiguration{
									Name:       "yet-another-kms",
									APIVersion: "v2",
									Timeout: &metav1.Duration{
										Duration: 2 * time.Second,
									},
									Endpoint: "unix:///tmp/anothertestprovider.sock",
								},
							},
						},
					},
				},
			},
			expectedErr:     "",
			expectedTimeout: 32 * time.Second,
		},
		{
			name: "with v1 and v2 provider",
			config: apiserverconfig.EncryptionConfiguration{
				Resources: []apiserverconfig.ResourceConfiguration{
					{
						Resources: []string{"secrets"},
						Providers: []apiserverconfig.ProviderConfiguration{
							{
								KMS: &apiserverconfig.KMSConfiguration{
									Name:       "kms",
									APIVersion: "v1",
									Timeout: &metav1.Duration{
										Duration: 1 * time.Second,
									},
									Endpoint: "unix:///tmp/testprovider.sock",
								},
							},
							{
								KMS: &apiserverconfig.KMSConfiguration{
									Name:       "another-kms",
									APIVersion: "v2",
									Timeout: &metav1.Duration{
										Duration: 1 * time.Second,
									},
									Endpoint: "unix:///tmp/anothertestprovider.sock",
								},
							},
						},
					},
					{
						Resources: []string{"configmaps"},
						Providers: []apiserverconfig.ProviderConfiguration{
							{
								KMS: &apiserverconfig.KMSConfiguration{
									Name:       "kms",
									APIVersion: "v1",
									Timeout: &metav1.Duration{
										Duration: 4 * time.Second,
									},
									Endpoint: "unix:///tmp/testprovider.sock",
								},
							},
							{
								KMS: &apiserverconfig.KMSConfiguration{
									Name:       "yet-another-kms",
									APIVersion: "v1",
									Timeout: &metav1.Duration{
										Duration: 2 * time.Second,
									},
									Endpoint: "unix:///tmp/anothertestprovider.sock",
								},
							},
						},
					},
				},
			},
			expectedErr:     "",
			expectedTimeout: 15 * time.Second,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			cacheSize := int32(1000)
			for _, resource := range testCase.config.Resources {
				for _, provider := range resource.Providers {
					if provider.KMS != nil {
						provider.KMS.CacheSize = &cacheSize
					}
				}
			}

			ctx, cancel := context.WithCancel(context.Background())
			cancel() // cancel this upfront so the kms v2 checks do not block

			_, _, kmsUsed, err := getTransformerOverridesAndKMSPluginHealthzCheckers(ctx, &testCase.config)

			if !strings.Contains(errString(err), testCase.expectedErr) {
				t.Fatalf("expecting error calling prefixTransformersAndProbes, expected: %s, got: %s", testCase.expectedErr, errString(err))
			}
			if len(testCase.expectedErr) == 0 {
				if kmsUsed == nil {
					t.Fatal("kmsUsed should not be nil")
				}

				if kmsUsed.kmsTimeoutSum != testCase.expectedTimeout {
					t.Fatalf("expected timeout %v, got %v", testCase.expectedTimeout, kmsUsed.kmsTimeoutSum)
				}
			}

		})
	}
}

func TestKMSPluginHealthz(t *testing.T) {
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.KMSv2, true)()

	kmsv2Probe := &kmsv2PluginProbe{
		name: "foo",
		ttl:  3 * time.Second,
	}
	keyID := "1"
	kmsv2Probe.keyID.Store(&keyID)

	testCases := []struct {
		desc    string
		config  string
		want    []healthChecker
		wantErr string
		kmsv2   bool
		kmsv1   bool
	}{
		{
			desc:    "Invalid config file path",
			config:  "invalid/path",
			want:    nil,
			wantErr: `error opening encryption provider configuration file "invalid/path"`,
		},
		{
			desc:    "Empty config file content",
			config:  "testdata/invalid-configs/kms/invalid-content.yaml",
			want:    nil,
			wantErr: `encryption provider configuration file "testdata/invalid-configs/kms/invalid-content.yaml" is empty`,
		},
		{
			desc:    "Unable to decode",
			config:  "testdata/invalid-configs/kms/invalid-gvk.yaml",
			want:    nil,
			wantErr: `error decoding encryption provider configuration file`,
		},
		{
			desc:    "Unexpected config type",
			config:  "testdata/invalid-configs/kms/invalid-config-type.yaml",
			want:    nil,
			wantErr: `no kind "EncryptionConfigurations" is registered for version "apiserver.config.k8s.io/v1"`,
		},
		{
			desc:   "Install Healthz",
			config: "testdata/valid-configs/kms/default-timeout.yaml",
			want: []healthChecker{
				&kmsPluginProbe{
					name: "foo",
					ttl:  3 * time.Second,
				},
			},
			kmsv1: true,
		},
		{
			desc:   "Install multiple healthz",
			config: "testdata/valid-configs/kms/multiple-providers.yaml",
			want: []healthChecker{
				&kmsPluginProbe{
					name: "foo",
					ttl:  3 * time.Second,
				},
				&kmsPluginProbe{
					name: "bar",
					ttl:  3 * time.Second,
				},
			},
			kmsv1: true,
		},
		{
			desc:   "No KMS Providers",
			config: "testdata/valid-configs/aes/aes-gcm.yaml",
		},
		{
			desc:   "Install multiple healthz with v1 and v2",
			config: "testdata/valid-configs/kms/multiple-providers-kmsv2.yaml",
			want: []healthChecker{
				kmsv2Probe,
				&kmsPluginProbe{
					name: "bar",
					ttl:  3 * time.Second,
				},
			},
			kmsv2: true,
			kmsv1: true,
		},
		{
			desc:    "Invalid API version",
			config:  "testdata/invalid-configs/kms/invalid-apiversion.yaml",
			want:    nil,
			wantErr: `resources[0].providers[0].kms.apiVersion: Invalid value: "v3": unsupported apiVersion apiVersion for KMS provider, only v1 and v2 are supported`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			config, _, err := loadConfig(tt.config, false)
			if errStr := errString(err); !strings.Contains(errStr, tt.wantErr) {
				t.Fatalf("unexpected error state got=%s want=%s", errStr, tt.wantErr)
			}
			if len(tt.wantErr) > 0 {
				return
			}

			ctx, cancel := context.WithCancel(context.Background())
			cancel() // cancel this upfront so the kms v2 healthz check poll does not run
			_, got, kmsUsed, err := getTransformerOverridesAndKMSPluginProbes(ctx, config)
			if err != nil {
				t.Fatal(err)
			}

			// unset fields that are not relevant to the test
			for i := range got {
				checker := got[i]
				switch p := checker.(type) {
				case *kmsPluginProbe:
					p.service = nil
					p.l = nil
					p.lastResponse = nil
				case *kmsv2PluginProbe:
					p.service = nil
					p.l = nil
					p.lastResponse = nil
					p.keyID = kmsv2Probe.keyID
				default:
					t.Fatalf("unexpected probe type %T", p)
				}
			}

			if tt.kmsv2 != kmsUsed.v2Used {
				t.Errorf("incorrect kms v2 detection: want=%v got=%v", tt.kmsv2, kmsUsed.v2Used)
			}
			if tt.kmsv1 != kmsUsed.v1Used {
				t.Errorf("incorrect kms v1 detection: want=%v got=%v", tt.kmsv1, kmsUsed.v1Used)
			}

			if d := cmp.Diff(tt.want, got,
				cmp.Comparer(func(a, b *kmsPluginProbe) bool {
					return *a == *b
				}),
				cmp.Comparer(func(a, b *kmsv2PluginProbe) bool {
					return *a == *b
				}),
			); d != "" {
				t.Fatalf("HealthzConfig mismatch (-want +got):\n%s", d)
			}
		})
	}
}

func TestKMSPluginHealthzTTL(t *testing.T) {
	ctx := testContext(t)

	service, _ := newMockEnvelopeService(ctx, "unix:///tmp/testprovider.sock", 3*time.Second)
	errService, _ := newMockErrorEnvelopeService("unix:///tmp/testprovider.sock", 3*time.Second)

	testCases := []struct {
		desc    string
		probe   *kmsPluginProbe
		wantTTL time.Duration
	}{
		{
			desc: "kms provider in good state",
			probe: &kmsPluginProbe{
				name:         "test",
				ttl:          kmsPluginHealthzNegativeTTL,
				service:      service,
				l:            &sync.Mutex{},
				lastResponse: &kmsPluginHealthzResponse{},
			},
			wantTTL: kmsPluginHealthzPositiveTTL,
		},
		{
			desc: "kms provider in bad state",
			probe: &kmsPluginProbe{
				name:         "test",
				ttl:          kmsPluginHealthzPositiveTTL,
				service:      errService,
				l:            &sync.Mutex{},
				lastResponse: &kmsPluginHealthzResponse{},
			},
			wantTTL: kmsPluginHealthzNegativeTTL,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			_ = tt.probe.check()
			if tt.probe.ttl != tt.wantTTL {
				t.Fatalf("want ttl %v, got ttl %v", tt.wantTTL, tt.probe.ttl)
			}
		})
	}
}

func TestKMSv2PluginHealthzTTL(t *testing.T) {
	ctx := testContext(t)

	service, _ := newMockEnvelopeKMSv2Service(ctx, "unix:///tmp/testprovider.sock", "providerName", 3*time.Second)
	errService, _ := newMockErrorEnvelopeKMSv2Service("unix:///tmp/testprovider.sock", 3*time.Second)

	testCases := []struct {
		desc    string
		probe   *kmsv2PluginProbe
		wantTTL time.Duration
	}{
		{
			desc: "kmsv2 provider in good state",
			probe: &kmsv2PluginProbe{
				name:         "test",
				ttl:          kmsPluginHealthzNegativeTTL,
				service:      service,
				l:            &sync.Mutex{},
				lastResponse: &kmsPluginHealthzResponse{},
			},
			wantTTL: kmsPluginHealthzPositiveTTL,
		},
		{
			desc: "kmsv2 provider in bad state",
			probe: &kmsv2PluginProbe{
				name:         "test",
				ttl:          kmsPluginHealthzPositiveTTL,
				service:      errService,
				l:            &sync.Mutex{},
				lastResponse: &kmsPluginHealthzResponse{},
			},
			wantTTL: kmsPluginHealthzNegativeTTL,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			_ = tt.probe.check(ctx)
			if tt.probe.ttl != tt.wantTTL {
				t.Fatalf("want ttl %v, got ttl %v", tt.wantTTL, tt.probe.ttl)
			}
		})
	}
}

func TestKMSv2InvalidKeyID(t *testing.T) {
	ctx := testContext(t)
	invalidKeyIDService, _ := newMockInvalidKeyIDEnvelopeKMSv2Service(ctx, "unix:///tmp/testprovider.sock", 3*time.Second, "")
	invalidLongKeyIDService, _ := newMockInvalidKeyIDEnvelopeKMSv2Service(ctx, "unix:///tmp/testprovider.sock", 3*time.Second, sampleInvalidKeyID)
	service, _ := newMockInvalidKeyIDEnvelopeKMSv2Service(ctx, "unix:///tmp/testprovider.sock", 3*time.Second, "1")

	testCases := []struct {
		desc    string
		probe   *kmsv2PluginProbe
		metrics []string
		want    string
	}{
		{
			desc: "kmsv2 provider returns an invalid empty keyID",
			probe: &kmsv2PluginProbe{
				name:         "test",
				ttl:          kmsPluginHealthzNegativeTTL,
				service:      invalidKeyIDService,
				l:            &sync.Mutex{},
				lastResponse: &kmsPluginHealthzResponse{},
			},
			metrics: []string{
				"apiserver_envelope_encryption_invalid_key_id_from_status_total",
			},
			want: `
			# HELP apiserver_envelope_encryption_invalid_key_id_from_status_total [ALPHA] Number of times an invalid keyID is returned by the Status RPC call split by error.
			# TYPE apiserver_envelope_encryption_invalid_key_id_from_status_total counter
			apiserver_envelope_encryption_invalid_key_id_from_status_total{error="empty",provider_name="test"} 1
			`,
		},
		{
			desc: "kmsv2 provider returns a valid keyID",
			probe: &kmsv2PluginProbe{
				name:         "test",
				ttl:          kmsPluginHealthzNegativeTTL,
				service:      service,
				l:            &sync.Mutex{},
				lastResponse: &kmsPluginHealthzResponse{},
			},
			metrics: []string{
				"apiserver_envelope_encryption_invalid_key_id_from_status_total",
			},
			want: ``,
		},
		{
			desc: "kmsv2 provider returns an invalid long keyID",
			probe: &kmsv2PluginProbe{
				name:         "test",
				ttl:          kmsPluginHealthzNegativeTTL,
				service:      invalidLongKeyIDService,
				l:            &sync.Mutex{},
				lastResponse: &kmsPluginHealthzResponse{},
			},
			metrics: []string{
				"apiserver_envelope_encryption_invalid_key_id_from_status_total",
			},
			want: `
			# HELP apiserver_envelope_encryption_invalid_key_id_from_status_total [ALPHA] Number of times an invalid keyID is returned by the Status RPC call split by error.
			# TYPE apiserver_envelope_encryption_invalid_key_id_from_status_total counter
			apiserver_envelope_encryption_invalid_key_id_from_status_total{error="too_long",provider_name="test"} 1
			`,
		},
	}

	metrics.InvalidKeyIDFromStatusTotal.Reset()
	metrics.RegisterMetrics()

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			defer metrics.InvalidKeyIDFromStatusTotal.Reset()
			_ = tt.probe.check(ctx)
			if err := testutil.GatherAndCompare(legacyregistry.DefaultGatherer, strings.NewReader(tt.want), tt.metrics...); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestCBCKeyRotationWithOverlappingProviders(t *testing.T) {
	testCBCKeyRotationWithProviders(
		t,
		"testdata/valid-configs/aes/aes-cbc-multiple-providers.json",
		"k8s:enc:aescbc:v1:1:",
		"testdata/valid-configs/aes/aes-cbc-multiple-providers-reversed.json",
		"k8s:enc:aescbc:v1:2:",
	)
}

func TestCBCKeyRotationWithoutOverlappingProviders(t *testing.T) {
	testCBCKeyRotationWithProviders(
		t,
		"testdata/valid-configs/aes/aes-cbc-multiple-keys.json",
		"k8s:enc:aescbc:v1:A:",
		"testdata/valid-configs/aes/aes-cbc-multiple-keys-reversed.json",
		"k8s:enc:aescbc:v1:B:",
	)
}

func testCBCKeyRotationWithProviders(t *testing.T, firstEncryptionConfig, firstPrefix, secondEncryptionConfig, secondPrefix string) {
	p := getTransformerFromEncryptionConfig(t, firstEncryptionConfig)

	ctx := testContext(t)
	dataCtx := value.DefaultContext([]byte("authenticated_data"))

	out, err := p.TransformToStorage(ctx, []byte("firstvalue"), dataCtx)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(out, []byte(firstPrefix)) {
		t.Fatalf("unexpected prefix: %q", out)
	}
	from, stale, err := p.TransformFromStorage(ctx, out, dataCtx)
	if err != nil {
		t.Fatal(err)
	}
	if stale || !bytes.Equal([]byte("firstvalue"), from) {
		t.Fatalf("unexpected data: %t %q", stale, from)
	}

	// verify changing the context fails storage
	_, _, err = p.TransformFromStorage(ctx, out, value.DefaultContext([]byte("incorrect_context")))
	if err != nil {
		t.Fatalf("CBC mode does not support authentication: %v", err)
	}

	// reverse the order, use the second key
	p = getTransformerFromEncryptionConfig(t, secondEncryptionConfig)
	from, stale, err = p.TransformFromStorage(ctx, out, dataCtx)
	if err != nil {
		t.Fatal(err)
	}
	if !stale || !bytes.Equal([]byte("firstvalue"), from) {
		t.Fatalf("unexpected data: %t %q", stale, from)
	}

	out, err = p.TransformToStorage(ctx, []byte("firstvalue"), dataCtx)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(out, []byte(secondPrefix)) {
		t.Fatalf("unexpected prefix: %q", out)
	}
	from, stale, err = p.TransformFromStorage(ctx, out, dataCtx)
	if err != nil {
		t.Fatal(err)
	}
	if stale || !bytes.Equal([]byte("firstvalue"), from) {
		t.Fatalf("unexpected data: %t %q", stale, from)
	}
}

func getTransformerFromEncryptionConfig(t *testing.T, encryptionConfigPath string) value.Transformer {
	ctx := testContext(t)

	t.Helper()
	encryptionConfiguration, err := LoadEncryptionConfig(ctx, encryptionConfigPath, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(encryptionConfiguration.Transformers) != 1 {
		t.Fatalf("input config does not have exactly one resource: %s", encryptionConfigPath)
	}
	for _, transformer := range encryptionConfiguration.Transformers {
		return transformer
	}
	panic("unreachable")
}

func TestIsKMSv2ProviderHealthyError(t *testing.T) {
	testCases := []struct {
		desc           string
		expectedErr    string
		statusResponse *kmsservice.StatusResponse
	}{
		{
			desc: "healthz status is not ok",
			statusResponse: &kmsservice.StatusResponse{
				Healthz: "unhealthy",
			},
			expectedErr: "got unexpected healthz status: unhealthy, expected KMSv2 API version v2alpha1, got , expected KMSv2 KeyID to be set, got ",
		},
		{
			desc: "version is not v2alpha1",
			statusResponse: &kmsservice.StatusResponse{
				Version: "v1beta1",
			},
			expectedErr: "got unexpected healthz status: , expected KMSv2 API version v2alpha1, got v1beta1, expected KMSv2 KeyID to be set, got ",
		},
		{
			desc: "missing keyID",
			statusResponse: &kmsservice.StatusResponse{
				Healthz: "ok",
				Version: "v2alpha1",
			},
			expectedErr: "expected KMSv2 KeyID to be set, got ",
		},
		{
			desc: "invalid long keyID",
			statusResponse: &kmsservice.StatusResponse{
				Healthz: "ok",
				Version: "v2alpha1",
				KeyID:   sampleInvalidKeyID,
			},
			expectedErr: "expected KMSv2 KeyID to be set, got ",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			err := isKMSv2ProviderHealthy("testplugin", tt.statusResponse)
			if !strings.Contains(errString(err), tt.expectedErr) {
				t.Errorf("expected err %q, got %q", tt.expectedErr, errString(err))
			}
		})
	}
}

func testContext(t *testing.T) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	return ctx
}

func errString(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()
}

func TestComputeEncryptionConfigHash(t *testing.T) {
	// hash the empty string to be sure that sha256 is being used
	expect := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	sum := computeEncryptionConfigHash([]byte(""))
	if expect != sum {
		t.Errorf("expected hash %q but got %q", expect, sum)
	}
}

func TestGetCurrentKeyID(t *testing.T) {
	ctx := testContext(t)
	kmsv2Probe := &kmsv2PluginProbe{
		name: "foo",
		ttl:  3 * time.Second,
	}
	testCases := []struct {
		desc        string
		keyID       string
		expectedErr string
	}{
		{
			desc:        "empty keyID",
			keyID:       "",
			expectedErr: "got unexpected empty keyID",
		},
		{
			desc:        "valid keyID",
			keyID:       "1",
			expectedErr: "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			kmsv2Probe.keyID.Store(&tt.keyID)
			_, err := kmsv2Probe.getCurrentKeyID(ctx)
			if errString(err) != tt.expectedErr {
				t.Errorf("expected err %q, got %q", tt.expectedErr, errString(err))
			}
		})
	}
}
