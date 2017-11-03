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
// source: k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2/v1/generated.proto
// DO NOT EDIT!

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2/v1/generated.proto

	It has these top-level messages:
		ReplicaSet
		ReplicaSetSpec
		ReplicaSetStatus
*/
package v1

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

func (m *ReplicaSet) Reset()                    { *m = ReplicaSet{} }
func (*ReplicaSet) ProtoMessage()               {}
func (*ReplicaSet) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *ReplicaSetSpec) Reset()                    { *m = ReplicaSetSpec{} }
func (*ReplicaSetSpec) ProtoMessage()               {}
func (*ReplicaSetSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *ReplicaSetStatus) Reset()                    { *m = ReplicaSetStatus{} }
func (*ReplicaSetStatus) ProtoMessage()               {}
func (*ReplicaSetStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*ReplicaSet)(nil), "k8s.io.apiserver.pkg.apis.example2.v1.ReplicaSet")
	proto.RegisterType((*ReplicaSetSpec)(nil), "k8s.io.apiserver.pkg.apis.example2.v1.ReplicaSetSpec")
	proto.RegisterType((*ReplicaSetStatus)(nil), "k8s.io.apiserver.pkg.apis.example2.v1.ReplicaSetStatus")
}
func (m *ReplicaSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaSet) MarshalTo(dAtA []byte) (int, error) {
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
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *ReplicaSetSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaSetSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Replicas != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.Replicas))
	}
	return i, nil
}

func (m *ReplicaSetStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaSetStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Replicas))
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
func (m *ReplicaSet) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ReplicaSetSpec) Size() (n int) {
	var l int
	_ = l
	if m.Replicas != nil {
		n += 1 + sovGenerated(uint64(*m.Replicas))
	}
	return n
}

func (m *ReplicaSetStatus) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Replicas))
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
func (this *ReplicaSet) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplicaSet{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "ReplicaSetSpec", "ReplicaSetSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "ReplicaSetStatus", "ReplicaSetStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReplicaSetSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplicaSetSpec{`,
		`Replicas:` + valueToStringGenerated(this.Replicas) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReplicaSetStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplicaSetStatus{`,
		`Replicas:` + fmt.Sprintf("%v", this.Replicas) + `,`,
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
func (m *ReplicaSet) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ReplicaSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaSet: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
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
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ReplicaSetSpec) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ReplicaSetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaSetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Replicas = &v
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
func (m *ReplicaSetStatus) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ReplicaSetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaSetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			m.Replicas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Replicas |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2/v1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0x87, 0x93, 0xb2, 0xac, 0x2a, 0xb3, 0x5a, 0xad, 0x72, 0xaa, 0x7a, 0x70, 0x51, 0x24, 0xa4,
	0x1e, 0xc0, 0x26, 0xcb, 0x5f, 0x71, 0x42, 0xb9, 0x03, 0x52, 0xf6, 0x80, 0xc4, 0x05, 0x1c, 0x77,
	0x48, 0x4d, 0x9a, 0xd8, 0xb2, 0x9d, 0x08, 0x6e, 0x3c, 0x02, 0x8f, 0xc1, 0xa3, 0xf4, 0xb8, 0xc7,
	0x3d, 0x55, 0x34, 0xbc, 0x08, 0xaa, 0x13, 0x12, 0xb1, 0xed, 0x0a, 0xb8, 0xe5, 0x67, 0xcf, 0xf7,
	0xcd, 0x64, 0x8c, 0x5e, 0xe7, 0xcf, 0x0d, 0x11, 0x92, 0xe6, 0x55, 0x0a, 0xba, 0x04, 0x0b, 0x86,
	0xd6, 0x50, 0x2e, 0xa4, 0xa6, 0xdd, 0x05, 0x53, 0xc2, 0x80, 0xae, 0x41, 0x53, 0x95, 0x67, 0x2e,
	0x51, 0xf8, 0xcc, 0x0a, 0xb5, 0x82, 0x73, 0x5a, 0x47, 0x34, 0x83, 0x12, 0x34, 0xb3, 0xb0, 0x20,
	0x4a, 0x4b, 0x2b, 0x83, 0x7b, 0x2d, 0x46, 0x7a, 0x8c, 0xa8, 0x3c, 0x73, 0x89, 0xfc, 0xc6, 0x48,
	0x1d, 0x4d, 0x1f, 0x64, 0xc2, 0x2e, 0xab, 0x94, 0x70, 0x59, 0xd0, 0x4c, 0x66, 0x92, 0x3a, 0x3a,
	0xad, 0x3e, 0xba, 0xe4, 0x82, 0xfb, 0x6a, 0xad, 0xd3, 0x70, 0x18, 0x86, 0x72, 0xa9, 0xe1, 0x40,
	0xe7, 0xe9, 0xe3, 0xa1, 0xa6, 0x60, 0x7c, 0x29, 0x4a, 0xd0, 0x5f, 0x86, 0x99, 0x0b, 0xb0, 0xec,
	0x10, 0x45, 0x6f, 0xa2, 0x74, 0x55, 0x5a, 0x51, 0xc0, 0x1e, 0xf0, 0xf4, 0x6f, 0x80, 0xe1, 0x4b,
	0x28, 0xd8, 0x1e, 0xf7, 0xe8, 0x26, 0xae, 0xb2, 0x62, 0x45, 0x45, 0x69, 0x8d, 0xd5, 0xd7, 0xa1,
	0xf0, 0xfb, 0x08, 0xa1, 0x04, 0xd4, 0x4a, 0x70, 0x76, 0x01, 0x36, 0xf8, 0x80, 0xc6, 0xbb, 0xff,
	0x58, 0x30, 0xcb, 0x26, 0xfe, 0x5d, 0x7f, 0x7e, 0xe7, 0xfc, 0x21, 0x19, 0xf6, 0xdd, 0x6b, 0x87,
	0x95, 0xef, 0xaa, 0x49, 0x1d, 0x91, 0x37, 0xe9, 0x27, 0xe0, 0xf6, 0x15, 0x58, 0x16, 0x07, 0xeb,
	0xcd, 0xcc, 0x6b, 0x36, 0x33, 0x34, 0x9c, 0x25, 0xbd, 0x35, 0x78, 0x8b, 0x8e, 0x8c, 0x02, 0x3e,
	0x19, 0x39, 0xfb, 0x13, 0xf2, 0x4f, 0xaf, 0x49, 0x86, 0x11, 0x2f, 0x14, 0xf0, 0xf8, 0xa4, 0x6b,
	0x71, 0xb4, 0x4b, 0x89, 0x13, 0x06, 0xef, 0xd1, 0xb1, 0xb1, 0xcc, 0x56, 0x66, 0x72, 0xcb, 0xa9,
	0x9f, 0xfd, 0xbf, 0xda, 0xe1, 0xf1, 0x69, 0x27, 0x3f, 0x6e, 0x73, 0xd2, 0x69, 0xc3, 0x17, 0xe8,
	0xf4, 0xcf, 0x31, 0x82, 0x39, 0x1a, 0xeb, 0xf6, 0xc4, 0xb8, 0x6d, 0xdd, 0x8e, 0x4f, 0x9a, 0xcd,
	0x6c, 0xdc, 0x55, 0x99, 0xa4, 0xbf, 0x0d, 0x5f, 0xa2, 0xb3, 0xeb, 0x7d, 0x82, 0xfb, 0x7b, 0xf4,
	0x59, 0xd7, 0xf9, 0x80, 0x21, 0x9e, 0xaf, 0xb7, 0xd8, 0xbb, 0xdc, 0x62, 0xef, 0x6a, 0x8b, 0xbd,
	0xaf, 0x0d, 0xf6, 0xd7, 0x0d, 0xf6, 0x2f, 0x1b, 0xec, 0x5f, 0x35, 0xd8, 0xff, 0xd1, 0x60, 0xff,
	0xdb, 0x4f, 0xec, 0xbd, 0x1b, 0xd5, 0xd1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x4a, 0x84,
	0xe4, 0x71, 0x03, 0x00, 0x00,
}
