// This is a generated file. Do not edit directly.

module k8s.io/api

go 1.13

require (
	github.com/gogo/protobuf v1.3.1
	github.com/stretchr/testify v1.4.0
	k8s.io/apimachinery v0.0.0
)

replace (
	golang.org/x/net => golang.org/x/net v0.0.0-20191004110552-13f9640d40b9
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	golang.org/x/xerrors => golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7
	gopkg.in/check.v1 => gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
)
