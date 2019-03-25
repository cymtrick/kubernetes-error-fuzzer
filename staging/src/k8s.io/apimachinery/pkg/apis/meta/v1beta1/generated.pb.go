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
		PartialObjectMetadata
		PartialObjectMetadataList
		TableOptions
*/
package v1beta1

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

func (m *PartialObjectMetadata) Reset()                    { *m = PartialObjectMetadata{} }
func (*PartialObjectMetadata) ProtoMessage()               {}
func (*PartialObjectMetadata) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *PartialObjectMetadataList) Reset()      { *m = PartialObjectMetadataList{} }
func (*PartialObjectMetadataList) ProtoMessage() {}
func (*PartialObjectMetadataList) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{1}
}

func (m *TableOptions) Reset()                    { *m = TableOptions{} }
func (*TableOptions) ProtoMessage()               {}
func (*TableOptions) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*PartialObjectMetadata)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1beta1.PartialObjectMetadata")
	proto.RegisterType((*PartialObjectMetadataList)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1beta1.PartialObjectMetadataList")
	proto.RegisterType((*TableOptions)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1beta1.TableOptions")
}
func (m *PartialObjectMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartialObjectMetadata) MarshalTo(dAtA []byte) (int, error) {
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
	return i, nil
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
	n2, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *TableOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.IncludeObject)))
	i += copy(dAtA[i:], m.IncludeObject)
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
func (m *PartialObjectMetadata) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
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

func (m *TableOptions) Size() (n int) {
	var l int
	_ = l
	l = len(m.IncludeObject)
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
func (this *PartialObjectMetadata) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PartialObjectMetadata{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PartialObjectMetadataList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PartialObjectMetadataList{`,
		`Items:` + strings.Replace(fmt.Sprintf("%v", this.Items), "PartialObjectMetadata", "PartialObjectMetadata", 1) + `,`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TableOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TableOptions{`,
		`IncludeObject:` + fmt.Sprintf("%v", this.IncludeObject) + `,`,
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
func (m *PartialObjectMetadata) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PartialObjectMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartialObjectMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
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
			m.Items = append(m.Items, &PartialObjectMetadata{})
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
func (m *TableOptions) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TableOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncludeObject", wireType)
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
			m.IncludeObject = IncludeObjectPolicy(dAtA[iNdEx:postIndex])
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
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x33, 0xca, 0xe2, 0xee, 0xac, 0x0b, 0x12, 0x11, 0xd6, 0x1e, 0x26, 0xcb, 0x9e, 0x2a,
	0xd8, 0x19, 0x5b, 0x44, 0x3c, 0x4a, 0x6e, 0x05, 0xa5, 0x25, 0x78, 0x12, 0x0f, 0x4e, 0x92, 0xd7,
	0x74, 0xcc, 0x9f, 0x09, 0x99, 0x49, 0xa1, 0x37, 0x3f, 0x82, 0x1f, 0xab, 0xc7, 0x1e, 0x7b, 0x90,
	0x60, 0xe3, 0xb7, 0xf0, 0x24, 0xf9, 0xa3, 0x4d, 0x6b, 0x65, 0x73, 0x9b, 0xf7, 0x79, 0x79, 0x7e,
	0x79, 0x9e, 0x37, 0xd8, 0x09, 0x5f, 0x2b, 0x2a, 0x24, 0x0b, 0x73, 0x17, 0xb2, 0x04, 0x34, 0x28,
	0xb6, 0x84, 0xc4, 0x97, 0x19, 0x6b, 0x17, 0x3c, 0x15, 0x31, 0xf7, 0x16, 0x22, 0x81, 0x6c, 0xc5,
	0xd2, 0x30, 0xa8, 0x04, 0xc5, 0x62, 0xd0, 0x9c, 0x2d, 0xc7, 0x2e, 0x68, 0x3e, 0x66, 0x01, 0x24,
	0x90, 0x71, 0x0d, 0x3e, 0x4d, 0x33, 0xa9, 0xa5, 0xf9, 0xac, 0xb1, 0xd2, 0xae, 0x95, 0xa6, 0x61,
	0x50, 0x09, 0x8a, 0x56, 0x56, 0xda, 0x5a, 0x07, 0xa3, 0x40, 0xe8, 0x45, 0xee, 0x52, 0x4f, 0xc6,
	0x2c, 0x90, 0x81, 0x64, 0x35, 0xc1, 0xcd, 0x3f, 0xd7, 0x53, 0x3d, 0xd4, 0xaf, 0x86, 0x3c, 0x78,
	0xd9, 0x27, 0xd4, 0x71, 0x9e, 0xc1, 0x7f, 0xab, 0x64, 0x79, 0xa2, 0x45, 0x0c, 0xff, 0x18, 0x5e,
	0xdd, 0x65, 0x50, 0xde, 0x02, 0x62, 0x7e, 0xec, 0xbb, 0x5d, 0xe1, 0x27, 0x73, 0x9e, 0x69, 0xc1,
	0xa3, 0x99, 0xfb, 0x05, 0x3c, 0xfd, 0x0e, 0x34, 0xf7, 0xb9, 0xe6, 0xe6, 0x27, 0x7c, 0x1e, 0xb7,
	0xef, 0x6b, 0x74, 0x83, 0x86, 0x97, 0x93, 0x17, 0xb4, 0xcf, 0x91, 0xe8, 0x9e, 0x63, 0x9b, 0xeb,
	0xc2, 0x32, 0xca, 0xc2, 0xc2, 0x7b, 0xcd, 0xf9, 0x4b, 0xbd, 0xfd, 0x8e, 0xf0, 0xd3, 0x93, 0xdf,
	0x7e, 0x2b, 0x94, 0x36, 0x39, 0x3e, 0x13, 0x1a, 0x62, 0x75, 0x8d, 0x6e, 0xee, 0x0f, 0x2f, 0x27,
	0x6f, 0x68, 0xef, 0x3f, 0x44, 0x4f, 0x42, 0xed, 0x8b, 0xb2, 0xb0, 0xce, 0xa6, 0x15, 0xd2, 0x69,
	0xc8, 0xe6, 0xc7, 0x4e, 0xc5, 0x7b, 0x75, 0x45, 0xda, 0xaf, 0x62, 0x15, 0xb0, 0x2e, 0xf8, 0xa8,
	0x2d, 0x78, 0xfe, 0x47, 0xe9, 0xd4, 0x73, 0xf1, 0xc3, 0xf7, 0xdc, 0x8d, 0x60, 0x96, 0x6a, 0x21,
	0x13, 0x65, 0x3a, 0xf8, 0x4a, 0x24, 0x5e, 0x94, 0xfb, 0xd0, 0x04, 0xab, 0xaf, 0x7a, 0x61, 0x3f,
	0x6f, 0x11, 0x57, 0xd3, 0xee, 0xf2, 0x57, 0x61, 0x3d, 0x3e, 0x10, 0xe6, 0x32, 0x12, 0xde, 0xca,
	0x39, 0x44, 0xd8, 0xa3, 0xf5, 0x8e, 0x18, 0x9b, 0x1d, 0x31, 0xb6, 0x3b, 0x62, 0x7c, 0x2d, 0x09,
	0x5a, 0x97, 0x04, 0x6d, 0x4a, 0x82, 0xb6, 0x25, 0x41, 0x3f, 0x4a, 0x82, 0xbe, 0xfd, 0x24, 0xc6,
	0x87, 0x07, 0xed, 0x61, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x78, 0xd8, 0x63, 0x3a, 0x03,
	0x00, 0x00,
}
