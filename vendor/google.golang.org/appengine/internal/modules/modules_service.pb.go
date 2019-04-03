// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google.golang.org/appengine/internal/modules/modules_service.proto

package modules

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

type ModulesServiceError_ErrorCode int32

const (
	ModulesServiceError_OK                ModulesServiceError_ErrorCode = 0
	ModulesServiceError_INVALID_MODULE    ModulesServiceError_ErrorCode = 1
	ModulesServiceError_INVALID_VERSION   ModulesServiceError_ErrorCode = 2
	ModulesServiceError_INVALID_INSTANCES ModulesServiceError_ErrorCode = 3
	ModulesServiceError_TRANSIENT_ERROR   ModulesServiceError_ErrorCode = 4
	ModulesServiceError_UNEXPECTED_STATE  ModulesServiceError_ErrorCode = 5
)

var ModulesServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INVALID_MODULE",
	2: "INVALID_VERSION",
	3: "INVALID_INSTANCES",
	4: "TRANSIENT_ERROR",
	5: "UNEXPECTED_STATE",
}
var ModulesServiceError_ErrorCode_value = map[string]int32{
	"OK":                0,
	"INVALID_MODULE":    1,
	"INVALID_VERSION":   2,
	"INVALID_INSTANCES": 3,
	"TRANSIENT_ERROR":   4,
	"UNEXPECTED_STATE":  5,
}

func (x ModulesServiceError_ErrorCode) Enum() *ModulesServiceError_ErrorCode {
	p := new(ModulesServiceError_ErrorCode)
	*p = x
	return p
}
func (x ModulesServiceError_ErrorCode) String() string {
	return proto.EnumName(ModulesServiceError_ErrorCode_name, int32(x))
}
func (x *ModulesServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ModulesServiceError_ErrorCode_value, data, "ModulesServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = ModulesServiceError_ErrorCode(value)
	return nil
}
func (ModulesServiceError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{0, 0}
}

type ModulesServiceError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModulesServiceError) Reset()         { *m = ModulesServiceError{} }
func (m *ModulesServiceError) String() string { return proto.CompactTextString(m) }
func (*ModulesServiceError) ProtoMessage()    {}
func (*ModulesServiceError) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{0}
}
func (m *ModulesServiceError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModulesServiceError.Unmarshal(m, b)
}
func (m *ModulesServiceError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModulesServiceError.Marshal(b, m, deterministic)
}
func (dst *ModulesServiceError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModulesServiceError.Merge(dst, src)
}
func (m *ModulesServiceError) XXX_Size() int {
	return xxx_messageInfo_ModulesServiceError.Size(m)
}
func (m *ModulesServiceError) XXX_DiscardUnknown() {
	xxx_messageInfo_ModulesServiceError.DiscardUnknown(m)
}

var xxx_messageInfo_ModulesServiceError proto.InternalMessageInfo

type GetModulesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetModulesRequest) Reset()         { *m = GetModulesRequest{} }
func (m *GetModulesRequest) String() string { return proto.CompactTextString(m) }
func (*GetModulesRequest) ProtoMessage()    {}
func (*GetModulesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{1}
}
func (m *GetModulesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModulesRequest.Unmarshal(m, b)
}
func (m *GetModulesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModulesRequest.Marshal(b, m, deterministic)
}
func (dst *GetModulesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModulesRequest.Merge(dst, src)
}
func (m *GetModulesRequest) XXX_Size() int {
	return xxx_messageInfo_GetModulesRequest.Size(m)
}
func (m *GetModulesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModulesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetModulesRequest proto.InternalMessageInfo

type GetModulesResponse struct {
	Module               []string `protobuf:"bytes,1,rep,name=module" json:"module,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetModulesResponse) Reset()         { *m = GetModulesResponse{} }
func (m *GetModulesResponse) String() string { return proto.CompactTextString(m) }
func (*GetModulesResponse) ProtoMessage()    {}
func (*GetModulesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{2}
}
func (m *GetModulesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModulesResponse.Unmarshal(m, b)
}
func (m *GetModulesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModulesResponse.Marshal(b, m, deterministic)
}
func (dst *GetModulesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModulesResponse.Merge(dst, src)
}
func (m *GetModulesResponse) XXX_Size() int {
	return xxx_messageInfo_GetModulesResponse.Size(m)
}
func (m *GetModulesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModulesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetModulesResponse proto.InternalMessageInfo

func (m *GetModulesResponse) GetModule() []string {
	if m != nil {
		return m.Module
	}
	return nil
}

type GetVersionsRequest struct {
	Module               *string  `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionsRequest) Reset()         { *m = GetVersionsRequest{} }
func (m *GetVersionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionsRequest) ProtoMessage()    {}
func (*GetVersionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{3}
}
func (m *GetVersionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionsRequest.Unmarshal(m, b)
}
func (m *GetVersionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionsRequest.Marshal(b, m, deterministic)
}
func (dst *GetVersionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionsRequest.Merge(dst, src)
}
func (m *GetVersionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetVersionsRequest.Size(m)
}
func (m *GetVersionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionsRequest proto.InternalMessageInfo

func (m *GetVersionsRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

type GetVersionsResponse struct {
	Version              []string `protobuf:"bytes,1,rep,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionsResponse) Reset()         { *m = GetVersionsResponse{} }
func (m *GetVersionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetVersionsResponse) ProtoMessage()    {}
func (*GetVersionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{4}
}
func (m *GetVersionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionsResponse.Unmarshal(m, b)
}
func (m *GetVersionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionsResponse.Marshal(b, m, deterministic)
}
func (dst *GetVersionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionsResponse.Merge(dst, src)
}
func (m *GetVersionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetVersionsResponse.Size(m)
}
func (m *GetVersionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionsResponse proto.InternalMessageInfo

func (m *GetVersionsResponse) GetVersion() []string {
	if m != nil {
		return m.Version
	}
	return nil
}

type GetDefaultVersionRequest struct {
	Module               *string  `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDefaultVersionRequest) Reset()         { *m = GetDefaultVersionRequest{} }
func (m *GetDefaultVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetDefaultVersionRequest) ProtoMessage()    {}
func (*GetDefaultVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{5}
}
func (m *GetDefaultVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDefaultVersionRequest.Unmarshal(m, b)
}
func (m *GetDefaultVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDefaultVersionRequest.Marshal(b, m, deterministic)
}
func (dst *GetDefaultVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDefaultVersionRequest.Merge(dst, src)
}
func (m *GetDefaultVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetDefaultVersionRequest.Size(m)
}
func (m *GetDefaultVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDefaultVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDefaultVersionRequest proto.InternalMessageInfo

func (m *GetDefaultVersionRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

type GetDefaultVersionResponse struct {
	Version              *string  `protobuf:"bytes,1,req,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDefaultVersionResponse) Reset()         { *m = GetDefaultVersionResponse{} }
func (m *GetDefaultVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetDefaultVersionResponse) ProtoMessage()    {}
func (*GetDefaultVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{6}
}
func (m *GetDefaultVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDefaultVersionResponse.Unmarshal(m, b)
}
func (m *GetDefaultVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDefaultVersionResponse.Marshal(b, m, deterministic)
}
func (dst *GetDefaultVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDefaultVersionResponse.Merge(dst, src)
}
func (m *GetDefaultVersionResponse) XXX_Size() int {
	return xxx_messageInfo_GetDefaultVersionResponse.Size(m)
}
func (m *GetDefaultVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDefaultVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDefaultVersionResponse proto.InternalMessageInfo

func (m *GetDefaultVersionResponse) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type GetNumInstancesRequest struct {
	Module               *string  `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version              *string  `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNumInstancesRequest) Reset()         { *m = GetNumInstancesRequest{} }
func (m *GetNumInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*GetNumInstancesRequest) ProtoMessage()    {}
func (*GetNumInstancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{7}
}
func (m *GetNumInstancesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNumInstancesRequest.Unmarshal(m, b)
}
func (m *GetNumInstancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNumInstancesRequest.Marshal(b, m, deterministic)
}
func (dst *GetNumInstancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNumInstancesRequest.Merge(dst, src)
}
func (m *GetNumInstancesRequest) XXX_Size() int {
	return xxx_messageInfo_GetNumInstancesRequest.Size(m)
}
func (m *GetNumInstancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNumInstancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNumInstancesRequest proto.InternalMessageInfo

func (m *GetNumInstancesRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *GetNumInstancesRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type GetNumInstancesResponse struct {
	Instances            *int64   `protobuf:"varint,1,req,name=instances" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNumInstancesResponse) Reset()         { *m = GetNumInstancesResponse{} }
func (m *GetNumInstancesResponse) String() string { return proto.CompactTextString(m) }
func (*GetNumInstancesResponse) ProtoMessage()    {}
func (*GetNumInstancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{8}
}
func (m *GetNumInstancesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNumInstancesResponse.Unmarshal(m, b)
}
func (m *GetNumInstancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNumInstancesResponse.Marshal(b, m, deterministic)
}
func (dst *GetNumInstancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNumInstancesResponse.Merge(dst, src)
}
func (m *GetNumInstancesResponse) XXX_Size() int {
	return xxx_messageInfo_GetNumInstancesResponse.Size(m)
}
func (m *GetNumInstancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNumInstancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNumInstancesResponse proto.InternalMessageInfo

func (m *GetNumInstancesResponse) GetInstances() int64 {
	if m != nil && m.Instances != nil {
		return *m.Instances
	}
	return 0
}

type SetNumInstancesRequest struct {
	Module               *string  `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version              *string  `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Instances            *int64   `protobuf:"varint,3,req,name=instances" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetNumInstancesRequest) Reset()         { *m = SetNumInstancesRequest{} }
func (m *SetNumInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*SetNumInstancesRequest) ProtoMessage()    {}
func (*SetNumInstancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{9}
}
func (m *SetNumInstancesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetNumInstancesRequest.Unmarshal(m, b)
}
func (m *SetNumInstancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetNumInstancesRequest.Marshal(b, m, deterministic)
}
func (dst *SetNumInstancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetNumInstancesRequest.Merge(dst, src)
}
func (m *SetNumInstancesRequest) XXX_Size() int {
	return xxx_messageInfo_SetNumInstancesRequest.Size(m)
}
func (m *SetNumInstancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetNumInstancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetNumInstancesRequest proto.InternalMessageInfo

func (m *SetNumInstancesRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *SetNumInstancesRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *SetNumInstancesRequest) GetInstances() int64 {
	if m != nil && m.Instances != nil {
		return *m.Instances
	}
	return 0
}

type SetNumInstancesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetNumInstancesResponse) Reset()         { *m = SetNumInstancesResponse{} }
func (m *SetNumInstancesResponse) String() string { return proto.CompactTextString(m) }
func (*SetNumInstancesResponse) ProtoMessage()    {}
func (*SetNumInstancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{10}
}
func (m *SetNumInstancesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetNumInstancesResponse.Unmarshal(m, b)
}
func (m *SetNumInstancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetNumInstancesResponse.Marshal(b, m, deterministic)
}
func (dst *SetNumInstancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetNumInstancesResponse.Merge(dst, src)
}
func (m *SetNumInstancesResponse) XXX_Size() int {
	return xxx_messageInfo_SetNumInstancesResponse.Size(m)
}
func (m *SetNumInstancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetNumInstancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetNumInstancesResponse proto.InternalMessageInfo

type StartModuleRequest struct {
	Module               *string  `protobuf:"bytes,1,req,name=module" json:"module,omitempty"`
	Version              *string  `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartModuleRequest) Reset()         { *m = StartModuleRequest{} }
func (m *StartModuleRequest) String() string { return proto.CompactTextString(m) }
func (*StartModuleRequest) ProtoMessage()    {}
func (*StartModuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{11}
}
func (m *StartModuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartModuleRequest.Unmarshal(m, b)
}
func (m *StartModuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartModuleRequest.Marshal(b, m, deterministic)
}
func (dst *StartModuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartModuleRequest.Merge(dst, src)
}
func (m *StartModuleRequest) XXX_Size() int {
	return xxx_messageInfo_StartModuleRequest.Size(m)
}
func (m *StartModuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartModuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartModuleRequest proto.InternalMessageInfo

func (m *StartModuleRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *StartModuleRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type StartModuleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartModuleResponse) Reset()         { *m = StartModuleResponse{} }
func (m *StartModuleResponse) String() string { return proto.CompactTextString(m) }
func (*StartModuleResponse) ProtoMessage()    {}
func (*StartModuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{12}
}
func (m *StartModuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartModuleResponse.Unmarshal(m, b)
}
func (m *StartModuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartModuleResponse.Marshal(b, m, deterministic)
}
func (dst *StartModuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartModuleResponse.Merge(dst, src)
}
func (m *StartModuleResponse) XXX_Size() int {
	return xxx_messageInfo_StartModuleResponse.Size(m)
}
func (m *StartModuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartModuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartModuleResponse proto.InternalMessageInfo

type StopModuleRequest struct {
	Module               *string  `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version              *string  `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopModuleRequest) Reset()         { *m = StopModuleRequest{} }
func (m *StopModuleRequest) String() string { return proto.CompactTextString(m) }
func (*StopModuleRequest) ProtoMessage()    {}
func (*StopModuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{13}
}
func (m *StopModuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopModuleRequest.Unmarshal(m, b)
}
func (m *StopModuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopModuleRequest.Marshal(b, m, deterministic)
}
func (dst *StopModuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopModuleRequest.Merge(dst, src)
}
func (m *StopModuleRequest) XXX_Size() int {
	return xxx_messageInfo_StopModuleRequest.Size(m)
}
func (m *StopModuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopModuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopModuleRequest proto.InternalMessageInfo

func (m *StopModuleRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *StopModuleRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type StopModuleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopModuleResponse) Reset()         { *m = StopModuleResponse{} }
func (m *StopModuleResponse) String() string { return proto.CompactTextString(m) }
func (*StopModuleResponse) ProtoMessage()    {}
func (*StopModuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{14}
}
func (m *StopModuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopModuleResponse.Unmarshal(m, b)
}
func (m *StopModuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopModuleResponse.Marshal(b, m, deterministic)
}
func (dst *StopModuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopModuleResponse.Merge(dst, src)
}
func (m *StopModuleResponse) XXX_Size() int {
	return xxx_messageInfo_StopModuleResponse.Size(m)
}
func (m *StopModuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopModuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopModuleResponse proto.InternalMessageInfo

type GetHostnameRequest struct {
	Module               *string  `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version              *string  `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Instance             *string  `protobuf:"bytes,3,opt,name=instance" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHostnameRequest) Reset()         { *m = GetHostnameRequest{} }
func (m *GetHostnameRequest) String() string { return proto.CompactTextString(m) }
func (*GetHostnameRequest) ProtoMessage()    {}
func (*GetHostnameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{15}
}
func (m *GetHostnameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHostnameRequest.Unmarshal(m, b)
}
func (m *GetHostnameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHostnameRequest.Marshal(b, m, deterministic)
}
func (dst *GetHostnameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHostnameRequest.Merge(dst, src)
}
func (m *GetHostnameRequest) XXX_Size() int {
	return xxx_messageInfo_GetHostnameRequest.Size(m)
}
func (m *GetHostnameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHostnameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHostnameRequest proto.InternalMessageInfo

func (m *GetHostnameRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *GetHostnameRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *GetHostnameRequest) GetInstance() string {
	if m != nil && m.Instance != nil {
		return *m.Instance
	}
	return ""
}

type GetHostnameResponse struct {
	Hostname             *string  `protobuf:"bytes,1,req,name=hostname" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHostnameResponse) Reset()         { *m = GetHostnameResponse{} }
func (m *GetHostnameResponse) String() string { return proto.CompactTextString(m) }
func (*GetHostnameResponse) ProtoMessage()    {}
func (*GetHostnameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_modules_service_9cd3bffe4e91c59a, []int{16}
}
func (m *GetHostnameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHostnameResponse.Unmarshal(m, b)
}
func (m *GetHostnameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHostnameResponse.Marshal(b, m, deterministic)
}
func (dst *GetHostnameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHostnameResponse.Merge(dst, src)
}
func (m *GetHostnameResponse) XXX_Size() int {
	return xxx_messageInfo_GetHostnameResponse.Size(m)
}
func (m *GetHostnameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHostnameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHostnameResponse proto.InternalMessageInfo

func (m *GetHostnameResponse) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func init() {
	proto.RegisterType((*ModulesServiceError)(nil), "appengine.ModulesServiceError")
	proto.RegisterType((*GetModulesRequest)(nil), "appengine.GetModulesRequest")
	proto.RegisterType((*GetModulesResponse)(nil), "appengine.GetModulesResponse")
	proto.RegisterType((*GetVersionsRequest)(nil), "appengine.GetVersionsRequest")
	proto.RegisterType((*GetVersionsResponse)(nil), "appengine.GetVersionsResponse")
	proto.RegisterType((*GetDefaultVersionRequest)(nil), "appengine.GetDefaultVersionRequest")
	proto.RegisterType((*GetDefaultVersionResponse)(nil), "appengine.GetDefaultVersionResponse")
	proto.RegisterType((*GetNumInstancesRequest)(nil), "appengine.GetNumInstancesRequest")
	proto.RegisterType((*GetNumInstancesResponse)(nil), "appengine.GetNumInstancesResponse")
	proto.RegisterType((*SetNumInstancesRequest)(nil), "appengine.SetNumInstancesRequest")
	proto.RegisterType((*SetNumInstancesResponse)(nil), "appengine.SetNumInstancesResponse")
	proto.RegisterType((*StartModuleRequest)(nil), "appengine.StartModuleRequest")
	proto.RegisterType((*StartModuleResponse)(nil), "appengine.StartModuleResponse")
	proto.RegisterType((*StopModuleRequest)(nil), "appengine.StopModuleRequest")
	proto.RegisterType((*StopModuleResponse)(nil), "appengine.StopModuleResponse")
	proto.RegisterType((*GetHostnameRequest)(nil), "appengine.GetHostnameRequest")
	proto.RegisterType((*GetHostnameResponse)(nil), "appengine.GetHostnameResponse")
}

func init() {
	proto.RegisterFile("google.golang.org/appengine/internal/modules/modules_service.proto", fileDescriptor_modules_service_9cd3bffe4e91c59a)
}

var fileDescriptor_modules_service_9cd3bffe4e91c59a = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x69, 0x02, 0xdb, 0xf2, 0x0e, 0x90, 0x3a, 0x5b, 0xd7, 0x4d, 0x1c, 0x50, 0x4e, 0x1c,
	0x50, 0x2b, 0x90, 0x10, 0xe7, 0xae, 0x35, 0x25, 0xb0, 0xa5, 0x28, 0xce, 0x2a, 0xc4, 0xa5, 0x0a,
	0xdb, 0x23, 0x8b, 0x94, 0xda, 0xc1, 0x76, 0x77, 0xe4, 0xbf, 0xe0, 0xff, 0x45, 0x4b, 0xed, 0xb6,
	0x81, 0x4e, 0x45, 0x68, 0xa7, 0xe4, 0x7d, 0xfe, 0xfc, 0x7b, 0x9f, 0x5f, 0xac, 0xc0, 0x59, 0x2e,
	0x44, 0x5e, 0x62, 0x2f, 0x17, 0x65, 0xc6, 0xf3, 0x9e, 0x90, 0x79, 0x3f, 0xab, 0x2a, 0xe4, 0x79,
	0xc1, 0xb1, 0x5f, 0x70, 0x8d, 0x92, 0x67, 0x65, 0x7f, 0x2e, 0xae, 0x17, 0x25, 0x2a, 0xfb, 0x9c,
	0x29, 0x94, 0xb7, 0xc5, 0x15, 0xf6, 0x2a, 0x29, 0xb4, 0x20, 0xde, 0x6a, 0x47, 0xf8, 0xab, 0x05,
	0xc1, 0xc5, 0xd2, 0xc4, 0x96, 0x1e, 0x2a, 0xa5, 0x90, 0xe1, 0x4f, 0xf0, 0xea, 0x97, 0xa1, 0xb8,
	0x46, 0xb2, 0x07, 0xce, 0xe4, 0x93, 0xff, 0x88, 0x10, 0x78, 0x1a, 0xc5, 0xd3, 0xc1, 0x79, 0x34,
	0x9a, 0x5d, 0x4c, 0x46, 0x97, 0xe7, 0xd4, 0x6f, 0x91, 0x00, 0x9e, 0x59, 0x6d, 0x4a, 0x13, 0x16,
	0x4d, 0x62, 0xdf, 0x21, 0x47, 0xd0, 0xb6, 0x62, 0x14, 0xb3, 0x74, 0x10, 0x0f, 0x29, 0xf3, 0xdd,
	0x3b, 0x6f, 0x9a, 0x0c, 0x62, 0x16, 0xd1, 0x38, 0x9d, 0xd1, 0x24, 0x99, 0x24, 0xfe, 0x63, 0x72,
	0x08, 0xfe, 0x65, 0x4c, 0xbf, 0x7c, 0xa6, 0xc3, 0x94, 0x8e, 0x66, 0x2c, 0x1d, 0xa4, 0xd4, 0x7f,
	0x12, 0x06, 0xd0, 0x1e, 0xa3, 0x36, 0xc9, 0x12, 0xfc, 0xb1, 0x40, 0xa5, 0xc3, 0x57, 0x40, 0x36,
	0x45, 0x55, 0x09, 0xae, 0x90, 0x74, 0x60, 0x6f, 0x79, 0xcc, 0x6e, 0xeb, 0x85, 0xfb, 0xd2, 0x4b,
	0x4c, 0x65, 0xdc, 0x53, 0x94, 0xaa, 0x10, 0xdc, 0x32, 0x1a, 0xee, 0xd6, 0x86, 0xbb, 0x0f, 0x41,
	0xc3, 0x6d, 0xe0, 0x5d, 0xd8, 0xbf, 0x5d, 0x6a, 0x86, 0x6e, 0xcb, 0xf0, 0x0d, 0x74, 0xc7, 0xa8,
	0x47, 0xf8, 0x3d, 0x5b, 0x94, 0x76, 0xdf, 0xae, 0x26, 0x6f, 0xe1, 0x64, 0xcb, 0x9e, 0x6d, 0xad,
	0x9c, 0xcd, 0x56, 0x1f, 0xa1, 0x33, 0x46, 0x1d, 0x2f, 0xe6, 0x11, 0x57, 0x3a, 0xe3, 0x57, 0xb8,
	0xeb, 0x34, 0x9b, 0x2c, 0xa7, 0x5e, 0x58, 0xb1, 0xde, 0xc1, 0xf1, 0x5f, 0x2c, 0x13, 0xe0, 0x39,
	0x78, 0x85, 0x15, 0xeb, 0x08, 0x6e, 0xb2, 0x16, 0xc2, 0x1b, 0xe8, 0xb0, 0x07, 0x0a, 0xd1, 0xec,
	0xe4, 0xfe, 0xd9, 0xe9, 0x04, 0x8e, 0xd9, 0xf6, 0x88, 0xe1, 0x7b, 0x20, 0x4c, 0x67, 0xd2, 0xdc,
	0x81, 0x6d, 0x01, 0x9c, 0xfb, 0x02, 0x34, 0x26, 0x7a, 0x04, 0x41, 0x83, 0x63, 0xf0, 0x14, 0xda,
	0x4c, 0x8b, 0xea, 0x7e, 0xfa, 0xbf, 0xcd, 0xf8, 0xf0, 0x2e, 0xe5, 0x1a, 0x63, 0xe0, 0xdf, 0xea,
	0xfb, 0xf8, 0x41, 0x28, 0xcd, 0xb3, 0xf9, 0xff, 0xd3, 0xc9, 0x29, 0x1c, 0xd8, 0x59, 0x75, 0xdd,
	0x7a, 0x69, 0x55, 0x87, 0xaf, 0xeb, 0x5b, 0xbc, 0xee, 0x61, 0xbe, 0xec, 0x29, 0x1c, 0xdc, 0x18,
	0xcd, 0x8c, 0x68, 0x55, 0x9f, 0x79, 0x5f, 0xf7, 0xcd, 0x5f, 0xe2, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x6e, 0xbc, 0xe0, 0x61, 0x5c, 0x04, 0x00, 0x00,
}
