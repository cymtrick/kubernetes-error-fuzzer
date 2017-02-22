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
// source: k8s.io/metrics/pkg/apis/metrics/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/metrics/pkg/apis/metrics/v1alpha1/generated.proto

	It has these top-level messages:
		ContainerMetrics
		NodeMetrics
		NodeMetricsList
		PodMetrics
		PodMetricsList
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_apimachinery_pkg_api_resource "k8s.io/apimachinery/pkg/api/resource"

import k8s_io_client_go_pkg_api_v1 "k8s.io/client-go/pkg/api/v1"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *ContainerMetrics) Reset()                    { *m = ContainerMetrics{} }
func (*ContainerMetrics) ProtoMessage()               {}
func (*ContainerMetrics) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *NodeMetrics) Reset()                    { *m = NodeMetrics{} }
func (*NodeMetrics) ProtoMessage()               {}
func (*NodeMetrics) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *NodeMetricsList) Reset()                    { *m = NodeMetricsList{} }
func (*NodeMetricsList) ProtoMessage()               {}
func (*NodeMetricsList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *PodMetrics) Reset()                    { *m = PodMetrics{} }
func (*PodMetrics) ProtoMessage()               {}
func (*PodMetrics) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *PodMetricsList) Reset()                    { *m = PodMetricsList{} }
func (*PodMetricsList) ProtoMessage()               {}
func (*PodMetricsList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{4} }

func init() {
	proto.RegisterType((*ContainerMetrics)(nil), "k8s.io.metrics.pkg.apis.metrics.v1alpha1.ContainerMetrics")
	proto.RegisterType((*NodeMetrics)(nil), "k8s.io.metrics.pkg.apis.metrics.v1alpha1.NodeMetrics")
	proto.RegisterType((*NodeMetricsList)(nil), "k8s.io.metrics.pkg.apis.metrics.v1alpha1.NodeMetricsList")
	proto.RegisterType((*PodMetrics)(nil), "k8s.io.metrics.pkg.apis.metrics.v1alpha1.PodMetrics")
	proto.RegisterType((*PodMetricsList)(nil), "k8s.io.metrics.pkg.apis.metrics.v1alpha1.PodMetricsList")
}
func (m *ContainerMetrics) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ContainerMetrics) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	if len(m.Usage) > 0 {
		for k := range m.Usage {
			data[i] = 0x12
			i++
			v := m.Usage[k]
			msgSize := (&v).Size()
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + msgSize + sovGenerated(uint64(msgSize))
			i = encodeVarintGenerated(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintGenerated(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64((&v).Size()))
			n1, err := (&v).MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n1
		}
	}
	return i, nil
}

func (m *NodeMetrics) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *NodeMetrics) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ObjectMeta.Size()))
	n2, err := m.ObjectMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Timestamp.Size()))
	n3, err := m.Timestamp.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Window.Size()))
	n4, err := m.Window.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Usage) > 0 {
		for k := range m.Usage {
			data[i] = 0x22
			i++
			v := m.Usage[k]
			msgSize := (&v).Size()
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + msgSize + sovGenerated(uint64(msgSize))
			i = encodeVarintGenerated(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintGenerated(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64((&v).Size()))
			n5, err := (&v).MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n5
		}
	}
	return i, nil
}

func (m *NodeMetricsList) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *NodeMetricsList) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ListMeta.Size()))
	n6, err := m.ListMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PodMetrics) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PodMetrics) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ObjectMeta.Size()))
	n7, err := m.ObjectMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Timestamp.Size()))
	n8, err := m.Timestamp.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Window.Size()))
	n9, err := m.Window.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n9
	if len(m.Containers) > 0 {
		for _, msg := range m.Containers {
			data[i] = 0x22
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PodMetricsList) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PodMetricsList) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ListMeta.Size()))
	n10, err := m.ListMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n10
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Generated(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ContainerMetrics) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Usage) > 0 {
		for k, v := range m.Usage {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + l + sovGenerated(uint64(l))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *NodeMetrics) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Timestamp.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Window.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Usage) > 0 {
		for k, v := range m.Usage {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + l + sovGenerated(uint64(l))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *NodeMetricsList) Size() (n int) {
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

func (m *PodMetrics) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Timestamp.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Window.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Containers) > 0 {
		for _, e := range m.Containers {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *PodMetricsList) Size() (n int) {
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
func (this *ContainerMetrics) String() string {
	if this == nil {
		return "nil"
	}
	keysForUsage := make([]string, 0, len(this.Usage))
	for k := range this.Usage {
		keysForUsage = append(keysForUsage, string(k))
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForUsage)
	mapStringForUsage := "k8s_io_client_go_pkg_api_v1.ResourceList{"
	for _, k := range keysForUsage {
		mapStringForUsage += fmt.Sprintf("%v: %v,", k, this.Usage[k8s_io_client_go_pkg_api_v1.ResourceName(k)])
	}
	mapStringForUsage += "}"
	s := strings.Join([]string{`&ContainerMetrics{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Usage:` + mapStringForUsage + `,`,
		`}`,
	}, "")
	return s
}
func (this *NodeMetrics) String() string {
	if this == nil {
		return "nil"
	}
	keysForUsage := make([]string, 0, len(this.Usage))
	for k := range this.Usage {
		keysForUsage = append(keysForUsage, string(k))
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForUsage)
	mapStringForUsage := "k8s_io_client_go_pkg_api_v1.ResourceList{"
	for _, k := range keysForUsage {
		mapStringForUsage += fmt.Sprintf("%v: %v,", k, this.Usage[k8s_io_client_go_pkg_api_v1.ResourceName(k)])
	}
	mapStringForUsage += "}"
	s := strings.Join([]string{`&NodeMetrics{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Timestamp:` + strings.Replace(strings.Replace(this.Timestamp.String(), "Time", "k8s_io_apimachinery_pkg_apis_meta_v1.Time", 1), `&`, ``, 1) + `,`,
		`Window:` + strings.Replace(strings.Replace(this.Window.String(), "Duration", "k8s_io_apimachinery_pkg_apis_meta_v1.Duration", 1), `&`, ``, 1) + `,`,
		`Usage:` + mapStringForUsage + `,`,
		`}`,
	}, "")
	return s
}
func (this *NodeMetricsList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NodeMetricsList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "NodeMetrics", "NodeMetrics", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodMetrics) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PodMetrics{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Timestamp:` + strings.Replace(strings.Replace(this.Timestamp.String(), "Time", "k8s_io_apimachinery_pkg_apis_meta_v1.Time", 1), `&`, ``, 1) + `,`,
		`Window:` + strings.Replace(strings.Replace(this.Window.String(), "Duration", "k8s_io_apimachinery_pkg_apis_meta_v1.Duration", 1), `&`, ``, 1) + `,`,
		`Containers:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Containers), "ContainerMetrics", "ContainerMetrics", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodMetricsList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PodMetricsList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "PodMetrics", "PodMetrics", 1), `&`, ``, 1) + `,`,
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
func (m *ContainerMetrics) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContainerMetrics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerMetrics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
				b := data[iNdEx]
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
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Usage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
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
				b := data[iNdEx]
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
				b := data[iNdEx]
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
			mapkey := k8s_io_client_go_pkg_api_v1.ResourceName(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapmsglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				mapmsglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if mapmsglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postmsgIndex := iNdEx + mapmsglen
			if mapmsglen < 0 {
				return ErrInvalidLengthGenerated
			}
			if postmsgIndex > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := &k8s_io_apimachinery_pkg_api_resource.Quantity{}
			if err := mapvalue.Unmarshal(data[iNdEx:postmsgIndex]); err != nil {
				return err
			}
			iNdEx = postmsgIndex
			if m.Usage == nil {
				m.Usage = make(k8s_io_client_go_pkg_api_v1.ResourceList)
			}
			m.Usage[k8s_io_client_go_pkg_api_v1.ResourceName(mapkey)] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
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
func (m *NodeMetrics) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeMetrics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeMetrics: illegal tag %d (wire type %d)", fieldNum, wire)
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
				b := data[iNdEx]
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
			if err := m.ObjectMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
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
			if err := m.Timestamp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Window", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
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
			if err := m.Window.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Usage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
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
				b := data[iNdEx]
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
				b := data[iNdEx]
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
			mapkey := k8s_io_client_go_pkg_api_v1.ResourceName(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapmsglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				mapmsglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if mapmsglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postmsgIndex := iNdEx + mapmsglen
			if mapmsglen < 0 {
				return ErrInvalidLengthGenerated
			}
			if postmsgIndex > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := &k8s_io_apimachinery_pkg_api_resource.Quantity{}
			if err := mapvalue.Unmarshal(data[iNdEx:postmsgIndex]); err != nil {
				return err
			}
			iNdEx = postmsgIndex
			if m.Usage == nil {
				m.Usage = make(k8s_io_client_go_pkg_api_v1.ResourceList)
			}
			m.Usage[k8s_io_client_go_pkg_api_v1.ResourceName(mapkey)] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
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
func (m *NodeMetricsList) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeMetricsList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeMetricsList: illegal tag %d (wire type %d)", fieldNum, wire)
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
				b := data[iNdEx]
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
			if err := m.ListMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
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
				b := data[iNdEx]
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
			m.Items = append(m.Items, NodeMetrics{})
			if err := m.Items[len(m.Items)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
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
func (m *PodMetrics) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodMetrics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodMetrics: illegal tag %d (wire type %d)", fieldNum, wire)
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
				b := data[iNdEx]
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
			if err := m.ObjectMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
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
			if err := m.Timestamp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Window", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
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
			if err := m.Window.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
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
				b := data[iNdEx]
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
			m.Containers = append(m.Containers, ContainerMetrics{})
			if err := m.Containers[len(m.Containers)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
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
func (m *PodMetricsList) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodMetricsList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodMetricsList: illegal tag %d (wire type %d)", fieldNum, wire)
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
				b := data[iNdEx]
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
			if err := m.ListMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
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
				b := data[iNdEx]
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
			m.Items = append(m.Items, PodMetrics{})
			if err := m.Items[len(m.Items)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
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
func skipGenerated(data []byte) (n int, err error) {
	l := len(data)
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
			b := data[iNdEx]
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
				if data[iNdEx-1] < 0x80 {
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
				b := data[iNdEx]
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
					b := data[iNdEx]
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
				next, err := skipGenerated(data[start:])
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

var fileDescriptorGenerated = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x54, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0x8d, 0x9b, 0xa4, 0x6a, 0x2f, 0x50, 0x8a, 0xa7, 0x2a, 0x83, 0x5b, 0x65, 0x8a, 0x8a, 0x7a,
	0xa6, 0xa5, 0xa0, 0xa8, 0x13, 0x32, 0xed, 0x80, 0x44, 0x0b, 0x58, 0x05, 0x44, 0x61, 0xe0, 0xe2,
	0x1c, 0xce, 0x91, 0xf8, 0xce, 0xf2, 0x9d, 0x53, 0x65, 0x43, 0xfc, 0x02, 0x24, 0x7e, 0x0f, 0x0b,
	0x53, 0x10, 0x12, 0xea, 0xc8, 0x14, 0x48, 0xfa, 0x2f, 0x98, 0x90, 0xcf, 0xe7, 0x38, 0xd4, 0x34,
	0xb5, 0x3a, 0xb0, 0xc0, 0x66, 0x7f, 0xf6, 0x7b, 0xdf, 0xf7, 0xde, 0xf7, 0xee, 0x40, 0xa3, 0xd3,
	0xe0, 0x90, 0x30, 0xd3, 0xc3, 0x22, 0x20, 0x0e, 0x37, 0xfd, 0x8e, 0x6b, 0x22, 0x9f, 0xf0, 0x49,
	0xa1, 0xb7, 0x89, 0xba, 0x7e, 0x1b, 0x6d, 0x9a, 0x2e, 0xa6, 0x38, 0x40, 0x02, 0xb7, 0xa0, 0x1f,
	0x30, 0xc1, 0xf4, 0x7a, 0x8c, 0x84, 0xea, 0x47, 0xe8, 0x77, 0x5c, 0x18, 0x21, 0x27, 0x85, 0x04,
	0x59, 0xdd, 0x70, 0x89, 0x68, 0x87, 0x4d, 0xe8, 0x30, 0xcf, 0x74, 0x99, 0xcb, 0x4c, 0x49, 0xd0,
	0x0c, 0x5f, 0xcb, 0x37, 0xf9, 0x22, 0x9f, 0x62, 0xe2, 0xea, 0xb6, 0x1a, 0x09, 0xf9, 0xc4, 0x43,
	0x4e, 0x9b, 0x50, 0x1c, 0xf4, 0x93, 0xb9, 0xcc, 0x00, 0x73, 0x16, 0x06, 0x0e, 0x3e, 0x3b, 0xce,
	0x4c, 0x94, 0x54, 0x83, 0xcc, 0x5e, 0x46, 0x44, 0xd5, 0x3c, 0x0f, 0x15, 0x84, 0x54, 0x10, 0x2f,
	0xdb, 0xe6, 0xce, 0x45, 0x00, 0xee, 0xb4, 0xb1, 0x87, 0x32, 0xb8, 0x5b, 0xe7, 0xe1, 0x42, 0x41,
	0xba, 0x26, 0xa1, 0x82, 0x8b, 0x20, 0x03, 0xba, 0xa1, 0x40, 0x4e, 0x97, 0x60, 0x2a, 0x36, 0x22,
	0xe7, 0x94, 0x0d, 0x59, 0x29, 0xb5, 0xd3, 0x39, 0xb0, 0x7c, 0x8f, 0x51, 0x81, 0x22, 0xee, 0xfd,
	0x78, 0x07, 0xfa, 0x1a, 0x28, 0x51, 0xe4, 0xe1, 0x15, 0x6d, 0x4d, 0xab, 0x2f, 0x5a, 0x57, 0x06,
	0xc3, 0xd5, 0xc2, 0x78, 0xb8, 0x5a, 0x3a, 0x40, 0x1e, 0xb6, 0xe5, 0x17, 0xfd, 0x93, 0x06, 0xca,
	0x21, 0x47, 0x2e, 0x5e, 0x99, 0x5b, 0x2b, 0xd6, 0x2b, 0x5b, 0x7b, 0x30, 0xef, 0x5e, 0xe1, 0xd9,
	0x6e, 0xf0, 0x49, 0xc4, 0xb3, 0x47, 0x45, 0xd0, 0xb7, 0xb0, 0x6a, 0x55, 0x96, 0xb5, 0x9f, 0xc3,
	0xd5, 0xfa, 0x0c, 0x2d, 0xd0, 0x56, 0x5b, 0x7d, 0x40, 0xb8, 0x78, 0xf7, 0x3d, 0xdf, 0xbf, 0x52,
	0x43, 0x3c, 0x7a, 0xb5, 0x0d, 0x40, 0xda, 0x5b, 0x5f, 0x06, 0xc5, 0x0e, 0xee, 0xc7, 0x9a, 0xed,
	0xe8, 0x51, 0xdf, 0x05, 0xe5, 0x1e, 0xea, 0x86, 0x91, 0x46, 0xad, 0x5e, 0xd9, 0x82, 0x89, 0xc6,
	0xe9, 0x6d, 0x24, 0x42, 0x61, 0x12, 0x31, 0xf8, 0x38, 0x44, 0x54, 0x10, 0xd1, 0xb7, 0x63, 0xf0,
	0xce, 0x5c, 0x43, 0xab, 0x7d, 0x2d, 0x81, 0xca, 0x01, 0x6b, 0xe1, 0xc4, 0xe0, 0x57, 0x60, 0x21,
	0xca, 0x56, 0x0b, 0x09, 0x24, 0x1b, 0x56, 0xb6, 0x6e, 0xce, 0x22, 0x97, 0x2e, 0x22, 0xd8, 0xdb,
	0x84, 0x0f, 0x9b, 0x6f, 0xb0, 0x23, 0xf6, 0xb1, 0x40, 0x96, 0xae, 0xbc, 0x02, 0x69, 0xcd, 0x9e,
	0xb0, 0xea, 0x2f, 0xc0, 0x62, 0x14, 0x2c, 0x2e, 0x90, 0xe7, 0xab, 0xf9, 0xd7, 0xf3, 0xb5, 0x38,
	0x24, 0x1e, 0xb6, 0xae, 0x2b, 0xf2, 0xc5, 0xc3, 0x84, 0xc4, 0x4e, 0xf9, 0xf4, 0xa7, 0x60, 0xfe,
	0x98, 0xd0, 0x16, 0x3b, 0x5e, 0x29, 0x5e, 0xec, 0x4c, 0xca, 0xbc, 0x1b, 0x06, 0x48, 0x10, 0x46,
	0xad, 0x25, 0xc5, 0x3e, 0xff, 0x4c, 0xb2, 0xd8, 0x8a, 0x4d, 0xff, 0x38, 0x49, 0x55, 0x49, 0xa6,
	0xea, 0x6e, 0xfe, 0x54, 0x4d, 0xb9, 0xfb, 0x2f, 0x04, 0xea, 0x8b, 0x06, 0xae, 0x4d, 0x49, 0x8e,
	0x06, 0xd6, 0x5f, 0x66, 0x42, 0x95, 0x73, 0x2f, 0x11, 0x5a, 0x46, 0x6a, 0x59, 0xb9, 0xb5, 0x90,
	0x54, 0xa6, 0x02, 0x75, 0x04, 0xca, 0x44, 0x60, 0x8f, 0xab, 0x03, 0x7f, 0xfb, 0x52, 0xab, 0xb1,
	0xae, 0x26, 0xfb, 0xb8, 0x1f, 0x71, 0xd9, 0x31, 0x65, 0xed, 0x43, 0x11, 0x80, 0x47, 0xac, 0xf5,
	0xff, 0x74, 0xcc, 0x3c, 0x1d, 0x14, 0x00, 0x27, 0xb9, 0x3b, 0xb9, 0x3a, 0x21, 0x3b, 0x97, 0xbf,
	0x77, 0x53, 0x8b, 0x26, 0x5f, 0xb8, 0x3d, 0xd5, 0xa1, 0xf6, 0x59, 0x03, 0x4b, 0xe9, 0x56, 0xfe,
	0x42, 0xc4, 0x9e, 0xff, 0x1e, 0xb1, 0xed, 0xfc, 0xda, 0xd2, 0x31, 0xff, 0x9c, 0x30, 0x6b, 0x7d,
	0x30, 0x32, 0x0a, 0x27, 0x23, 0xa3, 0xf0, 0x6d, 0x64, 0x14, 0xde, 0x8e, 0x0d, 0x6d, 0x30, 0x36,
	0xb4, 0x93, 0xb1, 0xa1, 0xfd, 0x18, 0x1b, 0xda, 0xfb, 0x53, 0xa3, 0x70, 0xb4, 0x90, 0x90, 0xfd,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x2d, 0x5f, 0xdb, 0xdd, 0x08, 0x00, 0x00,
}
