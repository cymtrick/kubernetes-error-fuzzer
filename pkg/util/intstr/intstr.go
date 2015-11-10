/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package intstr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/google/gofuzz"
)

// IntOrString is a type that can hold an int or a string.  When used in
// JSON or YAML marshalling and unmarshalling, it produces or consumes the
// inner type.  This allows you to have, for example, a JSON field that can
// accept a name or number.
type IntOrString struct {
	Type   Type
	IntVal int
	StrVal string
}

// Type represents the stored type of IntOrString.
type Type int

const (
	Int    Type = iota // The IntOrString holds an int.
	String             // The IntOrString holds a string.
)

// FromInt creates an IntOrString object with an int value.
func FromInt(val int) IntOrString {
	return IntOrString{Type: Int, IntVal: val}
}

// FromString creates an IntOrString object with a string value.
func FromString(val string) IntOrString {
	return IntOrString{Type: String, StrVal: val}
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (intstr *IntOrString) UnmarshalJSON(value []byte) error {
	if value[0] == '"' {
		intstr.Type = String
		return json.Unmarshal(value, &intstr.StrVal)
	}
	intstr.Type = Int
	return json.Unmarshal(value, &intstr.IntVal)
}

// String returns the string value, or the Itoa of the int value.
func (intstr *IntOrString) String() string {
	if intstr.Type == String {
		return intstr.StrVal
	}
	return strconv.Itoa(intstr.IntVal)
}

// MarshalJSON implements the json.Marshaller interface.
func (intstr IntOrString) MarshalJSON() ([]byte, error) {
	switch intstr.Type {
	case Int:
		return json.Marshal(intstr.IntVal)
	case String:
		return json.Marshal(intstr.StrVal)
	default:
		return []byte{}, fmt.Errorf("impossible IntOrString.Type")
	}
}

func (intstr *IntOrString) Fuzz(c fuzz.Continue) {
	if intstr == nil {
		return
	}
	if c.RandBool() {
		intstr.Type = Int
		c.Fuzz(&intstr.IntVal)
		intstr.StrVal = ""
	} else {
		intstr.Type = String
		intstr.IntVal = 0
		c.Fuzz(&intstr.StrVal)
	}
}
