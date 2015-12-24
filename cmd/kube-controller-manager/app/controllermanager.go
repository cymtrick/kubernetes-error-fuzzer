/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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
// CAUTION: If you update code in this file, you may need to also update code
//          in contrib/mesos/pkg/controllermanager/controllermanager.go
package app

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/http/pprof"
	"strconv"
	"time"

	"k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
	"k8s.io/kubernetes/pkg/api/unversioned"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/daemon"
	"k8s.io/kubernetes/pkg/controller/deployment"
	endpointcontroller "k8s.io/kubernetes/pkg/controller/endpoint"
	"k8s.io/kubernetes/pkg/controller/gc"
	"k8s.io/kubernetes/pkg/controller/job"
	namespacecontroller "k8s.io/kubernetes/pkg/controller/namespace"
	nodecontroller "k8s.io/kubernetes/pkg/controller/node"
	persistentvolumecontroller "k8s.io/kubernetes/pkg/controller/persistentvolume"
	"k8s.io/kubernetes/pkg/controller/podautoscaler"
	"k8s.io/kubernetes/pkg/controller/podautoscaler/metrics"
	replicationcontroller "k8s.io/kubernetes/pkg/controller/replication"
	resourcequotacontroller "k8s.io/kubernetes/pkg/controller/resourcequota"
	routecontroller "k8s.io/kubernetes/pkg/controller/route"
	servicecontroller "k8s.io/kubernetes/pkg/controller/service"
	serviceaccountcontroller "k8s.io/kubernetes/pkg/controller/serviceaccount"
	"k8s.io/kubernetes/pkg/healthz"
	"k8s.io/kubernetes/pkg/serviceaccount"
	"k8s.io/kubernetes/pkg/util"
	"k8s.io/kubernetes/pkg/util/wait"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// NewControllerManagerCommand creates a *cobra.Command object with default parameters
func NewControllerManagerCommand() *cobra.Command {
	s := options.NewCMServer()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use: "kube-controller-manager",
		Long: `The Kubernetes controller manager is a daemon that embeds
the core control loops shipped with Kubernetes. In applications of robotics and
automation, a control loop is a non-terminating loop that regulates the state of
the system. In Kubernetes, a controller is a control loop that watches the shared
state of the cluster through the apiserver and makes changes attempting to move the
current state towards the desired state. Examples of controllers that ship with
Kubernetes today are the replication controller, endpoints controller, namespace
controller, and serviceaccounts controller.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}

func ResyncPeriod(s *options.CMServer) func() time.Duration {
	return func() time.Duration {
		factor := rand.Float64() + 1
		return time.Duration(float64(s.MinResyncPeriod.Nanoseconds()) * factor)
	}
}

func clientForUserAgentOrDie(config client.Config, userAgent string) *client.Client {
	fullUserAgent := client.DefaultKubernetesUserAgent() + "/" + userAgent
	config.UserAgent = fullUserAgent
	kubeClient, err := client.New(&config)
	if err != nil {
		glog.Fatalf("Invalid API configuration: %v", err)
	}
	return kubeClient
}

// Run runs the CMServer.  This should never exit.
func Run(s *options.CMServer) error {
	kubeconfig, err := clientcmd.BuildConfigFromFlags(s.Master, s.Kubeconfig)
	if err != nil {
		return err
	}

	// Override kubeconfig qps/burst settings from flags
	kubeconfig.QPS = s.KubeAPIQPS
	kubeconfig.Burst = s.KubeAPIBurst

	kubeClient, err := client.New(kubeconfig)
	if err != nil {
		glog.Fatalf("Invalid API configuration: %v", err)
	}

	go func() {
		mux := http.NewServeMux()
		healthz.InstallHandler(mux)
		if s.EnableProfiling {
			mux.HandleFunc("/debug/pprof/", pprof.Index)
			mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		}
		mux.Handle("/metrics", prometheus.Handler())

		server := &http.Server{
			Addr:    net.JoinHostPort(s.Address.String(), strconv.Itoa(s.Port)),
			Handler: mux,
		}
		glog.Fatal(server.ListenAndServe())
	}()

	go endpointcontroller.NewEndpointController(clientForUserAgentOrDie(*kubeconfig, "endpoint-controller"), ResyncPeriod(s)).
		Run(s.ConcurrentEndpointSyncs, util.NeverStop)

	go replicationcontroller.NewReplicationManager(
		clientForUserAgentOrDie(*kubeconfig, "replication-controller"),
		ResyncPeriod(s),
		replicationcontroller.BurstReplicas,
	).Run(s.ConcurrentRCSyncs, util.NeverStop)

	if s.TerminatedPodGCThreshold > 0 {
		go gc.New(clientForUserAgentOrDie(*kubeconfig, "garbage-collector"), ResyncPeriod(s), s.TerminatedPodGCThreshold).
			Run(util.NeverStop)
	}

	cloud, err := cloudprovider.InitCloudProvider(s.CloudProvider, s.CloudConfigFile)
	if err != nil {
		glog.Fatalf("Cloud provider could not be initialized: %v", err)
	}

	nodeController := nodecontroller.NewNodeController(cloud, clientForUserAgentOrDie(*kubeconfig, "node-controller"),
		s.PodEvictionTimeout, util.NewTokenBucketRateLimiter(s.DeletingPodsQps, s.DeletingPodsBurst),
		util.NewTokenBucketRateLimiter(s.DeletingPodsQps, s.DeletingPodsBurst),
		s.NodeMonitorGracePeriod, s.NodeStartupGracePeriod, s.NodeMonitorPeriod, &s.ClusterCIDR, s.AllocateNodeCIDRs)
	nodeController.Run(s.NodeSyncPeriod)

	serviceController := servicecontroller.New(cloud, clientForUserAgentOrDie(*kubeconfig, "service-controller"), s.ClusterName)
	if err := serviceController.Run(s.ServiceSyncPeriod, s.NodeSyncPeriod); err != nil {
		glog.Errorf("Failed to start service controller: %v", err)
	}

	if s.AllocateNodeCIDRs {
		if cloud == nil {
			glog.Warning("allocate-node-cidrs is set, but no cloud provider specified. Will not manage routes.")
		} else if routes, ok := cloud.Routes(); !ok {
			glog.Warning("allocate-node-cidrs is set, but cloud provider does not support routes. Will not manage routes.")
		} else {
			routeController := routecontroller.New(routes, clientForUserAgentOrDie(*kubeconfig, "route-controller"), s.ClusterName, &s.ClusterCIDR)
			routeController.Run(s.NodeSyncPeriod)
		}
	} else {
		glog.Infof("allocate-node-cidrs set to %v, node controller not creating routes", s.AllocateNodeCIDRs)
	}

	go resourcequotacontroller.NewResourceQuotaController(
		clientForUserAgentOrDie(*kubeconfig, "resourcequota-controller"),
		controller.StaticResyncPeriodFunc(s.ResourceQuotaSyncPeriod)).Run(s.ConcurrentResourceQuotaSyncs, util.NeverStop)

	// If apiserver is not running we should wait for some time and fail only then. This is particularly
	// important when we start apiserver and controller manager at the same time.
	var versionStrings []string
	err = wait.PollImmediate(time.Second, 10*time.Second, func() (bool, error) {
		if versionStrings, err = client.ServerAPIVersions(kubeconfig); err == nil {
			return true, nil
		}
		glog.Errorf("Failed to get api versions from server: %v", err)
		return false, nil
	})
	if err != nil {
		glog.Fatalf("Failed to get api versions from server: %v", err)
	}
	versions := &unversioned.APIVersions{Versions: versionStrings}

	resourceMap, err := kubeClient.Discovery().ServerResources()
	if err != nil {
		glog.Fatalf("Failed to get supported resources from server: %v", err)
	}

	namespacecontroller.NewNamespaceController(clientForUserAgentOrDie(*kubeconfig, "namespace-controller"), versions, s.NamespaceSyncPeriod).Run()

	groupVersion := "extensions/v1beta1"
	resources, found := resourceMap[groupVersion]
	// TODO: this needs to be dynamic so users don't have to restart their controller manager if they change the apiserver
	if containsVersion(versions, groupVersion) && found {
		glog.Infof("Starting %s apis", groupVersion)
		if containsResource(resources, "horizontalpodautoscalers") {
			glog.Infof("Starting horizontal pod controller.")
			hpaClient := clientForUserAgentOrDie(*kubeconfig, "horizontal-pod-autoscaler")
			metricsClient := metrics.NewHeapsterMetricsClient(
				hpaClient,
				metrics.DefaultHeapsterNamespace,
				metrics.DefaultHeapsterScheme,
				metrics.DefaultHeapsterService,
				metrics.DefaultHeapsterPort,
			)
			podautoscaler.NewHorizontalController(hpaClient, hpaClient, hpaClient, metricsClient).
				Run(s.HorizontalPodAutoscalerSyncPeriod)
		}

		if containsResource(resources, "daemonsets") {
			glog.Infof("Starting daemon set controller")
			go daemon.NewDaemonSetsController(clientForUserAgentOrDie(*kubeconfig, "daemon-set-controller"), ResyncPeriod(s)).
				Run(s.ConcurrentDSCSyncs, util.NeverStop)
		}

		if containsResource(resources, "jobs") {
			glog.Infof("Starting job controller")
			go job.NewJobController(clientForUserAgentOrDie(*kubeconfig, "job-controller"), ResyncPeriod(s)).
				Run(s.ConcurrentJobSyncs, util.NeverStop)
		}

		if containsResource(resources, "deployments") {
			glog.Infof("Starting deployment controller")
			go deployment.NewDeploymentController(clientForUserAgentOrDie(*kubeconfig, "deployment-controller"), ResyncPeriod(s)).
				Run(s.ConcurrentDeploymentSyncs, util.NeverStop)
		}
	}

	volumePlugins := ProbeRecyclableVolumePlugins(s.VolumeConfigFlags)
	provisioner, err := NewVolumeProvisioner(cloud, s.VolumeConfigFlags)
	if err != nil {
		glog.Fatal("A Provisioner could not be created, but one was expected. Provisioning will not work. This functionality is considered an early Alpha version.")
	}

	pvclaimBinder := persistentvolumecontroller.NewPersistentVolumeClaimBinder(clientForUserAgentOrDie(*kubeconfig, "persistent-volume-binder"), s.PVClaimBinderSyncPeriod)
	pvclaimBinder.Run()

	pvRecycler, err := persistentvolumecontroller.NewPersistentVolumeRecycler(clientForUserAgentOrDie(*kubeconfig, "persistent-volume-recycler"), s.PVClaimBinderSyncPeriod, ProbeRecyclableVolumePlugins(s.VolumeConfigFlags), cloud)
	if err != nil {
		glog.Fatalf("Failed to start persistent volume recycler: %+v", err)
	}
	pvRecycler.Run()

	if provisioner != nil {
		pvController, err := persistentvolumecontroller.NewPersistentVolumeProvisionerController(persistentvolumecontroller.NewControllerClient(clientForUserAgentOrDie(*kubeconfig, "persistent-volume-provisioner")), s.PVClaimBinderSyncPeriod, volumePlugins, provisioner, cloud)
		if err != nil {
			glog.Fatalf("Failed to start persistent volume provisioner controller: %+v", err)
		}
		pvController.Run()
	}

	var rootCA []byte

	if s.RootCAFile != "" {
		rootCA, err = ioutil.ReadFile(s.RootCAFile)
		if err != nil {
			return fmt.Errorf("error reading root-ca-file at %s: %v", s.RootCAFile, err)
		}
		if _, err := util.CertsFromPEM(rootCA); err != nil {
			return fmt.Errorf("error parsing root-ca-file at %s: %v", s.RootCAFile, err)
		}
	} else {
		rootCA = kubeconfig.CAData
	}

	if len(s.ServiceAccountKeyFile) > 0 {
		privateKey, err := serviceaccount.ReadPrivateKey(s.ServiceAccountKeyFile)
		if err != nil {
			glog.Errorf("Error reading key for service account token controller: %v", err)
		} else {
			serviceaccountcontroller.NewTokensController(
				clientForUserAgentOrDie(*kubeconfig, "tokens-controller"),
				serviceaccountcontroller.TokensControllerOptions{
					TokenGenerator: serviceaccount.JWTTokenGenerator(privateKey),
					RootCA:         rootCA,
				},
			).Run()
		}
	}

	serviceaccountcontroller.NewServiceAccountsController(
		clientForUserAgentOrDie(*kubeconfig, "service-account-controller"),
		serviceaccountcontroller.DefaultServiceAccountsControllerOptions(),
	).Run()

	select {}
}

func containsVersion(versions *unversioned.APIVersions, version string) bool {
	for ix := range versions.Versions {
		if versions.Versions[ix] == version {
			return true
		}
	}
	return false
}

func containsResource(resources *unversioned.APIResourceList, resourceName string) bool {
	for ix := range resources.APIResources {
		resource := resources.APIResources[ix]
		if resource.Name == resourceName {
			return true
		}
	}
	return false
}
