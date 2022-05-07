/*
Copyright 2019 The Kubernetes Authors.

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

package nodename

import (
	"context"
	"reflect"
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	st "k8s.io/kubernetes/pkg/scheduler/testing"
)

func TestNodeName(t *testing.T) {
	tests := []struct {
		pod        *v1.Pod
		node       *v1.Node
		name       string
		wantStatus *framework.Status
	}{
		{
			pod:  &v1.Pod{},
			node: &v1.Node{},
			name: "no host specified",
		},
		{
			pod:  st.MakePod().Node("foo").Obj(),
			node: st.MakeNode().Name("foo").Obj(),
			name: "host matches",
		},
		{
			pod:        st.MakePod().Node("bar").Obj(),
			node:       st.MakeNode().Name("foo").Obj(),
			name:       "host doesn't match",
			wantStatus: framework.NewStatus(framework.UnschedulableAndUnresolvable, ErrReason),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			nodeInfo := framework.NewNodeInfo()
			nodeInfo.SetNode(test.node)

			p, _ := New(nil, nil)
			gotStatus := p.(framework.FilterPlugin).Filter(context.Background(), nil, test.pod, nodeInfo)
			if !reflect.DeepEqual(gotStatus, test.wantStatus) {
				t.Errorf("status does not match: %v, want: %v", gotStatus, test.wantStatus)
			}
		})
	}
}
