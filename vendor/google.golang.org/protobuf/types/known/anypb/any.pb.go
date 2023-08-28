// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/any.proto

// Package anypb contains generated types for google/protobuf/any.proto.
//
// The Any message is a dynamic representation of any other message value.
// It is functionally a tuple of the full name of the remote message type and
// the serialized bytes of the remote message value.
//
// # Constructing an Any
//
// An Any message containing another message value is constructed using New:
//
//	any, err := anypb.New(m)
//	if err != nil {
//		... // handle error
//	}
//	... // make use of any
//
// # Unmarshaling an Any
//
// With a populated Any message, the underlying message can be serialized into
// a remote concrete message value in a few ways.
//
// If the exact concrete type is known, then a new (or pre-existing) instance
// of that message can be passed to the UnmarshalTo method:
//
//	m := new(foopb.MyMessage)
//	if err := any.UnmarshalTo(m); err != nil {
//		... // handle error
//	}
//	... // make use of m
//
// If the exact concrete type is not known, then the UnmarshalNew method can be
// used to unmarshal the contents into a new instance of the remote message type:
//
//	m, err := any.UnmarshalNew()
//	if err != nil {
//		... // handle error
//	}
//	... // make use of m
//
// UnmarshalNew uses the global type registry to resolve the message type and
// construct a new instance of that message to unmarshal into. In order for a
// message type to appear in the global registry, the Go type representing that
// protobuf message type must be linked into the Go binary. For messages
// generated by protoc-gen-go, this is achieved through an import of the
// generated Go package representing a .proto file.
//
// A common pattern with UnmarshalNew is to use a type switch with the resulting
// proto.Message value:
//
//	switch m := m.(type) {
//	case *foopb.MyMessage:
//		... // make use of m as a *foopb.MyMessage
//	case *barpb.OtherMessage:
//		... // make use of m as a *barpb.OtherMessage
//	case *bazpb.SomeMessage:
//		... // make use of m as a *bazpb.SomeMessage
//	}
//
// This pattern ensures that the generated packages containing the message types
// listed in the case clauses are linked into the Go binary and therefore also
// registered in the global registry.
//
// # Type checking an Any
//
// In order to type check whether an Any message represents some other message,
// then use the MessageIs method:
//
//	if any.MessageIs((*foopb.MyMessage)(nil)) {
//		... // make use of any, knowing that it contains a foopb.MyMessage
//	}
//
// The MessageIs method can also be used with an allocated instance of the target
// message type if the intention is to unmarshal into it if the type matches:
//
//	m := new(foopb.MyMessage)
//	if any.MessageIs(m) {
//		if err := any.UnmarshalTo(m); err != nil {
//			... // handle error
//		}
//		... // make use of m
//	}
package anypb

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoregistry "google.golang.org/protobuf/reflect/protoregistry"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	strings "strings"
	sync "sync"
)

// `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//	Foo foo = ...;
//	Any any;
//	any.PackFrom(foo);
//	...
//	if (any.UnpackTo(&foo)) {
//	  ...
//	}
//
// Example 2: Pack and unpack a message in Java.
//
//	   Foo foo = ...;
//	   Any any = Any.pack(foo);
//	   ...
//	   if (any.is(Foo.class)) {
//	     foo = any.unpack(Foo.class);
//	   }
//	   // or ...
//	   if (any.isSameTypeAs(Foo.getDefaultInstance())) {
//	     foo = any.unpack(Foo.getDefaultInstance());
//	   }
//
//	Example 3: Pack and unpack a message in Python.
//
//	   foo = Foo(...)
//	   any = Any()
//	   any.Pack(foo)
//	   ...
//	   if any.Is(Foo.DESCRIPTOR):
//	     any.Unpack(foo)
//	     ...
//
//	Example 4: Pack and unpack a message in Go
//
//	    foo := &pb.Foo{...}
//	    any, err := anypb.New(foo)
//	    if err != nil {
//	      ...
//	    }
//	    ...
//	    foo := &pb.Foo{}
//	    if err := any.UnmarshalTo(foo); err != nil {
//	      ...
//	    }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//	package google.profile;
//	message Person {
//	  string first_name = 1;
//	  string last_name = 2;
//	}
//
//	{
//	  "@type": "type.googleapis.com/google.profile.Person",
//	  "firstName": <string>,
//	  "lastName": <string>
//	}
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//	{
//	  "@type": "type.googleapis.com/google.protobuf.Duration",
//	  "value": "1.212s"
//	}
type Any struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	//   - If no scheme is provided, `https` is assumed.
	//   - An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//     value in binary format, or produce an error.
	//   - Applications are allowed to cache lookup results based on the
	//     URL, or have them precompiled into a binary to avoid any
	//     lookup. Therefore, binary compatibility needs to be preserved
	//     on changes to types. (Use versioned type names to manage
	//     breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Must be a valid serialized protocol buffer of the above specified type.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

// New marshals src into a new Any instance.
func New(src proto.Message) (*Any, error) {
	dst := new(Any)
	if err := dst.MarshalFrom(src); err != nil {
		return nil, err
	}
	return dst, nil
}

// MarshalFrom marshals src into dst as the underlying message
// using the provided marshal options.
//
// If no options are specified, call dst.MarshalFrom instead.
func MarshalFrom(dst *Any, src proto.Message, opts proto.MarshalOptions) error {
	const urlPrefix = "type.googleapis.com/"
	if src == nil {
		return protoimpl.X.NewError("invalid nil source message")
	}
	b, err := opts.Marshal(src)
	if err != nil {
		return err
	}
	dst.TypeUrl = urlPrefix + string(src.ProtoReflect().Descriptor().FullName())
	dst.Value = b
	return nil
}

// UnmarshalTo unmarshals the underlying message from src into dst
// using the provided unmarshal options.
// It reports an error if dst is not of the right message type.
//
// If no options are specified, call src.UnmarshalTo instead.
func UnmarshalTo(src *Any, dst proto.Message, opts proto.UnmarshalOptions) error {
	if src == nil {
		return protoimpl.X.NewError("invalid nil source message")
	}
	if !src.MessageIs(dst) {
		got := dst.ProtoReflect().Descriptor().FullName()
		want := src.MessageName()
		return protoimpl.X.NewError("mismatched message type: got %q, want %q", got, want)
	}
	return opts.Unmarshal(src.GetValue(), dst)
}

// UnmarshalNew unmarshals the underlying message from src into dst,
// which is newly created message using a type resolved from the type URL.
// The message type is resolved according to opt.Resolver,
// which should implement protoregistry.MessageTypeResolver.
// It reports an error if the underlying message type could not be resolved.
//
// If no options are specified, call src.UnmarshalNew instead.
func UnmarshalNew(src *Any, opts proto.UnmarshalOptions) (dst proto.Message, err error) {
	if src.GetTypeUrl() == "" {
		return nil, protoimpl.X.NewError("invalid empty type URL")
	}
	if opts.Resolver == nil {
		opts.Resolver = protoregistry.GlobalTypes
	}
	r, ok := opts.Resolver.(protoregistry.MessageTypeResolver)
	if !ok {
		return nil, protoregistry.NotFound
	}
	mt, err := r.FindMessageByURL(src.GetTypeUrl())
	if err != nil {
		if err == protoregistry.NotFound {
			return nil, err
		}
		return nil, protoimpl.X.NewError("could not resolve %q: %v", src.GetTypeUrl(), err)
	}
	dst = mt.New().Interface()
	return dst, opts.Unmarshal(src.GetValue(), dst)
}

// MessageIs reports whether the underlying message is of the same type as m.
func (x *Any) MessageIs(m proto.Message) bool {
	if m == nil {
		return false
	}
	url := x.GetTypeUrl()
	name := string(m.ProtoReflect().Descriptor().FullName())
	if !strings.HasSuffix(url, name) {
		return false
	}
	return len(url) == len(name) || url[len(url)-len(name)-1] == '/'
}

// MessageName reports the full name of the underlying message,
// returning an empty string if invalid.
func (x *Any) MessageName() protoreflect.FullName {
	url := x.GetTypeUrl()
	name := protoreflect.FullName(url)
	if i := strings.LastIndexByte(url, '/'); i >= 0 {
		name = name[i+len("/"):]
	}
	if !name.IsValid() {
		return ""
	}
	return name
}

// MarshalFrom marshals m into x as the underlying message.
func (x *Any) MarshalFrom(m proto.Message) error {
	return MarshalFrom(x, m, proto.MarshalOptions{})
}

// UnmarshalTo unmarshals the contents of the underlying message of x into m.
// It resets m before performing the unmarshal operation.
// It reports an error if m is not of the right message type.
func (x *Any) UnmarshalTo(m proto.Message) error {
	return UnmarshalTo(x, m, proto.UnmarshalOptions{})
}

// UnmarshalNew unmarshals the contents of the underlying message of x into
// a newly allocated message of the specified type.
// It reports an error if the underlying message type could not be resolved.
func (x *Any) UnmarshalNew() (proto.Message, error) {
	return UnmarshalNew(x, proto.UnmarshalOptions{})
}

func (x *Any) Reset() {
	*x = Any{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_any_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Any) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Any) ProtoMessage() {}

func (x *Any) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_any_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Any.ProtoReflect.Descriptor instead.
func (*Any) Descriptor() ([]byte, []int) {
	return file_google_protobuf_any_proto_rawDescGZIP(), []int{0}
}

func (x *Any) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *Any) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_google_protobuf_any_proto protoreflect.FileDescriptor

var file_google_protobuf_any_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x36, 0x0a, 0x03,
	0x41, 0x6e, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x76, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x08, 0x41, 0x6e, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f,
	0x61, 0x6e, 0x79, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x42, 0xaa, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x57, 0x65,
	0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_protobuf_any_proto_rawDescOnce sync.Once
	file_google_protobuf_any_proto_rawDescData = file_google_protobuf_any_proto_rawDesc
)

func file_google_protobuf_any_proto_rawDescGZIP() []byte {
	file_google_protobuf_any_proto_rawDescOnce.Do(func() {
		file_google_protobuf_any_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_any_proto_rawDescData)
	})
	return file_google_protobuf_any_proto_rawDescData
}

var file_google_protobuf_any_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_protobuf_any_proto_goTypes = []interface{}{
	(*Any)(nil), // 0: google.protobuf.Any
}
var file_google_protobuf_any_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_protobuf_any_proto_init() }
func file_google_protobuf_any_proto_init() {
	if File_google_protobuf_any_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_protobuf_any_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Any); i {
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
			RawDescriptor: file_google_protobuf_any_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_protobuf_any_proto_goTypes,
		DependencyIndexes: file_google_protobuf_any_proto_depIdxs,
		MessageInfos:      file_google_protobuf_any_proto_msgTypes,
	}.Build()
	File_google_protobuf_any_proto = out.File
	file_google_protobuf_any_proto_rawDesc = nil
	file_google_protobuf_any_proto_goTypes = nil
	file_google_protobuf_any_proto_depIdxs = nil
}
