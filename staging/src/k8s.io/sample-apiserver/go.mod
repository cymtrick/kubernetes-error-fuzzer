// This is a generated file. Do not edit directly.

module k8s.io/sample-apiserver

go 1.14

require (
	github.com/go-openapi/spec v0.19.3
	github.com/google/gofuzz v1.1.0
	github.com/spf13/cobra v1.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/apiserver v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/code-generator v0.0.0
	k8s.io/component-base v0.0.0
	k8s.io/klog/v2 v2.1.0
	k8s.io/kube-openapi v0.0.0-20200427153329-656914f816f9
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200201011859-915c9c3d4ccf // pinned to release-branch.go1.14-std
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/apiserver => ../apiserver
	k8s.io/client-go => ../client-go
	k8s.io/code-generator => ../code-generator
	k8s.io/component-base => ../component-base
	k8s.io/sample-apiserver => ../sample-apiserver
)
