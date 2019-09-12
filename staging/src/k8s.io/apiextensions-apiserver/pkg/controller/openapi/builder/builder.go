/*
Copyright 2019 The Kubernetes Authors.

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

package builder

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/emicklei/go-restful"
	"github.com/go-openapi/spec"

	v1 "k8s.io/api/autoscaling/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/validation"
	structuralschema "k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	openapiv2 "k8s.io/apiextensions-apiserver/pkg/controller/openapi/v2"
	generatedopenapi "k8s.io/apiextensions-apiserver/pkg/generated/openapi"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/endpoints"
	"k8s.io/apiserver/pkg/endpoints/openapi"
	"k8s.io/apiserver/pkg/features"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	openapibuilder "k8s.io/kube-openapi/pkg/builder"
	"k8s.io/kube-openapi/pkg/common"
	"k8s.io/kube-openapi/pkg/util"
)

const (
	// Reference and Go types for built-in metadata
	objectMetaSchemaRef = "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
	listMetaType        = "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"
	typeMetaType        = "k8s.io/apimachinery/pkg/apis/meta/v1.TypeMeta"

	definitionPrefix = "#/definitions/"
)

var (
	swaggerPartialObjectMetadataDescriptions = metav1beta1.PartialObjectMetadata{}.SwaggerDoc()

	nameToken      = "{name}"
	namespaceToken = "{namespace}"
)

var definitions map[string]common.OpenAPIDefinition
var buildDefinitions sync.Once
var namer *openapi.DefinitionNamer

// Options contains builder options.
type Options struct {
	// Convert to OpenAPI v2.
	V2 bool

	// Strip defaults.
	StripDefaults bool

	// Strip value validation.
	StripValueValidation bool
}

// BuildSwagger builds swagger for the given crd in the given version
func BuildSwagger(crd *apiextensions.CustomResourceDefinition, version string, opts Options) (*spec.Swagger, error) {
	var schema *structuralschema.Structural
	s, err := apiextensions.GetSchemaForVersion(crd, version)
	if err != nil {
		return nil, err
	}

	if s != nil && s.OpenAPIV3Schema != nil {
		if !validation.SchemaHasInvalidTypes(s.OpenAPIV3Schema) {
			if ss, err := structuralschema.NewStructural(s.OpenAPIV3Schema); err == nil {
				// skip non-structural schemas
				schema = ss

				if opts.StripDefaults {
					schema = schema.StripDefaults()
				}
				if opts.StripValueValidation {
					schema = schema.StripValueValidations()
				}

				schema = schema.Unfold()
			}
		}
	}

	// TODO(roycaihw): remove the WebService templating below. The following logic
	// comes from function registerResourceHandlers() in k8s.io/apiserver.
	// Alternatives are either (ideally) refactoring registerResourceHandlers() to
	// reuse the code, or faking an APIInstaller for CR to feed to registerResourceHandlers().
	b := newBuilder(crd, version, schema, opts.V2)

	// Sample response types for building web service
	sample := &CRDCanonicalTypeNamer{
		group:   b.group,
		version: b.version,
		kind:    b.kind,
	}
	sampleList := &CRDCanonicalTypeNamer{
		group:   b.group,
		version: b.version,
		kind:    b.listKind,
	}
	status := &metav1.Status{}
	patch := &metav1.Patch{}
	scale := &v1.Scale{}

	routes := make([]*restful.RouteBuilder, 0)
	root := fmt.Sprintf("/apis/%s/%s/%s", b.group, b.version, b.plural)

	if b.namespaced {
		routes = append(routes, b.buildRoute(root, "", "GET", "list", "list", sampleList).Operation("list"+b.kind+"ForAllNamespaces"))
		root = fmt.Sprintf("/apis/%s/%s/namespaces/{namespace}/%s", b.group, b.version, b.plural)
	}
	routes = append(routes, b.buildRoute(root, "", "GET", "list", "list", sampleList))
	routes = append(routes, b.buildRoute(root, "", "POST", "post", "create", sample).Reads(sample))
	routes = append(routes, b.buildRoute(root, "", "DELETE", "deletecollection", "deletecollection", status))

	routes = append(routes, b.buildRoute(root, "/{name}", "GET", "get", "read", sample))
	routes = append(routes, b.buildRoute(root, "/{name}", "PUT", "put", "replace", sample).Reads(sample))
	routes = append(routes, b.buildRoute(root, "/{name}", "DELETE", "delete", "delete", status))
	routes = append(routes, b.buildRoute(root, "/{name}", "PATCH", "patch", "patch", sample).Reads(patch))

	subresources, err := apiextensions.GetSubresourcesForVersion(crd, version)
	if err != nil {
		return nil, err
	}
	if subresources != nil && subresources.Status != nil {
		routes = append(routes, b.buildRoute(root, "/{name}/status", "GET", "get", "read", sample))
		routes = append(routes, b.buildRoute(root, "/{name}/status", "PUT", "put", "replace", sample).Reads(sample))
		routes = append(routes, b.buildRoute(root, "/{name}/status", "PATCH", "patch", "patch", sample).Reads(patch))
	}
	if subresources != nil && subresources.Scale != nil {
		routes = append(routes, b.buildRoute(root, "/{name}/scale", "GET", "get", "read", scale))
		routes = append(routes, b.buildRoute(root, "/{name}/scale", "PUT", "put", "replace", scale).Reads(scale))
		routes = append(routes, b.buildRoute(root, "/{name}/scale", "PATCH", "patch", "patch", scale).Reads(patch))
	}

	for _, route := range routes {
		b.ws.Route(route)
	}

	openAPISpec, err := openapibuilder.BuildOpenAPISpec([]*restful.WebService{b.ws}, b.getOpenAPIConfig())
	if err != nil {
		return nil, err
	}

	return openAPISpec, nil
}

// Implements CanonicalTypeNamer
var _ = util.OpenAPICanonicalTypeNamer(&CRDCanonicalTypeNamer{})

// CRDCanonicalTypeNamer implements CanonicalTypeNamer interface for CRDs to
// seed kube-openapi canonical type name without Go types
type CRDCanonicalTypeNamer struct {
	group   string
	version string
	kind    string
}

// OpenAPICanonicalTypeName returns canonical type name for given CRD
func (c *CRDCanonicalTypeNamer) OpenAPICanonicalTypeName() string {
	return fmt.Sprintf("%s/%s.%s", c.group, c.version, c.kind)
}

// builder contains validation schema and basic naming information for a CRD in
// one version. The builder works to build a WebService that kube-openapi can
// consume.
type builder struct {
	schema     *spec.Schema
	listSchema *spec.Schema
	ws         *restful.WebService

	group    string
	version  string
	kind     string
	listKind string
	plural   string

	namespaced bool
}

// subresource is a handy method to get subresource name. Valid inputs are:
//     input                     output
//     ""                        ""
//     "/"                       ""
//     "/{name}"                 ""
//     "/{name}/scale"           "scale"
//     "/{name}/scale/foo"       invalid input
func subresource(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) <= 2 {
		return ""
	}
	if len(parts) == 3 {
		return parts[2]
	}
	// panic to alert on programming error
	panic("failed to parse subresource; invalid path")
}

func (b *builder) descriptionFor(path, operationVerb string) string {
	var article string
	switch operationVerb {
	case "list":
		article = " objects of kind "
	case "read", "replace":
		article = " the specified "
	case "patch":
		article = " the specified "
	case "create", "delete":
		article = endpoints.GetArticleForNoun(b.kind, " ")
	default:
		article = ""
	}

	var description string
	sub := subresource(path)
	if len(sub) > 0 {
		sub = " " + sub + " of"
	}
	switch operationVerb {
	case "patch":
		description = "partially update" + sub + article + b.kind
	case "deletecollection":
		// to match the text for built-in APIs
		if len(sub) > 0 {
			sub = sub + " a"
		}
		description = "delete collection of" + sub + " " + b.kind
	default:
		description = operationVerb + sub + article + b.kind
	}

	return description
}

// buildRoute returns a RouteBuilder for WebService to consume and builds path in swagger
//     action can be one of: GET, PUT, PATCH, POST, DELETE;
//     verb can be one of: list, read, replace, patch, create, delete, deletecollection;
//     sample is the sample Go type for response type.
func (b *builder) buildRoute(root, path, httpMethod, actionVerb, operationVerb string, sample interface{}) *restful.RouteBuilder {
	var namespaced string
	if b.namespaced {
		namespaced = "Namespaced"
	}
	route := b.ws.Method(httpMethod).
		Path(root+path).
		To(func(req *restful.Request, res *restful.Response) {}).
		Doc(b.descriptionFor(path, operationVerb)).
		Param(b.ws.QueryParameter("pretty", "If 'true', then the output is pretty printed.")).
		Operation(operationVerb+namespaced+b.kind+strings.Title(subresource(path))).
		Metadata(endpoints.ROUTE_META_GVK, metav1.GroupVersionKind{
			Group:   b.group,
			Version: b.version,
			Kind:    b.kind,
		}).
		Metadata(endpoints.ROUTE_META_ACTION, actionVerb).
		Produces("application/json", "application/yaml").
		Returns(http.StatusOK, "OK", sample).
		Writes(sample)
	if strings.Contains(root, namespaceToken) || strings.Contains(path, namespaceToken) {
		route.Param(b.ws.PathParameter("namespace", "object name and auth scope, such as for teams and projects").DataType("string"))
	}
	if strings.Contains(root, nameToken) || strings.Contains(path, nameToken) {
		route.Param(b.ws.PathParameter("name", "name of the "+b.kind).DataType("string"))
	}

	// Build consume media types
	if httpMethod == "PATCH" {
		supportedTypes := []string{
			string(types.JSONPatchType),
			string(types.MergePatchType),
		}
		if utilfeature.DefaultFeatureGate.Enabled(features.ServerSideApply) {
			supportedTypes = append(supportedTypes, string(types.ApplyPatchType))
		}

		route.Consumes(supportedTypes...)
	} else {
		route.Consumes(runtime.ContentTypeJSON, runtime.ContentTypeYAML)
	}

	// Build option parameters
	switch actionVerb {
	case "get":
		// TODO: CRD support for export is still under consideration
		endpoints.AddObjectParams(b.ws, route, &metav1.GetOptions{})
	case "list", "deletecollection":
		endpoints.AddObjectParams(b.ws, route, &metav1.ListOptions{})
	case "put", "patch":
		// TODO: PatchOption added in feature branch but not in master yet
		endpoints.AddObjectParams(b.ws, route, &metav1.UpdateOptions{})
	case "post":
		endpoints.AddObjectParams(b.ws, route, &metav1.CreateOptions{})
	case "delete":
		endpoints.AddObjectParams(b.ws, route, &metav1.DeleteOptions{})
		route.Reads(&metav1.DeleteOptions{}).ParameterNamed("body").Required(false)
	}

	// Build responses
	switch actionVerb {
	case "post":
		route.Returns(http.StatusAccepted, "Accepted", sample)
		route.Returns(http.StatusCreated, "Created", sample)
	case "delete":
		route.Returns(http.StatusAccepted, "Accepted", sample)
	case "put":
		route.Returns(http.StatusCreated, "Created", sample)
	}

	return route
}

// buildKubeNative builds input schema with Kubernetes' native object meta, type meta and
// extensions
func (b *builder) buildKubeNative(schema *structuralschema.Structural, v2 bool, crdPreserveUnknownFields bool) (ret *spec.Schema) {
	// only add properties if we have a schema. Otherwise, kubectl would (wrongly) assume additionalProperties=false
	// and forbid anything outside of apiVersion, kind and metadata. We have to fix kubectl to stop doing this, e.g. by
	// adding additionalProperties=true support to explicitly allow additional fields.
	// TODO: fix kubectl to understand additionalProperties=true
	if schema == nil || (v2 && (schema.XPreserveUnknownFields || crdPreserveUnknownFields)) {
		ret = &spec.Schema{
			SchemaProps: spec.SchemaProps{Type: []string{"object"}},
		}
		// no, we cannot add more properties here, not even TypeMeta/ObjectMeta because kubectl will complain about
		// unknown fields for anything else.
	} else {
		if v2 {
			schema = openapiv2.ToStructuralOpenAPIV2(schema)
		}
		ret = schema.ToGoOpenAPI()
		ret.SetProperty("metadata", *spec.RefSchema(objectMetaSchemaRef).
			WithDescription(swaggerPartialObjectMetadataDescriptions["metadata"]))
		addTypeMetaProperties(ret)
		addEmbeddedProperties(ret, v2)
	}
	ret.AddExtension(endpoints.ROUTE_META_GVK, []interface{}{
		map[string]interface{}{
			"group":   b.group,
			"version": b.version,
			"kind":    b.kind,
		},
	})

	return ret
}

func addEmbeddedProperties(s *spec.Schema, v2 bool) {
	if s == nil {
		return
	}

	for k := range s.Properties {
		v := s.Properties[k]
		addEmbeddedProperties(&v, v2)
		s.Properties[k] = v
	}
	if s.Items != nil {
		addEmbeddedProperties(s.Items.Schema, v2)
	}
	if s.AdditionalProperties != nil {
		addEmbeddedProperties(s.AdditionalProperties.Schema, v2)
	}

	if isTrue, ok := s.VendorExtensible.Extensions.GetBool("x-kubernetes-preserve-unknown-fields"); ok && isTrue && v2 {
		// don't add metadata properties if we're publishing to openapi v2 and are allowing unknown fields.
		// adding these metadata properties makes kubectl refuse to validate unknown fields.
		return
	}
	if isTrue, ok := s.VendorExtensible.Extensions.GetBool("x-kubernetes-embedded-resource"); ok && isTrue {
		s.SetProperty("apiVersion", withDescription(getDefinition(typeMetaType).SchemaProps.Properties["apiVersion"],
			"apiVersion defines the versioned schema of this representation of an object. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
		))
		s.SetProperty("kind", withDescription(getDefinition(typeMetaType).SchemaProps.Properties["kind"],
			"kind is a string value representing the type of this object. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
		))
		s.SetProperty("metadata", *spec.RefSchema(objectMetaSchemaRef).WithDescription(swaggerPartialObjectMetadataDescriptions["metadata"]))

		req := sets.NewString(s.Required...)
		if !req.Has("kind") {
			s.Required = append(s.Required, "kind")
		}
		if !req.Has("apiVersion") {
			s.Required = append(s.Required, "apiVersion")
		}
	}
}

// getDefinition gets definition for given Kubernetes type. This function is extracted from
// kube-openapi builder logic
func getDefinition(name string) spec.Schema {
	buildDefinitions.Do(buildDefinitionsFunc)
	return definitions[name].Schema
}

func withDescription(s spec.Schema, desc string) spec.Schema {
	return *s.WithDescription(desc)
}

func buildDefinitionsFunc() {
	namer = openapi.NewDefinitionNamer(runtime.NewScheme())
	definitions = generatedopenapi.GetOpenAPIDefinitions(func(name string) spec.Ref {
		defName, _ := namer.GetDefinitionName(name)
		return spec.MustCreateRef(definitionPrefix + common.EscapeJsonPointer(defName))
	})
}

// addTypeMetaProperties adds Kubernetes-specific type meta properties to input schema:
//     apiVersion and kind
func addTypeMetaProperties(s *spec.Schema) {
	s.SetProperty("apiVersion", getDefinition(typeMetaType).SchemaProps.Properties["apiVersion"])
	s.SetProperty("kind", getDefinition(typeMetaType).SchemaProps.Properties["kind"])
}

// buildListSchema builds the list kind schema for the CRD
func (b *builder) buildListSchema() *spec.Schema {
	name := definitionPrefix + util.ToRESTFriendlyName(fmt.Sprintf("%s/%s/%s", b.group, b.version, b.kind))
	doc := fmt.Sprintf("List of %s. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md", b.plural)
	s := new(spec.Schema).WithDescription(fmt.Sprintf("%s is a list of %s", b.listKind, b.kind)).
		WithRequired("items").
		SetProperty("items", *spec.ArrayProperty(spec.RefSchema(name)).WithDescription(doc)).
		SetProperty("metadata", getDefinition(listMetaType))
	addTypeMetaProperties(s)
	s.AddExtension(endpoints.ROUTE_META_GVK, []map[string]string{
		{
			"group":   b.group,
			"version": b.version,
			"kind":    b.listKind,
		},
	})
	return s
}

// getOpenAPIConfig builds config which wires up generated definitions for kube-openapi to consume
func (b *builder) getOpenAPIConfig() *common.Config {
	return &common.Config{
		ProtocolList: []string{"https"},
		Info: &spec.Info{
			InfoProps: spec.InfoProps{
				Title:   "Kubernetes CRD Swagger",
				Version: "v0.1.0",
			},
		},
		CommonResponses: map[int]spec.Response{
			401: {
				ResponseProps: spec.ResponseProps{
					Description: "Unauthorized",
				},
			},
		},
		GetOperationIDAndTags: openapi.GetOperationIDAndTags,
		GetDefinitionName: func(name string) (string, spec.Extensions) {
			buildDefinitions.Do(buildDefinitionsFunc)
			return namer.GetDefinitionName(name)
		},
		GetDefinitions: func(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
			def := generatedopenapi.GetOpenAPIDefinitions(ref)
			def[fmt.Sprintf("%s/%s.%s", b.group, b.version, b.kind)] = common.OpenAPIDefinition{
				Schema: *b.schema,
			}
			def[fmt.Sprintf("%s/%s.%s", b.group, b.version, b.listKind)] = common.OpenAPIDefinition{
				Schema: *b.listSchema,
			}
			return def
		},
	}
}

func newBuilder(crd *apiextensions.CustomResourceDefinition, version string, schema *structuralschema.Structural, v2 bool) *builder {
	b := &builder{
		schema: &spec.Schema{
			SchemaProps: spec.SchemaProps{Type: []string{"object"}},
		},
		listSchema: &spec.Schema{},
		ws:         &restful.WebService{},

		group:    crd.Spec.Group,
		version:  version,
		kind:     crd.Spec.Names.Kind,
		listKind: crd.Spec.Names.ListKind,
		plural:   crd.Spec.Names.Plural,
	}
	if crd.Spec.Scope == apiextensions.NamespaceScoped {
		b.namespaced = true
	}

	// Pre-build schema with Kubernetes native properties
	preserveUnknownFields := crd.Spec.PreserveUnknownFields != nil && *crd.Spec.PreserveUnknownFields
	b.schema = b.buildKubeNative(schema, v2, preserveUnknownFields)
	b.listSchema = b.buildListSchema()

	return b
}
