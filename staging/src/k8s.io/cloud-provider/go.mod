// This is a generated file. Do not edit directly.

module k8s.io/cloud-provider

go 1.15

require (
	github.com/google/go-cmp v0.4.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.4.0
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/component-base v0.0.0
	k8s.io/controller-manager v0.0.0
	k8s.io/klog/v2 v2.2.0
	k8s.io/utils v0.0.0-20200729134348-d5654de09c73
)

replace (
	github.com/Azure/go-autorest/autorest/mocks => github.com/Azure/go-autorest/autorest/mocks v0.4.0
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/apiserver => ../apiserver
	k8s.io/client-go => ../client-go
	k8s.io/cloud-provider => ../cloud-provider
	k8s.io/component-base => ../component-base
	k8s.io/controller-manager => ../controller-manager
)
