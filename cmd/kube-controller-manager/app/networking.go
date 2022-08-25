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

// Package app implements a server that runs a set of active
// components.  This includes replication controllers, service endpoints and
// nodes.
package app

import (
	"context"

	"k8s.io/controller-manager/controller"
	"k8s.io/kubernetes/pkg/controller/servicecidrs"
)

func startServiceCIDRsController(ctx context.Context, controllerContext ControllerContext) (controller.Interface, bool, error) {
	go servicecidrs.NewController(
		controllerContext.InformerFactory.Networking().V1alpha1().ServiceCIDRs(),
		controllerContext.InformerFactory.Networking().V1alpha1().IPAddresses(),
		controllerContext.ClientBuilder.ClientOrDie("service-cidrs-controller"),
	).Run(ctx, 5)
	// TODO use component config
	return nil, true, nil

}
