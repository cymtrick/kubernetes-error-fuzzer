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

package v1beta1

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg1_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkg3_types "k8s.io/apimachinery/pkg/types"
	pkg2_v1 "k8s.io/kubernetes/pkg/api/v1"
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
		var v0 pkg1_v1.TypeMeta
		var v1 pkg3_types.UID
		var v2 pkg2_v1.ObjectMeta
		var v3 time.Time
		_, _, _, _ = v0, v1, v2, v3
	}
}

func (x *TokenReview) CodecEncodeSelf(e *codec1978.Encoder) {
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
			var yyq2 [5]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = true
			yyq2[4] = true
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(5)
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
				yy13 := &x.Spec
				yy13.CodecEncodeSelf(e)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("spec"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy14 := &x.Spec
				yy14.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[4] {
					yy16 := &x.Status
					yy16.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[4] {
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

func (x *TokenReview) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym18 := z.DecBinary()
	_ = yym18
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct19 := r.ContainerType()
		if yyct19 == codecSelferValueTypeMap1234 {
			yyl19 := r.ReadMapStart()
			if yyl19 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl19, d)
			}
		} else if yyct19 == codecSelferValueTypeArray1234 {
			yyl19 := r.ReadArrayStart()
			if yyl19 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl19, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TokenReview) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys20Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys20Slc
	var yyhl20 bool = l >= 0
	for yyj20 := 0; ; yyj20++ {
		if yyhl20 {
			if yyj20 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys20Slc = r.DecodeBytes(yys20Slc, true, true)
		yys20 := string(yys20Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys20 {
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
				x.ObjectMeta = pkg2_v1.ObjectMeta{}
			} else {
				yyv23 := &x.ObjectMeta
				yyv23.CodecDecodeSelf(d)
			}
		case "spec":
			if r.TryDecodeAsNil() {
				x.Spec = TokenReviewSpec{}
			} else {
				yyv24 := &x.Spec
				yyv24.CodecDecodeSelf(d)
			}
		case "status":
			if r.TryDecodeAsNil() {
				x.Status = TokenReviewStatus{}
			} else {
				yyv25 := &x.Status
				yyv25.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys20)
		} // end switch yys20
	} // end for yyj20
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TokenReview) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj26 int
	var yyb26 bool
	var yyhl26 bool = l >= 0
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_v1.ObjectMeta{}
	} else {
		yyv29 := &x.ObjectMeta
		yyv29.CodecDecodeSelf(d)
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Spec = TokenReviewSpec{}
	} else {
		yyv30 := &x.Spec
		yyv30.CodecDecodeSelf(d)
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Status = TokenReviewStatus{}
	} else {
		yyv31 := &x.Status
		yyv31.CodecDecodeSelf(d)
	}
	for {
		yyj26++
		if yyhl26 {
			yyb26 = yyj26 > l
		} else {
			yyb26 = r.CheckBreak()
		}
		if yyb26 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj26-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *TokenReviewSpec) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym32 := z.EncBinary()
		_ = yym32
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep33 := !z.EncBinary()
			yy2arr33 := z.EncBasicHandle().StructToArray
			var yyq33 [1]bool
			_, _, _ = yysep33, yyq33, yy2arr33
			const yyr33 bool = false
			yyq33[0] = x.Token != ""
			var yynn33 int
			if yyr33 || yy2arr33 {
				r.EncodeArrayStart(1)
			} else {
				yynn33 = 0
				for _, b := range yyq33 {
					if b {
						yynn33++
					}
				}
				r.EncodeMapStart(yynn33)
				yynn33 = 0
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq33[0] {
					yym35 := z.EncBinary()
					_ = yym35
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Token))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq33[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("token"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym36 := z.EncBinary()
					_ = yym36
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Token))
					}
				}
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *TokenReviewSpec) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym37 := z.DecBinary()
	_ = yym37
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct38 := r.ContainerType()
		if yyct38 == codecSelferValueTypeMap1234 {
			yyl38 := r.ReadMapStart()
			if yyl38 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl38, d)
			}
		} else if yyct38 == codecSelferValueTypeArray1234 {
			yyl38 := r.ReadArrayStart()
			if yyl38 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl38, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TokenReviewSpec) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys39Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys39Slc
	var yyhl39 bool = l >= 0
	for yyj39 := 0; ; yyj39++ {
		if yyhl39 {
			if yyj39 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys39Slc = r.DecodeBytes(yys39Slc, true, true)
		yys39 := string(yys39Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys39 {
		case "token":
			if r.TryDecodeAsNil() {
				x.Token = ""
			} else {
				x.Token = string(r.DecodeString())
			}
		default:
			z.DecStructFieldNotFound(-1, yys39)
		} // end switch yys39
	} // end for yyj39
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TokenReviewSpec) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj41 int
	var yyb41 bool
	var yyhl41 bool = l >= 0
	yyj41++
	if yyhl41 {
		yyb41 = yyj41 > l
	} else {
		yyb41 = r.CheckBreak()
	}
	if yyb41 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Token = ""
	} else {
		x.Token = string(r.DecodeString())
	}
	for {
		yyj41++
		if yyhl41 {
			yyb41 = yyj41 > l
		} else {
			yyb41 = r.CheckBreak()
		}
		if yyb41 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj41-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *TokenReviewStatus) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym43 := z.EncBinary()
		_ = yym43
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep44 := !z.EncBinary()
			yy2arr44 := z.EncBasicHandle().StructToArray
			var yyq44 [3]bool
			_, _, _ = yysep44, yyq44, yy2arr44
			const yyr44 bool = false
			yyq44[0] = x.Authenticated != false
			yyq44[1] = true
			yyq44[2] = x.Error != ""
			var yynn44 int
			if yyr44 || yy2arr44 {
				r.EncodeArrayStart(3)
			} else {
				yynn44 = 0
				for _, b := range yyq44 {
					if b {
						yynn44++
					}
				}
				r.EncodeMapStart(yynn44)
				yynn44 = 0
			}
			if yyr44 || yy2arr44 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq44[0] {
					yym46 := z.EncBinary()
					_ = yym46
					if false {
					} else {
						r.EncodeBool(bool(x.Authenticated))
					}
				} else {
					r.EncodeBool(false)
				}
			} else {
				if yyq44[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("authenticated"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym47 := z.EncBinary()
					_ = yym47
					if false {
					} else {
						r.EncodeBool(bool(x.Authenticated))
					}
				}
			}
			if yyr44 || yy2arr44 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq44[1] {
					yy49 := &x.User
					yy49.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq44[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("user"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy50 := &x.User
					yy50.CodecEncodeSelf(e)
				}
			}
			if yyr44 || yy2arr44 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq44[2] {
					yym52 := z.EncBinary()
					_ = yym52
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Error))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq44[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("error"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym53 := z.EncBinary()
					_ = yym53
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Error))
					}
				}
			}
			if yyr44 || yy2arr44 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *TokenReviewStatus) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym54 := z.DecBinary()
	_ = yym54
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct55 := r.ContainerType()
		if yyct55 == codecSelferValueTypeMap1234 {
			yyl55 := r.ReadMapStart()
			if yyl55 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl55, d)
			}
		} else if yyct55 == codecSelferValueTypeArray1234 {
			yyl55 := r.ReadArrayStart()
			if yyl55 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl55, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TokenReviewStatus) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys56Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys56Slc
	var yyhl56 bool = l >= 0
	for yyj56 := 0; ; yyj56++ {
		if yyhl56 {
			if yyj56 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys56Slc = r.DecodeBytes(yys56Slc, true, true)
		yys56 := string(yys56Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys56 {
		case "authenticated":
			if r.TryDecodeAsNil() {
				x.Authenticated = false
			} else {
				x.Authenticated = bool(r.DecodeBool())
			}
		case "user":
			if r.TryDecodeAsNil() {
				x.User = UserInfo{}
			} else {
				yyv58 := &x.User
				yyv58.CodecDecodeSelf(d)
			}
		case "error":
			if r.TryDecodeAsNil() {
				x.Error = ""
			} else {
				x.Error = string(r.DecodeString())
			}
		default:
			z.DecStructFieldNotFound(-1, yys56)
		} // end switch yys56
	} // end for yyj56
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TokenReviewStatus) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj60 int
	var yyb60 bool
	var yyhl60 bool = l >= 0
	yyj60++
	if yyhl60 {
		yyb60 = yyj60 > l
	} else {
		yyb60 = r.CheckBreak()
	}
	if yyb60 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Authenticated = false
	} else {
		x.Authenticated = bool(r.DecodeBool())
	}
	yyj60++
	if yyhl60 {
		yyb60 = yyj60 > l
	} else {
		yyb60 = r.CheckBreak()
	}
	if yyb60 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.User = UserInfo{}
	} else {
		yyv62 := &x.User
		yyv62.CodecDecodeSelf(d)
	}
	yyj60++
	if yyhl60 {
		yyb60 = yyj60 > l
	} else {
		yyb60 = r.CheckBreak()
	}
	if yyb60 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Error = ""
	} else {
		x.Error = string(r.DecodeString())
	}
	for {
		yyj60++
		if yyhl60 {
			yyb60 = yyj60 > l
		} else {
			yyb60 = r.CheckBreak()
		}
		if yyb60 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj60-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *UserInfo) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym64 := z.EncBinary()
		_ = yym64
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep65 := !z.EncBinary()
			yy2arr65 := z.EncBasicHandle().StructToArray
			var yyq65 [4]bool
			_, _, _ = yysep65, yyq65, yy2arr65
			const yyr65 bool = false
			yyq65[0] = x.Username != ""
			yyq65[1] = x.UID != ""
			yyq65[2] = len(x.Groups) != 0
			yyq65[3] = len(x.Extra) != 0
			var yynn65 int
			if yyr65 || yy2arr65 {
				r.EncodeArrayStart(4)
			} else {
				yynn65 = 0
				for _, b := range yyq65 {
					if b {
						yynn65++
					}
				}
				r.EncodeMapStart(yynn65)
				yynn65 = 0
			}
			if yyr65 || yy2arr65 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq65[0] {
					yym67 := z.EncBinary()
					_ = yym67
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Username))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq65[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("username"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym68 := z.EncBinary()
					_ = yym68
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Username))
					}
				}
			}
			if yyr65 || yy2arr65 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq65[1] {
					yym70 := z.EncBinary()
					_ = yym70
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.UID))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq65[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("uid"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym71 := z.EncBinary()
					_ = yym71
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.UID))
					}
				}
			}
			if yyr65 || yy2arr65 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq65[2] {
					if x.Groups == nil {
						r.EncodeNil()
					} else {
						yym73 := z.EncBinary()
						_ = yym73
						if false {
						} else {
							z.F.EncSliceStringV(x.Groups, false, e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq65[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("groups"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Groups == nil {
						r.EncodeNil()
					} else {
						yym74 := z.EncBinary()
						_ = yym74
						if false {
						} else {
							z.F.EncSliceStringV(x.Groups, false, e)
						}
					}
				}
			}
			if yyr65 || yy2arr65 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq65[3] {
					if x.Extra == nil {
						r.EncodeNil()
					} else {
						yym76 := z.EncBinary()
						_ = yym76
						if false {
						} else {
							h.encMapstringExtraValue((map[string]ExtraValue)(x.Extra), e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq65[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("extra"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Extra == nil {
						r.EncodeNil()
					} else {
						yym77 := z.EncBinary()
						_ = yym77
						if false {
						} else {
							h.encMapstringExtraValue((map[string]ExtraValue)(x.Extra), e)
						}
					}
				}
			}
			if yyr65 || yy2arr65 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *UserInfo) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym78 := z.DecBinary()
	_ = yym78
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct79 := r.ContainerType()
		if yyct79 == codecSelferValueTypeMap1234 {
			yyl79 := r.ReadMapStart()
			if yyl79 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl79, d)
			}
		} else if yyct79 == codecSelferValueTypeArray1234 {
			yyl79 := r.ReadArrayStart()
			if yyl79 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl79, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *UserInfo) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys80Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys80Slc
	var yyhl80 bool = l >= 0
	for yyj80 := 0; ; yyj80++ {
		if yyhl80 {
			if yyj80 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys80Slc = r.DecodeBytes(yys80Slc, true, true)
		yys80 := string(yys80Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys80 {
		case "username":
			if r.TryDecodeAsNil() {
				x.Username = ""
			} else {
				x.Username = string(r.DecodeString())
			}
		case "uid":
			if r.TryDecodeAsNil() {
				x.UID = ""
			} else {
				x.UID = string(r.DecodeString())
			}
		case "groups":
			if r.TryDecodeAsNil() {
				x.Groups = nil
			} else {
				yyv83 := &x.Groups
				yym84 := z.DecBinary()
				_ = yym84
				if false {
				} else {
					z.F.DecSliceStringX(yyv83, false, d)
				}
			}
		case "extra":
			if r.TryDecodeAsNil() {
				x.Extra = nil
			} else {
				yyv85 := &x.Extra
				yym86 := z.DecBinary()
				_ = yym86
				if false {
				} else {
					h.decMapstringExtraValue((*map[string]ExtraValue)(yyv85), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys80)
		} // end switch yys80
	} // end for yyj80
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *UserInfo) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj87 int
	var yyb87 bool
	var yyhl87 bool = l >= 0
	yyj87++
	if yyhl87 {
		yyb87 = yyj87 > l
	} else {
		yyb87 = r.CheckBreak()
	}
	if yyb87 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Username = ""
	} else {
		x.Username = string(r.DecodeString())
	}
	yyj87++
	if yyhl87 {
		yyb87 = yyj87 > l
	} else {
		yyb87 = r.CheckBreak()
	}
	if yyb87 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.UID = ""
	} else {
		x.UID = string(r.DecodeString())
	}
	yyj87++
	if yyhl87 {
		yyb87 = yyj87 > l
	} else {
		yyb87 = r.CheckBreak()
	}
	if yyb87 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Groups = nil
	} else {
		yyv90 := &x.Groups
		yym91 := z.DecBinary()
		_ = yym91
		if false {
		} else {
			z.F.DecSliceStringX(yyv90, false, d)
		}
	}
	yyj87++
	if yyhl87 {
		yyb87 = yyj87 > l
	} else {
		yyb87 = r.CheckBreak()
	}
	if yyb87 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Extra = nil
	} else {
		yyv92 := &x.Extra
		yym93 := z.DecBinary()
		_ = yym93
		if false {
		} else {
			h.decMapstringExtraValue((*map[string]ExtraValue)(yyv92), d)
		}
	}
	for {
		yyj87++
		if yyhl87 {
			yyb87 = yyj87 > l
		} else {
			yyb87 = r.CheckBreak()
		}
		if yyb87 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj87-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x ExtraValue) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym94 := z.EncBinary()
		_ = yym94
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			h.encExtraValue((ExtraValue)(x), e)
		}
	}
}

func (x *ExtraValue) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym95 := z.DecBinary()
	_ = yym95
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		h.decExtraValue((*ExtraValue)(x), d)
	}
}

func (x codecSelfer1234) encMapstringExtraValue(v map[string]ExtraValue, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeMapStart(len(v))
	for yyk96, yyv96 := range v {
		z.EncSendContainerState(codecSelfer_containerMapKey1234)
		yym97 := z.EncBinary()
		_ = yym97
		if false {
		} else {
			r.EncodeString(codecSelferC_UTF81234, string(yyk96))
		}
		z.EncSendContainerState(codecSelfer_containerMapValue1234)
		if yyv96 == nil {
			r.EncodeNil()
		} else {
			yyv96.CodecEncodeSelf(e)
		}
	}
	z.EncSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x codecSelfer1234) decMapstringExtraValue(v *map[string]ExtraValue, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv98 := *v
	yyl98 := r.ReadMapStart()
	yybh98 := z.DecBasicHandle()
	if yyv98 == nil {
		yyrl98, _ := z.DecInferLen(yyl98, yybh98.MaxInitLen, 40)
		yyv98 = make(map[string]ExtraValue, yyrl98)
		*v = yyv98
	}
	var yymk98 string
	var yymv98 ExtraValue
	var yymg98 bool
	if yybh98.MapValueReset {
		yymg98 = true
	}
	if yyl98 > 0 {
		for yyj98 := 0; yyj98 < yyl98; yyj98++ {
			z.DecSendContainerState(codecSelfer_containerMapKey1234)
			if r.TryDecodeAsNil() {
				yymk98 = ""
			} else {
				yymk98 = string(r.DecodeString())
			}

			if yymg98 {
				yymv98 = yyv98[yymk98]
			} else {
				yymv98 = nil
			}
			z.DecSendContainerState(codecSelfer_containerMapValue1234)
			if r.TryDecodeAsNil() {
				yymv98 = nil
			} else {
				yyv100 := &yymv98
				yyv100.CodecDecodeSelf(d)
			}

			if yyv98 != nil {
				yyv98[yymk98] = yymv98
			}
		}
	} else if yyl98 < 0 {
		for yyj98 := 0; !r.CheckBreak(); yyj98++ {
			z.DecSendContainerState(codecSelfer_containerMapKey1234)
			if r.TryDecodeAsNil() {
				yymk98 = ""
			} else {
				yymk98 = string(r.DecodeString())
			}

			if yymg98 {
				yymv98 = yyv98[yymk98]
			} else {
				yymv98 = nil
			}
			z.DecSendContainerState(codecSelfer_containerMapValue1234)
			if r.TryDecodeAsNil() {
				yymv98 = nil
			} else {
				yyv102 := &yymv98
				yyv102.CodecDecodeSelf(d)
			}

			if yyv98 != nil {
				yyv98[yymk98] = yymv98
			}
		}
	} // else len==0: TODO: Should we clear map entries?
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x codecSelfer1234) encExtraValue(v ExtraValue, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv103 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yym104 := z.EncBinary()
		_ = yym104
		if false {
		} else {
			r.EncodeString(codecSelferC_UTF81234, string(yyv103))
		}
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decExtraValue(v *ExtraValue, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv105 := *v
	yyh105, yyl105 := z.DecSliceHelperStart()
	var yyc105 bool
	if yyl105 == 0 {
		if yyv105 == nil {
			yyv105 = []string{}
			yyc105 = true
		} else if len(yyv105) != 0 {
			yyv105 = yyv105[:0]
			yyc105 = true
		}
	} else if yyl105 > 0 {
		var yyrr105, yyrl105 int
		var yyrt105 bool
		if yyl105 > cap(yyv105) {

			yyrl105, yyrt105 = z.DecInferLen(yyl105, z.DecBasicHandle().MaxInitLen, 16)
			if yyrt105 {
				if yyrl105 <= cap(yyv105) {
					yyv105 = yyv105[:yyrl105]
				} else {
					yyv105 = make([]string, yyrl105)
				}
			} else {
				yyv105 = make([]string, yyrl105)
			}
			yyc105 = true
			yyrr105 = len(yyv105)
		} else if yyl105 != len(yyv105) {
			yyv105 = yyv105[:yyl105]
			yyc105 = true
		}
		yyj105 := 0
		for ; yyj105 < yyrr105; yyj105++ {
			yyh105.ElemContainerState(yyj105)
			if r.TryDecodeAsNil() {
				yyv105[yyj105] = ""
			} else {
				yyv105[yyj105] = string(r.DecodeString())
			}

		}
		if yyrt105 {
			for ; yyj105 < yyl105; yyj105++ {
				yyv105 = append(yyv105, "")
				yyh105.ElemContainerState(yyj105)
				if r.TryDecodeAsNil() {
					yyv105[yyj105] = ""
				} else {
					yyv105[yyj105] = string(r.DecodeString())
				}

			}
		}

	} else {
		yyj105 := 0
		for ; !r.CheckBreak(); yyj105++ {

			if yyj105 >= len(yyv105) {
				yyv105 = append(yyv105, "") // var yyz105 string
				yyc105 = true
			}
			yyh105.ElemContainerState(yyj105)
			if yyj105 < len(yyv105) {
				if r.TryDecodeAsNil() {
					yyv105[yyj105] = ""
				} else {
					yyv105[yyj105] = string(r.DecodeString())
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj105 < len(yyv105) {
			yyv105 = yyv105[:yyj105]
			yyc105 = true
		} else if yyj105 == 0 && yyv105 == nil {
			yyv105 = []string{}
			yyc105 = true
		}
	}
	yyh105.End()
	if yyc105 {
		*v = yyv105
	}
}
