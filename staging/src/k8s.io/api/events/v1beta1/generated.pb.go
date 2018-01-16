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
// source: k8s.io/kubernetes/vendor/k8s.io/api/events/v1beta1/generated.proto
// DO NOT EDIT!

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/api/events/v1beta1/generated.proto

	It has these top-level messages:
		Event
		EventList
		EventSeries
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_api_core_v1 "k8s.io/api/core/v1"

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

func (m *Event) Reset()                    { *m = Event{} }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *EventList) Reset()                    { *m = EventList{} }
func (*EventList) ProtoMessage()               {}
func (*EventList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *EventSeries) Reset()                    { *m = EventSeries{} }
func (*EventSeries) ProtoMessage()               {}
func (*EventSeries) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*Event)(nil), "k8s.io.api.events.v1beta1.Event")
	proto.RegisterType((*EventList)(nil), "k8s.io.api.events.v1beta1.EventList")
	proto.RegisterType((*EventSeries)(nil), "k8s.io.api.events.v1beta1.EventSeries")
}
func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
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
	i = encodeVarintGenerated(dAtA, i, uint64(m.EventTime.Size()))
	n2, err := m.EventTime.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.Series != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Series.Size()))
		n3, err := m.Series.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ReportingController)))
	i += copy(dAtA[i:], m.ReportingController)
	dAtA[i] = 0x2a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ReportingInstance)))
	i += copy(dAtA[i:], m.ReportingInstance)
	dAtA[i] = 0x32
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Action)))
	i += copy(dAtA[i:], m.Action)
	dAtA[i] = 0x3a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Reason)))
	i += copy(dAtA[i:], m.Reason)
	dAtA[i] = 0x42
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Regarding.Size()))
	n4, err := m.Regarding.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.Related != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Related.Size()))
		n5, err := m.Related.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	dAtA[i] = 0x52
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Note)))
	i += copy(dAtA[i:], m.Note)
	dAtA[i] = 0x5a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Type)))
	i += copy(dAtA[i:], m.Type)
	dAtA[i] = 0x62
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.DeprecatedSource.Size()))
	n6, err := m.DeprecatedSource.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	dAtA[i] = 0x6a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.DeprecatedFirstTimestamp.Size()))
	n7, err := m.DeprecatedFirstTimestamp.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x72
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.DeprecatedLastTimestamp.Size()))
	n8, err := m.DeprecatedLastTimestamp.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	dAtA[i] = 0x78
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.DeprecatedCount))
	return i, nil
}

func (m *EventList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n9, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n9
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

func (m *EventSeries) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSeries) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Count))
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.LastObservedTime.Size()))
	n10, err := m.LastObservedTime.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n10
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.State)))
	i += copy(dAtA[i:], m.State)
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
func (m *Event) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.EventTime.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.Series != nil {
		l = m.Series.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.ReportingController)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.ReportingInstance)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Action)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Reason)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Regarding.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.Related != nil {
		l = m.Related.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Note)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Type)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.DeprecatedSource.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.DeprecatedFirstTimestamp.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.DeprecatedLastTimestamp.Size()
	n += 1 + l + sovGenerated(uint64(l))
	n += 1 + sovGenerated(uint64(m.DeprecatedCount))
	return n
}

func (m *EventList) Size() (n int) {
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

func (m *EventSeries) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Count))
	l = m.LastObservedTime.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.State)
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
func (this *Event) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Event{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`EventTime:` + strings.Replace(strings.Replace(this.EventTime.String(), "MicroTime", "k8s_io_apimachinery_pkg_apis_meta_v1.MicroTime", 1), `&`, ``, 1) + `,`,
		`Series:` + strings.Replace(fmt.Sprintf("%v", this.Series), "EventSeries", "EventSeries", 1) + `,`,
		`ReportingController:` + fmt.Sprintf("%v", this.ReportingController) + `,`,
		`ReportingInstance:` + fmt.Sprintf("%v", this.ReportingInstance) + `,`,
		`Action:` + fmt.Sprintf("%v", this.Action) + `,`,
		`Reason:` + fmt.Sprintf("%v", this.Reason) + `,`,
		`Regarding:` + strings.Replace(strings.Replace(this.Regarding.String(), "ObjectReference", "k8s_io_api_core_v1.ObjectReference", 1), `&`, ``, 1) + `,`,
		`Related:` + strings.Replace(fmt.Sprintf("%v", this.Related), "ObjectReference", "k8s_io_api_core_v1.ObjectReference", 1) + `,`,
		`Note:` + fmt.Sprintf("%v", this.Note) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`DeprecatedSource:` + strings.Replace(strings.Replace(this.DeprecatedSource.String(), "EventSource", "k8s_io_api_core_v1.EventSource", 1), `&`, ``, 1) + `,`,
		`DeprecatedFirstTimestamp:` + strings.Replace(strings.Replace(this.DeprecatedFirstTimestamp.String(), "Time", "k8s_io_apimachinery_pkg_apis_meta_v1.Time", 1), `&`, ``, 1) + `,`,
		`DeprecatedLastTimestamp:` + strings.Replace(strings.Replace(this.DeprecatedLastTimestamp.String(), "Time", "k8s_io_apimachinery_pkg_apis_meta_v1.Time", 1), `&`, ``, 1) + `,`,
		`DeprecatedCount:` + fmt.Sprintf("%v", this.DeprecatedCount) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EventList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EventList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "Event", "Event", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EventSeries) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EventSeries{`,
		`Count:` + fmt.Sprintf("%v", this.Count) + `,`,
		`LastObservedTime:` + strings.Replace(strings.Replace(this.LastObservedTime.String(), "MicroTime", "k8s_io_apimachinery_pkg_apis_meta_v1.MicroTime", 1), `&`, ``, 1) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
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
func (m *Event) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field EventTime", wireType)
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
			if err := m.EventTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Series", wireType)
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
			if m.Series == nil {
				m.Series = &EventSeries{}
			}
			if err := m.Series.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportingController", wireType)
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
			m.ReportingController = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportingInstance", wireType)
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
			m.ReportingInstance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
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
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
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
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regarding", wireType)
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
			if err := m.Regarding.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Related", wireType)
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
			if m.Related == nil {
				m.Related = &k8s_io_api_core_v1.ObjectReference{}
			}
			if err := m.Related.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Note", wireType)
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
			m.Note = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
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
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedSource", wireType)
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
			if err := m.DeprecatedSource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedFirstTimestamp", wireType)
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
			if err := m.DeprecatedFirstTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedLastTimestamp", wireType)
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
			if err := m.DeprecatedLastTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedCount", wireType)
			}
			m.DeprecatedCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeprecatedCount |= (int32(b) & 0x7F) << shift
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
func (m *EventList) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventList: illegal tag %d (wire type %d)", fieldNum, wire)
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
			m.Items = append(m.Items, Event{})
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
func (m *EventSeries) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventSeries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSeries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastObservedTime", wireType)
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
			if err := m.LastObservedTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
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
			m.State = EventSeriesState(dAtA[iNdEx:postIndex])
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
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/events/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 814 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x16, 0x13, 0x4b, 0xb6, 0x56, 0x49, 0x2c, 0x6f, 0x0e, 0xde, 0xb8, 0x00, 0xa5, 0x3a, 0x40,
	0x60, 0x14, 0x08, 0x59, 0xa7, 0x45, 0xdb, 0x6b, 0x18, 0xbb, 0x45, 0x02, 0xbb, 0x01, 0xd6, 0x3e,
	0x15, 0x3d, 0x64, 0x45, 0x4d, 0xe8, 0xad, 0xa5, 0x5d, 0x62, 0x77, 0x29, 0xc0, 0xb7, 0x5e, 0x0a,
	0xf4, 0xd8, 0x67, 0xe8, 0x13, 0xf4, 0x31, 0x7c, 0xcc, 0x31, 0x27, 0xa1, 0x66, 0xdf, 0xa2, 0xa7,
	0x82, 0xcb, 0x95, 0x28, 0x8b, 0x16, 0xec, 0x22, 0x37, 0x72, 0xe6, 0xfb, 0x99, 0x19, 0x0e, 0x07,
	0x45, 0xe7, 0xdf, 0xe9, 0x80, 0xcb, 0xf0, 0x3c, 0x1b, 0x80, 0x12, 0x60, 0x40, 0x87, 0x13, 0x10,
	0x43, 0xa9, 0x42, 0x97, 0x60, 0x29, 0x0f, 0x61, 0x02, 0xc2, 0xe8, 0x70, 0xb2, 0x3f, 0x00, 0xc3,
	0xf6, 0xc3, 0x04, 0x04, 0x28, 0x66, 0x60, 0x18, 0xa4, 0x4a, 0x1a, 0x89, 0x9f, 0x94, 0xd0, 0x80,
	0xa5, 0x3c, 0x28, 0xa1, 0x81, 0x83, 0xee, 0x3c, 0x4f, 0xb8, 0x39, 0xcb, 0x06, 0x41, 0x2c, 0xc7,
	0x61, 0x22, 0x13, 0x19, 0x5a, 0xc6, 0x20, 0x7b, 0x6f, 0xdf, 0xec, 0x8b, 0x7d, 0x2a, 0x95, 0x76,
	0x76, 0x17, 0x4c, 0x63, 0xa9, 0x20, 0x9c, 0xd4, 0xdc, 0x76, 0xbe, 0xae, 0x30, 0x63, 0x16, 0x9f,
	0x71, 0x01, 0xea, 0x22, 0x4c, 0xcf, 0x93, 0x22, 0xa0, 0xc3, 0x31, 0x18, 0x76, 0x13, 0x2b, 0x5c,
	0xc5, 0x52, 0x99, 0x30, 0x7c, 0x0c, 0x35, 0xc2, 0x37, 0xb7, 0x11, 0x74, 0x7c, 0x06, 0x63, 0x56,
	0xe3, 0x7d, 0xb5, 0x8a, 0x97, 0x19, 0x3e, 0x0a, 0xb9, 0x30, 0xda, 0xa8, 0x65, 0xd2, 0xee, 0x9f,
	0x6d, 0xd4, 0x3c, 0x2c, 0x26, 0x87, 0xdf, 0xa1, 0x8d, 0xa2, 0x85, 0x21, 0x33, 0x8c, 0x78, 0x7d,
	0x6f, 0xaf, 0xf3, 0xe2, 0xcb, 0xa0, 0x1a, 0xef, 0x5c, 0x31, 0x48, 0xcf, 0x93, 0x22, 0xa0, 0x83,
	0x02, 0x1d, 0x4c, 0xf6, 0x83, 0xb7, 0x83, 0x5f, 0x20, 0x36, 0xc7, 0x60, 0x58, 0x84, 0x2f, 0xa7,
	0xbd, 0x46, 0x3e, 0xed, 0xa1, 0x2a, 0x46, 0xe7, 0xaa, 0xf8, 0x1d, 0x6a, 0xdb, 0x8f, 0x74, 0xca,
	0xc7, 0x40, 0xee, 0x59, 0x8b, 0xf0, 0x6e, 0x16, 0xc7, 0x3c, 0x56, 0xb2, 0xa0, 0x45, 0x5b, 0xce,
	0xa1, 0x7d, 0x38, 0x53, 0xa2, 0x95, 0x28, 0x7e, 0x83, 0x5a, 0x1a, 0x14, 0x07, 0x4d, 0xee, 0x5b,
	0xf9, 0x67, 0xc1, 0xca, 0x05, 0x09, 0xac, 0xc0, 0x89, 0x45, 0x47, 0x28, 0x9f, 0xf6, 0x5a, 0xe5,
	0x33, 0x75, 0x0a, 0xf8, 0x18, 0x3d, 0x56, 0x90, 0x4a, 0x65, 0xb8, 0x48, 0x5e, 0x49, 0x61, 0x94,
	0x1c, 0x8d, 0x40, 0x91, 0xb5, 0xbe, 0xb7, 0xd7, 0x8e, 0x3e, 0x73, 0x65, 0x3c, 0xa6, 0x75, 0x08,
	0xbd, 0x89, 0x87, 0x7f, 0x40, 0x5b, 0xf3, 0xf0, 0x6b, 0xa1, 0x0d, 0x13, 0x31, 0x90, 0xa6, 0x15,
	0x7b, 0xe2, 0xc4, 0xb6, 0xe8, 0x32, 0x80, 0xd6, 0x39, 0xf8, 0x19, 0x6a, 0xb1, 0xd8, 0x70, 0x29,
	0x48, 0xcb, 0xb2, 0x1f, 0x39, 0x76, 0xeb, 0xa5, 0x8d, 0x52, 0x97, 0x2d, 0x70, 0x0a, 0x98, 0x96,
	0x82, 0xac, 0x5f, 0xc7, 0x51, 0x1b, 0xa5, 0x2e, 0x8b, 0x4f, 0x51, 0x5b, 0x41, 0xc2, 0xd4, 0x90,
	0x8b, 0x84, 0x6c, 0xd8, 0xb1, 0x3d, 0x5d, 0x1c, 0x5b, 0xf1, 0x37, 0x54, 0x9f, 0x99, 0xc2, 0x7b,
	0x50, 0x20, 0xe2, 0x85, 0x2f, 0x41, 0x67, 0x6c, 0x5a, 0x09, 0xe1, 0x37, 0x68, 0x5d, 0xc1, 0xa8,
	0x58, 0x34, 0xd2, 0xbe, 0xbb, 0x66, 0x27, 0x9f, 0xf6, 0xd6, 0x69, 0xc9, 0xa3, 0x33, 0x01, 0xdc,
	0x47, 0x6b, 0x42, 0x1a, 0x20, 0xc8, 0xf6, 0xf1, 0xc0, 0xf9, 0xae, 0xfd, 0x28, 0x0d, 0x50, 0x9b,
	0x29, 0x10, 0xe6, 0x22, 0x05, 0xd2, 0xb9, 0x8e, 0x38, 0xbd, 0x48, 0x81, 0xda, 0x0c, 0x06, 0xd4,
	0x1d, 0x42, 0xaa, 0x20, 0x2e, 0x14, 0x4f, 0x64, 0xa6, 0x62, 0x20, 0x0f, 0x6c, 0x61, 0xbd, 0x9b,
	0x0a, 0x2b, 0x97, 0xc3, 0xc2, 0x22, 0xe2, 0xe4, 0xba, 0x07, 0x4b, 0x02, 0xb4, 0x26, 0x89, 0x7f,
	0xf7, 0x10, 0xa9, 0x82, 0xdf, 0x73, 0xa5, 0xed, 0x62, 0x6a, 0xc3, 0xc6, 0x29, 0x79, 0x68, 0xfd,
	0xbe, 0xb8, 0xdb, 0xca, 0xdb, 0x6d, 0xef, 0x3b, 0x6b, 0x72, 0xb0, 0x42, 0x93, 0xae, 0x74, 0xc3,
	0xbf, 0x79, 0x68, 0xbb, 0x4a, 0x1e, 0xb1, 0xc5, 0x4a, 0x1e, 0xfd, 0xef, 0x4a, 0x7a, 0xae, 0x92,
	0xed, 0x83, 0x9b, 0x25, 0xe9, 0x2a, 0x2f, 0xfc, 0x12, 0x6d, 0x56, 0xa9, 0x57, 0x32, 0x13, 0x86,
	0x6c, 0xf6, 0xbd, 0xbd, 0x66, 0xb4, 0xed, 0x24, 0x37, 0x0f, 0xae, 0xa7, 0xe9, 0x32, 0x7e, 0xf7,
	0x2f, 0x0f, 0x95, 0xff, 0xfb, 0x11, 0xd7, 0x06, 0xff, 0x5c, 0x3b, 0x54, 0xc1, 0xdd, 0x1a, 0x29,
	0xd8, 0xf6, 0x4c, 0x75, 0x9d, 0xf3, 0xc6, 0x2c, 0xb2, 0x70, 0xa4, 0x0e, 0x51, 0x93, 0x1b, 0x18,
	0x6b, 0x72, 0xaf, 0x7f, 0x7f, 0xaf, 0xf3, 0xa2, 0x7f, 0xdb, 0x05, 0x89, 0x1e, 0x3a, 0xb1, 0xe6,
	0xeb, 0x82, 0x46, 0x4b, 0xf6, 0x6e, 0xee, 0xa1, 0xce, 0xc2, 0x85, 0xc1, 0x4f, 0x51, 0x33, 0xb6,
	0xbd, 0x7b, 0xb6, 0xf7, 0x39, 0xa9, 0xec, 0xb8, 0xcc, 0xe1, 0x0c, 0x75, 0x47, 0x4c, 0x9b, 0xb7,
	0x03, 0x0d, 0x6a, 0x02, 0xc3, 0x4f, 0xb9, 0x93, 0xf3, 0xa5, 0x3d, 0x5a, 0x12, 0xa4, 0x35, 0x0b,
	0xfc, 0x2d, 0x6a, 0x6a, 0xc3, 0x0c, 0xd8, 0xa3, 0xd9, 0x8e, 0x3e, 0x9f, 0xd5, 0x76, 0x52, 0x04,
	0xff, 0x9d, 0xf6, 0xba, 0x0b, 0x8d, 0xd8, 0x18, 0x2d, 0xf1, 0xd1, 0xf3, 0xcb, 0x2b, 0xbf, 0xf1,
	0xe1, 0xca, 0x6f, 0x7c, 0xbc, 0xf2, 0x1b, 0xbf, 0xe6, 0xbe, 0x77, 0x99, 0xfb, 0xde, 0x87, 0xdc,
	0xf7, 0x3e, 0xe6, 0xbe, 0xf7, 0x77, 0xee, 0x7b, 0x7f, 0xfc, 0xe3, 0x37, 0x7e, 0x5a, 0x77, 0xf3,
	0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x69, 0xa9, 0x7b, 0x6e, 0xf2, 0x07, 0x00, 0x00,
}
