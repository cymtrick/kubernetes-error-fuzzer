/*
Copyright 2020 The Kubernetes Authors.

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

package apimachinery

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/test/e2e/framework"

	"github.com/onsi/ginkgo"
)

var _ = SIGDescribe("apigroup preferred version", func() {
	f := framework.NewDefaultFramework("apigroup-preferred-version")
	ginkgo.It("should validate PreferredVersion for each APIGroup", func() {

		// get list of APIGroup endpoints
		list := &metav1.APIGroupList{}
		err := f.ClientSet.Discovery().RESTClient().Get().AbsPath("/apis/").Do(context.TODO()).Into(list)

		framework.ExpectNoError(err, "Failed to find /apis/")

		for _, group := range list.Groups {
			framework.Logf("Checking APIGroup: %v", group.Name)

			// locate APIGroup endpoint
			checkGroup := &metav1.APIGroup{}
			apiPath := "/apis/" + group.Name + "/"
			err = f.ClientSet.Discovery().RESTClient().Get().AbsPath(apiPath).Do(context.TODO()).Into(checkGroup)
			framework.ExpectNoError(err, "Fail to access: %s", apiPath)

			framework.Logf("PreferredVersion.GroupVersion: %s", checkGroup.PreferredVersion.GroupVersion)
			framework.Logf("Versions found %v", checkGroup.Versions)

			// confirm that the PreferredVersion is a valid version
			match := false
			for _, version := range checkGroup.Versions {
				if version.GroupVersion == checkGroup.PreferredVersion.GroupVersion {
					framework.Logf("%s matches %s", version.GroupVersion, checkGroup.PreferredVersion.GroupVersion)
					match = true
					break
				}
			}
			framework.ExpectEqual(true, match, "failed to find a valid version for PreferredVersion")
		}
	})
})
