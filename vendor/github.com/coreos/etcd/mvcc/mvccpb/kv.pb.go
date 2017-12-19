// Code generated by protoc-gen-gogo.
// source: kv.proto
// DO NOT EDIT!

/*
	Package mvccpb is a generated protocol buffer package.

	It is generated from these files:
		kv.proto

	It has these top-level messages:
		KeyValue
		Event
*/
package mvccpb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event_EventType int32

const (
	PUT    Event_EventType = 0
	DELETE Event_EventType = 1
)

var Event_EventType_name = map[int32]string{
	0: "PUT",
	1: "DELETE",
}
var Event_EventType_value = map[string]int32{
	"PUT":    0,
	"DELETE": 1,
}

func (x Event_EventType) String() string {
	return proto.EnumName(Event_EventType_name, int32(x))
}
func (Event_EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorKv, []int{1, 0} }

type KeyValue struct {
	// key is the key in bytes. An empty key is not allowed.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// create_revision is the revision of last creation on this key.
	CreateRevision int64 `protobuf:"varint,2,opt,name=create_revision,json=createRevision,proto3" json:"create_revision,omitempty"`
	// mod_revision is the revision of last modification on this key.
	ModRevision int64 `protobuf:"varint,3,opt,name=mod_revision,json=modRevision,proto3" json:"mod_revision,omitempty"`
	// version is the version of the key. A deletion resets
	// the version to zero and any modification of the key
	// increases its version.
	Version int64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	// value is the value held by the key, in bytes.
	Value []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	// lease is the ID of the lease that attached to key.
	// When the attached lease expires, the key will be deleted.
	// If lease is 0, then no lease is attached to the key.
	Lease int64 `protobuf:"varint,6,opt,name=lease,proto3" json:"lease,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptorKv, []int{0} }

type Event struct {
	// type is the kind of event. If type is a PUT, it indicates
	// new data has been stored to the key. If type is a DELETE,
	// it indicates the key was deleted.
	Type Event_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=mvccpb.Event_EventType" json:"type,omitempty"`
	// kv holds the KeyValue for the event.
	// A PUT event contains current kv pair.
	// A PUT event with kv.Version=1 indicates the creation of a key.
	// A DELETE/EXPIRE event contains the deleted key with
	// its modification revision set to the revision of deletion.
	Kv *KeyValue `protobuf:"bytes,2,opt,name=kv" json:"kv,omitempty"`
	// prev_kv holds the key-value pair before the event happens.
	PrevKv *KeyValue `protobuf:"bytes,3,opt,name=prev_kv,json=prevKv" json:"prev_kv,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorKv, []int{1} }

func init() {
	proto.RegisterType((*KeyValue)(nil), "mvccpb.KeyValue")
	proto.RegisterType((*Event)(nil), "mvccpb.Event")
	proto.RegisterEnum("mvccpb.Event_EventType", Event_EventType_name, Event_EventType_value)
}
func (m *KeyValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKv(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if m.CreateRevision != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintKv(dAtA, i, uint64(m.CreateRevision))
	}
	if m.ModRevision != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintKv(dAtA, i, uint64(m.ModRevision))
	}
	if m.Version != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintKv(dAtA, i, uint64(m.Version))
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintKv(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.Lease != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintKv(dAtA, i, uint64(m.Lease))
	}
	return i, nil
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
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintKv(dAtA, i, uint64(m.Type))
	}
	if m.Kv != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintKv(dAtA, i, uint64(m.Kv.Size()))
		n1, err := m.Kv.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.PrevKv != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintKv(dAtA, i, uint64(m.PrevKv.Size()))
		n2, err := m.PrevKv.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64Kv(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Kv(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintKv(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *KeyValue) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovKv(uint64(l))
	}
	if m.CreateRevision != 0 {
		n += 1 + sovKv(uint64(m.CreateRevision))
	}
	if m.ModRevision != 0 {
		n += 1 + sovKv(uint64(m.ModRevision))
	}
	if m.Version != 0 {
		n += 1 + sovKv(uint64(m.Version))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovKv(uint64(l))
	}
	if m.Lease != 0 {
		n += 1 + sovKv(uint64(m.Lease))
	}
	return n
}

func (m *Event) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovKv(uint64(m.Type))
	}
	if m.Kv != nil {
		l = m.Kv.Size()
		n += 1 + l + sovKv(uint64(l))
	}
	if m.PrevKv != nil {
		l = m.PrevKv.Size()
		n += 1 + l + sovKv(uint64(l))
	}
	return n
}

func sovKv(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozKv(x uint64) (n int) {
	return sovKv(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKv
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
			return fmt.Errorf("proto: KeyValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
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
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateRevision", wireType)
			}
			m.CreateRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModRevision", wireType)
			}
			m.ModRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
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
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			m.Lease = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lease |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKv
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
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKv
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Event_EventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
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
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Kv == nil {
				m.Kv = &KeyValue{}
			}
			if err := m.Kv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
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
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrevKv == nil {
				m.PrevKv = &KeyValue{}
			}
			if err := m.PrevKv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKv
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
func skipKv(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKv
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
					return 0, ErrIntOverflowKv
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
					return 0, ErrIntOverflowKv
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
				return 0, ErrInvalidLengthKv
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKv
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
				next, err := skipKv(dAtA[start:])
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
	ErrInvalidLengthKv = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKv   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("kv.proto", fileDescriptorKv) }

var fileDescriptorKv = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4e, 0xc2, 0x40,
	0x14, 0x86, 0x3b, 0x14, 0x0a, 0x3e, 0x08, 0x36, 0x13, 0x12, 0x27, 0x2e, 0x26, 0x95, 0x8d, 0x18,
	0x13, 0x4c, 0xf0, 0x06, 0xc6, 0xae, 0x70, 0x61, 0x1a, 0x74, 0x4b, 0x4a, 0x79, 0x21, 0xa4, 0x94,
	0x69, 0x4a, 0x9d, 0xa4, 0x37, 0x71, 0xef, 0xde, 0x73, 0xb0, 0xe4, 0x08, 0x52, 0x2f, 0x62, 0xfa,
	0xc6, 0xe2, 0xc6, 0xcd, 0xe4, 0xfd, 0xff, 0xff, 0x65, 0xe6, 0x7f, 0x03, 0x9d, 0x58, 0x8f, 0xd3,
	0x4c, 0xe5, 0x8a, 0x3b, 0x89, 0x8e, 0xa2, 0x74, 0x71, 0x39, 0x58, 0xa9, 0x95, 0x22, 0xeb, 0xae,
	0x9a, 0x4c, 0x3a, 0xfc, 0x64, 0xd0, 0x99, 0x62, 0xf1, 0x1a, 0x6e, 0xde, 0x90, 0xbb, 0x60, 0xc7,
	0x58, 0x08, 0xe6, 0xb1, 0x51, 0x2f, 0xa8, 0x46, 0x7e, 0x0d, 0xe7, 0x51, 0x86, 0x61, 0x8e, 0xf3,
	0x0c, 0xf5, 0x7a, 0xb7, 0x56, 0x5b, 0xd1, 0xf0, 0xd8, 0xc8, 0x0e, 0xfa, 0xc6, 0x0e, 0x7e, 0x5d,
	0x7e, 0x05, 0xbd, 0x44, 0x2d, 0xff, 0x28, 0x9b, 0xa8, 0x6e, 0xa2, 0x96, 0x27, 0x44, 0x40, 0x5b,
	0x63, 0x46, 0x69, 0x93, 0xd2, 0x5a, 0xf2, 0x01, 0xb4, 0x74, 0x55, 0x40, 0xb4, 0xe8, 0x65, 0x23,
	0x2a, 0x77, 0x83, 0xe1, 0x0e, 0x85, 0x43, 0xb4, 0x11, 0xc3, 0x0f, 0x06, 0x2d, 0x5f, 0xe3, 0x36,
	0xe7, 0xb7, 0xd0, 0xcc, 0x8b, 0x14, 0xa9, 0x6e, 0x7f, 0x72, 0x31, 0x36, 0x7b, 0x8e, 0x29, 0x34,
	0xe7, 0xac, 0x48, 0x31, 0x20, 0x88, 0x7b, 0xd0, 0x88, 0x35, 0x75, 0xef, 0x4e, 0xdc, 0x1a, 0xad,
	0x17, 0x0f, 0x1a, 0xb1, 0xe6, 0x37, 0xd0, 0x4e, 0x33, 0xd4, 0xf3, 0x58, 0x53, 0xf9, 0xff, 0x30,
	0xa7, 0x02, 0xa6, 0x7a, 0xe8, 0xc1, 0xd9, 0xe9, 0x7e, 0xde, 0x06, 0xfb, 0xf9, 0x65, 0xe6, 0x5a,
	0x1c, 0xc0, 0x79, 0xf4, 0x9f, 0xfc, 0x99, 0xef, 0xb2, 0x07, 0xb1, 0x3f, 0x4a, 0xeb, 0x70, 0x94,
	0xd6, 0xbe, 0x94, 0xec, 0x50, 0x4a, 0xf6, 0x55, 0x4a, 0xf6, 0xfe, 0x2d, 0xad, 0x85, 0x43, 0xff,
	0x7e, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x45, 0x92, 0x5d, 0xa1, 0x01, 0x00, 0x00,
}
