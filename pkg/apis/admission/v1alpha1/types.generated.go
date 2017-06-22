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

package v1alpha1

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg4_v1 "k8s.io/api/authentication/v1"
	pkg1_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkg2_runtime "k8s.io/apimachinery/pkg/runtime"
	pkg5_types "k8s.io/apimachinery/pkg/types"
	pkg3_admission "k8s.io/apiserver/pkg/admission"
	"reflect"
	"runtime"
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
		var v0 pkg4_v1.UserInfo
		var v1 pkg1_v1.TypeMeta
		var v2 pkg2_runtime.RawExtension
		var v3 pkg5_types.UID
		var v4 pkg3_admission.Operation
		_, _, _, _, _ = v0, v1, v2, v3, v4
	}
}

func (x *AdmissionReview) CodecEncodeSelf(e *codec1978.Encoder) {
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
					yy10 := &x.Spec
					yy10.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("spec"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy12 := &x.Spec
					yy12.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[3] {
					yy15 := &x.Status
					yy15.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("status"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy17 := &x.Status
					yy17.CodecEncodeSelf(e)
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

func (x *AdmissionReview) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *AdmissionReview) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				yyv4 := &x.Kind
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				yyv6 := &x.APIVersion
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "spec":
			if r.TryDecodeAsNil() {
				x.Spec = AdmissionReviewSpec{}
			} else {
				yyv8 := &x.Spec
				yyv8.CodecDecodeSelf(d)
			}
		case "status":
			if r.TryDecodeAsNil() {
				x.Status = AdmissionReviewStatus{}
			} else {
				yyv9 := &x.Status
				yyv9.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *AdmissionReview) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj10 int
	var yyb10 bool
	var yyhl10 bool = l >= 0
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		yyv11 := &x.Kind
		yym12 := z.DecBinary()
		_ = yym12
		if false {
		} else {
			*((*string)(yyv11)) = r.DecodeString()
		}
	}
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		yyv13 := &x.APIVersion
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} else {
			*((*string)(yyv13)) = r.DecodeString()
		}
	}
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Spec = AdmissionReviewSpec{}
	} else {
		yyv15 := &x.Spec
		yyv15.CodecDecodeSelf(d)
	}
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Status = AdmissionReviewStatus{}
	} else {
		yyv16 := &x.Status
		yyv16.CodecDecodeSelf(d)
	}
	for {
		yyj10++
		if yyhl10 {
			yyb10 = yyj10 > l
		} else {
			yyb10 = r.CheckBreak()
		}
		if yyb10 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj10-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *AdmissionReviewSpec) CodecEncodeSelf(e *codec1978.Encoder) {
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
			var yyq2 [9]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = true
			yyq2[1] = true
			yyq2[2] = true
			yyq2[3] = x.Operation != ""
			yyq2[4] = x.Name != ""
			yyq2[5] = x.Namespace != ""
			yyq2[6] = true
			yyq2[7] = x.SubResource != ""
			yyq2[8] = true
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(9)
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
					yy4 := &x.Kind
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else if z.HasExtensions() && z.EncExt(yy4) {
					} else {
						z.EncFallback(yy4)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy6 := &x.Kind
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else if z.HasExtensions() && z.EncExt(yy6) {
					} else {
						z.EncFallback(yy6)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yy9 := &x.Object
					yym10 := z.EncBinary()
					_ = yym10
					if false {
					} else if z.HasExtensions() && z.EncExt(yy9) {
					} else if !yym10 && z.IsJSONHandle() {
						z.EncJSONMarshal(yy9)
					} else {
						z.EncFallback(yy9)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("object"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy11 := &x.Object
					yym12 := z.EncBinary()
					_ = yym12
					if false {
					} else if z.HasExtensions() && z.EncExt(yy11) {
					} else if !yym12 && z.IsJSONHandle() {
						z.EncJSONMarshal(yy11)
					} else {
						z.EncFallback(yy11)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yy14 := &x.OldObject
					yym15 := z.EncBinary()
					_ = yym15
					if false {
					} else if z.HasExtensions() && z.EncExt(yy14) {
					} else if !yym15 && z.IsJSONHandle() {
						z.EncJSONMarshal(yy14)
					} else {
						z.EncFallback(yy14)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("oldObject"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy16 := &x.OldObject
					yym17 := z.EncBinary()
					_ = yym17
					if false {
					} else if z.HasExtensions() && z.EncExt(yy16) {
					} else if !yym17 && z.IsJSONHandle() {
						z.EncJSONMarshal(yy16)
					} else {
						z.EncFallback(yy16)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[3] {
					yym19 := z.EncBinary()
					_ = yym19
					if false {
					} else if z.HasExtensions() && z.EncExt(x.Operation) {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Operation))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("operation"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym20 := z.EncBinary()
					_ = yym20
					if false {
					} else if z.HasExtensions() && z.EncExt(x.Operation) {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Operation))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[4] {
					yym22 := z.EncBinary()
					_ = yym22
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Name))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[4] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("name"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym23 := z.EncBinary()
					_ = yym23
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Name))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[5] {
					yym25 := z.EncBinary()
					_ = yym25
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Namespace))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[5] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("namespace"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym26 := z.EncBinary()
					_ = yym26
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Namespace))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[6] {
					yy28 := &x.Resource
					yym29 := z.EncBinary()
					_ = yym29
					if false {
					} else if z.HasExtensions() && z.EncExt(yy28) {
					} else {
						z.EncFallback(yy28)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[6] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("resource"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy30 := &x.Resource
					yym31 := z.EncBinary()
					_ = yym31
					if false {
					} else if z.HasExtensions() && z.EncExt(yy30) {
					} else {
						z.EncFallback(yy30)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[7] {
					yym33 := z.EncBinary()
					_ = yym33
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.SubResource))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[7] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("subResource"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym34 := z.EncBinary()
					_ = yym34
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.SubResource))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[8] {
					yy36 := &x.UserInfo
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
				if yyq2[8] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("userInfo"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy38 := &x.UserInfo
					yym39 := z.EncBinary()
					_ = yym39
					if false {
					} else if z.HasExtensions() && z.EncExt(yy38) {
					} else {
						z.EncFallback(yy38)
					}
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

func (x *AdmissionReviewSpec) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *AdmissionReviewSpec) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = pkg1_v1.GroupVersionKind{}
			} else {
				yyv4 := &x.Kind
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv4) {
				} else {
					z.DecFallback(yyv4, false)
				}
			}
		case "object":
			if r.TryDecodeAsNil() {
				x.Object = pkg2_runtime.RawExtension{}
			} else {
				yyv6 := &x.Object
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv6) {
				} else if !yym7 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(yyv6)
				} else {
					z.DecFallback(yyv6, false)
				}
			}
		case "oldObject":
			if r.TryDecodeAsNil() {
				x.OldObject = pkg2_runtime.RawExtension{}
			} else {
				yyv8 := &x.OldObject
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv8) {
				} else if !yym9 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(yyv8)
				} else {
					z.DecFallback(yyv8, false)
				}
			}
		case "operation":
			if r.TryDecodeAsNil() {
				x.Operation = ""
			} else {
				yyv10 := &x.Operation
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv10) {
				} else {
					*((*string)(yyv10)) = r.DecodeString()
				}
			}
		case "name":
			if r.TryDecodeAsNil() {
				x.Name = ""
			} else {
				yyv12 := &x.Name
				yym13 := z.DecBinary()
				_ = yym13
				if false {
				} else {
					*((*string)(yyv12)) = r.DecodeString()
				}
			}
		case "namespace":
			if r.TryDecodeAsNil() {
				x.Namespace = ""
			} else {
				yyv14 := &x.Namespace
				yym15 := z.DecBinary()
				_ = yym15
				if false {
				} else {
					*((*string)(yyv14)) = r.DecodeString()
				}
			}
		case "resource":
			if r.TryDecodeAsNil() {
				x.Resource = pkg1_v1.GroupVersionResource{}
			} else {
				yyv16 := &x.Resource
				yym17 := z.DecBinary()
				_ = yym17
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv16) {
				} else {
					z.DecFallback(yyv16, false)
				}
			}
		case "subResource":
			if r.TryDecodeAsNil() {
				x.SubResource = ""
			} else {
				yyv18 := &x.SubResource
				yym19 := z.DecBinary()
				_ = yym19
				if false {
				} else {
					*((*string)(yyv18)) = r.DecodeString()
				}
			}
		case "userInfo":
			if r.TryDecodeAsNil() {
				x.UserInfo = pkg4_v1.UserInfo{}
			} else {
				yyv20 := &x.UserInfo
				yym21 := z.DecBinary()
				_ = yym21
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv20) {
				} else {
					z.DecFallback(yyv20, false)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *AdmissionReviewSpec) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
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
		x.Kind = pkg1_v1.GroupVersionKind{}
	} else {
		yyv23 := &x.Kind
		yym24 := z.DecBinary()
		_ = yym24
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv23) {
		} else {
			z.DecFallback(yyv23, false)
		}
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
		x.Object = pkg2_runtime.RawExtension{}
	} else {
		yyv25 := &x.Object
		yym26 := z.DecBinary()
		_ = yym26
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv25) {
		} else if !yym26 && z.IsJSONHandle() {
			z.DecJSONUnmarshal(yyv25)
		} else {
			z.DecFallback(yyv25, false)
		}
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
		x.OldObject = pkg2_runtime.RawExtension{}
	} else {
		yyv27 := &x.OldObject
		yym28 := z.DecBinary()
		_ = yym28
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv27) {
		} else if !yym28 && z.IsJSONHandle() {
			z.DecJSONUnmarshal(yyv27)
		} else {
			z.DecFallback(yyv27, false)
		}
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
		x.Operation = ""
	} else {
		yyv29 := &x.Operation
		yym30 := z.DecBinary()
		_ = yym30
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv29) {
		} else {
			*((*string)(yyv29)) = r.DecodeString()
		}
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
		x.Name = ""
	} else {
		yyv31 := &x.Name
		yym32 := z.DecBinary()
		_ = yym32
		if false {
		} else {
			*((*string)(yyv31)) = r.DecodeString()
		}
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
		x.Namespace = ""
	} else {
		yyv33 := &x.Namespace
		yym34 := z.DecBinary()
		_ = yym34
		if false {
		} else {
			*((*string)(yyv33)) = r.DecodeString()
		}
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
		x.Resource = pkg1_v1.GroupVersionResource{}
	} else {
		yyv35 := &x.Resource
		yym36 := z.DecBinary()
		_ = yym36
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv35) {
		} else {
			z.DecFallback(yyv35, false)
		}
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
		x.SubResource = ""
	} else {
		yyv37 := &x.SubResource
		yym38 := z.DecBinary()
		_ = yym38
		if false {
		} else {
			*((*string)(yyv37)) = r.DecodeString()
		}
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
		x.UserInfo = pkg4_v1.UserInfo{}
	} else {
		yyv39 := &x.UserInfo
		yym40 := z.DecBinary()
		_ = yym40
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv39) {
		} else {
			z.DecFallback(yyv39, false)
		}
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

func (x *AdmissionReviewStatus) CodecEncodeSelf(e *codec1978.Encoder) {
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
			var yyq2 [2]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[1] = x.Result != nil
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(2)
			} else {
				yynn2 = 1
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
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeBool(bool(x.Allowed))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("allowed"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeBool(bool(x.Allowed))
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					if x.Result == nil {
						r.EncodeNil()
					} else {
						yym7 := z.EncBinary()
						_ = yym7
						if false {
						} else if z.HasExtensions() && z.EncExt(x.Result) {
						} else {
							z.EncFallback(x.Result)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("status"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Result == nil {
						r.EncodeNil()
					} else {
						yym8 := z.EncBinary()
						_ = yym8
						if false {
						} else if z.HasExtensions() && z.EncExt(x.Result) {
						} else {
							z.EncFallback(x.Result)
						}
					}
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

func (x *AdmissionReviewStatus) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *AdmissionReviewStatus) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "allowed":
			if r.TryDecodeAsNil() {
				x.Allowed = false
			} else {
				yyv4 := &x.Allowed
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*bool)(yyv4)) = r.DecodeBool()
				}
			}
		case "status":
			if r.TryDecodeAsNil() {
				if x.Result != nil {
					x.Result = nil
				}
			} else {
				if x.Result == nil {
					x.Result = new(pkg1_v1.Status)
				}
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else if z.HasExtensions() && z.DecExt(x.Result) {
				} else {
					z.DecFallback(x.Result, false)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *AdmissionReviewStatus) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj8 int
	var yyb8 bool
	var yyhl8 bool = l >= 0
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Allowed = false
	} else {
		yyv9 := &x.Allowed
		yym10 := z.DecBinary()
		_ = yym10
		if false {
		} else {
			*((*bool)(yyv9)) = r.DecodeBool()
		}
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		if x.Result != nil {
			x.Result = nil
		}
	} else {
		if x.Result == nil {
			x.Result = new(pkg1_v1.Status)
		}
		yym12 := z.DecBinary()
		_ = yym12
		if false {
		} else if z.HasExtensions() && z.DecExt(x.Result) {
		} else {
			z.DecFallback(x.Result, false)
		}
	}
	for {
		yyj8++
		if yyhl8 {
			yyb8 = yyj8 > l
		} else {
			yyb8 = r.CheckBreak()
		}
		if yyb8 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj8-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}
