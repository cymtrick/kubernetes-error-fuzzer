// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate-protos. DO NOT EDIT.

package fieldnum

// Field numbers for google.protobuf.Api.
const (
	Api_Name          = 1 // optional string
	Api_Methods       = 2 // repeated google.protobuf.Method
	Api_Options       = 3 // repeated google.protobuf.Option
	Api_Version       = 4 // optional string
	Api_SourceContext = 5 // optional google.protobuf.SourceContext
	Api_Mixins        = 6 // repeated google.protobuf.Mixin
	Api_Syntax        = 7 // optional google.protobuf.Syntax
)

// Field numbers for google.protobuf.Method.
const (
	Method_Name              = 1 // optional string
	Method_RequestTypeUrl    = 2 // optional string
	Method_RequestStreaming  = 3 // optional bool
	Method_ResponseTypeUrl   = 4 // optional string
	Method_ResponseStreaming = 5 // optional bool
	Method_Options           = 6 // repeated google.protobuf.Option
	Method_Syntax            = 7 // optional google.protobuf.Syntax
)

// Field numbers for google.protobuf.Mixin.
const (
	Mixin_Name = 1 // optional string
	Mixin_Root = 2 // optional string
)
