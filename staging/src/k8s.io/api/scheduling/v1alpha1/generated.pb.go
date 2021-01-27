/*
Copyright The Kubernetes Authors.

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

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1alpha1/generated.proto

package v1alpha1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"

	k8s_io_api_core_v1 "k8s.io/api/core/v1"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *PriorityClass) Reset()      { *m = PriorityClass{} }
func (*PriorityClass) ProtoMessage() {}
func (*PriorityClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_f033641dd0b95dce, []int{0}
}
func (m *PriorityClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriorityClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *PriorityClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriorityClass.Merge(m, src)
}
func (m *PriorityClass) XXX_Size() int {
	return m.Size()
}
func (m *PriorityClass) XXX_DiscardUnknown() {
	xxx_messageInfo_PriorityClass.DiscardUnknown(m)
}

var xxx_messageInfo_PriorityClass proto.InternalMessageInfo

func (m *PriorityClassList) Reset()      { *m = PriorityClassList{} }
func (*PriorityClassList) ProtoMessage() {}
func (*PriorityClassList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f033641dd0b95dce, []int{1}
}
func (m *PriorityClassList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriorityClassList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *PriorityClassList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriorityClassList.Merge(m, src)
}
func (m *PriorityClassList) XXX_Size() int {
	return m.Size()
}
func (m *PriorityClassList) XXX_DiscardUnknown() {
	xxx_messageInfo_PriorityClassList.DiscardUnknown(m)
}

var xxx_messageInfo_PriorityClassList proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PriorityClass)(nil), "k8s.io.api.scheduling.v1alpha1.PriorityClass")
	proto.RegisterType((*PriorityClassList)(nil), "k8s.io.api.scheduling.v1alpha1.PriorityClassList")
}

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1alpha1/generated.proto", fileDescriptor_f033641dd0b95dce)
}

var fileDescriptor_f033641dd0b95dce = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x3b, 0x5d, 0x0b, 0x75, 0x4a, 0xa1, 0x46, 0x84, 0xd0, 0xc3, 0xb4, 0x74, 0x2f, 0xbd,
	0xec, 0x8c, 0x5d, 0x54, 0x84, 0xbd, 0xd5, 0x85, 0x45, 0x50, 0x2c, 0x39, 0x78, 0x10, 0x0f, 0x4e,
	0xd3, 0x77, 0xd3, 0xb1, 0x49, 0x26, 0xcc, 0x4c, 0x02, 0xbd, 0xf9, 0x11, 0xfc, 0x52, 0x42, 0x8f,
	0x7b, 0xdc, 0x53, 0xb1, 0xf1, 0x23, 0x78, 0xf3, 0x24, 0x49, 0xd3, 0x4d, 0xdb, 0xf8, 0x67, 0x6f,
	0x99, 0xf7, 0xf9, 0x3d, 0xcf, 0xcc, 0x3c, 0x49, 0xf0, 0xd5, 0xe2, 0xa5, 0xa6, 0x42, 0xb2, 0x45,
	0x3c, 0x05, 0x15, 0x82, 0x01, 0xcd, 0x12, 0x08, 0x67, 0x52, 0xb1, 0x42, 0xe0, 0x91, 0x60, 0xda,
	0x9d, 0xc3, 0x2c, 0xf6, 0x45, 0xe8, 0xb1, 0x64, 0xc4, 0xfd, 0x68, 0xce, 0x47, 0xcc, 0x83, 0x10,
	0x14, 0x37, 0x30, 0xa3, 0x91, 0x92, 0x46, 0x5a, 0x64, 0xcb, 0x53, 0x1e, 0x09, 0x5a, 0xf2, 0x74,
	0xc7, 0x77, 0xcf, 0x3c, 0x61, 0xe6, 0xf1, 0x94, 0xba, 0x32, 0x60, 0x9e, 0xf4, 0x24, 0xcb, 0x6d,
	0xd3, 0xf8, 0x3a, 0x5f, 0xe5, 0x8b, 0xfc, 0x69, 0x1b, 0xd7, 0x1d, 0xec, 0x6d, 0xef, 0x4a, 0x05,
	0x2c, 0xa9, 0x6c, 0xd9, 0x7d, 0x56, 0x32, 0x01, 0x77, 0xe7, 0x22, 0x04, 0xb5, 0x64, 0xd1, 0xc2,
	0xcb, 0x06, 0x9a, 0x05, 0x60, 0xf8, 0x9f, 0x5c, 0xec, 0x6f, 0x2e, 0x15, 0x87, 0x46, 0x04, 0x50,
	0x31, 0xbc, 0xf8, 0x9f, 0x21, 0xbb, 0x6e, 0xc0, 0x8f, 0x7d, 0x83, 0x9f, 0x75, 0xdc, 0x9e, 0x28,
	0x21, 0x95, 0x30, 0xcb, 0x57, 0x3e, 0xd7, 0xda, 0xfa, 0x84, 0x9b, 0xd9, 0xa9, 0x66, 0xdc, 0x70,
	0x1b, 0xf5, 0xd1, 0xb0, 0x75, 0xfe, 0x94, 0x96, 0xb5, 0xdd, 0x85, 0xd3, 0x68, 0xe1, 0x65, 0x03,
	0x4d, 0x33, 0x9a, 0x26, 0x23, 0xfa, 0x6e, 0xfa, 0x19, 0x5c, 0xf3, 0x16, 0x0c, 0x1f, 0x5b, 0xab,
	0x75, 0xaf, 0x96, 0xae, 0x7b, 0xb8, 0x9c, 0x39, 0x77, 0xa9, 0xd6, 0x29, 0x6e, 0x24, 0xdc, 0x8f,
	0xc1, 0xae, 0xf7, 0xd1, 0xb0, 0x31, 0x6e, 0x17, 0x70, 0xe3, 0x7d, 0x36, 0x74, 0xb6, 0x9a, 0x75,
	0x81, 0xdb, 0x9e, 0x2f, 0xa7, 0xdc, 0xbf, 0x84, 0x6b, 0x1e, 0xfb, 0xc6, 0x3e, 0xe9, 0xa3, 0x61,
	0x73, 0xfc, 0xa4, 0x80, 0xdb, 0x57, 0xfb, 0xa2, 0x73, 0xc8, 0x5a, 0xcf, 0x71, 0x6b, 0x06, 0xda,
	0x55, 0x22, 0x32, 0x42, 0x86, 0xf6, 0x83, 0x3e, 0x1a, 0x3e, 0x1c, 0x3f, 0x2e, 0xac, 0xad, 0xcb,
	0x52, 0x72, 0xf6, 0x39, 0xcb, 0xc3, 0x9d, 0x48, 0x01, 0x04, 0xf9, 0x6a, 0x22, 0x7d, 0xe1, 0x2e,
	0xed, 0x46, 0xee, 0xbd, 0x48, 0xd7, 0xbd, 0xce, 0xe4, 0x48, 0xfb, 0xb5, 0xee, 0x9d, 0x56, 0xbf,
	0x00, 0x7a, 0x8c, 0x39, 0x95, 0xd0, 0xc1, 0x37, 0x84, 0x1f, 0x1d, 0xb4, 0xfe, 0x46, 0x68, 0x63,
	0x7d, 0xac, 0x34, 0x4f, 0xef, 0xd7, 0x7c, 0xe6, 0xce, 0x7b, 0xef, 0x14, 0x57, 0x6c, 0xee, 0x26,
	0x7b, 0xad, 0x3b, 0xb8, 0x21, 0x0c, 0x04, 0xda, 0xae, 0xf7, 0x4f, 0x86, 0xad, 0xf3, 0x33, 0xfa,
	0xef, 0x7f, 0x81, 0x1e, 0x9c, 0xaf, 0x7c, 0x49, 0xaf, 0xb3, 0x0c, 0x67, 0x1b, 0x35, 0xa6, 0xab,
	0x0d, 0xa9, 0xdd, 0x6c, 0x48, 0xed, 0x76, 0x43, 0x6a, 0x5f, 0x52, 0x82, 0x56, 0x29, 0x41, 0x37,
	0x29, 0x41, 0xb7, 0x29, 0x41, 0xdf, 0x53, 0x82, 0xbe, 0xfe, 0x20, 0xb5, 0x0f, 0xcd, 0x5d, 0xe6,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0x5c, 0x1a, 0x39, 0xc9, 0x03, 0x00, 0x00,
}

func (m *PriorityClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriorityClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PreemptionPolicy != nil {
		i -= len(*m.PreemptionPolicy)
		copy(dAtA[i:], *m.PreemptionPolicy)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.PreemptionPolicy)))
		i--
		dAtA[i] = 0x2a
	}
	i -= len(m.Description)
	copy(dAtA[i:], m.Description)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Description)))
	i--
	dAtA[i] = 0x22
	i--
	if m.GlobalDefault {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x18
	i = encodeVarintGenerated(dAtA, i, uint64(m.Value))
	i--
	dAtA[i] = 0x10
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PriorityClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClassList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriorityClassList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PriorityClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	n += 1 + sovGenerated(uint64(m.Value))
	n += 2
	l = len(m.Description)
	n += 1 + l + sovGenerated(uint64(l))
	if m.PreemptionPolicy != nil {
		l = len(*m.PreemptionPolicy)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *PriorityClassList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PriorityClass) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PriorityClass{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`GlobalDefault:` + fmt.Sprintf("%v", this.GlobalDefault) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`PreemptionPolicy:` + valueToStringGenerated(this.PreemptionPolicy) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PriorityClassList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]PriorityClass{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "PriorityClass", "PriorityClass", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&PriorityClassList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
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
func (m *PriorityClass) Unmarshal(dAtA []byte) error {
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
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PriorityClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriorityClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
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
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalDefault", wireType)
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
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.GlobalDefault = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreemptionPolicy", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := k8s_io_api_core_v1.PreemptionPolicy(dAtA[iNdEx:postIndex])
			m.PreemptionPolicy = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *PriorityClassList) Unmarshal(dAtA []byte) error {
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
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PriorityClassList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriorityClassList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
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
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
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
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, PriorityClass{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
	depth := 0
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
		case 1:
			iNdEx += 8
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
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
