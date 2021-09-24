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
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"k8s.io/klog/v2"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/quota/v1/generic"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/metadata"
	restclient "k8s.io/client-go/rest"
	cloudnodelifecyclecontroller "k8s.io/cloud-provider/controllers/nodelifecycle"
	routecontroller "k8s.io/cloud-provider/controllers/route"
	servicecontroller "k8s.io/cloud-provider/controllers/service"
	"k8s.io/component-base/metrics/prometheus/ratelimiter"
	"k8s.io/controller-manager/controller"
	csitrans "k8s.io/csi-translation-lib"
	"k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
	pkgcontroller "k8s.io/kubernetes/pkg/controller"
	endpointcontroller "k8s.io/kubernetes/pkg/controller/endpoint"
	"k8s.io/kubernetes/pkg/controller/garbagecollector"
	namespacecontroller "k8s.io/kubernetes/pkg/controller/namespace"
	nodeipamcontroller "k8s.io/kubernetes/pkg/controller/nodeipam"
	nodeipamconfig "k8s.io/kubernetes/pkg/controller/nodeipam/config"
	"k8s.io/kubernetes/pkg/controller/nodeipam/ipam"
	lifecyclecontroller "k8s.io/kubernetes/pkg/controller/nodelifecycle"
	"k8s.io/kubernetes/pkg/controller/podgc"
	replicationcontroller "k8s.io/kubernetes/pkg/controller/replication"
	resourcequotacontroller "k8s.io/kubernetes/pkg/controller/resourcequota"
	serviceaccountcontroller "k8s.io/kubernetes/pkg/controller/serviceaccount"
	"k8s.io/kubernetes/pkg/controller/storageversiongc"
	ttlcontroller "k8s.io/kubernetes/pkg/controller/ttl"
	"k8s.io/kubernetes/pkg/controller/ttlafterfinished"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach"
	"k8s.io/kubernetes/pkg/controller/volume/ephemeral"
	"k8s.io/kubernetes/pkg/controller/volume/expand"
	persistentvolumecontroller "k8s.io/kubernetes/pkg/controller/volume/persistentvolume"
	"k8s.io/kubernetes/pkg/controller/volume/pvcprotection"
	"k8s.io/kubernetes/pkg/controller/volume/pvprotection"
	"k8s.io/kubernetes/pkg/features"
	quotainstall "k8s.io/kubernetes/pkg/quota/v1/install"
	"k8s.io/kubernetes/pkg/volume/csimigration"
	netutils "k8s.io/utils/net"
)

const (
	// defaultNodeMaskCIDRIPv4 is default mask size for IPv4 node cidr
	defaultNodeMaskCIDRIPv4 = 24
	// defaultNodeMaskCIDRIPv6 is default mask size for IPv6 node cidr
	defaultNodeMaskCIDRIPv6 = 64
)

func startServiceController(ctx ControllerContext) (controller.Interface, bool, error) {
	serviceController, err := servicecontroller.New(
		ctx.Cloud,
		ctx.ClientBuilder.ClientOrDie("service-controller"),
		ctx.InformerFactory.Core().V1().Services(),
		ctx.InformerFactory.Core().V1().Nodes(),
		ctx.ComponentConfig.KubeCloudShared.ClusterName,
		utilfeature.DefaultFeatureGate,
	)
	if err != nil {
		// This error shouldn't fail. It lives like this as a legacy.
		klog.Errorf("Failed to start service controller: %v", err)
		return nil, false, nil
	}
	go serviceController.Run(ctx.Stop, int(ctx.ComponentConfig.ServiceController.ConcurrentServiceSyncs))
	return nil, true, nil
}

func startNodeIpamController(ctx ControllerContext) (controller.Interface, bool, error) {
	var serviceCIDR *net.IPNet
	var secondaryServiceCIDR *net.IPNet

	// should we start nodeIPAM
	if !ctx.ComponentConfig.KubeCloudShared.AllocateNodeCIDRs {
		return nil, false, nil
	}

	// failure: bad cidrs in config
	clusterCIDRs, dualStack, err := processCIDRs(ctx.ComponentConfig.KubeCloudShared.ClusterCIDR)
	if err != nil {
		return nil, false, err
	}

	// failure: more than one cidr but they are not configured as dual stack
	if len(clusterCIDRs) > 1 && !dualStack {
		return nil, false, fmt.Errorf("len of ClusterCIDRs==%v and they are not configured as dual stack (at least one from each IPFamily)", len(clusterCIDRs))
	}

	// failure: more than cidrs is not allowed even with dual stack
	if len(clusterCIDRs) > 2 {
		return nil, false, fmt.Errorf("len of clusters is:%v > more than max allowed of 2", len(clusterCIDRs))
	}

	// service cidr processing
	if len(strings.TrimSpace(ctx.ComponentConfig.NodeIPAMController.ServiceCIDR)) != 0 {
		_, serviceCIDR, err = netutils.ParseCIDRSloppy(ctx.ComponentConfig.NodeIPAMController.ServiceCIDR)
		if err != nil {
			klog.Warningf("Unsuccessful parsing of service CIDR %v: %v", ctx.ComponentConfig.NodeIPAMController.ServiceCIDR, err)
		}
	}

	if len(strings.TrimSpace(ctx.ComponentConfig.NodeIPAMController.SecondaryServiceCIDR)) != 0 {
		_, secondaryServiceCIDR, err = netutils.ParseCIDRSloppy(ctx.ComponentConfig.NodeIPAMController.SecondaryServiceCIDR)
		if err != nil {
			klog.Warningf("Unsuccessful parsing of service CIDR %v: %v", ctx.ComponentConfig.NodeIPAMController.SecondaryServiceCIDR, err)
		}
	}

	// the following checks are triggered if both serviceCIDR and secondaryServiceCIDR are provided
	if serviceCIDR != nil && secondaryServiceCIDR != nil {
		// should be dual stack (from different IPFamilies)
		dualstackServiceCIDR, err := netutils.IsDualStackCIDRs([]*net.IPNet{serviceCIDR, secondaryServiceCIDR})
		if err != nil {
			return nil, false, fmt.Errorf("failed to perform dualstack check on serviceCIDR and secondaryServiceCIDR error:%v", err)
		}
		if !dualstackServiceCIDR {
			return nil, false, fmt.Errorf("serviceCIDR and secondaryServiceCIDR are not dualstack (from different IPfamiles)")
		}
	}

	// only --node-cidr-mask-size-ipv4 and --node-cidr-mask-size-ipv6 supported with dual stack clusters.
	// --node-cidr-mask-size flag is incompatible with dual stack clusters.
	nodeCIDRMaskSizes, err := setNodeCIDRMaskSizes(ctx.ComponentConfig.NodeIPAMController, clusterCIDRs)
	if err != nil {
		return nil, false, err
	}

	nodeIpamController, err := nodeipamcontroller.NewNodeIpamController(
		ctx.InformerFactory.Core().V1().Nodes(),
		ctx.Cloud,
		ctx.ClientBuilder.ClientOrDie("node-controller"),
		clusterCIDRs,
		serviceCIDR,
		secondaryServiceCIDR,
		nodeCIDRMaskSizes,
		ipam.CIDRAllocatorType(ctx.ComponentConfig.KubeCloudShared.CIDRAllocatorType),
	)
	if err != nil {
		return nil, true, err
	}
	go nodeIpamController.Run(ctx.Stop)
	return nil, true, nil
}

func startNodeLifecycleController(ctx ControllerContext) (controller.Interface, bool, error) {
	lifecycleController, err := lifecyclecontroller.NewNodeLifecycleController(
		ctx.InformerFactory.Coordination().V1().Leases(),
		ctx.InformerFactory.Core().V1().Pods(),
		ctx.InformerFactory.Core().V1().Nodes(),
		ctx.InformerFactory.Apps().V1().DaemonSets(),
		// node lifecycle controller uses existing cluster role from node-controller
		ctx.ClientBuilder.ClientOrDie("node-controller"),
		ctx.ComponentConfig.KubeCloudShared.NodeMonitorPeriod.Duration,
		ctx.ComponentConfig.NodeLifecycleController.NodeStartupGracePeriod.Duration,
		ctx.ComponentConfig.NodeLifecycleController.NodeMonitorGracePeriod.Duration,
		ctx.ComponentConfig.NodeLifecycleController.PodEvictionTimeout.Duration,
		ctx.ComponentConfig.NodeLifecycleController.NodeEvictionRate,
		ctx.ComponentConfig.NodeLifecycleController.SecondaryNodeEvictionRate,
		ctx.ComponentConfig.NodeLifecycleController.LargeClusterSizeThreshold,
		ctx.ComponentConfig.NodeLifecycleController.UnhealthyZoneThreshold,
		ctx.ComponentConfig.NodeLifecycleController.EnableTaintManager,
	)
	if err != nil {
		return nil, true, err
	}
	go lifecycleController.Run(ctx.Stop)
	return nil, true, nil
}

func startCloudNodeLifecycleController(ctx ControllerContext) (controller.Interface, bool, error) {
	cloudNodeLifecycleController, err := cloudnodelifecyclecontroller.NewCloudNodeLifecycleController(
		ctx.InformerFactory.Core().V1().Nodes(),
		// cloud node lifecycle controller uses existing cluster role from node-controller
		ctx.ClientBuilder.ClientOrDie("node-controller"),
		ctx.Cloud,
		ctx.ComponentConfig.KubeCloudShared.NodeMonitorPeriod.Duration,
	)
	if err != nil {
		// the controller manager should continue to run if the "Instances" interface is not
		// supported, though it's unlikely for a cloud provider to not support it
		klog.Errorf("failed to start cloud node lifecycle controller: %v", err)
		return nil, false, nil
	}

	go cloudNodeLifecycleController.Run(ctx.Stop)
	return nil, true, nil
}

func startRouteController(ctx ControllerContext) (controller.Interface, bool, error) {
	if !ctx.ComponentConfig.KubeCloudShared.AllocateNodeCIDRs || !ctx.ComponentConfig.KubeCloudShared.ConfigureCloudRoutes {
		klog.Infof("Will not configure cloud provider routes for allocate-node-cidrs: %v, configure-cloud-routes: %v.", ctx.ComponentConfig.KubeCloudShared.AllocateNodeCIDRs, ctx.ComponentConfig.KubeCloudShared.ConfigureCloudRoutes)
		return nil, false, nil
	}
	if ctx.Cloud == nil {
		klog.Warning("configure-cloud-routes is set, but no cloud provider specified. Will not configure cloud provider routes.")
		return nil, false, nil
	}
	routes, ok := ctx.Cloud.Routes()
	if !ok {
		klog.Warning("configure-cloud-routes is set, but cloud provider does not support routes. Will not configure cloud provider routes.")
		return nil, false, nil
	}

	// failure: bad cidrs in config
	clusterCIDRs, dualStack, err := processCIDRs(ctx.ComponentConfig.KubeCloudShared.ClusterCIDR)
	if err != nil {
		return nil, false, err
	}

	// failure: more than one cidr but they are not configured as dual stack
	if len(clusterCIDRs) > 1 && !dualStack {
		return nil, false, fmt.Errorf("len of ClusterCIDRs==%v and they are not configured as dual stack (at least one from each IPFamily", len(clusterCIDRs))
	}

	// failure: more than cidrs is not allowed even with dual stack
	if len(clusterCIDRs) > 2 {
		return nil, false, fmt.Errorf("length of clusterCIDRs is:%v more than max allowed of 2", len(clusterCIDRs))
	}

	routeController := routecontroller.New(routes,
		ctx.ClientBuilder.ClientOrDie("route-controller"),
		ctx.InformerFactory.Core().V1().Nodes(),
		ctx.ComponentConfig.KubeCloudShared.ClusterName,
		clusterCIDRs)
	go routeController.Run(ctx.Stop, ctx.ComponentConfig.KubeCloudShared.RouteReconciliationPeriod.Duration)
	return nil, true, nil
}

func startPersistentVolumeBinderController(ctx ControllerContext) (controller.Interface, bool, error) {
	plugins, err := ProbeControllerVolumePlugins(ctx.Cloud, ctx.ComponentConfig.PersistentVolumeBinderController.VolumeConfiguration)
	if err != nil {
		return nil, true, fmt.Errorf("failed to probe volume plugins when starting persistentvolume controller: %v", err)
	}
	filteredDialOptions, err := options.ParseVolumeHostFilters(
		ctx.ComponentConfig.PersistentVolumeBinderController.VolumeHostCIDRDenylist,
		ctx.ComponentConfig.PersistentVolumeBinderController.VolumeHostAllowLocalLoopback)
	if err != nil {
		return nil, true, err
	}
	params := persistentvolumecontroller.ControllerParameters{
		KubeClient:                ctx.ClientBuilder.ClientOrDie("persistent-volume-binder"),
		SyncPeriod:                ctx.ComponentConfig.PersistentVolumeBinderController.PVClaimBinderSyncPeriod.Duration,
		VolumePlugins:             plugins,
		Cloud:                     ctx.Cloud,
		ClusterName:               ctx.ComponentConfig.KubeCloudShared.ClusterName,
		VolumeInformer:            ctx.InformerFactory.Core().V1().PersistentVolumes(),
		ClaimInformer:             ctx.InformerFactory.Core().V1().PersistentVolumeClaims(),
		ClassInformer:             ctx.InformerFactory.Storage().V1().StorageClasses(),
		PodInformer:               ctx.InformerFactory.Core().V1().Pods(),
		NodeInformer:              ctx.InformerFactory.Core().V1().Nodes(),
		EnableDynamicProvisioning: ctx.ComponentConfig.PersistentVolumeBinderController.VolumeConfiguration.EnableDynamicProvisioning,
		FilteredDialOptions:       filteredDialOptions,
	}
	volumeController, volumeControllerErr := persistentvolumecontroller.NewController(params)
	if volumeControllerErr != nil {
		return nil, true, fmt.Errorf("failed to construct persistentvolume controller: %v", volumeControllerErr)
	}
	go volumeController.Run(ctx.Stop)
	return nil, true, nil
}

func startAttachDetachController(ctx ControllerContext) (controller.Interface, bool, error) {
	if ctx.ComponentConfig.AttachDetachController.ReconcilerSyncLoopPeriod.Duration < time.Second {
		return nil, true, fmt.Errorf("duration time must be greater than one second as set via command line option reconcile-sync-loop-period")
	}

	csiNodeInformer := ctx.InformerFactory.Storage().V1().CSINodes()
	csiDriverInformer := ctx.InformerFactory.Storage().V1().CSIDrivers()

	plugins, err := ProbeAttachableVolumePlugins()
	if err != nil {
		return nil, true, fmt.Errorf("failed to probe volume plugins when starting attach/detach controller: %v", err)
	}

	filteredDialOptions, err := options.ParseVolumeHostFilters(
		ctx.ComponentConfig.PersistentVolumeBinderController.VolumeHostCIDRDenylist,
		ctx.ComponentConfig.PersistentVolumeBinderController.VolumeHostAllowLocalLoopback)
	if err != nil {
		return nil, true, err
	}

	attachDetachController, attachDetachControllerErr :=
		attachdetach.NewAttachDetachController(
			ctx.ClientBuilder.ClientOrDie("attachdetach-controller"),
			ctx.InformerFactory.Core().V1().Pods(),
			ctx.InformerFactory.Core().V1().Nodes(),
			ctx.InformerFactory.Core().V1().PersistentVolumeClaims(),
			ctx.InformerFactory.Core().V1().PersistentVolumes(),
			csiNodeInformer,
			csiDriverInformer,
			ctx.InformerFactory.Storage().V1().VolumeAttachments(),
			ctx.Cloud,
			plugins,
			GetDynamicPluginProber(ctx.ComponentConfig.PersistentVolumeBinderController.VolumeConfiguration),
			ctx.ComponentConfig.AttachDetachController.DisableAttachDetachReconcilerSync,
			ctx.ComponentConfig.AttachDetachController.ReconcilerSyncLoopPeriod.Duration,
			attachdetach.DefaultTimerConfig,
			filteredDialOptions,
		)
	if attachDetachControllerErr != nil {
		return nil, true, fmt.Errorf("failed to start attach/detach controller: %v", attachDetachControllerErr)
	}
	go attachDetachController.Run(ctx.Stop)
	return nil, true, nil
}

func startVolumeExpandController(ctx ControllerContext) (controller.Interface, bool, error) {
	if utilfeature.DefaultFeatureGate.Enabled(features.ExpandPersistentVolumes) {
		plugins, err := ProbeExpandableVolumePlugins(ctx.ComponentConfig.PersistentVolumeBinderController.VolumeConfiguration)
		if err != nil {
			return nil, true, fmt.Errorf("failed to probe volume plugins when starting volume expand controller: %v", err)
		}
		csiTranslator := csitrans.New()
		filteredDialOptions, err := options.ParseVolumeHostFilters(
			ctx.ComponentConfig.PersistentVolumeBinderController.VolumeHostCIDRDenylist,
			ctx.ComponentConfig.PersistentVolumeBinderController.VolumeHostAllowLocalLoopback)
		if err != nil {
			return nil, true, err
		}
		expandController, expandControllerErr := expand.NewExpandController(
			ctx.ClientBuilder.ClientOrDie("expand-controller"),
			ctx.InformerFactory.Core().V1().PersistentVolumeClaims(),
			ctx.InformerFactory.Core().V1().PersistentVolumes(),
			ctx.Cloud,
			plugins,
			csiTranslator,
			csimigration.NewPluginManager(csiTranslator, utilfeature.DefaultFeatureGate),
			filteredDialOptions,
		)

		if expandControllerErr != nil {
			return nil, true, fmt.Errorf("failed to start volume expand controller: %v", expandControllerErr)
		}
		go expandController.Run(ctx.Stop)
		return nil, true, nil
	}
	return nil, false, nil
}

func startEphemeralVolumeController(ctx ControllerContext) (controller.Interface, bool, error) {
	if utilfeature.DefaultFeatureGate.Enabled(features.GenericEphemeralVolume) {
		ephemeralController, err := ephemeral.NewController(
			ctx.ClientBuilder.ClientOrDie("ephemeral-volume-controller"),
			ctx.InformerFactory.Core().V1().Pods(),
			ctx.InformerFactory.Core().V1().PersistentVolumeClaims())
		if err != nil {
			return nil, true, fmt.Errorf("failed to start ephemeral volume controller: %v", err)
		}
		go ephemeralController.Run(int(ctx.ComponentConfig.EphemeralVolumeController.ConcurrentEphemeralVolumeSyncs), ctx.Stop)
		return nil, true, nil
	}
	return nil, false, nil
}

func startEndpointController(ctx ControllerContext) (controller.Interface, bool, error) {
	go endpointcontroller.NewEndpointController(
		ctx.InformerFactory.Core().V1().Pods(),
		ctx.InformerFactory.Core().V1().Services(),
		ctx.InformerFactory.Core().V1().Endpoints(),
		ctx.ClientBuilder.ClientOrDie("endpoint-controller"),
		ctx.ComponentConfig.EndpointController.EndpointUpdatesBatchPeriod.Duration,
	).Run(int(ctx.ComponentConfig.EndpointController.ConcurrentEndpointSyncs), ctx.Stop)
	return nil, true, nil
}

func startReplicationController(ctx ControllerContext) (controller.Interface, bool, error) {
	go replicationcontroller.NewReplicationManager(
		ctx.InformerFactory.Core().V1().Pods(),
		ctx.InformerFactory.Core().V1().ReplicationControllers(),
		ctx.ClientBuilder.ClientOrDie("replication-controller"),
		replicationcontroller.BurstReplicas,
	).Run(int(ctx.ComponentConfig.ReplicationController.ConcurrentRCSyncs), ctx.Stop)
	return nil, true, nil
}

func startPodGCController(ctx ControllerContext) (controller.Interface, bool, error) {
	go podgc.NewPodGC(
		ctx.ClientBuilder.ClientOrDie("pod-garbage-collector"),
		ctx.InformerFactory.Core().V1().Pods(),
		ctx.InformerFactory.Core().V1().Nodes(),
		int(ctx.ComponentConfig.PodGCController.TerminatedPodGCThreshold),
	).Run(ctx.Stop)
	return nil, true, nil
}

func startResourceQuotaController(ctx ControllerContext) (controller.Interface, bool, error) {
	resourceQuotaControllerClient := ctx.ClientBuilder.ClientOrDie("resourcequota-controller")
	resourceQuotaControllerDiscoveryClient := ctx.ClientBuilder.DiscoveryClientOrDie("resourcequota-controller")
	discoveryFunc := resourceQuotaControllerDiscoveryClient.ServerPreferredNamespacedResources
	listerFuncForResource := generic.ListerFuncForResourceFunc(ctx.InformerFactory.ForResource)
	quotaConfiguration := quotainstall.NewQuotaConfigurationForControllers(listerFuncForResource)

	resourceQuotaControllerOptions := &resourcequotacontroller.ControllerOptions{
		QuotaClient:               resourceQuotaControllerClient.CoreV1(),
		ResourceQuotaInformer:     ctx.InformerFactory.Core().V1().ResourceQuotas(),
		ResyncPeriod:              pkgcontroller.StaticResyncPeriodFunc(ctx.ComponentConfig.ResourceQuotaController.ResourceQuotaSyncPeriod.Duration),
		InformerFactory:           ctx.ObjectOrMetadataInformerFactory,
		ReplenishmentResyncPeriod: ctx.ResyncPeriod,
		DiscoveryFunc:             discoveryFunc,
		IgnoredResourcesFunc:      quotaConfiguration.IgnoredResources,
		InformersStarted:          ctx.InformersStarted,
		Registry:                  generic.NewRegistry(quotaConfiguration.Evaluators()),
	}
	if resourceQuotaControllerClient.CoreV1().RESTClient().GetRateLimiter() != nil {
		if err := ratelimiter.RegisterMetricAndTrackRateLimiterUsage("resource_quota_controller", resourceQuotaControllerClient.CoreV1().RESTClient().GetRateLimiter()); err != nil {
			return nil, true, err
		}
	}

	resourceQuotaController, err := resourcequotacontroller.NewController(resourceQuotaControllerOptions)
	if err != nil {
		return nil, false, err
	}
	go resourceQuotaController.Run(int(ctx.ComponentConfig.ResourceQuotaController.ConcurrentResourceQuotaSyncs), ctx.Stop)

	// Periodically the quota controller to detect new resource types
	go resourceQuotaController.Sync(discoveryFunc, 30*time.Second, ctx.Stop)

	return nil, true, nil
}

func startNamespaceController(ctx ControllerContext) (controller.Interface, bool, error) {
	// the namespace cleanup controller is very chatty.  It makes lots of discovery calls and then it makes lots of delete calls
	// the ratelimiter negatively affects its speed.  Deleting 100 total items in a namespace (that's only a few of each resource
	// including events), takes ~10 seconds by default.
	nsKubeconfig := ctx.ClientBuilder.ConfigOrDie("namespace-controller")
	nsKubeconfig.QPS *= 20
	nsKubeconfig.Burst *= 100
	namespaceKubeClient := clientset.NewForConfigOrDie(nsKubeconfig)
	return startModifiedNamespaceController(ctx, namespaceKubeClient, nsKubeconfig)
}

func startModifiedNamespaceController(ctx ControllerContext, namespaceKubeClient clientset.Interface, nsKubeconfig *restclient.Config) (controller.Interface, bool, error) {

	metadataClient, err := metadata.NewForConfig(nsKubeconfig)
	if err != nil {
		return nil, true, err
	}

	discoverResourcesFn := namespaceKubeClient.Discovery().ServerPreferredNamespacedResources

	namespaceController := namespacecontroller.NewNamespaceController(
		namespaceKubeClient,
		metadataClient,
		discoverResourcesFn,
		ctx.InformerFactory.Core().V1().Namespaces(),
		ctx.ComponentConfig.NamespaceController.NamespaceSyncPeriod.Duration,
		v1.FinalizerKubernetes,
	)
	go namespaceController.Run(int(ctx.ComponentConfig.NamespaceController.ConcurrentNamespaceSyncs), ctx.Stop)

	return nil, true, nil
}

func startServiceAccountController(ctx ControllerContext) (controller.Interface, bool, error) {
	sac, err := serviceaccountcontroller.NewServiceAccountsController(
		ctx.InformerFactory.Core().V1().ServiceAccounts(),
		ctx.InformerFactory.Core().V1().Namespaces(),
		ctx.ClientBuilder.ClientOrDie("service-account-controller"),
		serviceaccountcontroller.DefaultServiceAccountsControllerOptions(),
	)
	if err != nil {
		return nil, true, fmt.Errorf("error creating ServiceAccount controller: %v", err)
	}
	go sac.Run(1, ctx.Stop)
	return nil, true, nil
}

func startTTLController(ctx ControllerContext) (controller.Interface, bool, error) {
	go ttlcontroller.NewTTLController(
		ctx.InformerFactory.Core().V1().Nodes(),
		ctx.ClientBuilder.ClientOrDie("ttl-controller"),
	).Run(5, ctx.Stop)
	return nil, true, nil
}

func startGarbageCollectorController(ctx ControllerContext) (controller.Interface, bool, error) {
	if !ctx.ComponentConfig.GarbageCollectorController.EnableGarbageCollector {
		return nil, false, nil
	}

	gcClientset := ctx.ClientBuilder.ClientOrDie("generic-garbage-collector")
	discoveryClient := ctx.ClientBuilder.DiscoveryClientOrDie("generic-garbage-collector")

	config := ctx.ClientBuilder.ConfigOrDie("generic-garbage-collector")
	// Increase garbage collector controller's throughput: each object deletion takes two API calls,
	// so to get |config.QPS| deletion rate we need to allow 2x more requests for this controller.
	config.QPS *= 2
	metadataClient, err := metadata.NewForConfig(config)
	if err != nil {
		return nil, true, err
	}

	ignoredResources := make(map[schema.GroupResource]struct{})
	for _, r := range ctx.ComponentConfig.GarbageCollectorController.GCIgnoredResources {
		ignoredResources[schema.GroupResource{Group: r.Group, Resource: r.Resource}] = struct{}{}
	}
	garbageCollector, err := garbagecollector.NewGarbageCollector(
		gcClientset,
		metadataClient,
		ctx.RESTMapper,
		ignoredResources,
		ctx.ObjectOrMetadataInformerFactory,
		ctx.InformersStarted,
	)
	if err != nil {
		return nil, true, fmt.Errorf("failed to start the generic garbage collector: %v", err)
	}

	// Start the garbage collector.
	workers := int(ctx.ComponentConfig.GarbageCollectorController.ConcurrentGCSyncs)
	go garbageCollector.Run(workers, ctx.Stop)

	// Periodically refresh the RESTMapper with new discovery information and sync
	// the garbage collector.
	go garbageCollector.Sync(discoveryClient, 30*time.Second, ctx.Stop)

	return garbageCollector, true, nil
}

func startPVCProtectionController(ctx ControllerContext) (controller.Interface, bool, error) {
	pvcProtectionController, err := pvcprotection.NewPVCProtectionController(
		ctx.InformerFactory.Core().V1().PersistentVolumeClaims(),
		ctx.InformerFactory.Core().V1().Pods(),
		ctx.ClientBuilder.ClientOrDie("pvc-protection-controller"),
		utilfeature.DefaultFeatureGate.Enabled(features.StorageObjectInUseProtection),
		utilfeature.DefaultFeatureGate.Enabled(features.GenericEphemeralVolume),
	)
	if err != nil {
		return nil, true, fmt.Errorf("failed to start the pvc protection controller: %v", err)
	}
	go pvcProtectionController.Run(1, ctx.Stop)
	return nil, true, nil
}

func startPVProtectionController(ctx ControllerContext) (controller.Interface, bool, error) {
	go pvprotection.NewPVProtectionController(
		ctx.InformerFactory.Core().V1().PersistentVolumes(),
		ctx.ClientBuilder.ClientOrDie("pv-protection-controller"),
		utilfeature.DefaultFeatureGate.Enabled(features.StorageObjectInUseProtection),
	).Run(1, ctx.Stop)
	return nil, true, nil
}

func startTTLAfterFinishedController(ctx ControllerContext) (controller.Interface, bool, error) {
	if !utilfeature.DefaultFeatureGate.Enabled(features.TTLAfterFinished) {
		return nil, false, nil
	}
	go ttlafterfinished.New(
		ctx.InformerFactory.Batch().V1().Jobs(),
		ctx.ClientBuilder.ClientOrDie("ttl-after-finished-controller"),
	).Run(int(ctx.ComponentConfig.TTLAfterFinishedController.ConcurrentTTLSyncs), ctx.Stop)
	return nil, true, nil
}

// processCIDRs is a helper function that works on a comma separated cidrs and returns
// a list of typed cidrs
// a flag if cidrs represents a dual stack
// error if failed to parse any of the cidrs
func processCIDRs(cidrsList string) ([]*net.IPNet, bool, error) {
	cidrsSplit := strings.Split(strings.TrimSpace(cidrsList), ",")

	cidrs, err := netutils.ParseCIDRs(cidrsSplit)
	if err != nil {
		return nil, false, err
	}

	// if cidrs has an error then the previous call will fail
	// safe to ignore error checking on next call
	dualstack, _ := netutils.IsDualStackCIDRs(cidrs)

	return cidrs, dualstack, nil
}

// setNodeCIDRMaskSizes returns the IPv4 and IPv6 node cidr mask sizes to the value provided
// for --node-cidr-mask-size-ipv4 and --node-cidr-mask-size-ipv6 respectively. If value not provided,
// then it will return default IPv4 and IPv6 cidr mask sizes.
func setNodeCIDRMaskSizes(cfg nodeipamconfig.NodeIPAMControllerConfiguration, clusterCIDRs []*net.IPNet) ([]int, error) {

	sortedSizes := func(maskSizeIPv4, maskSizeIPv6 int) []int {
		nodeMaskCIDRs := make([]int, len(clusterCIDRs))

		for idx, clusterCIDR := range clusterCIDRs {
			if netutils.IsIPv6CIDR(clusterCIDR) {
				nodeMaskCIDRs[idx] = maskSizeIPv6
			} else {
				nodeMaskCIDRs[idx] = maskSizeIPv4
			}
		}
		return nodeMaskCIDRs
	}

	// --node-cidr-mask-size flag is incompatible with dual stack clusters.
	ipv4Mask, ipv6Mask := defaultNodeMaskCIDRIPv4, defaultNodeMaskCIDRIPv6
	isDualstack := len(clusterCIDRs) > 1

	// case one: cluster is dualstack (i.e, more than one cidr)
	if isDualstack {
		// if --node-cidr-mask-size then fail, user must configure the correct dual-stack mask sizes (or use default)
		if cfg.NodeCIDRMaskSize != 0 {
			return nil, errors.New("usage of --node-cidr-mask-size is not allowed with dual-stack clusters")

		}

		if cfg.NodeCIDRMaskSizeIPv4 != 0 {
			ipv4Mask = int(cfg.NodeCIDRMaskSizeIPv4)
		}
		if cfg.NodeCIDRMaskSizeIPv6 != 0 {
			ipv6Mask = int(cfg.NodeCIDRMaskSizeIPv6)
		}
		return sortedSizes(ipv4Mask, ipv6Mask), nil
	}

	maskConfigured := cfg.NodeCIDRMaskSize != 0
	maskV4Configured := cfg.NodeCIDRMaskSizeIPv4 != 0
	maskV6Configured := cfg.NodeCIDRMaskSizeIPv6 != 0
	isSingleStackIPv6 := netutils.IsIPv6CIDR(clusterCIDRs[0])

	// original flag is set
	if maskConfigured {
		// original mask flag is still the main reference.
		if maskV4Configured || maskV6Configured {
			return nil, errors.New("usage of --node-cidr-mask-size-ipv4 and --node-cidr-mask-size-ipv6 is not allowed if --node-cidr-mask-size is set. For dual-stack clusters please unset it and use IPFamily specific flags")
		}

		mask := int(cfg.NodeCIDRMaskSize)
		return sortedSizes(mask, mask), nil
	}

	if maskV4Configured {
		if isSingleStackIPv6 {
			return nil, errors.New("usage of --node-cidr-mask-size-ipv4 is not allowed for a single-stack IPv6 cluster")
		}

		ipv4Mask = int(cfg.NodeCIDRMaskSizeIPv4)
	}

	// !maskV4Configured && !maskConfigured && maskV6Configured
	if maskV6Configured {
		if !isSingleStackIPv6 {
			return nil, errors.New("usage of --node-cidr-mask-size-ipv6 is not allowed for a single-stack IPv4 cluster")
		}

		ipv6Mask = int(cfg.NodeCIDRMaskSizeIPv6)
	}
	return sortedSizes(ipv4Mask, ipv6Mask), nil
}
func startStorageVersionGCController(ctx ControllerContext) (controller.Interface, bool, error) {
	go storageversiongc.NewStorageVersionGC(
		ctx.ClientBuilder.ClientOrDie("storage-version-garbage-collector"),
		ctx.InformerFactory.Coordination().V1().Leases(),
		ctx.InformerFactory.Internal().V1alpha1().StorageVersions(),
	).Run(ctx.Stop)
	return nil, true, nil
}
