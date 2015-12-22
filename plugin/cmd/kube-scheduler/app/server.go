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

// Package app implements a Server object for running the scheduler.
package app

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strconv"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/record"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	"k8s.io/kubernetes/pkg/healthz"
	"k8s.io/kubernetes/pkg/master/ports"
	"k8s.io/kubernetes/pkg/util"
	"k8s.io/kubernetes/plugin/pkg/scheduler"
	_ "k8s.io/kubernetes/plugin/pkg/scheduler/algorithmprovider"
	schedulerapi "k8s.io/kubernetes/plugin/pkg/scheduler/api"
	latestschedulerapi "k8s.io/kubernetes/plugin/pkg/scheduler/api/latest"
	"k8s.io/kubernetes/plugin/pkg/scheduler/factory"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// SchedulerServer has all the context and params needed to run a Scheduler
type SchedulerServer struct {
	Port              int
	Address           net.IP
	AlgorithmProvider string
	PolicyConfigFile  string
	EnableProfiling   bool
	Master            string
	Kubeconfig        string
	BindPodsQPS       float32
	BindPodsBurst     int
	KubeAPIQPS        float32
	KubeAPIBurst      int
	SchedulerName     string
}

// NewSchedulerServer creates a new SchedulerServer with default parameters
func NewSchedulerServer() *SchedulerServer {
	s := SchedulerServer{
		Port:              ports.SchedulerPort,
		Address:           net.ParseIP("0.0.0.0"),
		AlgorithmProvider: factory.DefaultProvider,
		BindPodsQPS:       50.0,
		BindPodsBurst:     100,
		KubeAPIQPS:        50.0,
		KubeAPIBurst:      100,
		SchedulerName:     api.DefaultSchedulerName,
	}
	return &s
}

// NewSchedulerCommand creates a *cobra.Command object with default parameters
func NewSchedulerCommand() *cobra.Command {
	s := NewSchedulerServer()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use: "kube-scheduler",
		Long: `The Kubernetes scheduler is a policy-rich, topology-aware,
workload-specific function that significantly impacts availability, performance,
and capacity. The scheduler needs to take into account individual and collective
resource requirements, quality of service requirements, hardware/software/policy
constraints, affinity and anti-affinity specifications, data locality, inter-workload
interference, deadlines, and so on. Workload-specific requirements will be exposed
through the API as necessary.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}

// AddFlags adds flags for a specific SchedulerServer to the specified FlagSet
func (s *SchedulerServer) AddFlags(fs *pflag.FlagSet) {
	fs.IntVar(&s.Port, "port", s.Port, "The port that the scheduler's http service runs on")
	fs.IPVar(&s.Address, "address", s.Address, "The IP address to serve on (set to 0.0.0.0 for all interfaces)")
	fs.StringVar(&s.AlgorithmProvider, "algorithm-provider", s.AlgorithmProvider, "The scheduling algorithm provider to use, one of: "+factory.ListAlgorithmProviders())
	fs.StringVar(&s.PolicyConfigFile, "policy-config-file", s.PolicyConfigFile, "File with scheduler policy configuration")
	fs.BoolVar(&s.EnableProfiling, "profiling", true, "Enable profiling via web interface host:port/debug/pprof/")
	fs.StringVar(&s.Master, "master", s.Master, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.Kubeconfig, "kubeconfig", s.Kubeconfig, "Path to kubeconfig file with authorization and master location information.")
	fs.Float32Var(&s.BindPodsQPS, "bind-pods-qps", s.BindPodsQPS, "Number of bindings per second scheduler is allowed to continuously make")
	fs.IntVar(&s.BindPodsBurst, "bind-pods-burst", s.BindPodsBurst, "Number of bindings per second scheduler is allowed to make during bursts")
	fs.Float32Var(&s.KubeAPIQPS, "kube-api-qps", s.KubeAPIQPS, "QPS to use while talking with kubernetes apiserver")
	fs.IntVar(&s.KubeAPIBurst, "kube-api-burst", s.KubeAPIBurst, "Burst to use while talking with kubernetes apiserver")
	fs.StringVar(&s.SchedulerName, "scheduler-name", s.SchedulerName, "Name of the scheduler, used to select which pods will be processed by this scheduler, based on pod's annotation with key 'scheduler.alpha.kubernetes.io/name'")
}

// Run runs the specified SchedulerServer.  This should never exit.
func (s *SchedulerServer) Run(_ []string) error {
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

	configFactory := factory.NewConfigFactory(kubeClient, util.NewTokenBucketRateLimiter(s.BindPodsQPS, s.BindPodsBurst), s.SchedulerName)
	config, err := s.createConfig(configFactory)
	if err != nil {
		glog.Fatalf("Failed to create scheduler configuration: %v", err)
	}

	eventBroadcaster := record.NewBroadcaster()
	config.Recorder = eventBroadcaster.NewRecorder(api.EventSource{Component: s.SchedulerName})
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(kubeClient.Events(""))

	sched := scheduler.New(config)
	sched.Run()

	select {}
}

func (s *SchedulerServer) createConfig(configFactory *factory.ConfigFactory) (*scheduler.Config, error) {
	var policy schedulerapi.Policy
	var configData []byte

	if _, err := os.Stat(s.PolicyConfigFile); err == nil {
		configData, err = ioutil.ReadFile(s.PolicyConfigFile)
		if err != nil {
			return nil, fmt.Errorf("Unable to read policy config: %v", err)
		}
		err = latestschedulerapi.Codec.DecodeInto(configData, &policy)
		if err != nil {
			return nil, fmt.Errorf("Invalid configuration: %v", err)
		}

		return configFactory.CreateFromConfig(policy)
	}

	// if the config file isn't provided, use the specified (or default) provider
	// check of algorithm provider is registered and fail fast
	_, err := factory.GetAlgorithmProvider(s.AlgorithmProvider)
	if err != nil {
		return nil, err
	}

	return configFactory.CreateFromProvider(s.AlgorithmProvider)
}
