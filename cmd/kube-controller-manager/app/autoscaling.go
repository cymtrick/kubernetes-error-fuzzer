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

// Package app implements a server that runs a set of active
// components.  This includes replication controllers, service endpoints and
// nodes.
//
package app

import (
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	discocache "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/scale"
	"k8s.io/kubernetes/pkg/controller/podautoscaler"
	"k8s.io/kubernetes/pkg/controller/podautoscaler/metrics"
	resourceclient "k8s.io/metrics/pkg/client/clientset_generated/clientset/typed/metrics/v1beta1"
	"k8s.io/metrics/pkg/client/custom_metrics"
	"k8s.io/metrics/pkg/client/external_metrics"
)

func startHPAController(ctx ControllerContext) (bool, error) {
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "autoscaling", Version: "v1", Resource: "horizontalpodautoscalers"}] {
		return false, nil
	}

	if ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerUseRESTClients {
		// use the new-style clients if support for custom metrics is enabled
		return startHPAControllerWithRESTClient(ctx)
	}

	return startHPAControllerWithLegacyClient(ctx)
}

func startHPAControllerWithRESTClient(ctx ControllerContext) (bool, error) {
	clientConfig := ctx.ClientBuilder.ConfigOrDie("horizontal-pod-autoscaler")
	metricsClient := metrics.NewRESTMetricsClient(
		resourceclient.NewForConfigOrDie(clientConfig),
		custom_metrics.NewForConfigOrDie(clientConfig),
		external_metrics.NewForConfigOrDie(clientConfig),
	)
	return startHPAControllerWithMetricsClient(ctx, metricsClient)
}

func startHPAControllerWithLegacyClient(ctx ControllerContext) (bool, error) {
	hpaClient := ctx.ClientBuilder.ClientOrDie("horizontal-pod-autoscaler")
	metricsClient := metrics.NewHeapsterMetricsClient(
		hpaClient,
		metrics.DefaultHeapsterNamespace,
		metrics.DefaultHeapsterScheme,
		metrics.DefaultHeapsterService,
		metrics.DefaultHeapsterPort,
	)
	return startHPAControllerWithMetricsClient(ctx, metricsClient)
}

func startHPAControllerWithMetricsClient(ctx ControllerContext, metricsClient metrics.MetricsClient) (bool, error) {
	hpaClientGoClient := ctx.ClientBuilder.ClientGoClientOrDie("horizontal-pod-autoscaler")
	hpaClient := ctx.ClientBuilder.ClientOrDie("horizontal-pod-autoscaler")
	hpaClientConfig := ctx.ClientBuilder.ConfigOrDie("horizontal-pod-autoscaler")

	// TODO: we need something like deferred discovery REST mapper that calls invalidate
	// on cache misses.
	cachedDiscovery := discocache.NewMemCacheClient(hpaClientGoClient.Discovery())
	restMapper := discovery.NewDeferredDiscoveryRESTMapper(cachedDiscovery, apimeta.InterfacesForUnstructured)
	restMapper.Reset()
	// we don't use cached discovery because DiscoveryScaleKindResolver does its own caching,
	// so we want to re-fetch every time when we actually ask for it
	scaleKindResolver := scale.NewDiscoveryScaleKindResolver(hpaClientGoClient.Discovery())
	scaleClient, err := scale.NewForConfig(hpaClientConfig, restMapper, dynamic.LegacyAPIPathResolverFunc, scaleKindResolver)
	if err != nil {
		return false, err
	}

	replicaCalc := podautoscaler.NewReplicaCalculator(
		metricsClient,
		hpaClient.CoreV1(),
		ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerTolerance,
	)
	go podautoscaler.NewHorizontalController(
		hpaClientGoClient.CoreV1(),
		scaleClient,
		hpaClient.AutoscalingV1(),
		restMapper,
		replicaCalc,
		ctx.InformerFactory.Autoscaling().V1().HorizontalPodAutoscalers(),
		ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerSyncPeriod.Duration,
		ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerUpscaleForbiddenWindow.Duration,
		ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerDownscaleForbiddenWindow.Duration,
	).Run(ctx.Stop)
	return true, nil
}
