/*
Copyright 2018 The Kubernetes Authors.

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

package azure

import (
	"errors"
	"fmt"
	"os"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/azure"
	"k8s.io/kubernetes/test/e2e/framework"
)

func init() {
	framework.RegisterProvider("azure", NewProvider)
}

func NewProvider() (framework.ProviderInterface, error) {
	if framework.TestContext.CloudConfig.ConfigFile == "" {
		return nil, fmt.Errorf("config-file must be specified for Azure")
	}
	config, err := os.Open(framework.TestContext.CloudConfig.ConfigFile)
	if err != nil {
		framework.Logf("Couldn't open cloud provider configuration %s: %#v",
			framework.TestContext.CloudConfig.ConfigFile, err)
	}
	defer config.Close()
	azureCloud, err := azure.NewCloud(config)
	return &Provider{
		azureCloud: azureCloud.(*azure.Cloud),
	}, err
}

type Provider struct {
	framework.NullProvider

	azureCloud *azure.Cloud
}

func (p *Provider) DeleteNode(node *v1.Node) error {
	return errors.New("not implemented yet")
}

func (p *Provider) CreatePD(zone string) (string, error) {
	pdName := fmt.Sprintf("%s-%s", framework.TestContext.Prefix, string(uuid.NewUUID()))
	_, diskURI, _, err := p.azureCloud.CreateVolume(pdName, "" /* account */, "" /* sku */, "" /* location */, 1 /* sizeGb */)
	if err != nil {
		return "", err
	}
	return diskURI, nil
}

func (p *Provider) DeletePD(pdName string) error {
	if err := p.azureCloud.DeleteVolume(pdName); err != nil {
		framework.Logf("failed to delete Azure volume %q: %v", pdName, err)
		return err
	}
	return nil
}

func (p *Provider) EnableAndDisableInternalLB() (enable, disable func(svc *v1.Service)) {
	enable = func(svc *v1.Service) {
		svc.ObjectMeta.Annotations = map[string]string{azure.ServiceAnnotationLoadBalancerInternal: "true"}
	}
	disable = func(svc *v1.Service) {
		svc.ObjectMeta.Annotations = map[string]string{azure.ServiceAnnotationLoadBalancerInternal: "false"}
	}
	return
}
