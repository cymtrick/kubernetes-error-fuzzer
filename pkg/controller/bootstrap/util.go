/*
Copyright 2016 The Kubernetes Authors.

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

package bootstrap

import (
	"time"

	"github.com/golang/glog"

	"k8s.io/client-go/pkg/api/v1"
	bootstrapapi "k8s.io/kubernetes/pkg/bootstrap/api"
)

// getSecretString gets a string value from a secret.  If there is an error or
// if the key doesn't exist, an empty string is returned.
func getSecretString(secret *v1.Secret, key string) string {
	data, ok := secret.Data[key]
	if !ok {
		return ""
	}

	return string(data)
}

func validateSecretForSigning(secret *v1.Secret) (tokenID, tokenSecret string, ok bool) {
	tokenID = getSecretString(secret, bootstrapapi.BootstrapTokenIDKey)
	if len(tokenID) == 0 {
		glog.V(3).Infof("No %s key in %s/%s Secret", bootstrapapi.BootstrapTokenIDKey, secret.Namespace, secret.Name)
		return "", "", false
	}

	tokenSecret = getSecretString(secret, bootstrapapi.BootstrapTokenSecretKey)
	if len(tokenSecret) == 0 {
		glog.V(3).Infof("No %s key in %s/%s Secret", bootstrapapi.BootstrapTokenSecretKey, secret.Namespace, secret.Name)
		return "", "", false
	}

	// Ensure this secret hasn't expired.  The TokenCleaner should remove this
	// but if that isn't working or it hasn't gotten there yet we should check
	// here.
	if isSecretExpired(secret) {
		return "", "", false
	}

	// Make sure this secret can be used for signing
	okToSign := getSecretString(secret, bootstrapapi.BootstrapTokenUsageSigningKey)
	if okToSign != "true" {
		return "", "", false
	}

	return tokenID, tokenSecret, true
}

// isSecretExpired returns true if the Secret is expired.
func isSecretExpired(secret *v1.Secret) bool {
	expiration := getSecretString(secret, bootstrapapi.BootstrapTokenExpirationKey)
	if len(expiration) > 0 {
		expTime, err2 := time.Parse(time.RFC3339, expiration)
		if err2 != nil {
			glog.V(3).Infof("Unparseable expiration time (%s) in %s/%s Secret: %v. Treating as expired.",
				expiration, secret.Namespace, secret.Name, err2)
			return true
		}
		if time.Now().After(expTime) {
			glog.V(3).Infof("Expired bootstrap token in %s/%s Secret: %v",
				secret.Namespace, secret.Name, expiration)
			return true
		}
	}
	return false
}
