// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

// Code generated by ./scripts/makeOpenApiInfoDotGo.sh; DO NOT EDIT.

package kubernetesapi

import (
	"sigs.k8s.io/kustomize/kyaml/openapi/kubernetesapi/v1_21_2"
)

const Info = "{title:Kubernetes,version:v1.21.2}"

var OpenAPIMustAsset = map[string]func(string) []byte{
	"v1.21.2": v1_21_2.MustAsset,
}

const DefaultOpenAPI = "v1.21.2"
