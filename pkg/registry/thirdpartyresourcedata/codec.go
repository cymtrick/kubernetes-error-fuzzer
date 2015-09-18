/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package thirdpartyresourcedata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"k8s.io/kubernetes/pkg/api/latest"
	"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/experimental"
	"k8s.io/kubernetes/pkg/runtime"
)

type thirdPartyResourceDataMapper struct {
	mapper  meta.RESTMapper
	kind    string
	version string
	group   string
}

func (t *thirdPartyResourceDataMapper) isThirdPartyResource(resource string) bool {
	return resource == strings.ToLower(t.kind)+"s"
}

func (t *thirdPartyResourceDataMapper) GroupForResource(resource string) (string, error) {
	if t.isThirdPartyResource(resource) {
		return t.group, nil
	}
	return t.mapper.GroupForResource(resource)
}

func (t *thirdPartyResourceDataMapper) RESTMapping(kind string, versions ...string) (*meta.RESTMapping, error) {
	if len(versions) != 1 {
		return nil, fmt.Errorf("unexpected set of versions: %v", versions)
	}
	if versions[0] != t.version {
		return nil, fmt.Errorf("unknown version %s expected %s", versions[0], t.version)
	}
	if kind != "ThirdPartyResourceData" {
		return nil, fmt.Errorf("unknown kind %s expected %s", kind, t.kind)
	}
	mapping, err := t.mapper.RESTMapping("ThirdPartyResourceData", latest.GroupOrDie("experimental").Version)
	if err != nil {
		return nil, err
	}
	mapping.Codec = NewCodec(mapping.Codec, t.kind)
	return mapping, nil
}

func (t *thirdPartyResourceDataMapper) AliasesForResource(resource string) ([]string, bool) {
	return t.mapper.AliasesForResource(resource)
}

func (t *thirdPartyResourceDataMapper) ResourceSingularizer(resource string) (singular string, err error) {
	return t.mapper.ResourceSingularizer(resource)
}

func (t *thirdPartyResourceDataMapper) VersionAndKindForResource(resource string) (defaultVersion, kind string, err error) {
	if t.isThirdPartyResource(resource) {
		return t.version, t.kind, nil
	}
	return t.mapper.VersionAndKindForResource(resource)
}

func NewMapper(mapper meta.RESTMapper, kind, version, group string) meta.RESTMapper {
	return &thirdPartyResourceDataMapper{
		mapper:  mapper,
		kind:    kind,
		version: version,
		group:   group,
	}
}

type thirdPartyResourceDataCodec struct {
	delegate runtime.Codec
	kind     string
}

func NewCodec(codec runtime.Codec, kind string) runtime.Codec {
	return &thirdPartyResourceDataCodec{codec, kind}
}

func (t *thirdPartyResourceDataCodec) populate(objIn *experimental.ThirdPartyResourceData, data []byte) error {
	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		fmt.Printf("Invalid JSON:\n%s\n", string(data))
		return err
	}
	mapObj, ok := obj.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected object: %#v", obj)
	}
	return t.populateFromObject(objIn, mapObj, data)
}

func (t *thirdPartyResourceDataCodec) populateFromObject(objIn *experimental.ThirdPartyResourceData, mapObj map[string]interface{}, data []byte) error {
	typeMeta := unversioned.TypeMeta{}
	if err := json.Unmarshal(data, &typeMeta); err != nil {
		return err
	}
	if typeMeta.Kind != t.kind {
		return fmt.Errorf("unexpected kind: %s, expected %s", typeMeta.Kind, t.kind)
	}

	metadata, ok := mapObj["metadata"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected object for metadata: %#v", mapObj["metadata"])
	}

	metadataData, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(metadataData, &objIn.ObjectMeta); err != nil {
		return err
	}

	objIn.Data = data
	return nil
}

func (t *thirdPartyResourceDataCodec) Decode(data []byte) (runtime.Object, error) {
	result := &experimental.ThirdPartyResourceData{}
	if err := t.populate(result, data); err != nil {
		return nil, err
	}
	return result, nil
}

func (t *thirdPartyResourceDataCodec) DecodeToVersion(data []byte, version string) (runtime.Object, error) {
	// TODO: this is hacky, there must be a better way...
	obj, err := t.Decode(data)
	if err != nil {
		return nil, err
	}
	objData, err := t.delegate.Encode(obj)
	if err != nil {
		return nil, err
	}
	return t.delegate.DecodeToVersion(objData, version)
}

func (t *thirdPartyResourceDataCodec) DecodeInto(data []byte, obj runtime.Object) error {
	thirdParty, ok := obj.(*experimental.ThirdPartyResourceData)
	if !ok {
		return fmt.Errorf("unexpected object: %#v", obj)
	}
	return t.populate(thirdParty, data)
}

func (t *thirdPartyResourceDataCodec) DecodeIntoWithSpecifiedVersionKind(data []byte, obj runtime.Object, version, kind string) error {
	thirdParty, ok := obj.(*experimental.ThirdPartyResourceData)
	if !ok {
		return fmt.Errorf("unexpected object: %#v", obj)
	}

	if kind != "ThirdPartyResourceData" {
		return fmt.Errorf("unexpeceted kind: %s", kind)
	}

	var dataObj interface{}
	if err := json.Unmarshal(data, &dataObj); err != nil {
		return err
	}
	mapObj, ok := dataObj.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpcted object: %#v", dataObj)
	}
	if kindObj, found := mapObj["kind"]; !found {
		mapObj["kind"] = kind
	} else {
		kindStr, ok := kindObj.(string)
		if !ok {
			return fmt.Errorf("unexpected object for 'kind': %v", kindObj)
		}
		if kindStr != t.kind {
			return fmt.Errorf("kind doesn't match, expecting: %s, got %s", kind, kindStr)
		}
	}
	if versionObj, found := mapObj["apiVersion"]; !found {
		mapObj["apiVersion"] = version
	} else {
		versionStr, ok := versionObj.(string)
		if !ok {
			return fmt.Errorf("unexpected object for 'apiVersion': %v", versionObj)
		}
		if versionStr != version {
			return fmt.Errorf("version doesn't match, expecting: %s, got %s", version, versionStr)
		}
	}

	if err := t.populate(thirdParty, data); err != nil {
		return err
	}
	return nil
}

const template = `{
  "kind": "%s",
  "items": [ %s ]
}`

func encodeToJSON(obj *experimental.ThirdPartyResourceData, stream io.Writer) error {
	var objOut interface{}
	if err := json.Unmarshal(obj.Data, &objOut); err != nil {
		return err
	}
	objMap, ok := objOut.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected type: %v", objOut)
	}
	objMap["metadata"] = obj.ObjectMeta
	encoder := json.NewEncoder(stream)
	return encoder.Encode(objMap)
}

func (t *thirdPartyResourceDataCodec) Encode(obj runtime.Object) ([]byte, error) {
	buff := &bytes.Buffer{}
	if err := t.EncodeToStream(obj, buff); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func (t *thirdPartyResourceDataCodec) EncodeToStream(obj runtime.Object, stream io.Writer) (err error) {
	switch obj := obj.(type) {
	case *experimental.ThirdPartyResourceData:
		return encodeToJSON(obj, stream)
	case *experimental.ThirdPartyResourceDataList:
		// TODO: There must be a better way to do this...
		dataStrings := make([]string, len(obj.Items))
		for ix := range obj.Items {
			buff := &bytes.Buffer{}
			err := encodeToJSON(&obj.Items[ix], buff)
			if err != nil {
				return err
			}
			dataStrings[ix] = buff.String()
		}
		fmt.Fprintf(stream, template, t.kind+"List", strings.Join(dataStrings, ","))
		return nil
	case *unversioned.Status:
		return t.delegate.EncodeToStream(obj, stream)
	default:
		return fmt.Errorf("unexpected object to encode: %#v", obj)
	}
}

func NewObjectCreator(version string, delegate runtime.ObjectCreater) runtime.ObjectCreater {
	return &thirdPartyResourceDataCreator{version, delegate}
}

type thirdPartyResourceDataCreator struct {
	version  string
	delegate runtime.ObjectCreater
}

func (t *thirdPartyResourceDataCreator) New(version, kind string) (out runtime.Object, err error) {
	switch kind {
	case "ThirdPartyResourceData":
		if t.version != version {
			return nil, fmt.Errorf("unknown version %s for kind %s", version, kind)
		}
		return &experimental.ThirdPartyResourceData{}, nil
	case "ThirdPartyResourceDataList":
		if t.version != version {
			return nil, fmt.Errorf("unknown version %s for kind %s", version, kind)
		}
		return &experimental.ThirdPartyResourceDataList{}, nil
	default:
		return t.delegate.New(version, kind)
	}
}
