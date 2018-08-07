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

package options

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	apiserveroptions "k8s.io/apiserver/pkg/server/options"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	apiserverflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
	cloudcontrollerconfig "k8s.io/kubernetes/cmd/cloud-controller-manager/app/config"
	cmoptions "k8s.io/kubernetes/cmd/controller-manager/app/options"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/componentconfig"
	componentconfigv1alpha1 "k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/master/ports"
	// add the kubernetes feature gates
	_ "k8s.io/kubernetes/pkg/features"

	"github.com/golang/glog"
)

const (
	// CloudControllerManagerUserAgent is the userAgent name when starting cloud-controller managers.
	CloudControllerManagerUserAgent = "cloud-controller-manager"
)

// CloudControllerManagerOptions is the main context object for the controller manager.
type CloudControllerManagerOptions struct {
	CloudProvider     *cmoptions.CloudProviderOptions
	Debugging         *cmoptions.DebuggingOptions
	GenericComponent  *cmoptions.GenericComponentConfigOptions
	KubeCloudShared   *cmoptions.KubeCloudSharedOptions
	ServiceController *cmoptions.ServiceControllerOptions

	SecureServing *apiserveroptions.SecureServingOptions
	// TODO: remove insecure serving mode
	InsecureServing *apiserveroptions.DeprecatedInsecureServingOptions
	Authentication  *apiserveroptions.DelegatingAuthenticationOptions
	Authorization   *apiserveroptions.DelegatingAuthorizationOptions

	Master     string
	Kubeconfig string

	// NodeStatusUpdateFrequency is the frequency at which the controller updates nodes' status
	NodeStatusUpdateFrequency metav1.Duration
}

// NewCloudControllerManagerOptions creates a new ExternalCMServer with a default config.
func NewCloudControllerManagerOptions() (*CloudControllerManagerOptions, error) {
	componentConfig, err := NewDefaultComponentConfig(ports.InsecureCloudControllerManagerPort)
	if err != nil {
		return nil, err
	}

	s := CloudControllerManagerOptions{
		CloudProvider:    &cmoptions.CloudProviderOptions{},
		Debugging:        &cmoptions.DebuggingOptions{},
		GenericComponent: cmoptions.NewGenericComponentConfigOptions(componentConfig.GenericComponent),
		KubeCloudShared:  cmoptions.NewKubeCloudSharedOptions(componentConfig.KubeCloudShared),
		ServiceController: &cmoptions.ServiceControllerOptions{
			ConcurrentServiceSyncs: componentConfig.ServiceController.ConcurrentServiceSyncs,
		},
		SecureServing: apiserveroptions.NewSecureServingOptions(),
		InsecureServing: &apiserveroptions.DeprecatedInsecureServingOptions{
			BindAddress: net.ParseIP(componentConfig.KubeCloudShared.Address),
			BindPort:    int(componentConfig.KubeCloudShared.Port),
			BindNetwork: "tcp",
		},
		Authentication:            apiserveroptions.NewDelegatingAuthenticationOptions(),
		Authorization:             apiserveroptions.NewDelegatingAuthorizationOptions(),
		NodeStatusUpdateFrequency: componentConfig.NodeStatusUpdateFrequency,
	}

	s.Authentication.RemoteKubeConfigFileOptional = true
	s.Authorization.RemoteKubeConfigFileOptional = true
	s.Authorization.AlwaysAllowPaths = []string{"/healthz"}

	s.SecureServing.ServerCert.CertDirectory = "/var/run/kubernetes"
	s.SecureServing.ServerCert.PairName = "cloud-controller-manager"
	s.SecureServing.BindPort = ports.CloudControllerManagerPort

	return &s, nil
}

// NewDefaultComponentConfig returns cloud-controller manager configuration object.
func NewDefaultComponentConfig(insecurePort int32) (componentconfig.CloudControllerManagerConfiguration, error) {
	scheme := runtime.NewScheme()
	if err := componentconfigv1alpha1.AddToScheme(scheme); err != nil {
		return componentconfig.CloudControllerManagerConfiguration{}, err
	}
	if err := componentconfig.AddToScheme(scheme); err != nil {
		return componentconfig.CloudControllerManagerConfiguration{}, err
	}

	versioned := componentconfigv1alpha1.CloudControllerManagerConfiguration{}
	scheme.Default(&versioned)

	internal := componentconfig.CloudControllerManagerConfiguration{}
	if err := scheme.Convert(&versioned, &internal, nil); err != nil {
		return internal, err
	}
	internal.KubeCloudShared.Port = insecurePort
	return internal, nil
}

// Flags returns flags for a specific APIServer by section name
func (o *CloudControllerManagerOptions) Flags() (fss apiserverflag.NamedFlagSets) {
	o.CloudProvider.AddFlags(fss.FlagSet("cloud provider"))
	o.Debugging.AddFlags(fss.FlagSet("debugging"))
	o.GenericComponent.AddFlags(fss.FlagSet("generic"))
	o.KubeCloudShared.AddFlags(fss.FlagSet("generic"))
	o.ServiceController.AddFlags(fss.FlagSet("service controller"))

	o.SecureServing.AddFlags(fss.FlagSet("secure serving"))
	o.InsecureServing.AddUnqualifiedFlags(fss.FlagSet("insecure serving"))
	o.Authentication.AddFlags(fss.FlagSet("authentication"))
	o.Authorization.AddFlags(fss.FlagSet("authorization"))

	fs := fss.FlagSet("misc")
	fs.StringVar(&o.Master, "master", o.Master, "The address of the Kubernetes API server (overrides any value in kubeconfig).")
	fs.StringVar(&o.Kubeconfig, "kubeconfig", o.Kubeconfig, "Path to kubeconfig file with authorization and master location information.")
	fs.DurationVar(&o.NodeStatusUpdateFrequency.Duration, "node-status-update-frequency", o.NodeStatusUpdateFrequency.Duration, "Specifies how often the controller updates nodes' status.")

	utilfeature.DefaultFeatureGate.AddFlag(fss.FlagSet("generic"))

	return fss
}

// ApplyTo fills up cloud controller manager config with options.
func (o *CloudControllerManagerOptions) ApplyTo(c *cloudcontrollerconfig.Config, userAgent string) error {
	var err error
	if err = o.CloudProvider.ApplyTo(&c.ComponentConfig.CloudProvider); err != nil {
		return err
	}
	if err = o.Debugging.ApplyTo(&c.ComponentConfig.Debugging); err != nil {
		return err
	}
	if err = o.GenericComponent.ApplyTo(&c.ComponentConfig.GenericComponent); err != nil {
		return err
	}
	if err = o.KubeCloudShared.ApplyTo(&c.ComponentConfig.KubeCloudShared); err != nil {
		return err
	}
	if err = o.ServiceController.ApplyTo(&c.ComponentConfig.ServiceController); err != nil {
		return err
	}
	if err = o.SecureServing.ApplyTo(&c.SecureServing); err != nil {
		return err
	}
	if err = o.InsecureServing.ApplyTo(&c.InsecureServing); err != nil {
		return err
	}
	if o.SecureServing.BindPort != 0 || o.SecureServing.Listener != nil {
		if err = o.Authentication.ApplyTo(&c.Authentication, c.SecureServing, nil); err != nil {
			return err
		}
		if err = o.Authorization.ApplyTo(&c.Authorization); err != nil {
			return err
		}
	}

	c.Kubeconfig, err = clientcmd.BuildConfigFromFlags(o.Master, o.Kubeconfig)
	if err != nil {
		return err
	}
	c.Kubeconfig.ContentConfig.ContentType = o.GenericComponent.ContentType
	c.Kubeconfig.QPS = o.GenericComponent.KubeAPIQPS
	c.Kubeconfig.Burst = int(o.GenericComponent.KubeAPIBurst)

	c.Client, err = clientset.NewForConfig(restclient.AddUserAgent(c.Kubeconfig, userAgent))
	if err != nil {
		return err
	}

	c.LeaderElectionClient = clientset.NewForConfigOrDie(restclient.AddUserAgent(c.Kubeconfig, "leader-election"))

	c.EventRecorder = createRecorder(c.Client, userAgent)

	rootClientBuilder := controller.SimpleControllerClientBuilder{
		ClientConfig: c.Kubeconfig,
	}
	if c.ComponentConfig.KubeCloudShared.UseServiceAccountCredentials {
		c.ClientBuilder = controller.SAControllerClientBuilder{
			ClientConfig:         restclient.AnonymousClientConfig(c.Kubeconfig),
			CoreClient:           c.Client.CoreV1(),
			AuthenticationClient: c.Client.AuthenticationV1(),
			Namespace:            "kube-system",
		}
	} else {
		c.ClientBuilder = rootClientBuilder
	}
	c.VersionedClient = rootClientBuilder.ClientOrDie("shared-informers")
	c.SharedInformers = informers.NewSharedInformerFactory(c.VersionedClient, resyncPeriod(c)())

	// sync back to component config
	// TODO: find more elegant way than syncing back the values.
	c.ComponentConfig.KubeCloudShared.Port = int32(o.InsecureServing.BindPort)
	c.ComponentConfig.KubeCloudShared.Address = o.InsecureServing.BindAddress.String()

	c.ComponentConfig.NodeStatusUpdateFrequency = o.NodeStatusUpdateFrequency

	return nil
}

// Validate is used to validate config before launching the cloud controller manager
func (o *CloudControllerManagerOptions) Validate() error {
	errors := []error{}

	errors = append(errors, o.CloudProvider.Validate()...)
	errors = append(errors, o.Debugging.Validate()...)
	errors = append(errors, o.GenericComponent.Validate()...)
	errors = append(errors, o.KubeCloudShared.Validate()...)
	errors = append(errors, o.ServiceController.Validate()...)
	errors = append(errors, o.SecureServing.Validate()...)
	errors = append(errors, o.InsecureServing.Validate()...)
	errors = append(errors, o.Authentication.Validate()...)
	errors = append(errors, o.Authorization.Validate()...)

	if len(o.CloudProvider.Name) == 0 {
		errors = append(errors, fmt.Errorf("--cloud-provider cannot be empty"))
	}

	return utilerrors.NewAggregate(errors)
}

// resyncPeriod computes the time interval a shared informer waits before resyncing with the api server
func resyncPeriod(c *cloudcontrollerconfig.Config) func() time.Duration {
	return func() time.Duration {
		factor := rand.Float64() + 1
		return time.Duration(float64(c.ComponentConfig.GenericComponent.MinResyncPeriod.Nanoseconds()) * factor)
	}
}

// Config return a cloud controller manager config objective
func (o *CloudControllerManagerOptions) Config() (*cloudcontrollerconfig.Config, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	if err := o.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	c := &cloudcontrollerconfig.Config{}
	if err := o.ApplyTo(c, CloudControllerManagerUserAgent); err != nil {
		return nil, err
	}

	return c, nil
}

func createRecorder(kubeClient clientset.Interface, userAgent string) record.EventRecorder {
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	// TODO: remove dependence on the legacyscheme
	return eventBroadcaster.NewRecorder(legacyscheme.Scheme, v1.EventSource{Component: userAgent})
}
