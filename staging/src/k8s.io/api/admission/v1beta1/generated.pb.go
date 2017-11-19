/*
Copyright 2017 The Kubernetes Authors.

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
// source: k8s.io/kubernetes/vendor/k8s.io/api/admission/v1beta1/generated.proto
// DO NOT EDIT!

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/api/admission/v1beta1/generated.proto

	It has these top-level messages:
		AdmissionRequest
		AdmissionResponse
		AdmissionReview
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_apimachinery_pkg_apis_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

import k8s_io_apimachinery_pkg_types "k8s.io/apimachinery/pkg/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *AdmissionRequest) Reset()                    { *m = AdmissionRequest{} }
func (*AdmissionRequest) ProtoMessage()               {}
func (*AdmissionRequest) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *AdmissionResponse) Reset()                    { *m = AdmissionResponse{} }
func (*AdmissionResponse) ProtoMessage()               {}
func (*AdmissionResponse) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *AdmissionReview) Reset()                    { *m = AdmissionReview{} }
func (*AdmissionReview) ProtoMessage()               {}
func (*AdmissionReview) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*AdmissionRequest)(nil), "k8s.io.api.admission.v1beta1.AdmissionRequest")
	proto.RegisterType((*AdmissionResponse)(nil), "k8s.io.api.admission.v1beta1.AdmissionResponse")
	proto.RegisterType((*AdmissionReview)(nil), "k8s.io.api.admission.v1beta1.AdmissionReview")
}
func (m *AdmissionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdmissionRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.UID)))
	i += copy(dAtA[i:], m.UID)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Kind.Size()))
	n1, err := m.Kind.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Resource.Size()))
	n2, err := m.Resource.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.SubResource)))
	i += copy(dAtA[i:], m.SubResource)
	dAtA[i] = 0x2a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	dAtA[i] = 0x32
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Namespace)))
	i += copy(dAtA[i:], m.Namespace)
	dAtA[i] = 0x3a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Operation)))
	i += copy(dAtA[i:], m.Operation)
	dAtA[i] = 0x42
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.UserInfo.Size()))
	n3, err := m.UserInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x4a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Object.Size()))
	n4, err := m.Object.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x52
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.OldObject.Size()))
	n5, err := m.OldObject.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	return i, nil
}

func (m *AdmissionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdmissionResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.UID)))
	i += copy(dAtA[i:], m.UID)
	dAtA[i] = 0x10
	i++
	if m.Allowed {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if m.Result != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Result.Size()))
		n6, err := m.Result.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.Patch != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Patch)))
		i += copy(dAtA[i:], m.Patch)
	}
	if m.PatchType != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.PatchType)))
		i += copy(dAtA[i:], *m.PatchType)
	}
	return i, nil
}

func (m *AdmissionReview) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdmissionReview) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Request != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Request.Size()))
		n7, err := m.Request.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.Response != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Response.Size()))
		n8, err := m.Response.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AdmissionRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.UID)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Kind.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Resource.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.SubResource)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Namespace)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Operation)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.UserInfo.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Object.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.OldObject.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *AdmissionResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.UID)
	n += 1 + l + sovGenerated(uint64(l))
	n += 2
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Patch != nil {
		l = len(m.Patch)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.PatchType != nil {
		l = len(*m.PatchType)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *AdmissionReview) Size() (n int) {
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Response != nil {
		l = m.Response.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AdmissionRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AdmissionRequest{`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`Kind:` + strings.Replace(strings.Replace(this.Kind.String(), "GroupVersionKind", "k8s_io_apimachinery_pkg_apis_meta_v1.GroupVersionKind", 1), `&`, ``, 1) + `,`,
		`Resource:` + strings.Replace(strings.Replace(this.Resource.String(), "GroupVersionResource", "k8s_io_apimachinery_pkg_apis_meta_v1.GroupVersionResource", 1), `&`, ``, 1) + `,`,
		`SubResource:` + fmt.Sprintf("%v", this.SubResource) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Operation:` + fmt.Sprintf("%v", this.Operation) + `,`,
		`UserInfo:` + strings.Replace(strings.Replace(this.UserInfo.String(), "UserInfo", "k8s_io_api_authentication_v1.UserInfo", 1), `&`, ``, 1) + `,`,
		`Object:` + strings.Replace(strings.Replace(this.Object.String(), "RawExtension", "k8s_io_apimachinery_pkg_runtime.RawExtension", 1), `&`, ``, 1) + `,`,
		`OldObject:` + strings.Replace(strings.Replace(this.OldObject.String(), "RawExtension", "k8s_io_apimachinery_pkg_runtime.RawExtension", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AdmissionResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AdmissionResponse{`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`Allowed:` + fmt.Sprintf("%v", this.Allowed) + `,`,
		`Result:` + strings.Replace(fmt.Sprintf("%v", this.Result), "Status", "k8s_io_apimachinery_pkg_apis_meta_v1.Status", 1) + `,`,
		`Patch:` + valueToStringGenerated(this.Patch) + `,`,
		`PatchType:` + valueToStringGenerated(this.PatchType) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AdmissionReview) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AdmissionReview{`,
		`Request:` + strings.Replace(fmt.Sprintf("%v", this.Request), "AdmissionRequest", "AdmissionRequest", 1) + `,`,
		`Response:` + strings.Replace(fmt.Sprintf("%v", this.Response), "AdmissionResponse", "AdmissionResponse", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AdmissionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdmissionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdmissionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = k8s_io_apimachinery_pkg_types.UID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Kind.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubResource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubResource = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operation = Operation(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UserInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Object", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Object.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldObject", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OldObject.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AdmissionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdmissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdmissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = k8s_io_apimachinery_pkg_types.UID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Allowed = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &k8s_io_apimachinery_pkg_apis_meta_v1.Status{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Patch", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Patch = append(m.Patch[:0], dAtA[iNdEx:postIndex]...)
			if m.Patch == nil {
				m.Patch = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PatchType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := PatchType(dAtA[iNdEx:postIndex])
			m.PatchType = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AdmissionReview) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdmissionReview: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdmissionReview: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &AdmissionRequest{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Response == nil {
				m.Response = &AdmissionResponse{}
			}
			if err := m.Response.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/admission/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0x8e, 0x21, 0x7f, 0x9e, 0xa0, 0x0b, 0xcc, 0xdd, 0x58, 0xd1, 0x95, 0xc3, 0x65, 0x71, 0xc5,
	0x95, 0x60, 0x5c, 0x68, 0x8b, 0x50, 0xd5, 0x0d, 0x16, 0xa8, 0x42, 0x95, 0x00, 0x0d, 0xa4, 0x6a,
	0xbb, 0xa8, 0x34, 0x71, 0x86, 0x64, 0x9a, 0xd8, 0xe3, 0x7a, 0xc6, 0xa1, 0xec, 0xfa, 0x08, 0x7d,
	0x93, 0x3e, 0x44, 0x37, 0x2c, 0x59, 0xb2, 0x8a, 0x4a, 0xfa, 0x00, 0xdd, 0xb3, 0xaa, 0x3c, 0x1e,
	0xc7, 0x29, 0x34, 0x2d, 0xad, 0xba, 0xca, 0x9c, 0x73, 0xbe, 0xef, 0x3b, 0xf1, 0x77, 0xce, 0x0c,
	0xd8, 0xed, 0x6d, 0x09, 0xc4, 0xb8, 0xd3, 0x8b, 0x5b, 0x34, 0x0a, 0xa8, 0xa4, 0xc2, 0x19, 0xd0,
	0xa0, 0xcd, 0x23, 0x47, 0x17, 0x48, 0xc8, 0x1c, 0xd2, 0xf6, 0x99, 0x10, 0x8c, 0x07, 0xce, 0x60,
	0xbd, 0x45, 0x25, 0x59, 0x77, 0x3a, 0x34, 0xa0, 0x11, 0x91, 0xb4, 0x8d, 0xc2, 0x88, 0x4b, 0x0e,
	0xff, 0x49, 0xd1, 0x88, 0x84, 0x0c, 0x8d, 0xd1, 0x48, 0xa3, 0xeb, 0x6b, 0x1d, 0x26, 0xbb, 0x71,
	0x0b, 0x79, 0xdc, 0x77, 0x3a, 0xbc, 0xc3, 0x1d, 0x45, 0x6a, 0xc5, 0x27, 0x2a, 0x52, 0x81, 0x3a,
	0xa5, 0x62, 0xf5, 0xd5, 0xc9, 0xd6, 0xb1, 0xec, 0xd2, 0x40, 0x32, 0x8f, 0xc8, 0xb4, 0xff, 0xcd,
	0xd6, 0xf5, 0x07, 0x39, 0xda, 0x27, 0x5e, 0x97, 0x05, 0x34, 0x3a, 0x73, 0xc2, 0x5e, 0x27, 0x49,
	0x08, 0xc7, 0xa7, 0x92, 0x7c, 0x8f, 0xe5, 0x4c, 0x63, 0x45, 0x71, 0x20, 0x99, 0x4f, 0x6f, 0x11,
	0x36, 0x7f, 0x46, 0x10, 0x5e, 0x97, 0xfa, 0xe4, 0x16, 0xef, 0xfe, 0x34, 0x5e, 0x2c, 0x59, 0xdf,
	0x61, 0x81, 0x14, 0x32, 0xba, 0x49, 0x5a, 0xfe, 0x52, 0x02, 0x0b, 0xdb, 0x99, 0x8d, 0x98, 0xbe,
	0x89, 0xa9, 0x90, 0xd0, 0x05, 0xb3, 0x31, 0x6b, 0x5b, 0xc6, 0x92, 0xb1, 0x62, 0xba, 0xf7, 0xce,
	0x87, 0x8d, 0xc2, 0x68, 0xd8, 0x98, 0x6d, 0xee, 0xed, 0x5c, 0x0f, 0x1b, 0xff, 0x4e, 0xeb, 0x22,
	0xcf, 0x42, 0x2a, 0x50, 0x73, 0x6f, 0x07, 0x27, 0x64, 0xf8, 0x1c, 0x14, 0x7b, 0x2c, 0x68, 0x5b,
	0x33, 0x4b, 0xc6, 0x4a, 0x6d, 0x63, 0x13, 0xe5, 0x63, 0x1b, 0xd3, 0x50, 0xd8, 0xeb, 0x24, 0x09,
	0x81, 0x12, 0xef, 0xd0, 0x60, 0x1d, 0x3d, 0x89, 0x78, 0x1c, 0x3e, 0xa3, 0x51, 0xf2, 0x67, 0x9e,
	0xb2, 0xa0, 0xed, 0xce, 0xe9, 0xe6, 0xc5, 0x24, 0xc2, 0x4a, 0x11, 0x76, 0x41, 0x35, 0xa2, 0x82,
	0xc7, 0x91, 0x47, 0xad, 0x59, 0xa5, 0xfe, 0xe8, 0xd7, 0xd5, 0xb1, 0x56, 0x70, 0x17, 0x74, 0x87,
	0x6a, 0x96, 0xc1, 0x63, 0x75, 0xf8, 0x10, 0xd4, 0x44, 0xdc, 0xca, 0x0a, 0x56, 0x51, 0xf9, 0xf1,
	0xb7, 0x26, 0xd4, 0x8e, 0xf2, 0x12, 0x9e, 0xc4, 0xc1, 0x25, 0x50, 0x0c, 0x88, 0x4f, 0xad, 0x92,
	0xc2, 0x8f, 0x3f, 0x61, 0x9f, 0xf8, 0x14, 0xab, 0x0a, 0x74, 0x80, 0x99, 0xfc, 0x8a, 0x90, 0x78,
	0xd4, 0x2a, 0x2b, 0xd8, 0xa2, 0x86, 0x99, 0xfb, 0x59, 0x01, 0xe7, 0x18, 0xf8, 0x18, 0x98, 0x3c,
	0x4c, 0x06, 0xc7, 0x78, 0x60, 0x55, 0x14, 0xc1, 0xce, 0x08, 0x07, 0x59, 0xe1, 0x7a, 0x32, 0xc0,
	0x39, 0x01, 0x1e, 0x83, 0x6a, 0x2c, 0x68, 0xb4, 0x17, 0x9c, 0x70, 0xab, 0xaa, 0x1c, 0xfb, 0x0f,
	0x4d, 0x5e, 0xa3, 0x6f, 0x36, 0x3f, 0x71, 0xaa, 0xa9, 0xd1, 0xb9, 0x3b, 0x59, 0x06, 0x8f, 0x95,
	0x60, 0x13, 0x94, 0x79, 0xeb, 0x35, 0xf5, 0xa4, 0x65, 0x2a, 0xcd, 0xb5, 0xa9, 0x53, 0xd0, 0x8b,
	0x8b, 0x30, 0x39, 0xdd, 0x7d, 0x2b, 0x69, 0x90, 0x0c, 0xc0, 0xfd, 0x4b, 0x4b, 0x97, 0x0f, 0x94,
	0x08, 0xd6, 0x62, 0xf0, 0x15, 0x30, 0x79, 0xbf, 0x9d, 0x26, 0x2d, 0xf0, 0x3b, 0xca, 0x63, 0x2b,
	0x0f, 0x32, 0x1d, 0x9c, 0x4b, 0x2e, 0x7f, 0x98, 0x01, 0x8b, 0x13, 0x1b, 0x2f, 0x42, 0x1e, 0x08,
	0xfa, 0x47, 0x56, 0xfe, 0x7f, 0x50, 0x21, 0xfd, 0x3e, 0x3f, 0xa5, 0xe9, 0xd6, 0x57, 0xdd, 0x79,
	0xad, 0x53, 0xd9, 0x4e, 0xd3, 0x38, 0xab, 0xc3, 0x43, 0x50, 0x16, 0x92, 0xc8, 0x58, 0xe8, 0x0d,
	0x5e, 0xbd, 0xdb, 0x06, 0x1f, 0x29, 0x8e, 0x0b, 0x12, 0xdb, 0x30, 0x15, 0x71, 0x5f, 0x62, 0xad,
	0x03, 0x1b, 0xa0, 0x14, 0x12, 0xe9, 0x75, 0xd5, 0x96, 0xce, 0xb9, 0xe6, 0x68, 0xd8, 0x28, 0x1d,
	0x26, 0x09, 0x9c, 0xe6, 0xe1, 0x16, 0x30, 0xd5, 0xe1, 0xf8, 0x2c, 0xcc, 0x56, 0xb3, 0x9e, 0x98,
	0x74, 0x98, 0x25, 0xaf, 0x27, 0x03, 0x9c, 0x83, 0x97, 0x3f, 0x1a, 0x60, 0x7e, 0xc2, 0xb1, 0x01,
	0xa3, 0xa7, 0xb0, 0x09, 0x2a, 0x51, 0xfa, 0x5a, 0x28, 0xcf, 0x6a, 0x1b, 0x08, 0xfd, 0xe8, 0x61,
	0x46, 0x37, 0xdf, 0x18, 0xb7, 0x96, 0xf8, 0xa2, 0x03, 0x9c, 0x69, 0xc1, 0x17, 0xea, 0x6e, 0xab,
	0x91, 0xe8, 0x97, 0xc3, 0xb9, 0xb3, 0x6e, 0x4a, 0x73, 0xe7, 0xf4, 0x65, 0x56, 0x11, 0x1e, 0xcb,
	0xb9, 0x6b, 0xe7, 0x57, 0x76, 0xe1, 0xe2, 0xca, 0x2e, 0x5c, 0x5e, 0xd9, 0x85, 0x77, 0x23, 0xdb,
	0x38, 0x1f, 0xd9, 0xc6, 0xc5, 0xc8, 0x36, 0x2e, 0x47, 0xb6, 0xf1, 0x69, 0x64, 0x1b, 0xef, 0x3f,
	0xdb, 0x85, 0x97, 0x15, 0x2d, 0xfc, 0x35, 0x00, 0x00, 0xff, 0xff, 0x76, 0x21, 0xd5, 0x35, 0xaf,
	0x06, 0x00, 0x00,
}
