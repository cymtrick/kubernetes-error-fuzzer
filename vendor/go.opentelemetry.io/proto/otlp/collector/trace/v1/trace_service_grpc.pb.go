// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.3
// source: opentelemetry/proto/collector/trace/v1/trace_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TraceServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, in *ExportTraceServiceRequest, opts ...grpc.CallOption) (*ExportTraceServiceResponse, error)
}

type traceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceServiceClient(cc grpc.ClientConnInterface) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) Export(ctx context.Context, in *ExportTraceServiceRequest, opts ...grpc.CallOption) (*ExportTraceServiceResponse, error) {
	out := new(ExportTraceServiceResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.collector.trace.v1.TraceService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceServiceServer is the server API for TraceService service.
// All implementations must embed UnimplementedTraceServiceServer
// for forward compatibility
type TraceServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *ExportTraceServiceRequest) (*ExportTraceServiceResponse, error)
	mustEmbedUnimplementedTraceServiceServer()
}

// UnimplementedTraceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (UnimplementedTraceServiceServer) Export(context.Context, *ExportTraceServiceRequest) (*ExportTraceServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}
func (UnimplementedTraceServiceServer) mustEmbedUnimplementedTraceServiceServer() {}

// UnsafeTraceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TraceServiceServer will
// result in compilation errors.
type UnsafeTraceServiceServer interface {
	mustEmbedUnimplementedTraceServiceServer()
}

func RegisterTraceServiceServer(s grpc.ServiceRegistrar, srv TraceServiceServer) {
	s.RegisterService(&TraceService_ServiceDesc, srv)
}

func _TraceService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportTraceServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.collector.trace.v1.TraceService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceServiceServer).Export(ctx, req.(*ExportTraceServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TraceService_ServiceDesc is the grpc.ServiceDesc for TraceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TraceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.collector.trace.v1.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _TraceService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/collector/trace/v1/trace_service.proto",
}
