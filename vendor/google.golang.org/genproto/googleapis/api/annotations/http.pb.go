// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/http.proto

package annotations // import "google.golang.org/genproto/googleapis/api/annotations"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Defines the HTTP configuration for an API service. It contains a list of
// [HttpRule][google.api.HttpRule], each specifying the mapping of an RPC method
// to one or more HTTP REST API methods.
type Http struct {
	// A list of HTTP configuration rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*HttpRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	// When set to true, URL path parameters will be fully URI-decoded except in
	// cases of single segment matches in reserved expansion, where "%2F" will be
	// left encoded.
	//
	// The default behavior is to not decode RFC 6570 reserved characters in multi
	// segment matches.
	FullyDecodeReservedExpansion bool     `protobuf:"varint,2,opt,name=fully_decode_reserved_expansion,json=fullyDecodeReservedExpansion,proto3" json:"fully_decode_reserved_expansion,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *Http) Reset()         { *m = Http{} }
func (m *Http) String() string { return proto.CompactTextString(m) }
func (*Http) ProtoMessage()    {}
func (*Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_5af6bbacbb935ee3, []int{0}
}
func (m *Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http.Unmarshal(m, b)
}
func (m *Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http.Marshal(b, m, deterministic)
}
func (dst *Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http.Merge(dst, src)
}
func (m *Http) XXX_Size() int {
	return xxx_messageInfo_Http.Size(m)
}
func (m *Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Http proto.InternalMessageInfo

func (m *Http) GetRules() []*HttpRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *Http) GetFullyDecodeReservedExpansion() bool {
	if m != nil {
		return m.FullyDecodeReservedExpansion
	}
	return false
}

// # gRPC Transcoding
//
// gRPC Transcoding is a feature for mapping between a gRPC method and one or
// more HTTP REST endpoints. It allows developers to build a single API service
// that supports both gRPC APIs and REST APIs. Many systems, including [Google
// APIs](https://github.com/googleapis/googleapis),
// [Cloud Endpoints](https://cloud.google.com/endpoints), [gRPC
// Gateway](https://github.com/grpc-ecosystem/grpc-gateway),
// and [Envoy](https://github.com/envoyproxy/envoy) proxy support this feature
// and use it for large scale production services.
//
// `HttpRule` defines the schema of the gRPC/REST mapping. The mapping specifies
// how different portions of the gRPC request message are mapped to the URL
// path, URL query parameters, and HTTP request body. It also controls how the
// gRPC response message is mapped to the HTTP response body. `HttpRule` is
// typically specified as an `google.api.http` annotation on the gRPC method.
//
// Each mapping specifies a URL path template and an HTTP method. The path
// template may refer to one or more fields in the gRPC request message, as long
// as each field is a non-repeated field with a primitive (non-message) type.
// The path template controls how fields of the request message are mapped to
// the URL path.
//
// Example:
//
//     service Messaging {
//       rpc GetMessage(GetMessageRequest) returns (Message) {
//         option (google.api.http) = {
//             get: "/v1/{name=messages/*}"
//         };
//       }
//     }
//     message GetMessageRequest {
//       string name = 1; // Mapped to URL path.
//     }
//     message Message {
//       string text = 1; // The resource content.
//     }
//
// This enables an HTTP REST to gRPC mapping as below:
//
// HTTP | gRPC
// -----|-----
// `GET /v1/messages/123456`  | `GetMessage(name: "messages/123456")`
//
// Any fields in the request message which are not bound by the path template
// automatically become HTTP query parameters if there is no HTTP request body.
// For example:
//
//     service Messaging {
//       rpc GetMessage(GetMessageRequest) returns (Message) {
//         option (google.api.http) = {
//             get:"/v1/messages/{message_id}"
//         };
//       }
//     }
//     message GetMessageRequest {
//       message SubMessage {
//         string subfield = 1;
//       }
//       string message_id = 1; // Mapped to URL path.
//       int64 revision = 2;    // Mapped to URL query parameter `revision`.
//       SubMessage sub = 3;    // Mapped to URL query parameter `sub.subfield`.
//     }
//
// This enables a HTTP JSON to RPC mapping as below:
//
// HTTP | gRPC
// -----|-----
// `GET /v1/messages/123456?revision=2&sub.subfield=foo` |
// `GetMessage(message_id: "123456" revision: 2 sub: SubMessage(subfield:
// "foo"))`
//
// Note that fields which are mapped to URL query parameters must have a
// primitive type or a repeated primitive type or a non-repeated message type.
// In the case of a repeated type, the parameter can be repeated in the URL
// as `...?param=A&param=B`. In the case of a message type, each field of the
// message is mapped to a separate parameter, such as
// `...?foo.a=A&foo.b=B&foo.c=C`.
//
// For HTTP methods that allow a request body, the `body` field
// specifies the mapping. Consider a REST update method on the
// message resource collection:
//
//     service Messaging {
//       rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
//         option (google.api.http) = {
//           patch: "/v1/messages/{message_id}"
//           body: "message"
//         };
//       }
//     }
//     message UpdateMessageRequest {
//       string message_id = 1; // mapped to the URL
//       Message message = 2;   // mapped to the body
//     }
//
// The following HTTP JSON to RPC mapping is enabled, where the
// representation of the JSON in the request body is determined by
// protos JSON encoding:
//
// HTTP | gRPC
// -----|-----
// `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
// "123456" message { text: "Hi!" })`
//
// The special name `*` can be used in the body mapping to define that
// every field not bound by the path template should be mapped to the
// request body.  This enables the following alternative definition of
// the update method:
//
//     service Messaging {
//       rpc UpdateMessage(Message) returns (Message) {
//         option (google.api.http) = {
//           patch: "/v1/messages/{message_id}"
//           body: "*"
//         };
//       }
//     }
//     message Message {
//       string message_id = 1;
//       string text = 2;
//     }
//
//
// The following HTTP JSON to RPC mapping is enabled:
//
// HTTP | gRPC
// -----|-----
// `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
// "123456" text: "Hi!")`
//
// Note that when using `*` in the body mapping, it is not possible to
// have HTTP parameters, as all fields not bound by the path end in
// the body. This makes this option more rarely used in practice when
// defining REST APIs. The common usage of `*` is in custom methods
// which don't use the URL at all for transferring data.
//
// It is possible to define multiple HTTP methods for one RPC by using
// the `additional_bindings` option. Example:
//
//     service Messaging {
//       rpc GetMessage(GetMessageRequest) returns (Message) {
//         option (google.api.http) = {
//           get: "/v1/messages/{message_id}"
//           additional_bindings {
//             get: "/v1/users/{user_id}/messages/{message_id}"
//           }
//         };
//       }
//     }
//     message GetMessageRequest {
//       string message_id = 1;
//       string user_id = 2;
//     }
//
// This enables the following two alternative HTTP JSON to RPC mappings:
//
// HTTP | gRPC
// -----|-----
// `GET /v1/messages/123456` | `GetMessage(message_id: "123456")`
// `GET /v1/users/me/messages/123456` | `GetMessage(user_id: "me" message_id:
// "123456")`
//
// ## Rules for HTTP mapping
//
// 1. Leaf request fields (recursive expansion nested messages in the request
//    message) are classified into three categories:
//    - Fields referred by the path template. They are passed via the URL path.
//    - Fields referred by the [HttpRule.body][google.api.HttpRule.body]. They are passed via the HTTP
//      request body.
//    - All other fields are passed via the URL query parameters, and the
//      parameter name is the field path in the request message. A repeated
//      field can be represented as multiple query parameters under the same
//      name.
//  2. If [HttpRule.body][google.api.HttpRule.body] is "*", there is no URL query parameter, all fields
//     are passed via URL path and HTTP request body.
//  3. If [HttpRule.body][google.api.HttpRule.body] is omitted, there is no HTTP request body, all
//     fields are passed via URL path and URL query parameters.
//
// ### Path template syntax
//
//     Template = "/" Segments [ Verb ] ;
//     Segments = Segment { "/" Segment } ;
//     Segment  = "*" | "**" | LITERAL | Variable ;
//     Variable = "{" FieldPath [ "=" Segments ] "}" ;
//     FieldPath = IDENT { "." IDENT } ;
//     Verb     = ":" LITERAL ;
//
// The syntax `*` matches a single URL path segment. The syntax `**` matches
// zero or more URL path segments, which must be the last part of the URL path
// except the `Verb`.
//
// The syntax `Variable` matches part of the URL path as specified by its
// template. A variable template must not contain other variables. If a variable
// matches a single path segment, its template may be omitted, e.g. `{var}`
// is equivalent to `{var=*}`.
//
// The syntax `LITERAL` matches literal text in the URL path. If the `LITERAL`
// contains any reserved character, such characters should be percent-encoded
// before the matching.
//
// If a variable contains exactly one path segment, such as `"{var}"` or
// `"{var=*}"`, when such a variable is expanded into a URL path on the client
// side, all characters except `[-_.~0-9a-zA-Z]` are percent-encoded. The
// server side does the reverse decoding. Such variables show up in the
// [Discovery
// Document](https://developers.google.com/discovery/v1/reference/apis) as
// `{var}`.
//
// If a variable contains multiple path segments, such as `"{var=foo/*}"`
// or `"{var=**}"`, when such a variable is expanded into a URL path on the
// client side, all characters except `[-_.~/0-9a-zA-Z]` are percent-encoded.
// The server side does the reverse decoding, except "%2F" and "%2f" are left
// unchanged. Such variables show up in the
// [Discovery
// Document](https://developers.google.com/discovery/v1/reference/apis) as
// `{+var}`.
//
// ## Using gRPC API Service Configuration
//
// gRPC API Service Configuration (service config) is a configuration language
// for configuring a gRPC service to become a user-facing product. The
// service config is simply the YAML representation of the `google.api.Service`
// proto message.
//
// As an alternative to annotating your proto file, you can configure gRPC
// transcoding in your service config YAML files. You do this by specifying a
// `HttpRule` that maps the gRPC method to a REST endpoint, achieving the same
// effect as the proto annotation. This can be particularly useful if you
// have a proto that is reused in multiple services. Note that any transcoding
// specified in the service config will override any matching transcoding
// configuration in the proto.
//
// Example:
//
//     http:
//       rules:
//         # Selects a gRPC method and applies HttpRule to it.
//         - selector: example.v1.Messaging.GetMessage
//           get: /v1/messages/{message_id}/{sub.subfield}
//
// ## Special notes
//
// When gRPC Transcoding is used to map a gRPC to JSON REST endpoints, the
// proto to JSON conversion must follow the [proto3
// specification](https://developers.google.com/protocol-buffers/docs/proto3#json).
//
// While the single segment variable follows the semantics of
// [RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.2 Simple String
// Expansion, the multi segment variable **does not** follow RFC 6570 Section
// 3.2.3 Reserved Expansion. The reason is that the Reserved Expansion
// does not expand special characters like `?` and `#`, which would lead
// to invalid URLs. As the result, gRPC Transcoding uses a custom encoding
// for multi segment variables.
//
// The path variables **must not** refer to any repeated or mapped field,
// because client libraries are not capable of handling such variable expansion.
//
// The path variables **must not** capture the leading "/" character. The reason
// is that the most common use case "{var}" does not capture the leading "/"
// character. For consistency, all path variables must share the same behavior.
//
// Repeated message fields must not be mapped to URL query parameters, because
// no client library can support such complicated mapping.
//
// If an API needs to use a JSON array for request or response body, it can map
// the request or response body to a repeated field. However, some gRPC
// Transcoding implementations may not support this feature.
type HttpRule struct {
	// Selects a method to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Determines the URL pattern is matched by this rules. This pattern can be
	// used with any of the {get|put|post|delete|patch} methods. A custom method
	// can be defined using the 'custom' field.
	//
	// Types that are valid to be assigned to Pattern:
	//	*HttpRule_Get
	//	*HttpRule_Put
	//	*HttpRule_Post
	//	*HttpRule_Delete
	//	*HttpRule_Patch
	//	*HttpRule_Custom
	Pattern isHttpRule_Pattern `protobuf_oneof:"pattern"`
	// The name of the request field whose value is mapped to the HTTP request
	// body, or `*` for mapping all request fields not captured by the path
	// pattern to the HTTP body, or omitted for not having any HTTP request body.
	//
	// NOTE: the referred field must be present at the top-level of the request
	// message type.
	Body string `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	// Optional. The name of the response field whose value is mapped to the HTTP
	// response body. When omitted, the entire response message will be used
	// as the HTTP response body.
	//
	// NOTE: The referred field must be present at the top-level of the response
	// message type.
	ResponseBody string `protobuf:"bytes,12,opt,name=response_body,json=responseBody,proto3" json:"response_body,omitempty"`
	// Additional HTTP bindings for the selector. Nested bindings must
	// not contain an `additional_bindings` field themselves (that is,
	// the nesting may only be one level deep).
	AdditionalBindings   []*HttpRule `protobuf:"bytes,11,rep,name=additional_bindings,json=additionalBindings,proto3" json:"additional_bindings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HttpRule) Reset()         { *m = HttpRule{} }
func (m *HttpRule) String() string { return proto.CompactTextString(m) }
func (*HttpRule) ProtoMessage()    {}
func (*HttpRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_5af6bbacbb935ee3, []int{1}
}
func (m *HttpRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpRule.Unmarshal(m, b)
}
func (m *HttpRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpRule.Marshal(b, m, deterministic)
}
func (dst *HttpRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRule.Merge(dst, src)
}
func (m *HttpRule) XXX_Size() int {
	return xxx_messageInfo_HttpRule.Size(m)
}
func (m *HttpRule) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRule.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRule proto.InternalMessageInfo

func (m *HttpRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

type isHttpRule_Pattern interface {
	isHttpRule_Pattern()
}

type HttpRule_Get struct {
	Get string `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}

type HttpRule_Put struct {
	Put string `protobuf:"bytes,3,opt,name=put,proto3,oneof"`
}

type HttpRule_Post struct {
	Post string `protobuf:"bytes,4,opt,name=post,proto3,oneof"`
}

type HttpRule_Delete struct {
	Delete string `protobuf:"bytes,5,opt,name=delete,proto3,oneof"`
}

type HttpRule_Patch struct {
	Patch string `protobuf:"bytes,6,opt,name=patch,proto3,oneof"`
}

type HttpRule_Custom struct {
	Custom *CustomHttpPattern `protobuf:"bytes,8,opt,name=custom,proto3,oneof"`
}

func (*HttpRule_Get) isHttpRule_Pattern() {}

func (*HttpRule_Put) isHttpRule_Pattern() {}

func (*HttpRule_Post) isHttpRule_Pattern() {}

func (*HttpRule_Delete) isHttpRule_Pattern() {}

func (*HttpRule_Patch) isHttpRule_Pattern() {}

func (*HttpRule_Custom) isHttpRule_Pattern() {}

func (m *HttpRule) GetPattern() isHttpRule_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *HttpRule) GetGet() string {
	if x, ok := m.GetPattern().(*HttpRule_Get); ok {
		return x.Get
	}
	return ""
}

func (m *HttpRule) GetPut() string {
	if x, ok := m.GetPattern().(*HttpRule_Put); ok {
		return x.Put
	}
	return ""
}

func (m *HttpRule) GetPost() string {
	if x, ok := m.GetPattern().(*HttpRule_Post); ok {
		return x.Post
	}
	return ""
}

func (m *HttpRule) GetDelete() string {
	if x, ok := m.GetPattern().(*HttpRule_Delete); ok {
		return x.Delete
	}
	return ""
}

func (m *HttpRule) GetPatch() string {
	if x, ok := m.GetPattern().(*HttpRule_Patch); ok {
		return x.Patch
	}
	return ""
}

func (m *HttpRule) GetCustom() *CustomHttpPattern {
	if x, ok := m.GetPattern().(*HttpRule_Custom); ok {
		return x.Custom
	}
	return nil
}

func (m *HttpRule) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *HttpRule) GetResponseBody() string {
	if m != nil {
		return m.ResponseBody
	}
	return ""
}

func (m *HttpRule) GetAdditionalBindings() []*HttpRule {
	if m != nil {
		return m.AdditionalBindings
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HttpRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HttpRule_OneofMarshaler, _HttpRule_OneofUnmarshaler, _HttpRule_OneofSizer, []interface{}{
		(*HttpRule_Get)(nil),
		(*HttpRule_Put)(nil),
		(*HttpRule_Post)(nil),
		(*HttpRule_Delete)(nil),
		(*HttpRule_Patch)(nil),
		(*HttpRule_Custom)(nil),
	}
}

func _HttpRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HttpRule)
	// pattern
	switch x := m.Pattern.(type) {
	case *HttpRule_Get:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Get)
	case *HttpRule_Put:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Put)
	case *HttpRule_Post:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Post)
	case *HttpRule_Delete:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Delete)
	case *HttpRule_Patch:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Patch)
	case *HttpRule_Custom:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Custom); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HttpRule.Pattern has unexpected type %T", x)
	}
	return nil
}

func _HttpRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HttpRule)
	switch tag {
	case 2: // pattern.get
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Get{x}
		return true, err
	case 3: // pattern.put
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Put{x}
		return true, err
	case 4: // pattern.post
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Post{x}
		return true, err
	case 5: // pattern.delete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Delete{x}
		return true, err
	case 6: // pattern.patch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Patch{x}
		return true, err
	case 8: // pattern.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CustomHttpPattern)
		err := b.DecodeMessage(msg)
		m.Pattern = &HttpRule_Custom{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HttpRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HttpRule)
	// pattern
	switch x := m.Pattern.(type) {
	case *HttpRule_Get:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Get)))
		n += len(x.Get)
	case *HttpRule_Put:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Put)))
		n += len(x.Put)
	case *HttpRule_Post:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Post)))
		n += len(x.Post)
	case *HttpRule_Delete:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Delete)))
		n += len(x.Delete)
	case *HttpRule_Patch:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Patch)))
		n += len(x.Patch)
	case *HttpRule_Custom:
		s := proto.Size(x.Custom)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A custom pattern is used for defining custom HTTP verb.
type CustomHttpPattern struct {
	// The name of this custom HTTP verb.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// The path matched by this custom verb.
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomHttpPattern) Reset()         { *m = CustomHttpPattern{} }
func (m *CustomHttpPattern) String() string { return proto.CompactTextString(m) }
func (*CustomHttpPattern) ProtoMessage()    {}
func (*CustomHttpPattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_5af6bbacbb935ee3, []int{2}
}
func (m *CustomHttpPattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomHttpPattern.Unmarshal(m, b)
}
func (m *CustomHttpPattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomHttpPattern.Marshal(b, m, deterministic)
}
func (dst *CustomHttpPattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomHttpPattern.Merge(dst, src)
}
func (m *CustomHttpPattern) XXX_Size() int {
	return xxx_messageInfo_CustomHttpPattern.Size(m)
}
func (m *CustomHttpPattern) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomHttpPattern.DiscardUnknown(m)
}

var xxx_messageInfo_CustomHttpPattern proto.InternalMessageInfo

func (m *CustomHttpPattern) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *CustomHttpPattern) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*Http)(nil), "google.api.Http")
	proto.RegisterType((*HttpRule)(nil), "google.api.HttpRule")
	proto.RegisterType((*CustomHttpPattern)(nil), "google.api.CustomHttpPattern")
}

func init() { proto.RegisterFile("google/api/http.proto", fileDescriptor_http_5af6bbacbb935ee3) }

var fileDescriptor_http_5af6bbacbb935ee3 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x49, 0x9b, 0x76, 0xdb, 0xe9, 0x82, 0x84, 0x59, 0x90, 0x85, 0x40, 0x54, 0xe5, 0x52,
	0x71, 0x48, 0xa5, 0xe5, 0xc0, 0x61, 0x4f, 0x1b, 0xa8, 0x58, 0x6e, 0x55, 0x8e, 0x5c, 0x22, 0x37,
	0x1e, 0x52, 0x83, 0xd7, 0xb6, 0xe2, 0x09, 0xa2, 0xaf, 0xc3, 0x63, 0xf1, 0x24, 0x1c, 0x91, 0x9d,
	0x84, 0x56, 0x42, 0xe2, 0x36, 0xf3, 0xff, 0x9f, 0xa7, 0x7f, 0x27, 0x03, 0x4f, 0x6b, 0x6b, 0x6b,
	0x8d, 0x1b, 0xe1, 0xd4, 0xe6, 0x40, 0xe4, 0x32, 0xd7, 0x58, 0xb2, 0x0c, 0x3a, 0x39, 0x13, 0x4e,
	0xad, 0x8e, 0x90, 0xde, 0x11, 0x39, 0xf6, 0x06, 0x26, 0x4d, 0xab, 0xd1, 0xf3, 0x64, 0x39, 0x5e,
	0x2f, 0xae, 0xaf, 0xb2, 0x13, 0x93, 0x05, 0xa0, 0x68, 0x35, 0x16, 0x1d, 0xc2, 0xb6, 0xf0, 0xea,
	0x4b, 0xab, 0xf5, 0xb1, 0x94, 0x58, 0x59, 0x89, 0x65, 0x83, 0x1e, 0x9b, 0xef, 0x28, 0x4b, 0xfc,
	0xe1, 0x84, 0xf1, 0xca, 0x1a, 0x3e, 0x5a, 0x26, 0xeb, 0x59, 0xf1, 0x22, 0x62, 0x1f, 0x22, 0x55,
	0xf4, 0xd0, 0x76, 0x60, 0x56, 0xbf, 0x46, 0x30, 0x1b, 0x46, 0xb3, 0xe7, 0x30, 0xf3, 0xa8, 0xb1,
	0x22, 0xdb, 0xf0, 0x64, 0x99, 0xac, 0xe7, 0xc5, 0xdf, 0x9e, 0x31, 0x18, 0xd7, 0x48, 0x71, 0xe6,
	0xfc, 0xee, 0x41, 0x11, 0x9a, 0xa0, 0xb9, 0x96, 0xf8, 0x78, 0xd0, 0x5c, 0x4b, 0xec, 0x0a, 0x52,
	0x67, 0x3d, 0xf1, 0xb4, 0x17, 0x63, 0xc7, 0x38, 0x4c, 0x25, 0x6a, 0x24, 0xe4, 0x93, 0x5e, 0xef,
	0x7b, 0xf6, 0x0c, 0x26, 0x4e, 0x50, 0x75, 0xe0, 0xd3, 0xde, 0xe8, 0x5a, 0xf6, 0x0e, 0xa6, 0x55,
	0xeb, 0xc9, 0xde, 0xf3, 0xd9, 0x32, 0x59, 0x2f, 0xae, 0x5f, 0x9e, 0x2f, 0xe3, 0x7d, 0x74, 0x42,
	0xee, 0x9d, 0x20, 0xc2, 0xc6, 0x84, 0x81, 0x1d, 0xce, 0x18, 0xa4, 0x7b, 0x2b, 0x8f, 0xfc, 0x22,
	0xfe, 0x81, 0x58, 0xb3, 0xd7, 0xf0, 0xb0, 0x41, 0xef, 0xac, 0xf1, 0x58, 0x46, 0xf3, 0x32, 0x9a,
	0x97, 0x83, 0x98, 0x07, 0x68, 0x0b, 0x4f, 0x84, 0x94, 0x8a, 0x94, 0x35, 0x42, 0x97, 0x7b, 0x65,
	0xa4, 0x32, 0xb5, 0xe7, 0x8b, 0xff, 0x7c, 0x0b, 0x76, 0x7a, 0x90, 0xf7, 0x7c, 0x3e, 0x87, 0x0b,
	0xd7, 0x85, 0x5a, 0xdd, 0xc0, 0xe3, 0x7f, 0x92, 0x86, 0x7c, 0xdf, 0x94, 0x91, 0xfd, 0x82, 0x63,
	0x1d, 0x34, 0x27, 0xe8, 0xd0, 0x6d, 0xb7, 0x88, 0x75, 0xfe, 0x15, 0x1e, 0x55, 0xf6, 0xfe, 0xec,
	0x67, 0xf3, 0x79, 0x1c, 0x13, 0xae, 0x67, 0x97, 0x7c, 0xbe, 0xed, 0x8d, 0xda, 0x6a, 0x61, 0xea,
	0xcc, 0x36, 0xf5, 0xa6, 0x46, 0x13, 0x6f, 0x6b, 0xd3, 0x59, 0xc2, 0x29, 0x1f, 0xaf, 0x4e, 0x18,
	0x63, 0x49, 0x84, 0x98, 0xfe, 0xe6, 0xac, 0xfe, 0x9d, 0x24, 0x3f, 0x47, 0xe9, 0xc7, 0xdb, 0xdd,
	0xa7, 0xfd, 0x34, 0xbe, 0x7b, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0xae, 0xde, 0xa1, 0xd0, 0xac,
	0x02, 0x00, 0x00,
}
