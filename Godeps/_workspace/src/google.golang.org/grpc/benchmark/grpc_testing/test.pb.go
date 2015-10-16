// Code generated by protoc-gen-go.
// source: src/google.golang.org/grpc/test/grpc_testing/test.proto
// DO NOT EDIT!

/*
Package grpc_testing is a generated protocol buffer package.

It is generated from these files:
	src/google.golang.org/grpc/test/grpc_testing/test.proto

It has these top-level messages:
	Empty
	Payload
	SimpleRequest
	SimpleResponse
	StreamingInputCallRequest
	StreamingInputCallResponse
	ResponseParameters
	StreamingOutputCallRequest
	StreamingOutputCallResponse
*/
package grpc_testing

import proto "github.com/golang/protobuf/proto"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// The type of payload that should be returned.
type PayloadType int32

const (
	// Compressable text format.
	PayloadType_COMPRESSABLE PayloadType = 0
	// Uncompressable binary format.
	PayloadType_UNCOMPRESSABLE PayloadType = 1
	// Randomly chosen from all other formats defined in this enum.
	PayloadType_RANDOM PayloadType = 2
)

var PayloadType_name = map[int32]string{
	0: "COMPRESSABLE",
	1: "UNCOMPRESSABLE",
	2: "RANDOM",
}
var PayloadType_value = map[string]int32{
	"COMPRESSABLE":   0,
	"UNCOMPRESSABLE": 1,
	"RANDOM":         2,
}

func (x PayloadType) Enum() *PayloadType {
	p := new(PayloadType)
	*p = x
	return p
}
func (x PayloadType) String() string {
	return proto.EnumName(PayloadType_name, int32(x))
}
func (x *PayloadType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PayloadType_value, data, "PayloadType")
	if err != nil {
		return err
	}
	*x = PayloadType(value)
	return nil
}

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}

// A block of data, to simply increase gRPC message size.
type Payload struct {
	// The type of data in body.
	Type *PayloadType `protobuf:"varint,1,opt,name=type,enum=grpc.testing.PayloadType" json:"type,omitempty"`
	// Primary contents of payload.
	Body             []byte `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}

func (m *Payload) GetType() PayloadType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return PayloadType_COMPRESSABLE
}

func (m *Payload) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// Unary request.
type SimpleRequest struct {
	// Desired payload type in the response from the server.
	// If response_type is RANDOM, server randomly chooses one from other formats.
	ResponseType *PayloadType `protobuf:"varint,1,opt,name=response_type,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`
	// Desired payload size in the response from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	ResponseSize *int32 `protobuf:"varint,2,opt,name=response_size" json:"response_size,omitempty"`
	// Optional input payload sent along with the request.
	Payload *Payload `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	// Whether SimpleResponse should include username.
	FillUsername *bool `protobuf:"varint,4,opt,name=fill_username" json:"fill_username,omitempty"`
	// Whether SimpleResponse should include OAuth scope.
	FillOauthScope   *bool  `protobuf:"varint,5,opt,name=fill_oauth_scope" json:"fill_oauth_scope,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SimpleRequest) Reset()         { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()    {}

func (m *SimpleRequest) GetResponseType() PayloadType {
	if m != nil && m.ResponseType != nil {
		return *m.ResponseType
	}
	return PayloadType_COMPRESSABLE
}

func (m *SimpleRequest) GetResponseSize() int32 {
	if m != nil && m.ResponseSize != nil {
		return *m.ResponseSize
	}
	return 0
}

func (m *SimpleRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SimpleRequest) GetFillUsername() bool {
	if m != nil && m.FillUsername != nil {
		return *m.FillUsername
	}
	return false
}

func (m *SimpleRequest) GetFillOauthScope() bool {
	if m != nil && m.FillOauthScope != nil {
		return *m.FillOauthScope
	}
	return false
}

// Unary response, as configured by the request.
type SimpleResponse struct {
	// Payload to increase message size.
	Payload *Payload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	// The user the request came from, for verifying authentication was
	// successful when the client expected it.
	Username *string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// OAuth scope.
	OauthScope       *string `protobuf:"bytes,3,opt,name=oauth_scope" json:"oauth_scope,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}

func (m *SimpleResponse) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SimpleResponse) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *SimpleResponse) GetOauthScope() string {
	if m != nil && m.OauthScope != nil {
		return *m.OauthScope
	}
	return ""
}

// Client-streaming request.
type StreamingInputCallRequest struct {
	// Optional input payload sent along with the request.
	Payload          *Payload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StreamingInputCallRequest) Reset()         { *m = StreamingInputCallRequest{} }
func (m *StreamingInputCallRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingInputCallRequest) ProtoMessage()    {}

func (m *StreamingInputCallRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Client-streaming response.
type StreamingInputCallResponse struct {
	// Aggregated size of payloads received from the client.
	AggregatedPayloadSize *int32 `protobuf:"varint,1,opt,name=aggregated_payload_size" json:"aggregated_payload_size,omitempty"`
	XXX_unrecognized      []byte `json:"-"`
}

func (m *StreamingInputCallResponse) Reset()         { *m = StreamingInputCallResponse{} }
func (m *StreamingInputCallResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingInputCallResponse) ProtoMessage()    {}

func (m *StreamingInputCallResponse) GetAggregatedPayloadSize() int32 {
	if m != nil && m.AggregatedPayloadSize != nil {
		return *m.AggregatedPayloadSize
	}
	return 0
}

// Configuration for a particular response.
type ResponseParameters struct {
	// Desired payload sizes in responses from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	Size *int32 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	// Desired interval between consecutive responses in the response stream in
	// microseconds.
	IntervalUs       *int32 `protobuf:"varint,2,opt,name=interval_us" json:"interval_us,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ResponseParameters) Reset()         { *m = ResponseParameters{} }
func (m *ResponseParameters) String() string { return proto.CompactTextString(m) }
func (*ResponseParameters) ProtoMessage()    {}

func (m *ResponseParameters) GetSize() int32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *ResponseParameters) GetIntervalUs() int32 {
	if m != nil && m.IntervalUs != nil {
		return *m.IntervalUs
	}
	return 0
}

// Server-streaming request.
type StreamingOutputCallRequest struct {
	// Desired payload type in the response from the server.
	// If response_type is RANDOM, the payload from each response in the stream
	// might be of different types. This is to simulate a mixed type of payload
	// stream.
	ResponseType *PayloadType `protobuf:"varint,1,opt,name=response_type,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`
	// Configuration for each expected response message.
	ResponseParameters []*ResponseParameters `protobuf:"bytes,2,rep,name=response_parameters" json:"response_parameters,omitempty"`
	// Optional input payload sent along with the request.
	Payload          *Payload `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StreamingOutputCallRequest) Reset()         { *m = StreamingOutputCallRequest{} }
func (m *StreamingOutputCallRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingOutputCallRequest) ProtoMessage()    {}

func (m *StreamingOutputCallRequest) GetResponseType() PayloadType {
	if m != nil && m.ResponseType != nil {
		return *m.ResponseType
	}
	return PayloadType_COMPRESSABLE
}

func (m *StreamingOutputCallRequest) GetResponseParameters() []*ResponseParameters {
	if m != nil {
		return m.ResponseParameters
	}
	return nil
}

func (m *StreamingOutputCallRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Server-streaming response, as configured by the request and parameters.
type StreamingOutputCallResponse struct {
	// Payload to increase response size.
	Payload          *Payload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StreamingOutputCallResponse) Reset()         { *m = StreamingOutputCallResponse{} }
func (m *StreamingOutputCallResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingOutputCallResponse) ProtoMessage()    {}

func (m *StreamingOutputCallResponse) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("grpc.testing.PayloadType", PayloadType_name, PayloadType_value)
}

// Client API for TestService service

type TestServiceClient interface {
	// One empty request followed by one empty response.
	EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// One request followed by a sequence of responses (streamed download).
	// The server returns the payload with client desired type and sizes.
	StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error)
	// A sequence of requests followed by one response (streamed upload).
	// The server returns the aggregated size of client payload as the result.
	StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error)
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error)
	// A sequence of requests followed by a sequence of responses.
	// The server buffers all the client requests and then serves them in order. A
	// stream of responses are returned to the client when the server starts with
	// first request.
	HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_HalfDuplexCallClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/grpc.testing.TestService/EmptyCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/grpc.testing.TestService/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/grpc.testing.TestService/StreamingOutputCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingOutputCallClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_StreamingOutputCallClient interface {
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceStreamingOutputCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingOutputCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[1], c.cc, "/grpc.testing.TestService/StreamingInputCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingInputCallClient{stream}
	return x, nil
}

type TestService_StreamingInputCallClient interface {
	Send(*StreamingInputCallRequest) error
	CloseAndRecv() (*StreamingInputCallResponse, error)
	grpc.ClientStream
}

type testServiceStreamingInputCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingInputCallClient) Send(m *StreamingInputCallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceStreamingInputCallClient) CloseAndRecv() (*StreamingInputCallResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamingInputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[2], c.cc, "/grpc.testing.TestService/FullDuplexCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceFullDuplexCallClient{stream}
	return x, nil
}

type TestService_FullDuplexCallClient interface {
	Send(*StreamingOutputCallRequest) error
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceFullDuplexCallClient struct {
	grpc.ClientStream
}

func (x *testServiceFullDuplexCallClient) Send(m *StreamingOutputCallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceFullDuplexCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_HalfDuplexCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[3], c.cc, "/grpc.testing.TestService/HalfDuplexCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceHalfDuplexCallClient{stream}
	return x, nil
}

type TestService_HalfDuplexCallClient interface {
	Send(*StreamingOutputCallRequest) error
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceHalfDuplexCallClient struct {
	grpc.ClientStream
}

func (x *testServiceHalfDuplexCallClient) Send(m *StreamingOutputCallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceHalfDuplexCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	// One empty request followed by one empty response.
	EmptyCall(context.Context, *Empty) (*Empty, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// One request followed by a sequence of responses (streamed download).
	// The server returns the payload with client desired type and sizes.
	StreamingOutputCall(*StreamingOutputCallRequest, TestService_StreamingOutputCallServer) error
	// A sequence of requests followed by one response (streamed upload).
	// The server returns the aggregated size of client payload as the result.
	StreamingInputCall(TestService_StreamingInputCallServer) error
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(TestService_FullDuplexCallServer) error
	// A sequence of requests followed by a sequence of responses.
	// The server buffers all the client requests and then serves them in order. A
	// stream of responses are returned to the client when the server starts with
	// first request.
	HalfDuplexCall(TestService_HalfDuplexCallServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_EmptyCall_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Empty)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(TestServiceServer).EmptyCall(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _TestService_UnaryCall_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(SimpleRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(TestServiceServer).UnaryCall(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _TestService_StreamingOutputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamingOutputCallRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).StreamingOutputCall(m, &testServiceStreamingOutputCallServer{stream})
}

type TestService_StreamingOutputCallServer interface {
	Send(*StreamingOutputCallResponse) error
	grpc.ServerStream
}

type testServiceStreamingOutputCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingOutputCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_StreamingInputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).StreamingInputCall(&testServiceStreamingInputCallServer{stream})
}

type TestService_StreamingInputCallServer interface {
	SendAndClose(*StreamingInputCallResponse) error
	Recv() (*StreamingInputCallRequest, error)
	grpc.ServerStream
}

type testServiceStreamingInputCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingInputCallServer) SendAndClose(m *StreamingInputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceStreamingInputCallServer) Recv() (*StreamingInputCallRequest, error) {
	m := new(StreamingInputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_FullDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).FullDuplexCall(&testServiceFullDuplexCallServer{stream})
}

type TestService_FullDuplexCallServer interface {
	Send(*StreamingOutputCallResponse) error
	Recv() (*StreamingOutputCallRequest, error)
	grpc.ServerStream
}

type testServiceFullDuplexCallServer struct {
	grpc.ServerStream
}

func (x *testServiceFullDuplexCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceFullDuplexCallServer) Recv() (*StreamingOutputCallRequest, error) {
	m := new(StreamingOutputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_HalfDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).HalfDuplexCall(&testServiceHalfDuplexCallServer{stream})
}

type TestService_HalfDuplexCallServer interface {
	Send(*StreamingOutputCallResponse) error
	Recv() (*StreamingOutputCallRequest, error)
	grpc.ServerStream
}

type testServiceHalfDuplexCallServer struct {
	grpc.ServerStream
}

func (x *testServiceHalfDuplexCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceHalfDuplexCallServer) Recv() (*StreamingOutputCallRequest, error) {
	m := new(StreamingOutputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyCall",
			Handler:    _TestService_EmptyCall_Handler,
		},
		{
			MethodName: "UnaryCall",
			Handler:    _TestService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingOutputCall",
			Handler:       _TestService_StreamingOutputCall_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamingInputCall",
			Handler:       _TestService_StreamingInputCall_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FullDuplexCall",
			Handler:       _TestService_FullDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "HalfDuplexCall",
			Handler:       _TestService_HalfDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}
