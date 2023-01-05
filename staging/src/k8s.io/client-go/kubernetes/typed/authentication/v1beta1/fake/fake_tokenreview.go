/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "k8s.io/api/authentication/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testing "k8s.io/client-go/testing"
)

// FakeTokenReviews implements TokenReviewInterface
type FakeTokenReviews struct {
	Fake *FakeAuthenticationV1beta1
}

var tokenreviewsResource = v1beta1.SchemeGroupVersion.WithResource("tokenreviews")

var tokenreviewsKind = v1beta1.SchemeGroupVersion.WithKind("TokenReview")

// Create takes the representation of a tokenReview and creates it.  Returns the server's representation of the tokenReview, and an error, if there is any.
func (c *FakeTokenReviews) Create(ctx context.Context, tokenReview *v1beta1.TokenReview, opts v1.CreateOptions) (result *v1beta1.TokenReview, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(tokenreviewsResource, tokenReview), &v1beta1.TokenReview{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TokenReview), err
}
