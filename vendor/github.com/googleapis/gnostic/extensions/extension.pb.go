// Copyright 2017 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: extensions/extension.proto

package gnostic_extension_v1

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The version number of Gnostic.
type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major int32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor int32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch int32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	// A suffix for alpha, beta or rc release, e.g., "alpha-1", "rc2". It should
	// be empty for mainline stable releases.
	Suffix string `protobuf:"bytes,4,opt,name=suffix,proto3" json:"suffix,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_extension_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_extension_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_extensions_extension_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetMajor() int32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Version) GetMinor() int32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Version) GetPatch() int32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

func (x *Version) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

// An encoded Request is written to the ExtensionHandler's stdin.
type ExtensionHandlerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The extension to process.
	Wrapper *Wrapper `protobuf:"bytes,1,opt,name=wrapper,proto3" json:"wrapper,omitempty"`
	// The version number of Gnostic.
	CompilerVersion *Version `protobuf:"bytes,2,opt,name=compiler_version,json=compilerVersion,proto3" json:"compiler_version,omitempty"`
}

func (x *ExtensionHandlerRequest) Reset() {
	*x = ExtensionHandlerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_extension_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionHandlerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionHandlerRequest) ProtoMessage() {}

func (x *ExtensionHandlerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_extension_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionHandlerRequest.ProtoReflect.Descriptor instead.
func (*ExtensionHandlerRequest) Descriptor() ([]byte, []int) {
	return file_extensions_extension_proto_rawDescGZIP(), []int{1}
}

func (x *ExtensionHandlerRequest) GetWrapper() *Wrapper {
	if x != nil {
		return x.Wrapper
	}
	return nil
}

func (x *ExtensionHandlerRequest) GetCompilerVersion() *Version {
	if x != nil {
		return x.CompilerVersion
	}
	return nil
}

// The extensions writes an encoded ExtensionHandlerResponse to stdout.
type ExtensionHandlerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// true if the extension is handled by the extension handler; false otherwise
	Handled bool `protobuf:"varint,1,opt,name=handled,proto3" json:"handled,omitempty"`
	// Error message(s).  If non-empty, the extension handling failed.
	// The extension handler process should exit with status code zero
	// even if it reports an error in this way.
	//
	// This should be used to indicate errors which prevent the extension from
	// operating as intended.  Errors which indicate a problem in gnostic
	// itself -- such as the input Document being unparseable -- should be
	// reported by writing a message to stderr and exiting with a non-zero
	// status code.
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	// text output
	Value *any.Any `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExtensionHandlerResponse) Reset() {
	*x = ExtensionHandlerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_extension_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionHandlerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionHandlerResponse) ProtoMessage() {}

func (x *ExtensionHandlerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_extension_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionHandlerResponse.ProtoReflect.Descriptor instead.
func (*ExtensionHandlerResponse) Descriptor() ([]byte, []int) {
	return file_extensions_extension_proto_rawDescGZIP(), []int{2}
}

func (x *ExtensionHandlerResponse) GetHandled() bool {
	if x != nil {
		return x.Handled
	}
	return false
}

func (x *ExtensionHandlerResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *ExtensionHandlerResponse) GetValue() *any.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type Wrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// version of the OpenAPI specification in which this extension was written.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Name of the extension.
	ExtensionName string `protobuf:"bytes,2,opt,name=extension_name,json=extensionName,proto3" json:"extension_name,omitempty"`
	// YAML-formatted extension value.
	Yaml string `protobuf:"bytes,3,opt,name=yaml,proto3" json:"yaml,omitempty"`
}

func (x *Wrapper) Reset() {
	*x = Wrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_extension_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wrapper) ProtoMessage() {}

func (x *Wrapper) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_extension_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wrapper.ProtoReflect.Descriptor instead.
func (*Wrapper) Descriptor() ([]byte, []int) {
	return file_extensions_extension_proto_rawDescGZIP(), []int{3}
}

func (x *Wrapper) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Wrapper) GetExtensionName() string {
	if x != nil {
		return x.ExtensionName
	}
	return ""
}

func (x *Wrapper) GetYaml() string {
	if x != nil {
		return x.Yaml
	}
	return ""
}

var File_extensions_extension_proto protoreflect.FileDescriptor

var file_extensions_extension_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x6e,
	0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d,
	0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75,
	0x66, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66,
	0x69, 0x78, 0x22, 0x9c, 0x01, 0x0a, 0x17, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37,
	0x0a, 0x07, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x07,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x78, 0x0a, 0x18, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12,
	0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5e, 0x0a, 0x07, 0x57,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x42, 0x4b, 0x0a, 0x0e, 0x6f,
	0x72, 0x67, 0x2e, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x47,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x01, 0x5a, 0x1f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x67, 0x6e,
	0x6f, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x4e, 0x58, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extensions_extension_proto_rawDescOnce sync.Once
	file_extensions_extension_proto_rawDescData = file_extensions_extension_proto_rawDesc
)

func file_extensions_extension_proto_rawDescGZIP() []byte {
	file_extensions_extension_proto_rawDescOnce.Do(func() {
		file_extensions_extension_proto_rawDescData = protoimpl.X.CompressGZIP(file_extensions_extension_proto_rawDescData)
	})
	return file_extensions_extension_proto_rawDescData
}

var file_extensions_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_extensions_extension_proto_goTypes = []interface{}{
	(*Version)(nil),                  // 0: gnostic.extension.v1.Version
	(*ExtensionHandlerRequest)(nil),  // 1: gnostic.extension.v1.ExtensionHandlerRequest
	(*ExtensionHandlerResponse)(nil), // 2: gnostic.extension.v1.ExtensionHandlerResponse
	(*Wrapper)(nil),                  // 3: gnostic.extension.v1.Wrapper
	(*any.Any)(nil),                  // 4: google.protobuf.Any
}
var file_extensions_extension_proto_depIdxs = []int32{
	3, // 0: gnostic.extension.v1.ExtensionHandlerRequest.wrapper:type_name -> gnostic.extension.v1.Wrapper
	0, // 1: gnostic.extension.v1.ExtensionHandlerRequest.compiler_version:type_name -> gnostic.extension.v1.Version
	4, // 2: gnostic.extension.v1.ExtensionHandlerResponse.value:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_extensions_extension_proto_init() }
func file_extensions_extension_proto_init() {
	if File_extensions_extension_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extensions_extension_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_extension_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionHandlerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_extension_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionHandlerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_extension_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wrapper); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_extensions_extension_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extensions_extension_proto_goTypes,
		DependencyIndexes: file_extensions_extension_proto_depIdxs,
		MessageInfos:      file_extensions_extension_proto_msgTypes,
	}.Build()
	File_extensions_extension_proto = out.File
	file_extensions_extension_proto_rawDesc = nil
	file_extensions_extension_proto_goTypes = nil
	file_extensions_extension_proto_depIdxs = nil
}
