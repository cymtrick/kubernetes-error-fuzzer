//go:build example
// +build example

/*
Copyright 2023 The Kubernetes Authors.

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

package gomega

// The tests below will fail and therefore are excluded from
// normal "make test" via the "example" build tag. To run
// the tests and check the output, use "go test -tags example ."

import (
	"context"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"k8s.io/kubernetes/test/utils/ktesting"
)

func TestGomega(t *testing.T) {
	tCtx := ktesting.Init(t)

	gomega.NewWithT(tCtx).Eventually(tCtx, func(ctx context.Context) int {
		// TODO: tCtx = ktesting.WithContext(tCtx, ctx)
		// Or some dedicated tCtx.Eventually?

		return 42
	}).WithPolling(time.Second).Should(gomega.Equal(1))
}
