// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

// Size returns the size in bytes of the wire-format encoding of m.
func Size(m Message) int {
	return MarshalOptions{}.Size(m)
}

// Size returns the size in bytes of the wire-format encoding of m.
func (o MarshalOptions) Size(m Message) int {
	// Treat a nil message interface as an empty message; nothing to output.
	if m == nil {
		return 0
	}

	return sizeMessage(m.ProtoReflect())
}

func sizeMessage(m protoreflect.Message) (size int) {
	methods := protoMethods(m)
	if methods != nil && methods.Size != nil {
		out := methods.Size(protoiface.SizeInput{
			Message: m,
		})
		return out.Size
	}
	if methods != nil && methods.Marshal != nil {
		// This is not efficient, but we don't have any choice.
		// This case is mainly used for legacy types with a Marshal method.
		out, _ := methods.Marshal(protoiface.MarshalInput{
			Message: m,
		})
		return len(out.Buf)
	}
	return sizeMessageSlow(m)
}

func sizeMessageSlow(m protoreflect.Message) (size int) {
	if messageset.IsMessageSet(m.Descriptor()) {
		return sizeMessageSet(m)
	}
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		size += sizeField(fd, v)
		return true
	})
	size += len(m.GetUnknown())
	return size
}

func sizeField(fd protoreflect.FieldDescriptor, value protoreflect.Value) (size int) {
	num := fd.Number()
	switch {
	case fd.IsList():
		return sizeList(num, fd, value.List())
	case fd.IsMap():
		return sizeMap(num, fd, value.Map())
	default:
		return protowire.SizeTag(num) + sizeSingular(num, fd.Kind(), value)
	}
}

func sizeList(num protowire.Number, fd protoreflect.FieldDescriptor, list protoreflect.List) (size int) {
	if fd.IsPacked() && list.Len() > 0 {
		content := 0
		for i, llen := 0, list.Len(); i < llen; i++ {
			content += sizeSingular(num, fd.Kind(), list.Get(i))
		}
		return protowire.SizeTag(num) + protowire.SizeBytes(content)
	}

	for i, llen := 0, list.Len(); i < llen; i++ {
		size += protowire.SizeTag(num) + sizeSingular(num, fd.Kind(), list.Get(i))
	}
	return size
}

func sizeMap(num protowire.Number, fd protoreflect.FieldDescriptor, mapv protoreflect.Map) (size int) {
	mapv.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		size += protowire.SizeTag(num)
		size += protowire.SizeBytes(sizeField(fd.MapKey(), key.Value()) + sizeField(fd.MapValue(), value))
		return true
	})
	return size
}
