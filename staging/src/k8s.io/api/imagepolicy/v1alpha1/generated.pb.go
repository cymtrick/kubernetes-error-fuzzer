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

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/vendor/k8s.io/api/imagepolicy/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/api/imagepolicy/v1alpha1/generated.proto

	It has these top-level messages:
		ImageReview
		ImageReviewContainerSpec
		ImageReviewSpec
		ImageReviewStatus
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

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

func (m *ImageReview) Reset()                    { *m = ImageReview{} }
func (*ImageReview) ProtoMessage()               {}
func (*ImageReview) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *ImageReviewContainerSpec) Reset()      { *m = ImageReviewContainerSpec{} }
func (*ImageReviewContainerSpec) ProtoMessage() {}
func (*ImageReviewContainerSpec) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{1}
}

func (m *ImageReviewSpec) Reset()                    { *m = ImageReviewSpec{} }
func (*ImageReviewSpec) ProtoMessage()               {}
func (*ImageReviewSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *ImageReviewStatus) Reset()                    { *m = ImageReviewStatus{} }
func (*ImageReviewStatus) ProtoMessage()               {}
func (*ImageReviewStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func init() {
	proto.RegisterType((*ImageReview)(nil), "k8s.io.api.imagepolicy.v1alpha1.ImageReview")
	proto.RegisterType((*ImageReviewContainerSpec)(nil), "k8s.io.api.imagepolicy.v1alpha1.ImageReviewContainerSpec")
	proto.RegisterType((*ImageReviewSpec)(nil), "k8s.io.api.imagepolicy.v1alpha1.ImageReviewSpec")
	proto.RegisterType((*ImageReviewStatus)(nil), "k8s.io.api.imagepolicy.v1alpha1.ImageReviewStatus")
}
func (m *ImageReview) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImageReview) MarshalTo(dAtA []byte) (int, error) {
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

func (m *ImageReviewContainerSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImageReviewContainerSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Image)))
	i += copy(dAtA[i:], m.Image)
	return i, nil
}

func (m *ImageReviewSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImageReviewSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Containers) > 0 {
		for _, msg := range m.Containers {
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
	if len(m.Annotations) > 0 {
		keysForAnnotations := make([]string, 0, len(m.Annotations))
		for k := range m.Annotations {
			keysForAnnotations = append(keysForAnnotations, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForAnnotations)
		for _, k := range keysForAnnotations {
			dAtA[i] = 0x12
			i++
			v := m.Annotations[string(k)]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Namespace)))
	i += copy(dAtA[i:], m.Namespace)
	return i, nil
}

func (m *ImageReviewStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImageReviewStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	if m.Allowed {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Reason)))
	i += copy(dAtA[i:], m.Reason)
	if len(m.AuditAnnotations) > 0 {
		keysForAuditAnnotations := make([]string, 0, len(m.AuditAnnotations))
		for k := range m.AuditAnnotations {
			keysForAuditAnnotations = append(keysForAuditAnnotations, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForAuditAnnotations)
		for _, k := range keysForAuditAnnotations {
			dAtA[i] = 0x1a
			i++
			v := m.AuditAnnotations[string(k)]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
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
func (m *ImageReview) Size() (n int) {
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

func (m *ImageReviewContainerSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.Image)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ImageReviewSpec) Size() (n int) {
	var l int
	_ = l
	if len(m.Containers) > 0 {
		for _, e := range m.Containers {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Annotations) > 0 {
		for k, v := range m.Annotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	l = len(m.Namespace)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ImageReviewStatus) Size() (n int) {
	var l int
	_ = l
	n += 2
	l = len(m.Reason)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.AuditAnnotations) > 0 {
		for k, v := range m.AuditAnnotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
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
func (this *ImageReview) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ImageReview{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "ImageReviewSpec", "ImageReviewSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "ImageReviewStatus", "ImageReviewStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ImageReviewContainerSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ImageReviewContainerSpec{`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ImageReviewSpec) String() string {
	if this == nil {
		return "nil"
	}
	keysForAnnotations := make([]string, 0, len(this.Annotations))
	for k := range this.Annotations {
		keysForAnnotations = append(keysForAnnotations, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForAnnotations)
	mapStringForAnnotations := "map[string]string{"
	for _, k := range keysForAnnotations {
		mapStringForAnnotations += fmt.Sprintf("%v: %v,", k, this.Annotations[k])
	}
	mapStringForAnnotations += "}"
	s := strings.Join([]string{`&ImageReviewSpec{`,
		`Containers:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Containers), "ImageReviewContainerSpec", "ImageReviewContainerSpec", 1), `&`, ``, 1) + `,`,
		`Annotations:` + mapStringForAnnotations + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ImageReviewStatus) String() string {
	if this == nil {
		return "nil"
	}
	keysForAuditAnnotations := make([]string, 0, len(this.AuditAnnotations))
	for k := range this.AuditAnnotations {
		keysForAuditAnnotations = append(keysForAuditAnnotations, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForAuditAnnotations)
	mapStringForAuditAnnotations := "map[string]string{"
	for _, k := range keysForAuditAnnotations {
		mapStringForAuditAnnotations += fmt.Sprintf("%v: %v,", k, this.AuditAnnotations[k])
	}
	mapStringForAuditAnnotations += "}"
	s := strings.Join([]string{`&ImageReviewStatus{`,
		`Allowed:` + fmt.Sprintf("%v", this.Allowed) + `,`,
		`Reason:` + fmt.Sprintf("%v", this.Reason) + `,`,
		`AuditAnnotations:` + mapStringForAuditAnnotations + `,`,
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
func (m *ImageReview) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ImageReview: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReview: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *ImageReviewContainerSpec) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ImageReviewContainerSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReviewContainerSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
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
			m.Image = string(dAtA[iNdEx:postIndex])
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
func (m *ImageReviewSpec) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ImageReviewSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReviewSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Containers", wireType)
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
			m.Containers = append(m.Containers, ImageReviewContainerSpec{})
			if err := m.Containers[len(m.Containers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
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
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Annotations == nil {
				m.Annotations = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthGenerated
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Annotations[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Annotations[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 3:
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
func (m *ImageReviewStatus) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ImageReviewStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReviewStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
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
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuditAnnotations", wireType)
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
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.AuditAnnotations == nil {
				m.AuditAnnotations = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthGenerated
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.AuditAnnotations[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.AuditAnnotations[mapkey] = mapvalue
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
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/imagepolicy/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0xe3, 0xa4, 0xff, 0x32, 0xf9, 0x3e, 0x9a, 0x0e, 0x20, 0x59, 0x59, 0xb8, 0x55, 0x90,
	0x50, 0x59, 0x30, 0x43, 0x2b, 0x84, 0x0a, 0x0b, 0x50, 0x5c, 0x21, 0x95, 0x05, 0x20, 0x0d, 0xbb,
	0xae, 0x98, 0x38, 0x17, 0xc7, 0x24, 0x9e, 0xb1, 0x3c, 0xe3, 0x94, 0xec, 0x78, 0x02, 0xc4, 0x1b,
	0xf0, 0x22, 0x3c, 0x40, 0x97, 0x5d, 0x76, 0x55, 0x51, 0xb3, 0xe4, 0x25, 0x90, 0xc7, 0x4e, 0x6c,
	0x92, 0x22, 0x94, 0x9d, 0xef, 0xbd, 0x73, 0x7e, 0xf7, 0xcc, 0xf1, 0xa0, 0x93, 0xd1, 0x91, 0x22,
	0x81, 0xa4, 0xa3, 0xa4, 0x0f, 0xb1, 0x00, 0x0d, 0x8a, 0x4e, 0x40, 0x0c, 0x64, 0x4c, 0x8b, 0x01,
	0x8f, 0x02, 0x1a, 0x84, 0xdc, 0x87, 0x48, 0x8e, 0x03, 0x6f, 0x4a, 0x27, 0x07, 0x7c, 0x1c, 0x0d,
	0xf9, 0x01, 0xf5, 0x41, 0x40, 0xcc, 0x35, 0x0c, 0x48, 0x14, 0x4b, 0x2d, 0xf1, 0x6e, 0x2e, 0x20,
	0x3c, 0x0a, 0x48, 0x45, 0x40, 0x66, 0x82, 0xce, 0x43, 0x3f, 0xd0, 0xc3, 0xa4, 0x4f, 0x3c, 0x19,
	0x52, 0x5f, 0xfa, 0x92, 0x1a, 0x5d, 0x3f, 0xf9, 0x60, 0x2a, 0x53, 0x98, 0xaf, 0x9c, 0xd7, 0x79,
	0x5c, 0x1a, 0x08, 0xb9, 0x37, 0x0c, 0x04, 0xc4, 0x53, 0x1a, 0x8d, 0xfc, 0xac, 0xa1, 0x68, 0x08,
	0x9a, 0xd3, 0xc9, 0x92, 0x8b, 0x0e, 0xfd, 0x9b, 0x2a, 0x4e, 0x84, 0x0e, 0x42, 0x58, 0x12, 0x3c,
	0xf9, 0x97, 0x40, 0x79, 0x43, 0x08, 0xf9, 0xa2, 0xae, 0xfb, 0xad, 0x8e, 0x5a, 0xaf, 0xb2, 0x6b,
	0x32, 0x98, 0x04, 0x70, 0x86, 0xdf, 0xa3, 0xad, 0xcc, 0xd3, 0x80, 0x6b, 0x6e, 0x5b, 0x7b, 0xd6,
	0x7e, 0xeb, 0xf0, 0x11, 0x29, 0x13, 0x99, 0xa3, 0x49, 0x34, 0xf2, 0xb3, 0x86, 0x22, 0xd9, 0x69,
	0x32, 0x39, 0x20, 0x6f, 0xfb, 0x1f, 0xc1, 0xd3, 0xaf, 0x41, 0x73, 0x17, 0x9f, 0x5f, 0xed, 0xd6,
	0xd2, 0xab, 0x5d, 0x54, 0xf6, 0xd8, 0x9c, 0x8a, 0x19, 0x5a, 0x53, 0x11, 0x78, 0x76, 0x7d, 0x89,
	0x7e, 0x63, 0xde, 0xa4, 0xe2, 0xee, 0x5d, 0x04, 0x9e, 0xfb, 0x5f, 0x41, 0x5f, 0xcb, 0x2a, 0x66,
	0x58, 0xf8, 0x14, 0x6d, 0x28, 0xcd, 0x75, 0xa2, 0xec, 0x86, 0xa1, 0x1e, 0xae, 0x44, 0x35, 0x4a,
	0xf7, 0x56, 0xc1, 0xdd, 0xc8, 0x6b, 0x56, 0x10, 0xbb, 0x2f, 0x90, 0x5d, 0x39, 0x7c, 0x2c, 0x85,
	0xe6, 0x59, 0x04, 0xd9, 0x76, 0x7c, 0x0f, 0xad, 0x1b, 0xba, 0x89, 0xaa, 0xe9, 0xfe, 0x5f, 0x20,
	0xd6, 0x73, 0x41, 0x3e, 0xeb, 0xfe, 0xaa, 0xa3, 0xed, 0x85, 0x4b, 0xe0, 0x10, 0x21, 0x6f, 0x46,
	0x52, 0xb6, 0xb5, 0xd7, 0xd8, 0x6f, 0x1d, 0x3e, 0x5d, 0xc5, 0xf4, 0x1f, 0x3e, 0xca, 0xc4, 0xe7,
	0x6d, 0xc5, 0x2a, 0x0b, 0xf0, 0x27, 0xd4, 0xe2, 0x42, 0x48, 0xcd, 0x75, 0x20, 0x85, 0xb2, 0xeb,
	0x66, 0x5f, 0x6f, 0xd5, 0xe8, 0x49, 0xaf, 0x64, 0xbc, 0x14, 0x3a, 0x9e, 0xba, 0xb7, 0x8b, 0xbd,
	0xad, 0xca, 0x84, 0x55, 0x57, 0x61, 0x8a, 0x9a, 0x82, 0x87, 0xa0, 0x22, 0xee, 0x81, 0xf9, 0x39,
	0x4d, 0x77, 0xa7, 0x10, 0x35, 0xdf, 0xcc, 0x06, 0xac, 0x3c, 0xd3, 0x79, 0x8e, 0xda, 0x8b, 0x6b,
	0x70, 0x1b, 0x35, 0x46, 0x30, 0xcd, 0x43, 0x66, 0xd9, 0x27, 0xbe, 0x83, 0xd6, 0x27, 0x7c, 0x9c,
	0x80, 0x79, 0x45, 0x4d, 0x96, 0x17, 0xcf, 0xea, 0x47, 0x56, 0xf7, 0x7b, 0x1d, 0xed, 0x2c, 0xfd,
	0x5c, 0xfc, 0x00, 0x6d, 0xf2, 0xf1, 0x58, 0x9e, 0xc1, 0xc0, 0x50, 0xb6, 0xdc, 0xed, 0xc2, 0xc4,
	0x66, 0x2f, 0x6f, 0xb3, 0xd9, 0x1c, 0xdf, 0x47, 0x1b, 0x31, 0x70, 0x25, 0x45, 0xce, 0x2e, 0xdf,
	0x05, 0x33, 0x5d, 0x56, 0x4c, 0xf1, 0x17, 0x0b, 0xb5, 0x79, 0x32, 0x08, 0x74, 0xc5, 0xae, 0xdd,
	0x30, 0xc9, 0x9e, 0xac, 0xfe, 0xfc, 0x48, 0x6f, 0x01, 0x95, 0x07, 0x6c, 0x17, 0xcb, 0xdb, 0x8b,
	0x63, 0xb6, 0xb4, 0xbb, 0x73, 0x8c, 0xee, 0xde, 0x08, 0x59, 0x25, 0x3e, 0x97, 0x9c, 0x5f, 0x3b,
	0xb5, 0x8b, 0x6b, 0xa7, 0x76, 0x79, 0xed, 0xd4, 0x3e, 0xa7, 0x8e, 0x75, 0x9e, 0x3a, 0xd6, 0x45,
	0xea, 0x58, 0x97, 0xa9, 0x63, 0xfd, 0x48, 0x1d, 0xeb, 0xeb, 0x4f, 0xa7, 0x76, 0xba, 0x35, 0xbb,
	0xc8, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0x16, 0x48, 0xa2, 0x79, 0x05, 0x00, 0x00,
}
