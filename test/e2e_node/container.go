/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package e2e_node

import (
	"errors"
	"fmt"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

//One pod one container
type ConformanceContainer struct {
	Container api.Container
	Client    *client.Client
	Status    string
}

type ConformanceContainerEqualMatcher struct {
	Expected interface{}
}

func CContainerEqual(expected interface{}) types.GomegaMatcher {
	return &ConformanceContainerEqualMatcher{
		Expected: expected,
	}
}

func (matcher *ConformanceContainerEqualMatcher) Match(actual interface{}) (bool, error) {
	if actual == nil && matcher.Expected == nil {
		return false, fmt.Errorf("Refusing to compare <nil> to <nil>.\nBe explicit and use BeNil() instead.  This is to avoid mistakes where both sides of an assertion are erroneously uninitialized.")
	}
	val := api.Semantic.DeepDerivative(matcher.Expected, actual)
	return val, nil
}

func (matcher *ConformanceContainerEqualMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to equal", matcher.Expected)
}

func (matcher *ConformanceContainerEqualMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to equal", matcher.Expected)
}

func (cc *ConformanceContainer) Create() error {
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			//Same with the container name
			Name:      cc.Container.Name,
			Namespace: api.NamespaceDefault,
		},
		Spec: api.PodSpec{
			NodeName:      *nodeName,
			RestartPolicy: api.RestartPolicyNever,
			Containers: []api.Container{
				cc.Container,
			},
		},
	}

	_, err := cc.Client.Pods(api.NamespaceDefault).Create(pod)
	return err
}

func (cc *ConformanceContainer) IsStarted() bool {
	if ccontainer, err := cc.Get(); err != nil {
		return false
	} else if ccontainer.Status == "running" || ccontainer.Status == "terminated" {
		return true
	} else {
		return false
	}
}

//Same with 'delete'
func (cc *ConformanceContainer) Stop() error {
	return cc.Client.Pods(api.NamespaceDefault).Delete(cc.Container.Name, &api.DeleteOptions{})
}

func (cc *ConformanceContainer) Delete() error {
	return cc.Client.Pods(api.NamespaceDefault).Delete(cc.Container.Name, &api.DeleteOptions{})
}

func (cc *ConformanceContainer) Get() (ConformanceContainer, error) {
	pod, err := cc.Client.Pods(api.NamespaceDefault).Get(cc.Container.Name)
	if err != nil {
		return ConformanceContainer{}, err
	}

	containers := pod.Spec.Containers
	if containers == nil || len(containers) != 1 {
		return ConformanceContainer{}, errors.New("Failed to get container")
	}

	cstatuss := pod.Status.ContainerStatuses
	if cstatuss == nil || len(cstatuss) != 1 {
		return ConformanceContainer{}, errors.New("Failed to get container status")
	}

	var status string
	if cstatuss[0].State.Running != nil {
		status = "running"
	} else if cstatuss[0].State.Terminated != nil {
		status = "terminated"
	} else {
		status = "waiting"
	}

	return ConformanceContainer{containers[0], cc.Client, status}, nil
}
