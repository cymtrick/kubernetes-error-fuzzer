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
// source: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto

	It has these top-level messages:
		PartialObjectMetadataList
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_apimachinery_pkg_apis_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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

func (m *PartialObjectMetadataList) Reset()      { *m = PartialObjectMetadataList{} }
func (*PartialObjectMetadataList) ProtoMessage() {}
func (*PartialObjectMetadataList) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{0}
}

func init() {
	proto.RegisterType((*PartialObjectMetadataList)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1beta1.PartialObjectMetadataList")
}
func (m *PartialObjectMetadataList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartialObjectMetadataList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n1, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
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
func (m *PartialObjectMetadataList) Size() (n int) {
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
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
func (this *PartialObjectMetadataList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PartialObjectMetadataList{`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "PartialObjectMetadata", "k8s_io_apimachinery_pkg_apis_meta_v1.PartialObjectMetadata", 1), `&`, ``, 1) + `,`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
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
func (m *PartialObjectMetadataList) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PartialObjectMetadataList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartialObjectMetadataList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
			m.Items = append(m.Items, k8s_io_apimachinery_pkg_apis_meta_v1.PartialObjectMetadata{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
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
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0x9b, 0xf7, 0x65, 0x38, 0x3a, 0x04, 0xd9, 0x69, 0xee, 0x90, 0x0d, 0x4f, 0xf3, 0xb0,
	0x84, 0x0d, 0x11, 0xc1, 0xdb, 0x6e, 0x82, 0xa2, 0xec, 0x28, 0x1e, 0x4c, 0xbb, 0xc7, 0x2e, 0xd6,
	0x34, 0x25, 0x79, 0x3a, 0xf0, 0xe6, 0x47, 0xf0, 0x63, 0xed, 0xb8, 0xe3, 0x40, 0x18, 0xae, 0x7e,
	0x11, 0x49, 0x57, 0x45, 0xa6, 0x62, 0x6f, 0x7d, 0xfe, 0xcd, 0xef, 0x97, 0x7f, 0x12, 0x7f, 0x1c,
	0x9f, 0x58, 0x26, 0x35, 0x8f, 0xb3, 0x00, 0x4c, 0x02, 0x08, 0x96, 0xcf, 0x20, 0x99, 0x68, 0xc3,
	0xcb, 0x1f, 0x22, 0x95, 0x4a, 0x84, 0x53, 0x99, 0x80, 0x79, 0xe4, 0x69, 0x1c, 0xb9, 0xc0, 0x72,
	0x05, 0x28, 0xf8, 0x6c, 0x10, 0x00, 0x8a, 0x01, 0x8f, 0x20, 0x01, 0x23, 0x10, 0x26, 0x2c, 0x35,
	0x1a, 0x75, 0xf3, 0x70, 0x83, 0xb2, 0xaf, 0x28, 0x4b, 0xe3, 0xc8, 0x05, 0x96, 0x39, 0x94, 0x95,
	0x68, 0xbb, 0x1f, 0x49, 0x9c, 0x66, 0x01, 0x0b, 0xb5, 0xe2, 0x91, 0x8e, 0x34, 0x2f, 0x0c, 0x41,
	0x76, 0x57, 0x4c, 0xc5, 0x50, 0x7c, 0x6d, 0xcc, 0xed, 0xa3, 0x2a, 0xa5, 0xb6, 0xfb, 0xb4, 0x7f,
	0x3d, 0x8a, 0xc9, 0x12, 0x94, 0x0a, 0xbe, 0x01, 0xc7, 0x7f, 0x01, 0x36, 0x9c, 0x82, 0x12, 0xdb,
	0xdc, 0xc1, 0x0b, 0xf1, 0xf7, 0xaf, 0x84, 0x41, 0x29, 0x1e, 0x2e, 0x83, 0x7b, 0x08, 0xf1, 0x02,
	0x50, 0x4c, 0x04, 0x8a, 0x73, 0x69, 0xb1, 0x79, 0xeb, 0xd7, 0x24, 0x82, 0xb2, 0x2d, 0xd2, 0xfd,
	0xdf, 0x6b, 0x0c, 0x4f, 0x59, 0x95, 0x6b, 0x62, 0x3f, 0xfa, 0x46, 0xbb, 0xf3, 0x55, 0xc7, 0xcb,
	0x57, 0x9d, 0xda, 0x99, 0x33, 0x8e, 0x37, 0xe2, 0xe6, 0x8d, 0x5f, 0x57, 0xe5, 0x8a, 0xd6, 0xbf,
	0x2e, 0xe9, 0x35, 0x86, 0xac, 0xda, 0x26, 0xae, 0x9f, 0x73, 0x8f, 0xf6, 0x4a, 0x6f, 0xfd, 0x23,
	0x19, 0x7f, 0x1a, 0x47, 0xfd, 0xf9, 0x9a, 0x7a, 0x8b, 0x35, 0xf5, 0x96, 0x6b, 0xea, 0x3d, 0xe5,
	0x94, 0xcc, 0x73, 0x4a, 0x16, 0x39, 0x25, 0xcb, 0x9c, 0x92, 0xd7, 0x9c, 0x92, 0xe7, 0x37, 0xea,
	0x5d, 0xef, 0x94, 0x4f, 0xfb, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x10, 0x2f, 0x48, 0xbd, 0x5a, 0x02,
	0x00, 0x00,
}
