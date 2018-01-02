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
// source: k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1alpha1/generated.proto

	It has these top-level messages:
		PriorityClass
		PriorityClassList
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

func (m *PriorityClass) Reset()                    { *m = PriorityClass{} }
func (*PriorityClass) ProtoMessage()               {}
func (*PriorityClass) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *PriorityClassList) Reset()                    { *m = PriorityClassList{} }
func (*PriorityClassList) ProtoMessage()               {}
func (*PriorityClassList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func init() {
	proto.RegisterType((*PriorityClass)(nil), "k8s.io.api.scheduling.v1alpha1.PriorityClass")
	proto.RegisterType((*PriorityClassList)(nil), "k8s.io.api.scheduling.v1alpha1.PriorityClassList")
}
func (m *PriorityClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClass) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x10
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Value))
	dAtA[i] = 0x18
	i++
	if m.GlobalDefault {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Description)))
	i += copy(dAtA[i:], m.Description)
	return i, nil
}

func (m *PriorityClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClassList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n2, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
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
func (m *PriorityClass) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	n += 1 + sovGenerated(uint64(m.Value))
	n += 2
	l = len(m.Description)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *PriorityClassList) Size() (n int) {
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
func (this *PriorityClass) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PriorityClass{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`GlobalDefault:` + fmt.Sprintf("%v", this.GlobalDefault) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PriorityClassList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PriorityClassList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "PriorityClass", "PriorityClass", 1), `&`, ``, 1) + `,`,
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
			wire |= (uint64(b) & 0x7F) << shift
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
				m.Value |= (int32(b) & 0x7F) << shift
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
				v |= (int(b) & 0x7F) << shift
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
			m.Description = string(dAtA[iNdEx:postIndex])
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
			wire |= (uint64(b) & 0x7F) << shift
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
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x33, 0x5d, 0x0b, 0x75, 0x4a, 0x41, 0x23, 0x42, 0xe8, 0x61, 0x36, 0xac, 0x97, 0x5c,
	0x76, 0xc6, 0xae, 0x7f, 0x10, 0xbc, 0xc5, 0x85, 0x45, 0x50, 0x94, 0x1c, 0x3c, 0x88, 0x07, 0x27,
	0xc9, 0x6c, 0x3a, 0x36, 0xc9, 0x84, 0x99, 0x37, 0x81, 0xbd, 0x79, 0xf6, 0xe4, 0x97, 0x12, 0x7a,
	0xdc, 0xe3, 0x9e, 0x16, 0x1b, 0xbf, 0x88, 0x24, 0x4d, 0x37, 0xad, 0x65, 0xd5, 0x5b, 0xe6, 0x79,
	0x9f, 0xdf, 0x33, 0xf3, 0x3e, 0x04, 0x9f, 0x2d, 0x5e, 0x18, 0x2a, 0x15, 0x5b, 0x94, 0xa1, 0xd0,
	0xb9, 0x00, 0x61, 0x58, 0x25, 0xf2, 0x58, 0x69, 0xd6, 0x0d, 0x78, 0x21, 0x99, 0x89, 0xe6, 0x22,
	0x2e, 0x53, 0x99, 0x27, 0xac, 0x9a, 0xf1, 0xb4, 0x98, 0xf3, 0x19, 0x4b, 0x44, 0x2e, 0x34, 0x07,
	0x11, 0xd3, 0x42, 0x2b, 0x50, 0x36, 0x59, 0xfb, 0x29, 0x2f, 0x24, 0xed, 0xfd, 0x74, 0xe3, 0x9f,
	0x1e, 0x27, 0x12, 0xe6, 0x65, 0x48, 0x23, 0x95, 0xb1, 0x44, 0x25, 0x8a, 0xb5, 0x58, 0x58, 0x9e,
	0xb7, 0xa7, 0xf6, 0xd0, 0x7e, 0xad, 0xe3, 0xa6, 0x4f, 0xfb, 0xeb, 0x33, 0x1e, 0xcd, 0x65, 0x2e,
	0xf4, 0x05, 0x2b, 0x16, 0x49, 0x23, 0x18, 0x96, 0x09, 0xe0, 0xac, 0xda, 0x7b, 0xc4, 0x94, 0xdd,
	0x46, 0xe9, 0x32, 0x07, 0x99, 0x89, 0x3d, 0xe0, 0xf9, 0xbf, 0x80, 0x66, 0x95, 0x8c, 0xef, 0x71,
	0x4f, 0x6e, 0xe3, 0x4a, 0x90, 0x29, 0x93, 0x39, 0x18, 0xd0, 0x7f, 0x42, 0x47, 0xdf, 0x06, 0x78,
	0xf2, 0x5e, 0x4b, 0xa5, 0x25, 0x5c, 0xbc, 0x4a, 0xb9, 0x31, 0xf6, 0x67, 0x3c, 0x6a, 0x56, 0x89,
	0x39, 0x70, 0x07, 0xb9, 0xc8, 0x1b, 0x9f, 0x3c, 0xa6, 0x7d, 0x8f, 0x37, 0xc9, 0xb4, 0x58, 0x24,
	0x8d, 0x60, 0x68, 0xe3, 0xa6, 0xd5, 0x8c, 0xbe, 0x0b, 0xbf, 0x88, 0x08, 0xde, 0x0a, 0xe0, 0xbe,
	0xbd, 0xbc, 0x3e, 0xb4, 0xea, 0xeb, 0x43, 0xdc, 0x6b, 0xc1, 0x4d, 0xaa, 0xfd, 0x08, 0x0f, 0x2b,
	0x9e, 0x96, 0xc2, 0x19, 0xb8, 0xc8, 0x1b, 0xfa, 0x93, 0xce, 0x3c, 0xfc, 0xd0, 0x88, 0xc1, 0x7a,
	0x66, 0xbf, 0xc4, 0x93, 0x24, 0x55, 0x21, 0x4f, 0x4f, 0xc5, 0x39, 0x2f, 0x53, 0x70, 0x0e, 0x5c,
	0xe4, 0x8d, 0xfc, 0x87, 0x9d, 0x79, 0x72, 0xb6, 0x3d, 0x0c, 0x76, 0xbd, 0xf6, 0x33, 0x3c, 0x8e,
	0x85, 0x89, 0xb4, 0x2c, 0x40, 0xaa, 0xdc, 0xb9, 0xe3, 0x22, 0xef, 0xae, 0xff, 0xa0, 0x43, 0xc7,
	0xa7, 0xfd, 0x28, 0xd8, 0xf6, 0x1d, 0xfd, 0x40, 0xf8, 0xfe, 0x4e, 0x19, 0x6f, 0xa4, 0x01, 0xfb,
	0xd3, 0x5e, 0x21, 0xf4, 0xff, 0x0a, 0x69, 0xe8, 0xb6, 0x8e, 0x7b, 0xdd, 0xcd, 0xa3, 0x8d, 0xb2,
	0x55, 0x46, 0x80, 0x87, 0x12, 0x44, 0x66, 0x9c, 0x81, 0x7b, 0xe0, 0x8d, 0x4f, 0x8e, 0xe9, 0xdf,
	0xff, 0x59, 0xba, 0xf3, 0xbe, 0xbe, 0xbb, 0xd7, 0x4d, 0x46, 0xb0, 0x8e, 0xf2, 0xe9, 0x72, 0x45,
	0xac, 0xcb, 0x15, 0xb1, 0xae, 0x56, 0xc4, 0xfa, 0x5a, 0x13, 0xb4, 0xac, 0x09, 0xba, 0xac, 0x09,
	0xba, 0xaa, 0x09, 0xfa, 0x59, 0x13, 0xf4, 0xfd, 0x17, 0xb1, 0x3e, 0x8e, 0x36, 0x99, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x44, 0x05, 0xba, 0x7b, 0x71, 0x03, 0x00, 0x00,
}
