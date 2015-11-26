// Code generated by protoc-gen-gogo.
// source: gogovanity.proto
// DO NOT EDIT!

/*
	Package vanity is a generated protocol buffer package.

	It is generated from these files:
		gogovanity.proto

	It has these top-level messages:
		B
*/
package vanity

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type B struct {
	String_ *string `protobuf:"bytes,1,opt,name=String" json:"String,omitempty"`
	Int64   int64   `protobuf:"varint,2,opt,name=Int64" json:"Int64"`
	Int32   *int32  `protobuf:"varint,3,opt,name=Int32,def=1234" json:"Int32,omitempty"`
}

func (m *B) Reset()      { *m = B{} }
func (*B) ProtoMessage() {}

const Default_B_Int32 int32 = 1234

func (m *B) GetString_() string {
	if m != nil && m.String_ != nil {
		return *m.String_
	}
	return ""
}

func (m *B) GetInt64() int64 {
	if m != nil {
		return m.Int64
	}
	return 0
}

func (m *B) GetInt32() int32 {
	if m != nil && m.Int32 != nil {
		return *m.Int32
	}
	return Default_B_Int32
}

func init() {
	proto.RegisterType((*B)(nil), "vanity.B")
}
func (this *B) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*B)
	if !ok {
		that2, ok := that.(B)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.String_ != nil && that1.String_ != nil {
		if *this.String_ != *that1.String_ {
			return false
		}
	} else if this.String_ != nil {
		return false
	} else if that1.String_ != nil {
		return false
	}
	if this.Int64 != that1.Int64 {
		return false
	}
	if this.Int32 != nil && that1.Int32 != nil {
		if *this.Int32 != *that1.Int32 {
			return false
		}
	} else if this.Int32 != nil {
		return false
	} else if that1.Int32 != nil {
		return false
	}
	return true
}
func (this *B) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&vanity.B{")
	if this.String_ != nil {
		s = append(s, "String_: "+valueToGoStringGogovanity(this.String_, "string")+",\n")
	}
	s = append(s, "Int64: "+fmt.Sprintf("%#v", this.Int64)+",\n")
	if this.Int32 != nil {
		s = append(s, "Int32: "+valueToGoStringGogovanity(this.Int32, "int32")+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGogovanity(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringGogovanity(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *B) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *B) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.String_ != nil {
		data[i] = 0xa
		i++
		i = encodeVarintGogovanity(data, i, uint64(len(*m.String_)))
		i += copy(data[i:], *m.String_)
	}
	data[i] = 0x10
	i++
	i = encodeVarintGogovanity(data, i, uint64(m.Int64))
	if m.Int32 != nil {
		data[i] = 0x18
		i++
		i = encodeVarintGogovanity(data, i, uint64(*m.Int32))
	}
	return i, nil
}

func encodeFixed64Gogovanity(data []byte, offset int, v uint64) int {
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
func encodeFixed32Gogovanity(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGogovanity(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *B) Size() (n int) {
	var l int
	_ = l
	if m.String_ != nil {
		l = len(*m.String_)
		n += 1 + l + sovGogovanity(uint64(l))
	}
	n += 1 + sovGogovanity(uint64(m.Int64))
	if m.Int32 != nil {
		n += 1 + sovGogovanity(uint64(*m.Int32))
	}
	return n
}

func sovGogovanity(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGogovanity(x uint64) (n int) {
	return sovGogovanity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *B) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&B{`,
		`String_:` + valueToStringGogovanity(this.String_) + `,`,
		`Int64:` + fmt.Sprintf("%v", this.Int64) + `,`,
		`Int32:` + valueToStringGogovanity(this.Int32) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGogovanity(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *B) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGogovanity
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
			return fmt.Errorf("proto: B: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: B: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field String_", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogovanity
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
				return ErrInvalidLengthGogovanity
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.String_ = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64", wireType)
			}
			m.Int64 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogovanity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Int64 |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogovanity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Int32 = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGogovanity(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGogovanity
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
func skipGogovanity(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGogovanity
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
					return 0, ErrIntOverflowGogovanity
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
					return 0, ErrIntOverflowGogovanity
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
				return 0, ErrInvalidLengthGogovanity
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGogovanity
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
				next, err := skipGogovanity(data[start:])
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
	ErrInvalidLengthGogovanity = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGogovanity   = fmt.Errorf("proto: integer overflow")
)
