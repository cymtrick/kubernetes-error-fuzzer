// This is a generated file. Do not edit directly.

module k8s.io/sample-controller

go 1.12

require (
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/code-generator v0.0.0
	k8s.io/klog v0.0.0-20190306015804-8e90cee79f82
)

replace (
	github.com/gogo/protobuf => github.com/gogo/protobuf v0.0.0-20171007142547-342cbe0a0415
	github.com/onsi/ginkgo => github.com/onsi/ginkgo v0.0.0-20170318221715-67b9df7f55fe
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190209173611-3b5209105503
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190313210603-aa82965741a9
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/client-go => ../client-go
	k8s.io/code-generator => ../code-generator
	k8s.io/sample-controller => ../sample-controller
)
