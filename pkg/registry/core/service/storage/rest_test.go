/*
Copyright 2014 The Kubernetes Authors.

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

package storage

import (
	"net"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	etcd3testing "k8s.io/apiserver/pkg/storage/etcd3/testing"
	svctest "k8s.io/kubernetes/pkg/api/service/testing"
	api "k8s.io/kubernetes/pkg/apis/core"
	endpointstore "k8s.io/kubernetes/pkg/registry/core/endpoint/storage"
	"k8s.io/kubernetes/pkg/registry/core/service/ipallocator"
	"k8s.io/kubernetes/pkg/registry/core/service/portallocator"
	"k8s.io/kubernetes/pkg/registry/registrytest"
	netutils "k8s.io/utils/net"
)

func NewTestREST(t *testing.T, ipFamilies []api.IPFamily) (*GenericREST, *etcd3testing.EtcdTestServer) {
	etcdStorage, server := registrytest.NewEtcdStorage(t, "")

	var rPrimary ipallocator.Interface
	var rSecondary ipallocator.Interface

	if len(ipFamilies) < 1 || len(ipFamilies) > 2 {
		t.Fatalf("unexpected ipfamilies passed: %v", ipFamilies)
	}
	for i, family := range ipFamilies {
		var r ipallocator.Interface
		var err error
		switch family {
		case api.IPv4Protocol:
			r, err = ipallocator.NewInMemory(makeIPNet(t))
			if err != nil {
				t.Fatalf("cannot create CIDR Range %v", err)
			}
		case api.IPv6Protocol:
			r, err = ipallocator.NewInMemory(makeIPNet6(t))
			if err != nil {
				t.Fatalf("cannot create CIDR Range %v", err)
			}
		}
		switch i {
		case 0:
			rPrimary = r
		case 1:
			rSecondary = r
		}
	}

	portRange := utilnet.PortRange{Base: 30000, Size: 1000}
	portAllocator, err := portallocator.NewInMemory(portRange)
	if err != nil {
		t.Fatalf("cannot create port allocator %v", err)
	}

	ipAllocators := map[api.IPFamily]ipallocator.Interface{
		rPrimary.IPFamily(): rPrimary,
	}
	if rSecondary != nil {
		ipAllocators[rSecondary.IPFamily()] = rSecondary
	}

	restOptions := generic.RESTOptions{
		StorageConfig:           etcdStorage.ForResource(schema.GroupResource{Resource: "services"}),
		Decorator:               generic.UndecoratedStorage,
		DeleteCollectionWorkers: 1,
		ResourcePrefix:          "services",
	}
	endpoints, err := endpointstore.NewREST(generic.RESTOptions{
		StorageConfig:  etcdStorage,
		Decorator:      generic.UndecoratedStorage,
		ResourcePrefix: "endpoints",
	})

	rest, _, _, err := NewREST(restOptions, api.IPv4Protocol, ipAllocators, portAllocator, endpoints, nil, nil)
	if err != nil {
		t.Fatalf("unexpected error from REST storage: %v", err)
	}

	return rest, server
}

func makeIPNet(t *testing.T) *net.IPNet {
	_, net, err := netutils.ParseCIDRSloppy("1.2.3.0/24")
	if err != nil {
		t.Error(err)
	}
	return net
}
func makeIPNet6(t *testing.T) *net.IPNet {
	_, net, err := netutils.ParseCIDRSloppy("2000::/108")
	if err != nil {
		t.Error(err)
	}
	return net
}

// this is local because it's not fully fleshed out enough for general use.
func makePod(name string, ips ...string) api.Pod {
	p := api.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: metav1.NamespaceDefault,
		},
		Spec: api.PodSpec{
			RestartPolicy: api.RestartPolicyAlways,
			DNSPolicy:     api.DNSDefault,
			Containers:    []api.Container{{Name: "ctr", Image: "img", ImagePullPolicy: api.PullIfNotPresent, TerminationMessagePolicy: api.TerminationMessageReadFile}},
		},
		Status: api.PodStatus{
			PodIPs: []api.PodIP{},
		},
	}

	for _, ip := range ips {
		p.Status.PodIPs = append(p.Status.PodIPs, api.PodIP{IP: ip})
	}

	return p
}

// Validate the internalTrafficPolicy field when set to "Cluster" then updated to "Local"
func TestServiceRegistryInternalTrafficPolicyClusterThenLocal(t *testing.T) {
	ctx := genericapirequest.NewDefaultContext()
	storage, server := NewTestREST(t, []api.IPFamily{api.IPv4Protocol})
	defer server.Terminate(t)
	svc := svctest.MakeService("internal-traffic-policy-cluster",
		svctest.SetInternalTrafficPolicy(api.ServiceInternalTrafficPolicyCluster),
	)
	obj, err := storage.Create(ctx, svc, rest.ValidateAllObjectFunc, &metav1.CreateOptions{})
	if obj == nil || err != nil {
		t.Errorf("Unexpected failure creating service %v", err)
	}

	createdSvc := obj.(*api.Service)
	if *createdSvc.Spec.InternalTrafficPolicy != api.ServiceInternalTrafficPolicyCluster {
		t.Errorf("Expecting internalTrafficPolicy field to have value Cluster, got: %s", *createdSvc.Spec.InternalTrafficPolicy)
	}

	update := createdSvc.DeepCopy()
	local := api.ServiceInternalTrafficPolicyLocal
	update.Spec.InternalTrafficPolicy = &local

	updatedSvc, _, errUpdate := storage.Update(ctx, update.Name, rest.DefaultUpdatedObjectInfo(update), rest.ValidateAllObjectFunc, rest.ValidateAllObjectUpdateFunc, false, &metav1.UpdateOptions{})
	if errUpdate != nil {
		t.Fatalf("unexpected error during update %v", errUpdate)
	}
	updatedService := updatedSvc.(*api.Service)
	if *updatedService.Spec.InternalTrafficPolicy != api.ServiceInternalTrafficPolicyLocal {
		t.Errorf("Expected internalTrafficPolicy to be Local, got: %s", *updatedService.Spec.InternalTrafficPolicy)
	}
}

// Validate the internalTrafficPolicy field when set to "Local" and then updated to "Cluster"
func TestServiceRegistryInternalTrafficPolicyLocalThenCluster(t *testing.T) {
	ctx := genericapirequest.NewDefaultContext()
	storage, server := NewTestREST(t, []api.IPFamily{api.IPv4Protocol})
	defer server.Terminate(t)
	svc := svctest.MakeService("internal-traffic-policy-cluster",
		svctest.SetInternalTrafficPolicy(api.ServiceInternalTrafficPolicyLocal),
	)
	obj, err := storage.Create(ctx, svc, rest.ValidateAllObjectFunc, &metav1.CreateOptions{})
	if obj == nil || err != nil {
		t.Errorf("Unexpected failure creating service %v", err)
	}

	createdSvc := obj.(*api.Service)
	if *createdSvc.Spec.InternalTrafficPolicy != api.ServiceInternalTrafficPolicyLocal {
		t.Errorf("Expecting internalTrafficPolicy field to have value Local, got: %s", *createdSvc.Spec.InternalTrafficPolicy)
	}

	update := createdSvc.DeepCopy()
	cluster := api.ServiceInternalTrafficPolicyCluster
	update.Spec.InternalTrafficPolicy = &cluster

	updatedSvc, _, errUpdate := storage.Update(ctx, update.Name, rest.DefaultUpdatedObjectInfo(update), rest.ValidateAllObjectFunc, rest.ValidateAllObjectUpdateFunc, false, &metav1.UpdateOptions{})
	if errUpdate != nil {
		t.Fatalf("unexpected error during update %v", errUpdate)
	}
	updatedService := updatedSvc.(*api.Service)
	if *updatedService.Spec.InternalTrafficPolicy != api.ServiceInternalTrafficPolicyCluster {
		t.Errorf("Expected internalTrafficPolicy to be Cluster, got: %s", *updatedService.Spec.InternalTrafficPolicy)
	}
}
