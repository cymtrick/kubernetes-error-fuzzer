// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/health/v1/health.proto

package grpc_health_v1 // import "google.golang.org/grpc/health/grpc_health_v1"

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_health_85731b6c49265086, []int{1, 0}
}

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_85731b6c49265086, []int{0}
}
func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (dst *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(dst, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=grpc.health.v1.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_85731b6c49265086, []int{1}
}
func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (dst *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(dst, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterType((*HealthCheckRequest)(nil), "grpc.health.v1.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "grpc.health.v1.HealthCheckResponse")
	proto.RegisterEnum("grpc.health.v1.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/grpc.health.v1.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.health.v1.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.health.v1.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/health/v1/health.proto",
}

func init() { proto.RegisterFile("grpc/health/v1/health.proto", fileDescriptor_health_85731b6c49265086) }

var fileDescriptor_health_85731b6c49265086 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x2f, 0x2a, 0x48,
	0xd6, 0xcf, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0xd0, 0x2f, 0x33, 0x84, 0xb2, 0xf4, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0xf8, 0x40, 0x92, 0x7a, 0x50, 0xa1, 0x32, 0x43, 0x25, 0x3d, 0x2e, 0x21, 0x0f,
	0x30, 0xc7, 0x39, 0x23, 0x35, 0x39, 0x3b, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x82,
	0x8b, 0xbd, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08,
	0xc6, 0x55, 0x9a, 0xc3, 0xc8, 0x25, 0x8c, 0xa2, 0xa1, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8,
	0x93, 0x8b, 0xad, 0xb8, 0x24, 0xb1, 0xa4, 0xb4, 0x18, 0xac, 0x81, 0xcf, 0xc8, 0x50, 0x0f, 0xd5,
	0x22, 0x3d, 0x2c, 0x9a, 0xf4, 0x82, 0x41, 0x86, 0xe6, 0xa5, 0x07, 0x83, 0x35, 0x06, 0x41, 0x0d,
	0x50, 0xb2, 0xe2, 0xe2, 0x45, 0x91, 0x10, 0xe2, 0xe6, 0x62, 0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f,
	0xf7, 0x13, 0x60, 0x00, 0x71, 0x82, 0x5d, 0x83, 0xc2, 0x3c, 0xfd, 0xdc, 0x05, 0x18, 0x85, 0xf8,
	0xb9, 0xb8, 0xfd, 0xfc, 0x43, 0xe2, 0x61, 0x02, 0x4c, 0x46, 0x51, 0x5c, 0x6c, 0x10, 0x8b, 0x84,
	0x02, 0xb8, 0x58, 0xc1, 0x96, 0x09, 0x29, 0xe1, 0x75, 0x09, 0xd8, 0xbf, 0x52, 0xca, 0x44, 0xb8,
	0xd6, 0x29, 0x91, 0x4b, 0x30, 0x33, 0x1f, 0x4d, 0xa1, 0x13, 0x37, 0x44, 0x65, 0x00, 0x28, 0x70,
	0x03, 0x18, 0xa3, 0x74, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xd2, 0xf3, 0x73, 0x12, 0xf3,
	0xd2, 0xf5, 0xf2, 0x8b, 0xd2, 0xf5, 0x91, 0x63, 0x03, 0xc4, 0x8e, 0x87, 0xb0, 0xe3, 0xcb, 0x0c,
	0x57, 0x31, 0xf1, 0xb9, 0x83, 0x4c, 0x83, 0x18, 0xa1, 0x17, 0x66, 0x98, 0xc4, 0x06, 0x8e, 0x24,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x66, 0x81, 0xcb, 0xc3, 0x01, 0x00, 0x00,
}
