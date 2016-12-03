/*
Copyright 2016 The Kubernetes Authors.

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

// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package testgroup

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg2_api "k8s.io/kubernetes/pkg/api"
	pkg1_v1 "k8s.io/kubernetes/pkg/apis/meta/v1"
	pkg3_types "k8s.io/kubernetes/pkg/types"
	"reflect"
	"runtime"
	time "time"
)

const (
	// ----- content types ----
	codecSelferC_UTF81234 = 1
	codecSelferC_RAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234 = 10
	codecSelferValueTypeMap1234   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey1234    = 2
	codecSelfer_containerMapValue1234  = 3
	codecSelfer_containerMapEnd1234    = 4
	codecSelfer_containerArrayElem1234 = 6
	codecSelfer_containerArrayEnd1234  = 7
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg2_api.ObjectMeta
		var v1 pkg1_v1.TypeMeta
		var v2 pkg3_types.UID
		var v3 time.Time
		_, _, _, _ = v0, v1, v2, v3
	}
}

func (x *TestType) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [4]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = true
			yyq2[3] = true
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(4)
			} else {
				yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yy10 := &x.ObjectMeta
					yy10.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy11 := &x.ObjectMeta
					yy11.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[3] {
					yy13 := &x.Status
					yy13.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("status"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy14 := &x.Status
					yy14.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *TestType) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym15 := z.DecBinary()
	_ = yym15
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct16 := r.ContainerType()
		if yyct16 == codecSelferValueTypeMap1234 {
			yyl16 := r.ReadMapStart()
			if yyl16 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl16, d)
			}
		} else if yyct16 == codecSelferValueTypeArray1234 {
			yyl16 := r.ReadArrayStart()
			if yyl16 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl16, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TestType) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys17Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys17Slc
	var yyhl17 bool = l >= 0
	for yyj17 := 0; ; yyj17++ {
		if yyhl17 {
			if yyj17 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys17Slc = r.DecodeBytes(yys17Slc, true, true)
		yys17 := string(yys17Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys17 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ObjectMeta = pkg2_api.ObjectMeta{}
			} else {
				yyv20 := &x.ObjectMeta
				yyv20.CodecDecodeSelf(d)
			}
		case "status":
			if r.TryDecodeAsNil() {
				x.Status = TestTypeStatus{}
			} else {
				yyv21 := &x.Status
				yyv21.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys17)
		} // end switch yys17
	} // end for yyj17
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TestType) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj22 int
	var yyb22 bool
	var yyhl22 bool = l >= 0
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_api.ObjectMeta{}
	} else {
		yyv25 := &x.ObjectMeta
		yyv25.CodecDecodeSelf(d)
	}
	yyj22++
	if yyhl22 {
		yyb22 = yyj22 > l
	} else {
		yyb22 = r.CheckBreak()
	}
	if yyb22 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Status = TestTypeStatus{}
	} else {
		yyv26 := &x.Status
		yyv26.CodecDecodeSelf(d)
	}
	for {
		yyj22++
		if yyhl22 {
			yyb22 = yyj22 > l
		} else {
			yyb22 = r.CheckBreak()
		}
		if yyb22 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj22-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *TestTypeList) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym27 := z.EncBinary()
		_ = yym27
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep28 := !z.EncBinary()
			yy2arr28 := z.EncBasicHandle().StructToArray
			var yyq28 [4]bool
			_, _, _ = yysep28, yyq28, yy2arr28
			const yyr28 bool = false
			yyq28[0] = x.Kind != ""
			yyq28[1] = x.APIVersion != ""
			yyq28[2] = true
			var yynn28 int
			if yyr28 || yy2arr28 {
				r.EncodeArrayStart(4)
			} else {
				yynn28 = 1
				for _, b := range yyq28 {
					if b {
						yynn28++
					}
				}
				r.EncodeMapStart(yynn28)
				yynn28 = 0
			}
			if yyr28 || yy2arr28 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq28[0] {
					yym30 := z.EncBinary()
					_ = yym30
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq28[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym31 := z.EncBinary()
					_ = yym31
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr28 || yy2arr28 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq28[1] {
					yym33 := z.EncBinary()
					_ = yym33
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq28[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym34 := z.EncBinary()
					_ = yym34
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr28 || yy2arr28 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq28[2] {
					yy36 := &x.ListMeta
					yym37 := z.EncBinary()
					_ = yym37
					if false {
					} else if z.HasExtensions() && z.EncExt(yy36) {
					} else {
						z.EncFallback(yy36)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq28[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy38 := &x.ListMeta
					yym39 := z.EncBinary()
					_ = yym39
					if false {
					} else if z.HasExtensions() && z.EncExt(yy38) {
					} else {
						z.EncFallback(yy38)
					}
				}
			}
			if yyr28 || yy2arr28 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym41 := z.EncBinary()
					_ = yym41
					if false {
					} else {
						h.encSliceTestType(([]TestType)(x.Items), e)
					}
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("items"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym42 := z.EncBinary()
					_ = yym42
					if false {
					} else {
						h.encSliceTestType(([]TestType)(x.Items), e)
					}
				}
			}
			if yyr28 || yy2arr28 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *TestTypeList) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym43 := z.DecBinary()
	_ = yym43
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct44 := r.ContainerType()
		if yyct44 == codecSelferValueTypeMap1234 {
			yyl44 := r.ReadMapStart()
			if yyl44 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl44, d)
			}
		} else if yyct44 == codecSelferValueTypeArray1234 {
			yyl44 := r.ReadArrayStart()
			if yyl44 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl44, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TestTypeList) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys45Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys45Slc
	var yyhl45 bool = l >= 0
	for yyj45 := 0; ; yyj45++ {
		if yyhl45 {
			if yyj45 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys45Slc = r.DecodeBytes(yys45Slc, true, true)
		yys45 := string(yys45Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys45 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ListMeta = pkg1_v1.ListMeta{}
			} else {
				yyv48 := &x.ListMeta
				yym49 := z.DecBinary()
				_ = yym49
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv48) {
				} else {
					z.DecFallback(yyv48, false)
				}
			}
		case "items":
			if r.TryDecodeAsNil() {
				x.Items = nil
			} else {
				yyv50 := &x.Items
				yym51 := z.DecBinary()
				_ = yym51
				if false {
				} else {
					h.decSliceTestType((*[]TestType)(yyv50), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys45)
		} // end switch yys45
	} // end for yyj45
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TestTypeList) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj52 int
	var yyb52 bool
	var yyhl52 bool = l >= 0
	yyj52++
	if yyhl52 {
		yyb52 = yyj52 > l
	} else {
		yyb52 = r.CheckBreak()
	}
	if yyb52 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj52++
	if yyhl52 {
		yyb52 = yyj52 > l
	} else {
		yyb52 = r.CheckBreak()
	}
	if yyb52 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj52++
	if yyhl52 {
		yyb52 = yyj52 > l
	} else {
		yyb52 = r.CheckBreak()
	}
	if yyb52 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ListMeta = pkg1_v1.ListMeta{}
	} else {
		yyv55 := &x.ListMeta
		yym56 := z.DecBinary()
		_ = yym56
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv55) {
		} else {
			z.DecFallback(yyv55, false)
		}
	}
	yyj52++
	if yyhl52 {
		yyb52 = yyj52 > l
	} else {
		yyb52 = r.CheckBreak()
	}
	if yyb52 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Items = nil
	} else {
		yyv57 := &x.Items
		yym58 := z.DecBinary()
		_ = yym58
		if false {
		} else {
			h.decSliceTestType((*[]TestType)(yyv57), d)
		}
	}
	for {
		yyj52++
		if yyhl52 {
			yyb52 = yyj52 > l
		} else {
			yyb52 = r.CheckBreak()
		}
		if yyb52 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj52-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *TestTypeStatus) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym59 := z.EncBinary()
		_ = yym59
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep60 := !z.EncBinary()
			yy2arr60 := z.EncBasicHandle().StructToArray
			var yyq60 [1]bool
			_, _, _ = yysep60, yyq60, yy2arr60
			const yyr60 bool = false
			var yynn60 int
			if yyr60 || yy2arr60 {
				r.EncodeArrayStart(1)
			} else {
				yynn60 = 1
				for _, b := range yyq60 {
					if b {
						yynn60++
					}
				}
				r.EncodeMapStart(yynn60)
				yynn60 = 0
			}
			if yyr60 || yy2arr60 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym62 := z.EncBinary()
				_ = yym62
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Blah))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("Blah"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym63 := z.EncBinary()
				_ = yym63
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Blah))
				}
			}
			if yyr60 || yy2arr60 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *TestTypeStatus) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym64 := z.DecBinary()
	_ = yym64
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct65 := r.ContainerType()
		if yyct65 == codecSelferValueTypeMap1234 {
			yyl65 := r.ReadMapStart()
			if yyl65 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl65, d)
			}
		} else if yyct65 == codecSelferValueTypeArray1234 {
			yyl65 := r.ReadArrayStart()
			if yyl65 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl65, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TestTypeStatus) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys66Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys66Slc
	var yyhl66 bool = l >= 0
	for yyj66 := 0; ; yyj66++ {
		if yyhl66 {
			if yyj66 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys66Slc = r.DecodeBytes(yys66Slc, true, true)
		yys66 := string(yys66Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys66 {
		case "Blah":
			if r.TryDecodeAsNil() {
				x.Blah = ""
			} else {
				x.Blah = string(r.DecodeString())
			}
		default:
			z.DecStructFieldNotFound(-1, yys66)
		} // end switch yys66
	} // end for yyj66
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TestTypeStatus) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj68 int
	var yyb68 bool
	var yyhl68 bool = l >= 0
	yyj68++
	if yyhl68 {
		yyb68 = yyj68 > l
	} else {
		yyb68 = r.CheckBreak()
	}
	if yyb68 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Blah = ""
	} else {
		x.Blah = string(r.DecodeString())
	}
	for {
		yyj68++
		if yyhl68 {
			yyb68 = yyj68 > l
		} else {
			yyb68 = r.CheckBreak()
		}
		if yyb68 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj68-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) encSliceTestType(v []TestType, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv70 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yy71 := &yyv70
		yy71.CodecEncodeSelf(e)
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decSliceTestType(v *[]TestType, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv72 := *v
	yyh72, yyl72 := z.DecSliceHelperStart()
	var yyc72 bool
	if yyl72 == 0 {
		if yyv72 == nil {
			yyv72 = []TestType{}
			yyc72 = true
		} else if len(yyv72) != 0 {
			yyv72 = yyv72[:0]
			yyc72 = true
		}
	} else if yyl72 > 0 {
		var yyrr72, yyrl72 int
		var yyrt72 bool
		if yyl72 > cap(yyv72) {

			yyrg72 := len(yyv72) > 0
			yyv272 := yyv72
			yyrl72, yyrt72 = z.DecInferLen(yyl72, z.DecBasicHandle().MaxInitLen, 272)
			if yyrt72 {
				if yyrl72 <= cap(yyv72) {
					yyv72 = yyv72[:yyrl72]
				} else {
					yyv72 = make([]TestType, yyrl72)
				}
			} else {
				yyv72 = make([]TestType, yyrl72)
			}
			yyc72 = true
			yyrr72 = len(yyv72)
			if yyrg72 {
				copy(yyv72, yyv272)
			}
		} else if yyl72 != len(yyv72) {
			yyv72 = yyv72[:yyl72]
			yyc72 = true
		}
		yyj72 := 0
		for ; yyj72 < yyrr72; yyj72++ {
			yyh72.ElemContainerState(yyj72)
			if r.TryDecodeAsNil() {
				yyv72[yyj72] = TestType{}
			} else {
				yyv73 := &yyv72[yyj72]
				yyv73.CodecDecodeSelf(d)
			}

		}
		if yyrt72 {
			for ; yyj72 < yyl72; yyj72++ {
				yyv72 = append(yyv72, TestType{})
				yyh72.ElemContainerState(yyj72)
				if r.TryDecodeAsNil() {
					yyv72[yyj72] = TestType{}
				} else {
					yyv74 := &yyv72[yyj72]
					yyv74.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj72 := 0
		for ; !r.CheckBreak(); yyj72++ {

			if yyj72 >= len(yyv72) {
				yyv72 = append(yyv72, TestType{}) // var yyz72 TestType
				yyc72 = true
			}
			yyh72.ElemContainerState(yyj72)
			if yyj72 < len(yyv72) {
				if r.TryDecodeAsNil() {
					yyv72[yyj72] = TestType{}
				} else {
					yyv75 := &yyv72[yyj72]
					yyv75.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj72 < len(yyv72) {
			yyv72 = yyv72[:yyj72]
			yyc72 = true
		} else if yyj72 == 0 && yyv72 == nil {
			yyv72 = []TestType{}
			yyc72 = true
		}
	}
	yyh72.End()
	if yyc72 {
		*v = yyv72
	}
}
