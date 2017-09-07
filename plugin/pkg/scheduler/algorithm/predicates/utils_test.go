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

package predicates

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// ExampleUtils is a https://blog.golang.org/examples styled unit test.
func ExampleFindLabelsInSet() {
	labelSubset := labels.Set{}
	labelSubset["label1"] = "value1"
	labelSubset["label2"] = "value2"
	// Lets make believe that these pods are on the cluster.
	// Utility functions will inspect their labels, filter them, and so on.
	nsPods := []*v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "pod1",
				Namespace: "ns1",
				Labels: map[string]string{
					"label1": "wontSeeThis",
					"label2": "wontSeeThis",
					"label3": "will_see_this",
				},
			},
		}, // first pod which will be used via the utilities
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "pod2",
				Namespace: "ns1",
			},
		},

		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pod3ThatWeWontSee",
			},
		},
	}
	fmt.Println(FindLabelsInSet([]string{"label1", "label2", "label3"}, nsPods[0].ObjectMeta.Labels)["label3"])
	AddUnsetLabelsToMap(labelSubset, []string{"label1", "label2", "label3"}, nsPods[0].ObjectMeta.Labels)
	fmt.Println(labelSubset)

	for _, pod := range FilterPodsByNamespace(nsPods, "ns1") {
		fmt.Print(pod.Name, ",")
	}
	// Output:
	// will_see_this
	// label1=value1,label2=value2,label3=will_see_this
	// pod1,pod2,
}

func Test_decode(t *testing.T) {
	tests := []struct {
		name string
		args string
		want *hostPortInfo
	}{
		{
			name: "test1",
			args: "UDP/127.0.0.1/80",
			want: &hostPortInfo{
				protocol: "UDP",
				hostIP:   "127.0.0.1",
				hostPort: "80",
			},
		},
		{
			name: "test2",
			args: "TCP/127.0.0.1/80",
			want: &hostPortInfo{
				protocol: "TCP",
				hostIP:   "127.0.0.1",
				hostPort: "80",
			},
		},
		{
			name: "test3",
			args: "TCP/0.0.0.0/80",
			want: &hostPortInfo{
				protocol: "TCP",
				hostIP:   "0.0.0.0",
				hostPort: "80",
			},
		},
	}

	for _, tt := range tests {
		if got := decode(tt.args); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("test name = %v, decode() = %v, want %v", tt.name, got, tt.want)
		}

	}
}
