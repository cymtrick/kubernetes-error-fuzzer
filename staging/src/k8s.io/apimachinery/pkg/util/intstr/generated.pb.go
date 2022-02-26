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
// source: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/util/intstr/generated.proto

package intstr

import (
	fmt "fmt"

	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
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

func (m *IntOrString) Reset()      { *m = IntOrString{} }
func (*IntOrString) ProtoMessage() {}
func (*IntOrString) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e046ae3ce6121c, []int{0}
}
func (m *IntOrString) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IntOrString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IntOrString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntOrString.Merge(m, src)
}
func (m *IntOrString) XXX_Size() int {
	return m.Size()
}
func (m *IntOrString) XXX_DiscardUnknown() {
	xxx_messageInfo_IntOrString.DiscardUnknown(m)
}

var xxx_messageInfo_IntOrString proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IntOrString)(nil), "k8s.io.apimachinery.pkg.util.intstr.IntOrString")
}

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/util/intstr/generated.proto", fileDescriptor_94e046ae3ce6121c)
}

var fileDescriptor_94e046ae3ce6121c = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4a, 0x03, 0x31,
	0x1c, 0xc6, 0x13, 0x5b, 0x8b, 0x9e, 0xe0, 0x50, 0x1c, 0x8a, 0x43, 0x7a, 0x58, 0x90, 0x5b, 0x4c,
	0x56, 0x71, 0xec, 0x56, 0x10, 0x84, 0x56, 0x1c, 0xdc, 0xee, 0xda, 0x98, 0x86, 0x6b, 0x93, 0x90,
	0xfb, 0x9f, 0x70, 0x5b, 0x1f, 0x41, 0x37, 0x47, 0x1f, 0xe7, 0xc6, 0x8e, 0x1d, 0xa4, 0x78, 0xf1,
	0x2d, 0x9c, 0xe4, 0x72, 0x07, 0x3a, 0x3a, 0x25, 0xdf, 0xf7, 0xfd, 0x7e, 0x19, 0x12, 0xdc, 0xa6,
	0xd7, 0x19, 0x95, 0x9a, 0xa5, 0x79, 0xc2, 0xad, 0xe2, 0xc0, 0x33, 0xf6, 0xcc, 0xd5, 0x42, 0x5b,
	0xd6, 0x0e, 0xb1, 0x91, 0xeb, 0x78, 0xbe, 0x94, 0x8a, 0xdb, 0x82, 0x99, 0x54, 0xb0, 0x1c, 0xe4,
	0x8a, 0x49, 0x05, 0x19, 0x58, 0x26, 0xb8, 0xe2, 0x36, 0x06, 0xbe, 0xa0, 0xc6, 0x6a, 0xd0, 0xfd,
	0x51, 0x23, 0xd1, 0xbf, 0x12, 0x35, 0xa9, 0xa0, 0xb5, 0x44, 0x1b, 0xe9, 0xfc, 0x4a, 0x48, 0x58,
	0xe6, 0x09, 0x9d, 0xeb, 0x35, 0x13, 0x5a, 0x68, 0xe6, 0xdd, 0x24, 0x7f, 0xf2, 0xc9, 0x07, 0x7f,
	0x6b, 0xde, 0xbc, 0x78, 0xc5, 0xc1, 0xc9, 0x44, 0xc1, 0x9d, 0x9d, 0x81, 0x95, 0x4a, 0xf4, 0xa3,
	0xa0, 0x0b, 0x85, 0xe1, 0x03, 0x1c, 0xe2, 0xa8, 0x33, 0x3e, 0x2b, 0xf7, 0x43, 0xe4, 0xf6, 0xc3,
	0xee, 0x7d, 0x61, 0xf8, 0x77, 0x7b, 0x4e, 0x3d, 0xd1, 0xbf, 0x0c, 0x7a, 0x52, 0xc1, 0x43, 0xbc,
	0x1a, 0x1c, 0x84, 0x38, 0x3a, 0x1c, 0x9f, 0xb6, 0x6c, 0x6f, 0xe2, 0xdb, 0x69, 0xbb, 0xd6, 0x5c,
	0x06, 0xb6, 0xe6, 0x3a, 0x21, 0x8e, 0x8e, 0x7f, 0xb9, 0x99, 0x6f, 0xa7, 0xed, 0x7a, 0x73, 0xf4,
	0xf6, 0x3e, 0x44, 0x9b, 0x8f, 0x10, 0x8d, 0x27, 0x65, 0x45, 0xd0, 0xb6, 0x22, 0x68, 0x57, 0x11,
	0xb4, 0x71, 0x04, 0x97, 0x8e, 0xe0, 0xad, 0x23, 0x78, 0xe7, 0x08, 0xfe, 0x74, 0x04, 0xbf, 0x7c,
	0x11, 0xf4, 0x38, 0xfa, 0xc7, 0x17, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xc4, 0xf0, 0xa0,
	0x81, 0x01, 0x00, 0x00,
}

func (m *IntOrString) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IntOrString) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IntOrString) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.StrVal)
	copy(dAtA[i:], m.StrVal)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.StrVal)))
	i--
	dAtA[i] = 0x1a
	i = encodeVarintGenerated(dAtA, i, uint64(m.IntVal))
	i--
	dAtA[i] = 0x10
	i = encodeVarintGenerated(dAtA, i, uint64(m.Type))
	i--
	dAtA[i] = 0x8
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
func (m *IntOrString) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Type))
	n += 1 + sovGenerated(uint64(m.IntVal))
	l = len(m.StrVal)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IntOrString) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IntOrString: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IntOrString: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntVal", wireType)
			}
			m.IntVal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IntVal |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrVal", wireType)
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
			m.StrVal = string(dAtA[iNdEx:postIndex])
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
