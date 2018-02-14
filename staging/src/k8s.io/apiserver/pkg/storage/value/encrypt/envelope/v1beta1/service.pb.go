/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by protoc-gen-gogo.
// source: service.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	VersionRequest
	VersionResponse
	DecryptRequest
	DecryptResponse
	EncryptRequest
	EncryptResponse
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type VersionRequest struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *VersionRequest) Reset()                    { *m = VersionRequest{} }
func (m *VersionRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()               {}
func (*VersionRequest) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{0} }

func (m *VersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type VersionResponse struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Name of the KMS provider.
	RuntimeName string `protobuf:"bytes,2,opt,name=runtime_name,json=runtimeName,proto3" json:"runtime_name,omitempty"`
	// Version of the KMS provider. The string must be semver-compatible.
	RuntimeVersion string `protobuf:"bytes,3,opt,name=runtime_version,json=runtimeVersion,proto3" json:"runtime_version,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{1} }

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionResponse) GetRuntimeName() string {
	if m != nil {
		return m.RuntimeName
	}
	return ""
}

func (m *VersionResponse) GetRuntimeVersion() string {
	if m != nil {
		return m.RuntimeVersion
	}
	return ""
}

type DecryptRequest struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The data to be decrypted.
	Cipher []byte `protobuf:"bytes,2,opt,name=cipher,proto3" json:"cipher,omitempty"`
}

func (m *DecryptRequest) Reset()                    { *m = DecryptRequest{} }
func (m *DecryptRequest) String() string            { return proto.CompactTextString(m) }
func (*DecryptRequest) ProtoMessage()               {}
func (*DecryptRequest) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{2} }

func (m *DecryptRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DecryptRequest) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
}

type DecryptResponse struct {
	// The decrypted data.
	Plain []byte `protobuf:"bytes,1,opt,name=plain,proto3" json:"plain,omitempty"`
}

func (m *DecryptResponse) Reset()                    { *m = DecryptResponse{} }
func (m *DecryptResponse) String() string            { return proto.CompactTextString(m) }
func (*DecryptResponse) ProtoMessage()               {}
func (*DecryptResponse) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{3} }

func (m *DecryptResponse) GetPlain() []byte {
	if m != nil {
		return m.Plain
	}
	return nil
}

type EncryptRequest struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The data to be encrypted.
	Plain []byte `protobuf:"bytes,2,opt,name=plain,proto3" json:"plain,omitempty"`
}

func (m *EncryptRequest) Reset()                    { *m = EncryptRequest{} }
func (m *EncryptRequest) String() string            { return proto.CompactTextString(m) }
func (*EncryptRequest) ProtoMessage()               {}
func (*EncryptRequest) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{4} }

func (m *EncryptRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *EncryptRequest) GetPlain() []byte {
	if m != nil {
		return m.Plain
	}
	return nil
}

type EncryptResponse struct {
	// The encrypted data.
	Cipher []byte `protobuf:"bytes,1,opt,name=cipher,proto3" json:"cipher,omitempty"`
}

func (m *EncryptResponse) Reset()                    { *m = EncryptResponse{} }
func (m *EncryptResponse) String() string            { return proto.CompactTextString(m) }
func (*EncryptResponse) ProtoMessage()               {}
func (*EncryptResponse) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{5} }

func (m *EncryptResponse) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionRequest)(nil), "v1beta1.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "v1beta1.VersionResponse")
	proto.RegisterType((*DecryptRequest)(nil), "v1beta1.DecryptRequest")
	proto.RegisterType((*DecryptResponse)(nil), "v1beta1.DecryptResponse")
	proto.RegisterType((*EncryptRequest)(nil), "v1beta1.EncryptRequest")
	proto.RegisterType((*EncryptResponse)(nil), "v1beta1.EncryptResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KMSService service

type KMSServiceClient interface {
	// Version returns the runtime name and runtime version of the KMS provider.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Execute decryption operation in KMS provider.
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
	// Execute encryption operation in KMS provider.
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
}

type kMSServiceClient struct {
	cc *grpc.ClientConn
}

func NewKMSServiceClient(cc *grpc.ClientConn) KMSServiceClient {
	return &kMSServiceClient{cc}
}

func (c *kMSServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/v1beta1.KMSService/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kMSServiceClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := grpc.Invoke(ctx, "/v1beta1.KMSService/Decrypt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kMSServiceClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := grpc.Invoke(ctx, "/v1beta1.KMSService/Encrypt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KMSService service

type KMSServiceServer interface {
	// Version returns the runtime name and runtime version of the KMS provider.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// Execute decryption operation in KMS provider.
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
	// Execute encryption operation in KMS provider.
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
}

func RegisterKMSServiceServer(s *grpc.Server, srv KMSServiceServer) {
	s.RegisterService(&_KMSService_serviceDesc, srv)
}

func _KMSService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KMSServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1beta1.KMSService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KMSServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KMSService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KMSServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1beta1.KMSService/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KMSServiceServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KMSService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KMSServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1beta1.KMSService/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KMSServiceServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KMSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1beta1.KMSService",
	HandlerType: (*KMSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _KMSService_Version_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _KMSService_Decrypt_Handler,
		},
		{
			MethodName: "Encrypt",
			Handler:    _KMSService_Encrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptorService) }

var fileDescriptorService = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xde, 0xae, 0xb8, 0xc5, 0xb1, 0xb6, 0x10, 0x44, 0x8b, 0x27, 0xcd, 0x65, 0xd5, 0x43, 0x61,
	0xf5, 0x2e, 0x22, 0x7a, 0x12, 0x3d, 0x74, 0xc1, 0xab, 0x74, 0xcb, 0x80, 0x01, 0x9b, 0xc6, 0x24,
	0x5b, 0xf1, 0x1d, 0x7d, 0x28, 0xb1, 0x99, 0xd6, 0xb4, 0x22, 0xee, 0x71, 0x26, 0xdf, 0xdf, 0xcc,
	0x04, 0xf6, 0x0c, 0xea, 0x46, 0x94, 0x98, 0x29, 0x5d, 0xdb, 0x9a, 0x85, 0xcd, 0x62, 0x85, 0xb6,
	0x58, 0xf0, 0x73, 0x88, 0x9f, 0x50, 0x1b, 0x51, 0xcb, 0x1c, 0xdf, 0xd6, 0x68, 0x2c, 0x4b, 0x21,
	0x6c, 0x5c, 0x27, 0x0d, 0x8e, 0x83, 0xd3, 0x9d, 0xbc, 0x2b, 0xf9, 0x3b, 0x24, 0x3d, 0xd6, 0xa8,
	0x5a, 0x1a, 0xfc, 0x1b, 0xcc, 0x4e, 0x20, 0xd2, 0x6b, 0x69, 0x45, 0x85, 0xcf, 0xb2, 0xa8, 0x30,
	0x9d, 0xb6, 0xcf, 0xbb, 0xd4, 0x7b, 0x2c, 0x2a, 0x64, 0x73, 0x48, 0x3a, 0x48, 0x27, 0xb2, 0xd5,
	0xa2, 0x62, 0x6a, 0x93, 0x1b, 0xbf, 0x81, 0xf8, 0x16, 0x4b, 0xfd, 0xa1, 0xec, 0xbf, 0x21, 0xd9,
	0x01, 0xcc, 0x4a, 0xa1, 0x5e, 0x50, 0xb7, 0x8e, 0x51, 0x4e, 0x15, 0x9f, 0x43, 0xd2, 0x6b, 0x50,
	0xf8, 0x7d, 0xd8, 0x56, 0xaf, 0x85, 0x70, 0x12, 0x51, 0xee, 0x0a, 0x7e, 0x0d, 0xf1, 0x9d, 0xdc,
	0xd0, 0xac, 0x57, 0x98, 0xfa, 0x0a, 0x67, 0x90, 0xf4, 0x0a, 0x64, 0xf5, 0x93, 0x2a, 0xf0, 0x53,
	0x5d, 0x7c, 0x06, 0x00, 0xf7, 0x0f, 0xcb, 0xa5, 0x3b, 0x0e, 0xbb, 0x82, 0x90, 0x66, 0x66, 0x87,
	0x19, 0x9d, 0x28, 0x1b, 0xde, 0xe7, 0x28, 0xfd, 0xfd, 0xe0, 0x4c, 0xf8, 0xe4, 0x9b, 0x4f, 0x43,
	0x7a, 0xfc, 0xe1, 0xea, 0x3c, 0xfe, 0x68, 0x1f, 0x8e, 0x4f, 0xc9, 0x3d, 0xfe, 0x70, 0x1b, 0x1e,
	0x7f, 0x34, 0x24, 0x9f, 0xac, 0x66, 0xed, 0xef, 0xba, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x28,
	0x69, 0xfc, 0xea, 0x6e, 0x02, 0x00, 0x00,
}
