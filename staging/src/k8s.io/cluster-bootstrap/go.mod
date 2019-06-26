// This is a generated file. Do not edit directly.

module k8s.io/cluster-bootstrap

go 1.12

require (
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8 // indirect
	gopkg.in/square/go-jose.v2 v2.2.2
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/klog v0.3.1
)

replace (
	github.com/stretchr/objx => github.com/stretchr/objx v0.1.1
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20181025213731-e84da0312774
	golang.org/x/net => golang.org/x/net v0.0.0-20190206173232-65e2d4e15006
	golang.org/x/sync => golang.org/x/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190209173611-3b5209105503
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190313210603-aa82965741a9
	gopkg.in/check.v1 => gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/cluster-bootstrap => ../cluster-bootstrap
)
