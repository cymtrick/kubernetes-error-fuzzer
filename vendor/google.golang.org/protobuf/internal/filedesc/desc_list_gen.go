// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate-types. DO NOT EDIT.

package filedesc

import (
	"fmt"
	"sync"

	"google.golang.org/protobuf/internal/descfmt"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Enums struct {
	List   []Enum
	once   sync.Once
	byName map[protoreflect.Name]*Enum // protected by once
}

func (p *Enums) Len() int {
	return len(p.List)
}
func (p *Enums) Get(i int) protoreflect.EnumDescriptor {
	return &p.List[i]
}
func (p *Enums) ByName(s protoreflect.Name) protoreflect.EnumDescriptor {
	if d := p.lazyInit().byName[s]; d != nil {
		return d
	}
	return nil
}
func (p *Enums) Format(s fmt.State, r rune) {
	descfmt.FormatList(s, r, p)
}
func (p *Enums) ProtoInternal(pragma.DoNotImplement) {}
func (p *Enums) lazyInit() *Enums {
	p.once.Do(func() {
		if len(p.List) > 0 {
			p.byName = make(map[protoreflect.Name]*Enum, len(p.List))
			for i := range p.List {
				d := &p.List[i]
				if _, ok := p.byName[d.Name()]; !ok {
					p.byName[d.Name()] = d
				}
			}
		}
	})
	return p
}

type EnumValues struct {
	List   []EnumValue
	once   sync.Once
	byName map[protoreflect.Name]*EnumValue       // protected by once
	byNum  map[protoreflect.EnumNumber]*EnumValue // protected by once
}

func (p *EnumValues) Len() int {
	return len(p.List)
}
func (p *EnumValues) Get(i int) protoreflect.EnumValueDescriptor {
	return &p.List[i]
}
func (p *EnumValues) ByName(s protoreflect.Name) protoreflect.EnumValueDescriptor {
	if d := p.lazyInit().byName[s]; d != nil {
		return d
	}
	return nil
}
func (p *EnumValues) ByNumber(n protoreflect.EnumNumber) protoreflect.EnumValueDescriptor {
	if d := p.lazyInit().byNum[n]; d != nil {
		return d
	}
	return nil
}
func (p *EnumValues) Format(s fmt.State, r rune) {
	descfmt.FormatList(s, r, p)
}
func (p *EnumValues) ProtoInternal(pragma.DoNotImplement) {}
func (p *EnumValues) lazyInit() *EnumValues {
	p.once.Do(func() {
		if len(p.List) > 0 {
			p.byName = make(map[protoreflect.Name]*EnumValue, len(p.List))
			p.byNum = make(map[protoreflect.EnumNumber]*EnumValue, len(p.List))
			for i := range p.List {
				d := &p.List[i]
				if _, ok := p.byName[d.Name()]; !ok {
					p.byName[d.Name()] = d
				}
				if _, ok := p.byNum[d.Number()]; !ok {
					p.byNum[d.Number()] = d
				}
			}
		}
	})
	return p
}

type Messages struct {
	List   []Message
	once   sync.Once
	byName map[protoreflect.Name]*Message // protected by once
}

func (p *Messages) Len() int {
	return len(p.List)
}
func (p *Messages) Get(i int) protoreflect.MessageDescriptor {
	return &p.List[i]
}
func (p *Messages) ByName(s protoreflect.Name) protoreflect.MessageDescriptor {
	if d := p.lazyInit().byName[s]; d != nil {
		return d
	}
	return nil
}
func (p *Messages) Format(s fmt.State, r rune) {
	descfmt.FormatList(s, r, p)
}
func (p *Messages) ProtoInternal(pragma.DoNotImplement) {}
func (p *Messages) lazyInit() *Messages {
	p.once.Do(func() {
		if len(p.List) > 0 {
			p.byName = make(map[protoreflect.Name]*Message, len(p.List))
			for i := range p.List {
				d := &p.List[i]
				if _, ok := p.byName[d.Name()]; !ok {
					p.byName[d.Name()] = d
				}
			}
		}
	})
	return p
}

type Fields struct {
	List   []Field
	once   sync.Once
	byName map[protoreflect.Name]*Field        // protected by once
	byJSON map[string]*Field                   // protected by once
	byNum  map[protoreflect.FieldNumber]*Field // protected by once
}

func (p *Fields) Len() int {
	return len(p.List)
}
func (p *Fields) Get(i int) protoreflect.FieldDescriptor {
	return &p.List[i]
}
func (p *Fields) ByName(s protoreflect.Name) protoreflect.FieldDescriptor {
	if d := p.lazyInit().byName[s]; d != nil {
		return d
	}
	return nil
}
func (p *Fields) ByJSONName(s string) protoreflect.FieldDescriptor {
	if d := p.lazyInit().byJSON[s]; d != nil {
		return d
	}
	return nil
}
func (p *Fields) ByNumber(n protoreflect.FieldNumber) protoreflect.FieldDescriptor {
	if d := p.lazyInit().byNum[n]; d != nil {
		return d
	}
	return nil
}
func (p *Fields) Format(s fmt.State, r rune) {
	descfmt.FormatList(s, r, p)
}
func (p *Fields) ProtoInternal(pragma.DoNotImplement) {}
func (p *Fields) lazyInit() *Fields {
	p.once.Do(func() {
		if len(p.List) > 0 {
			p.byName = make(map[protoreflect.Name]*Field, len(p.List))
			p.byJSON = make(map[string]*Field, len(p.List))
			p.byNum = make(map[protoreflect.FieldNumber]*Field, len(p.List))
			for i := range p.List {
				d := &p.List[i]
				if _, ok := p.byName[d.Name()]; !ok {
					p.byName[d.Name()] = d
				}
				if _, ok := p.byJSON[d.JSONName()]; !ok {
					p.byJSON[d.JSONName()] = d
				}
				if _, ok := p.byNum[d.Number()]; !ok {
					p.byNum[d.Number()] = d
				}
			}
		}
	})
	return p
}

type Oneofs struct {
	List   []Oneof
	once   sync.Once
	byName map[protoreflect.Name]*Oneof // protected by once
}

func (p *Oneofs) Len() int {
	return len(p.List)
}
func (p *Oneofs) Get(i int) protoreflect.OneofDescriptor {
	return &p.List[i]
}
func (p *Oneofs) ByName(s protoreflect.Name) protoreflect.OneofDescriptor {
	if d := p.lazyInit().byName[s]; d != nil {
		return d
	}
	return nil
}
func (p *Oneofs) Format(s fmt.State, r rune) {
	descfmt.FormatList(s, r, p)
}
func (p *Oneofs) ProtoInternal(pragma.DoNotImplement) {}
func (p *Oneofs) lazyInit() *Oneofs {
	p.once.Do(func() {
		if len(p.List) > 0 {
			p.byName = make(map[protoreflect.Name]*Oneof, len(p.List))
			for i := range p.List {
				d := &p.List[i]
				if _, ok := p.byName[d.Name()]; !ok {
					p.byName[d.Name()] = d
				}
			}
		}
	})
	return p
}

type Extensions struct {
	List   []Extension
	once   sync.Once
	byName map[protoreflect.Name]*Extension // protected by once
}

func (p *Extensions) Len() int {
	return len(p.List)
}
func (p *Extensions) Get(i int) protoreflect.ExtensionDescriptor {
	return &p.List[i]
}
func (p *Extensions) ByName(s protoreflect.Name) protoreflect.ExtensionDescriptor {
	if d := p.lazyInit().byName[s]; d != nil {
		return d
	}
	return nil
}
func (p *Extensions) Format(s fmt.State, r rune) {
	descfmt.FormatList(s, r, p)
}
func (p *Extensions) ProtoInternal(pragma.DoNotImplement) {}
func (p *Extensions) lazyInit() *Extensions {
	p.once.Do(func() {
		if len(p.List) > 0 {
			p.byName = make(map[protoreflect.Name]*Extension, len(p.List))
			for i := range p.List {
				d := &p.List[i]
				if _, ok := p.byName[d.Name()]; !ok {
					p.byName[d.Name()] = d
				}
			}
		}
	})
	return p
}

type Services struct {
	List   []Service
	once   sync.Once
	byName map[protoreflect.Name]*Service // protected by once
}

func (p *Services) Len() int {
	return len(p.List)
}
func (p *Services) Get(i int) protoreflect.ServiceDescriptor {
	return &p.List[i]
}
func (p *Services) ByName(s protoreflect.Name) protoreflect.ServiceDescriptor {
	if d := p.lazyInit().byName[s]; d != nil {
		return d
	}
	return nil
}
func (p *Services) Format(s fmt.State, r rune) {
	descfmt.FormatList(s, r, p)
}
func (p *Services) ProtoInternal(pragma.DoNotImplement) {}
func (p *Services) lazyInit() *Services {
	p.once.Do(func() {
		if len(p.List) > 0 {
			p.byName = make(map[protoreflect.Name]*Service, len(p.List))
			for i := range p.List {
				d := &p.List[i]
				if _, ok := p.byName[d.Name()]; !ok {
					p.byName[d.Name()] = d
				}
			}
		}
	})
	return p
}

type Methods struct {
	List   []Method
	once   sync.Once
	byName map[protoreflect.Name]*Method // protected by once
}

func (p *Methods) Len() int {
	return len(p.List)
}
func (p *Methods) Get(i int) protoreflect.MethodDescriptor {
	return &p.List[i]
}
func (p *Methods) ByName(s protoreflect.Name) protoreflect.MethodDescriptor {
	if d := p.lazyInit().byName[s]; d != nil {
		return d
	}
	return nil
}
func (p *Methods) Format(s fmt.State, r rune) {
	descfmt.FormatList(s, r, p)
}
func (p *Methods) ProtoInternal(pragma.DoNotImplement) {}
func (p *Methods) lazyInit() *Methods {
	p.once.Do(func() {
		if len(p.List) > 0 {
			p.byName = make(map[protoreflect.Name]*Method, len(p.List))
			for i := range p.List {
				d := &p.List[i]
				if _, ok := p.byName[d.Name()]; !ok {
					p.byName[d.Name()] = d
				}
			}
		}
	})
	return p
}
