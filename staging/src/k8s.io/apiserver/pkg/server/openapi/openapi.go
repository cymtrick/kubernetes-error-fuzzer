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

package openapi

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/go-openapi/spec"

	"k8s.io/apimachinery/pkg/openapi"
	"k8s.io/apimachinery/pkg/util/json"
	genericmux "k8s.io/apiserver/pkg/server/mux"
	"k8s.io/apiserver/pkg/util/trie"
)

const (
	OpenAPIVersion = "2.0"
)

type openAPI struct {
	config       *openapi.Config
	swagger      *spec.Swagger
	protocolList []string
	servePath    string
	definitions  map[string]openapi.OpenAPIDefinition
}

// RegisterOpenAPIService registers a handler to provides standard OpenAPI specification.
func RegisterOpenAPIService(servePath string, webServices []*restful.WebService, config *openapi.Config, mux *genericmux.PathRecorderMux) (err error) {
	o := openAPI{
		config:    config,
		servePath: servePath,
		swagger: &spec.Swagger{
			SwaggerProps: spec.SwaggerProps{
				Swagger:     OpenAPIVersion,
				Definitions: spec.Definitions{},
				Paths:       &spec.Paths{Paths: map[string]spec.PathItem{}},
				Info:        config.Info,
			},
		},
	}

	err = o.init(webServices)
	if err != nil {
		return err
	}

	mux.UnlistedHandleFunc(servePath, func(w http.ResponseWriter, r *http.Request) {
		resp := restful.NewResponse(w)
		if r.URL.Path != servePath {
			resp.WriteErrorString(http.StatusNotFound, "Path not found!")
		}
		// TODO: we can cache json string and return it here.
		resp.WriteAsJson(o.swagger)
	})
	return nil
}

func (o *openAPI) init(webServices []*restful.WebService) error {
	if o.config.GetOperationIDAndTags == nil {
		o.config.GetOperationIDAndTags = func(_ string, r *restful.Route) (string, []string, error) {
			return r.Operation, nil, nil
		}
	}
	if o.config.GetDefinitionName == nil {
		o.config.GetDefinitionName = func(_, name string) (string, spec.Extensions) {
			return name[strings.LastIndex(name, "/")+1:], nil
		}
	}
	o.definitions = o.config.GetDefinitions(func(name string) spec.Ref {
		defName, _ := o.config.GetDefinitionName(o.servePath, name)
		return spec.MustCreateRef("#/definitions/" + openapi.EscapeJsonPointer(defName))
	})
	if o.config.CommonResponses == nil {
		o.config.CommonResponses = map[int]spec.Response{}
	}
	err := o.buildPaths(webServices)
	if err != nil {
		return err
	}
	if o.config.SecurityDefinitions != nil {
		o.swagger.SecurityDefinitions = *o.config.SecurityDefinitions
		o.swagger.Security = o.config.DefaultSecurity
	}
	if o.config.PostProcessSpec != nil {
		o.swagger, err = o.config.PostProcessSpec(o.swagger)
		if err != nil {
			return err
		}
	}
	return nil
}

func getCanonicalizeTypeName(t reflect.Type) string {
	if t.PkgPath() == "" {
		return t.Name()
	}
	path := t.PkgPath()
	if strings.Contains(path, "/vendor/") {
		path = path[strings.Index(path, "/vendor/")+len("/vendor/"):]
	}
	return path + "." + t.Name()
}

func (o *openAPI) buildDefinitionRecursively(name string) error {
	uniqueName, extensions := o.config.GetDefinitionName(o.servePath, name)
	if _, ok := o.swagger.Definitions[uniqueName]; ok {
		return nil
	}
	if item, ok := o.definitions[name]; ok {
		schema := spec.Schema{
			VendorExtensible:   item.Schema.VendorExtensible,
			SchemaProps:        item.Schema.SchemaProps,
			SwaggerSchemaProps: item.Schema.SwaggerSchemaProps,
		}
		if extensions != nil {
			schema.Extensions = spec.Extensions{}
			for k, v := range extensions {
				schema.Extensions[k] = v
			}
		}
		o.swagger.Definitions[uniqueName] = schema
		for _, v := range item.Dependencies {
			if err := o.buildDefinitionRecursively(v); err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("cannot find model definition for %v. If you added a new type, you may need to add +k8s:openapi-gen=true to the package or type and run code-gen again.", name)
	}
	return nil
}

// buildDefinitionForType build a definition for a given type and return a referable name to it's definition.
// This is the main function that keep track of definitions used in this spec and is depend on code generated
// by k8s.io/kubernetes/cmd/libs/go2idl/openapi-gen.
func (o *openAPI) buildDefinitionForType(sample interface{}) (string, error) {
	t := reflect.TypeOf(sample)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	name := getCanonicalizeTypeName(t)
	if err := o.buildDefinitionRecursively(name); err != nil {
		return "", err
	}
	defName, _ := o.config.GetDefinitionName(o.servePath, name)
	return "#/definitions/" + openapi.EscapeJsonPointer(defName), nil
}

// buildPaths builds OpenAPI paths using go-restful's web services.
func (o *openAPI) buildPaths(webServices []*restful.WebService) error {
	pathsToIgnore := trie.New(o.config.IgnorePrefixes)
	duplicateOpId := make(map[string]string)
	for _, w := range webServices {
		rootPath := w.RootPath()
		if pathsToIgnore.HasPrefix(rootPath) {
			continue
		}
		commonParams, err := o.buildParameters(w.PathParameters())
		if err != nil {
			return err
		}
		for path, routes := range groupRoutesByPath(w.Routes()) {
			// go-swagger has special variable definition {$NAME:*} that can only be
			// used at the end of the path and it is not recognized by OpenAPI.
			if strings.HasSuffix(path, ":*}") {
				path = path[:len(path)-3] + "}"
			}
			if pathsToIgnore.HasPrefix(path) {
				continue
			}
			// Aggregating common parameters make API spec (and generated clients) simpler
			inPathCommonParamsMap, err := o.findCommonParameters(routes)
			if err != nil {
				return err
			}
			pathItem, exists := o.swagger.Paths.Paths[path]
			if exists {
				return fmt.Errorf("duplicate webservice route has been found for path: %v", path)
			}
			pathItem = spec.PathItem{
				PathItemProps: spec.PathItemProps{
					Parameters: make([]spec.Parameter, 0),
				},
			}
			// add web services's parameters as well as any parameters appears in all ops, as common parameters
			pathItem.Parameters = append(pathItem.Parameters, commonParams...)
			for _, p := range inPathCommonParamsMap {
				pathItem.Parameters = append(pathItem.Parameters, p)
			}
			sortParameters(pathItem.Parameters)
			for _, route := range routes {
				op, err := o.buildOperations(route, inPathCommonParamsMap)
				sortParameters(op.Parameters)
				if err != nil {
					return err
				}
				dpath, exists := duplicateOpId[op.ID]
				if exists {
					return fmt.Errorf("Duplicate Operation ID %v for path %v and %v.", op.ID, dpath, path)
				} else {
					duplicateOpId[op.ID] = path
				}
				switch strings.ToUpper(route.Method) {
				case "GET":
					pathItem.Get = op
				case "POST":
					pathItem.Post = op
				case "HEAD":
					pathItem.Head = op
				case "PUT":
					pathItem.Put = op
				case "DELETE":
					pathItem.Delete = op
				case "OPTIONS":
					pathItem.Options = op
				case "PATCH":
					pathItem.Patch = op
				}
			}
			o.swagger.Paths.Paths[path] = pathItem
		}
	}
	return nil
}

// buildOperations builds operations for each webservice path
func (o *openAPI) buildOperations(route restful.Route, inPathCommonParamsMap map[interface{}]spec.Parameter) (ret *spec.Operation, err error) {
	ret = &spec.Operation{
		OperationProps: spec.OperationProps{
			Description: route.Doc,
			Consumes:    route.Consumes,
			Produces:    route.Produces,
			Schemes:     o.config.ProtocolList,
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: make(map[int]spec.Response),
				},
			},
		},
	}
	if ret.ID, ret.Tags, err = o.config.GetOperationIDAndTags(o.servePath, &route); err != nil {
		return ret, err
	}

	// Build responses
	for _, resp := range route.ResponseErrors {
		ret.Responses.StatusCodeResponses[resp.Code], err = o.buildResponse(resp.Model, resp.Message)
		if err != nil {
			return ret, err
		}
	}
	// If there is no response but a write sample, assume that write sample is an http.StatusOK response.
	if len(ret.Responses.StatusCodeResponses) == 0 && route.WriteSample != nil {
		ret.Responses.StatusCodeResponses[http.StatusOK], err = o.buildResponse(route.WriteSample, "OK")
		if err != nil {
			return ret, err
		}
	}
	for code, resp := range o.config.CommonResponses {
		if _, exists := ret.Responses.StatusCodeResponses[code]; !exists {
			ret.Responses.StatusCodeResponses[code] = resp
		}
	}
	// If there is still no response, use default response provided.
	if len(ret.Responses.StatusCodeResponses) == 0 {
		ret.Responses.Default = o.config.DefaultResponse
	}

	// Build non-common Parameters
	ret.Parameters = make([]spec.Parameter, 0)
	for _, param := range route.ParameterDocs {
		if _, isCommon := inPathCommonParamsMap[mapKeyFromParam(param)]; !isCommon {
			openAPIParam, err := o.buildParameter(param.Data(), route.ReadSample)
			if err != nil {
				return ret, err
			}
			ret.Parameters = append(ret.Parameters, openAPIParam)
		}
	}
	return ret, nil
}

func (o *openAPI) buildResponse(model interface{}, description string) (spec.Response, error) {
	schema, err := o.toSchema(model)
	if err != nil {
		return spec.Response{}, err
	}
	return spec.Response{
		ResponseProps: spec.ResponseProps{
			Description: description,
			Schema:      schema,
		},
	}, nil
}

func (o *openAPI) findCommonParameters(routes []restful.Route) (map[interface{}]spec.Parameter, error) {
	commonParamsMap := make(map[interface{}]spec.Parameter, 0)
	paramOpsCountByName := make(map[interface{}]int, 0)
	paramNameKindToDataMap := make(map[interface{}]restful.ParameterData, 0)
	for _, route := range routes {
		routeParamDuplicateMap := make(map[interface{}]bool)
		s := ""
		for _, param := range route.ParameterDocs {
			m, _ := json.Marshal(param.Data())
			s += string(m) + "\n"
			key := mapKeyFromParam(param)
			if routeParamDuplicateMap[key] {
				msg, _ := json.Marshal(route.ParameterDocs)
				return commonParamsMap, fmt.Errorf("duplicate parameter %v for route %v, %v.", param.Data().Name, string(msg), s)
			}
			routeParamDuplicateMap[key] = true
			paramOpsCountByName[key]++
			paramNameKindToDataMap[key] = param.Data()
		}
	}
	for key, count := range paramOpsCountByName {
		paramData := paramNameKindToDataMap[key]
		if count == len(routes) && paramData.Kind != restful.BodyParameterKind {
			openAPIParam, err := o.buildParameter(paramData, nil)
			if err != nil {
				return commonParamsMap, err
			}
			commonParamsMap[key] = openAPIParam
		}
	}
	return commonParamsMap, nil
}

func (o *openAPI) toSchema(model interface{}) (_ *spec.Schema, err error) {
	if openAPIType, openAPIFormat := openapi.GetOpenAPITypeFormat(getCanonicalizeTypeName(reflect.TypeOf(model))); openAPIType != "" {
		return &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:   []string{openAPIType},
				Format: openAPIFormat,
			},
		}, nil
	} else {
		ref, err := o.buildDefinitionForType(model)
		if err != nil {
			return nil, err
		}
		return &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Ref: spec.MustCreateRef(ref),
			},
		}, nil
	}
}

func (o *openAPI) buildParameter(restParam restful.ParameterData, bodySample interface{}) (ret spec.Parameter, err error) {
	ret = spec.Parameter{
		ParamProps: spec.ParamProps{
			Name:        restParam.Name,
			Description: restParam.Description,
			Required:    restParam.Required,
		},
	}
	switch restParam.Kind {
	case restful.BodyParameterKind:
		if bodySample != nil {
			ret.In = "body"
			ret.Schema, err = o.toSchema(bodySample)
			return ret, err
		} else {
			// There is not enough information in the body parameter to build the definition.
			// Body parameter has a data type that is a short name but we need full package name
			// of the type to create a definition.
			return ret, fmt.Errorf("restful body parameters are not supported: %v", restParam.DataType)
		}
	case restful.PathParameterKind:
		ret.In = "path"
		if !restParam.Required {
			return ret, fmt.Errorf("path parameters should be marked at required for parameter %v", restParam)
		}
	case restful.QueryParameterKind:
		ret.In = "query"
	case restful.HeaderParameterKind:
		ret.In = "header"
	case restful.FormParameterKind:
		ret.In = "form"
	default:
		return ret, fmt.Errorf("unknown restful operation kind : %v", restParam.Kind)
	}
	openAPIType, openAPIFormat := openapi.GetOpenAPITypeFormat(restParam.DataType)
	if openAPIType == "" {
		return ret, fmt.Errorf("non-body Restful parameter type should be a simple type, but got : %v", restParam.DataType)
	}
	ret.Type = openAPIType
	ret.Format = openAPIFormat
	ret.UniqueItems = !restParam.AllowMultiple
	return ret, nil
}

func (o *openAPI) buildParameters(restParam []*restful.Parameter) (ret []spec.Parameter, err error) {
	ret = make([]spec.Parameter, len(restParam))
	for i, v := range restParam {
		ret[i], err = o.buildParameter(v.Data(), nil)
		if err != nil {
			return ret, err
		}
	}
	return ret, nil
}
