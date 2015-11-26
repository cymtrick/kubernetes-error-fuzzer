/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package testing

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg2_api "k8s.io/kubernetes/pkg/api"
	pkg1_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
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
		var v1 pkg1_unversioned.TypeMeta
		var v2 pkg3_types.UID
		var v3 time.Time
		_, _, _, _ = v0, v1, v2, v3
	}
}

func (x *Simple) CodecEncodeSelf(e *codec1978.Encoder) {
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
			yyq2[3] = x.Other != ""
			yyq2[4] = len(x.Labels) != 0
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
				yy10 := &x.ObjectMeta
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else if z.HasExtensions() && z.EncExt(yy10) {
				} else {
					z.EncFallback(yy10)
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("metadata"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy12 := &x.ObjectMeta
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else if z.HasExtensions() && z.EncExt(yy12) {
				} else {
					z.EncFallback(yy12)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[3] {
					yym15 := z.EncBinary()
					_ = yym15
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Other))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("other"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym16 := z.EncBinary()
					_ = yym16
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Other))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[4] {
					if x.Labels == nil {
						r.EncodeNil()
					} else {
						yym18 := z.EncBinary()
						_ = yym18
						if false {
						} else {
							z.F.EncMapStringStringV(x.Labels, false, e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[4] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("labels"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Labels == nil {
						r.EncodeNil()
					} else {
						yym19 := z.EncBinary()
						_ = yym19
						if false {
						} else {
							z.F.EncMapStringStringV(x.Labels, false, e)
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

func (x *Simple) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym20 := z.DecBinary()
	_ = yym20
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct21 := r.ContainerType()
		if yyct21 == codecSelferValueTypeMap1234 {
			yyl21 := r.ReadMapStart()
			if yyl21 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl21, d)
			}
		} else if yyct21 == codecSelferValueTypeArray1234 {
			yyl21 := r.ReadArrayStart()
			if yyl21 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl21, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *Simple) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys22Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys22Slc
	var yyhl22 bool = l >= 0
	for yyj22 := 0; ; yyj22++ {
		if yyhl22 {
			if yyj22 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys22Slc = r.DecodeBytes(yys22Slc, true, true)
		yys22 := string(yys22Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys22 {
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
				yyv25 := &x.ObjectMeta
				yym26 := z.DecBinary()
				_ = yym26
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv25) {
				} else {
					z.DecFallback(yyv25, false)
				}
			}
		case "other":
			if r.TryDecodeAsNil() {
				x.Other = ""
			} else {
				x.Other = string(r.DecodeString())
			}
		case "labels":
			if r.TryDecodeAsNil() {
				x.Labels = nil
			} else {
				yyv28 := &x.Labels
				yym29 := z.DecBinary()
				_ = yym29
				if false {
				} else {
					z.F.DecMapStringStringX(yyv28, false, d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys22)
		} // end switch yys22
	} // end for yyj22
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *Simple) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj30 int
	var yyb30 bool
	var yyhl30 bool = l >= 0
	yyj30++
	if yyhl30 {
		yyb30 = yyj30 > l
	} else {
		yyb30 = r.CheckBreak()
	}
	if yyb30 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj30++
	if yyhl30 {
		yyb30 = yyj30 > l
	} else {
		yyb30 = r.CheckBreak()
	}
	if yyb30 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj30++
	if yyhl30 {
		yyb30 = yyj30 > l
	} else {
		yyb30 = r.CheckBreak()
	}
	if yyb30 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_api.ObjectMeta{}
	} else {
		yyv33 := &x.ObjectMeta
		yym34 := z.DecBinary()
		_ = yym34
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv33) {
		} else {
			z.DecFallback(yyv33, false)
		}
	}
	yyj30++
	if yyhl30 {
		yyb30 = yyj30 > l
	} else {
		yyb30 = r.CheckBreak()
	}
	if yyb30 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Other = ""
	} else {
		x.Other = string(r.DecodeString())
	}
	yyj30++
	if yyhl30 {
		yyb30 = yyj30 > l
	} else {
		yyb30 = r.CheckBreak()
	}
	if yyb30 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Labels = nil
	} else {
		yyv36 := &x.Labels
		yym37 := z.DecBinary()
		_ = yym37
		if false {
		} else {
			z.F.DecMapStringStringX(yyv36, false, d)
		}
	}
	for {
		yyj30++
		if yyhl30 {
			yyb30 = yyj30 > l
		} else {
			yyb30 = r.CheckBreak()
		}
		if yyb30 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj30-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *SimpleRoot) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym38 := z.EncBinary()
		_ = yym38
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep39 := !z.EncBinary()
			yy2arr39 := z.EncBasicHandle().StructToArray
			var yyq39 [5]bool
			_, _, _ = yysep39, yyq39, yy2arr39
			const yyr39 bool = false
			yyq39[0] = x.Kind != ""
			yyq39[1] = x.APIVersion != ""
			yyq39[3] = x.Other != ""
			yyq39[4] = len(x.Labels) != 0
			var yynn39 int
			if yyr39 || yy2arr39 {
				r.EncodeArrayStart(5)
			} else {
				yynn39 = 1
				for _, b := range yyq39 {
					if b {
						yynn39++
					}
				}
				r.EncodeMapStart(yynn39)
				yynn39 = 0
			}
			if yyr39 || yy2arr39 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq39[0] {
					yym41 := z.EncBinary()
					_ = yym41
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq39[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym42 := z.EncBinary()
					_ = yym42
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr39 || yy2arr39 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq39[1] {
					yym44 := z.EncBinary()
					_ = yym44
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq39[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym45 := z.EncBinary()
					_ = yym45
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr39 || yy2arr39 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yy47 := &x.ObjectMeta
				yym48 := z.EncBinary()
				_ = yym48
				if false {
				} else if z.HasExtensions() && z.EncExt(yy47) {
				} else {
					z.EncFallback(yy47)
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("metadata"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy49 := &x.ObjectMeta
				yym50 := z.EncBinary()
				_ = yym50
				if false {
				} else if z.HasExtensions() && z.EncExt(yy49) {
				} else {
					z.EncFallback(yy49)
				}
			}
			if yyr39 || yy2arr39 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq39[3] {
					yym52 := z.EncBinary()
					_ = yym52
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Other))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq39[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("other"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym53 := z.EncBinary()
					_ = yym53
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Other))
					}
				}
			}
			if yyr39 || yy2arr39 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq39[4] {
					if x.Labels == nil {
						r.EncodeNil()
					} else {
						yym55 := z.EncBinary()
						_ = yym55
						if false {
						} else {
							z.F.EncMapStringStringV(x.Labels, false, e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq39[4] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("labels"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Labels == nil {
						r.EncodeNil()
					} else {
						yym56 := z.EncBinary()
						_ = yym56
						if false {
						} else {
							z.F.EncMapStringStringV(x.Labels, false, e)
						}
					}
				}
			}
			if yyr39 || yy2arr39 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *SimpleRoot) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym57 := z.DecBinary()
	_ = yym57
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct58 := r.ContainerType()
		if yyct58 == codecSelferValueTypeMap1234 {
			yyl58 := r.ReadMapStart()
			if yyl58 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl58, d)
			}
		} else if yyct58 == codecSelferValueTypeArray1234 {
			yyl58 := r.ReadArrayStart()
			if yyl58 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl58, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *SimpleRoot) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys59Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys59Slc
	var yyhl59 bool = l >= 0
	for yyj59 := 0; ; yyj59++ {
		if yyhl59 {
			if yyj59 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys59Slc = r.DecodeBytes(yys59Slc, true, true)
		yys59 := string(yys59Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys59 {
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
				yyv62 := &x.ObjectMeta
				yym63 := z.DecBinary()
				_ = yym63
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv62) {
				} else {
					z.DecFallback(yyv62, false)
				}
			}
		case "other":
			if r.TryDecodeAsNil() {
				x.Other = ""
			} else {
				x.Other = string(r.DecodeString())
			}
		case "labels":
			if r.TryDecodeAsNil() {
				x.Labels = nil
			} else {
				yyv65 := &x.Labels
				yym66 := z.DecBinary()
				_ = yym66
				if false {
				} else {
					z.F.DecMapStringStringX(yyv65, false, d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys59)
		} // end switch yys59
	} // end for yyj59
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *SimpleRoot) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj67 int
	var yyb67 bool
	var yyhl67 bool = l >= 0
	yyj67++
	if yyhl67 {
		yyb67 = yyj67 > l
	} else {
		yyb67 = r.CheckBreak()
	}
	if yyb67 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj67++
	if yyhl67 {
		yyb67 = yyj67 > l
	} else {
		yyb67 = r.CheckBreak()
	}
	if yyb67 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj67++
	if yyhl67 {
		yyb67 = yyj67 > l
	} else {
		yyb67 = r.CheckBreak()
	}
	if yyb67 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_api.ObjectMeta{}
	} else {
		yyv70 := &x.ObjectMeta
		yym71 := z.DecBinary()
		_ = yym71
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv70) {
		} else {
			z.DecFallback(yyv70, false)
		}
	}
	yyj67++
	if yyhl67 {
		yyb67 = yyj67 > l
	} else {
		yyb67 = r.CheckBreak()
	}
	if yyb67 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Other = ""
	} else {
		x.Other = string(r.DecodeString())
	}
	yyj67++
	if yyhl67 {
		yyb67 = yyj67 > l
	} else {
		yyb67 = r.CheckBreak()
	}
	if yyb67 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Labels = nil
	} else {
		yyv73 := &x.Labels
		yym74 := z.DecBinary()
		_ = yym74
		if false {
		} else {
			z.F.DecMapStringStringX(yyv73, false, d)
		}
	}
	for {
		yyj67++
		if yyhl67 {
			yyb67 = yyj67 > l
		} else {
			yyb67 = r.CheckBreak()
		}
		if yyb67 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj67-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *SimpleGetOptions) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym75 := z.EncBinary()
		_ = yym75
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep76 := !z.EncBinary()
			yy2arr76 := z.EncBasicHandle().StructToArray
			var yyq76 [5]bool
			_, _, _ = yysep76, yyq76, yy2arr76
			const yyr76 bool = false
			yyq76[0] = x.Kind != ""
			yyq76[1] = x.APIVersion != ""
			var yynn76 int
			if yyr76 || yy2arr76 {
				r.EncodeArrayStart(5)
			} else {
				yynn76 = 3
				for _, b := range yyq76 {
					if b {
						yynn76++
					}
				}
				r.EncodeMapStart(yynn76)
				yynn76 = 0
			}
			if yyr76 || yy2arr76 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq76[0] {
					yym78 := z.EncBinary()
					_ = yym78
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq76[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym79 := z.EncBinary()
					_ = yym79
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr76 || yy2arr76 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq76[1] {
					yym81 := z.EncBinary()
					_ = yym81
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq76[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym82 := z.EncBinary()
					_ = yym82
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr76 || yy2arr76 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym84 := z.EncBinary()
				_ = yym84
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Param1))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("param1"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym85 := z.EncBinary()
				_ = yym85
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Param1))
				}
			}
			if yyr76 || yy2arr76 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym87 := z.EncBinary()
				_ = yym87
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Param2))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("param2"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym88 := z.EncBinary()
				_ = yym88
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Param2))
				}
			}
			if yyr76 || yy2arr76 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym90 := z.EncBinary()
				_ = yym90
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Path))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("atAPath"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym91 := z.EncBinary()
				_ = yym91
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Path))
				}
			}
			if yyr76 || yy2arr76 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *SimpleGetOptions) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym92 := z.DecBinary()
	_ = yym92
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct93 := r.ContainerType()
		if yyct93 == codecSelferValueTypeMap1234 {
			yyl93 := r.ReadMapStart()
			if yyl93 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl93, d)
			}
		} else if yyct93 == codecSelferValueTypeArray1234 {
			yyl93 := r.ReadArrayStart()
			if yyl93 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl93, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *SimpleGetOptions) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys94Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys94Slc
	var yyhl94 bool = l >= 0
	for yyj94 := 0; ; yyj94++ {
		if yyhl94 {
			if yyj94 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys94Slc = r.DecodeBytes(yys94Slc, true, true)
		yys94 := string(yys94Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys94 {
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
		case "param1":
			if r.TryDecodeAsNil() {
				x.Param1 = ""
			} else {
				x.Param1 = string(r.DecodeString())
			}
		case "param2":
			if r.TryDecodeAsNil() {
				x.Param2 = ""
			} else {
				x.Param2 = string(r.DecodeString())
			}
		case "atAPath":
			if r.TryDecodeAsNil() {
				x.Path = ""
			} else {
				x.Path = string(r.DecodeString())
			}
		default:
			z.DecStructFieldNotFound(-1, yys94)
		} // end switch yys94
	} // end for yyj94
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *SimpleGetOptions) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj100 int
	var yyb100 bool
	var yyhl100 bool = l >= 0
	yyj100++
	if yyhl100 {
		yyb100 = yyj100 > l
	} else {
		yyb100 = r.CheckBreak()
	}
	if yyb100 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj100++
	if yyhl100 {
		yyb100 = yyj100 > l
	} else {
		yyb100 = r.CheckBreak()
	}
	if yyb100 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj100++
	if yyhl100 {
		yyb100 = yyj100 > l
	} else {
		yyb100 = r.CheckBreak()
	}
	if yyb100 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Param1 = ""
	} else {
		x.Param1 = string(r.DecodeString())
	}
	yyj100++
	if yyhl100 {
		yyb100 = yyj100 > l
	} else {
		yyb100 = r.CheckBreak()
	}
	if yyb100 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Param2 = ""
	} else {
		x.Param2 = string(r.DecodeString())
	}
	yyj100++
	if yyhl100 {
		yyb100 = yyj100 > l
	} else {
		yyb100 = r.CheckBreak()
	}
	if yyb100 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Path = ""
	} else {
		x.Path = string(r.DecodeString())
	}
	for {
		yyj100++
		if yyhl100 {
			yyb100 = yyj100 > l
		} else {
			yyb100 = r.CheckBreak()
		}
		if yyb100 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj100-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *SimpleList) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym106 := z.EncBinary()
		_ = yym106
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep107 := !z.EncBinary()
			yy2arr107 := z.EncBasicHandle().StructToArray
			var yyq107 [4]bool
			_, _, _ = yysep107, yyq107, yy2arr107
			const yyr107 bool = false
			yyq107[0] = x.Kind != ""
			yyq107[1] = x.APIVersion != ""
			yyq107[3] = len(x.Items) != 0
			var yynn107 int
			if yyr107 || yy2arr107 {
				r.EncodeArrayStart(4)
			} else {
				yynn107 = 1
				for _, b := range yyq107 {
					if b {
						yynn107++
					}
				}
				r.EncodeMapStart(yynn107)
				yynn107 = 0
			}
			if yyr107 || yy2arr107 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq107[0] {
					yym109 := z.EncBinary()
					_ = yym109
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq107[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym110 := z.EncBinary()
					_ = yym110
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr107 || yy2arr107 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq107[1] {
					yym112 := z.EncBinary()
					_ = yym112
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq107[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym113 := z.EncBinary()
					_ = yym113
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr107 || yy2arr107 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yy115 := &x.ListMeta
				yym116 := z.EncBinary()
				_ = yym116
				if false {
				} else if z.HasExtensions() && z.EncExt(yy115) {
				} else {
					z.EncFallback(yy115)
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("metadata"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy117 := &x.ListMeta
				yym118 := z.EncBinary()
				_ = yym118
				if false {
				} else if z.HasExtensions() && z.EncExt(yy117) {
				} else {
					z.EncFallback(yy117)
				}
			}
			if yyr107 || yy2arr107 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq107[3] {
					if x.Items == nil {
						r.EncodeNil()
					} else {
						yym120 := z.EncBinary()
						_ = yym120
						if false {
						} else {
							h.encSliceSimple(([]Simple)(x.Items), e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq107[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("items"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Items == nil {
						r.EncodeNil()
					} else {
						yym121 := z.EncBinary()
						_ = yym121
						if false {
						} else {
							h.encSliceSimple(([]Simple)(x.Items), e)
						}
					}
				}
			}
			if yyr107 || yy2arr107 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *SimpleList) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym122 := z.DecBinary()
	_ = yym122
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct123 := r.ContainerType()
		if yyct123 == codecSelferValueTypeMap1234 {
			yyl123 := r.ReadMapStart()
			if yyl123 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl123, d)
			}
		} else if yyct123 == codecSelferValueTypeArray1234 {
			yyl123 := r.ReadArrayStart()
			if yyl123 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl123, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *SimpleList) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys124Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys124Slc
	var yyhl124 bool = l >= 0
	for yyj124 := 0; ; yyj124++ {
		if yyhl124 {
			if yyj124 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys124Slc = r.DecodeBytes(yys124Slc, true, true)
		yys124 := string(yys124Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys124 {
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
				x.ListMeta = pkg1_unversioned.ListMeta{}
			} else {
				yyv127 := &x.ListMeta
				yym128 := z.DecBinary()
				_ = yym128
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv127) {
				} else {
					z.DecFallback(yyv127, false)
				}
			}
		case "items":
			if r.TryDecodeAsNil() {
				x.Items = nil
			} else {
				yyv129 := &x.Items
				yym130 := z.DecBinary()
				_ = yym130
				if false {
				} else {
					h.decSliceSimple((*[]Simple)(yyv129), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys124)
		} // end switch yys124
	} // end for yyj124
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *SimpleList) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj131 int
	var yyb131 bool
	var yyhl131 bool = l >= 0
	yyj131++
	if yyhl131 {
		yyb131 = yyj131 > l
	} else {
		yyb131 = r.CheckBreak()
	}
	if yyb131 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj131++
	if yyhl131 {
		yyb131 = yyj131 > l
	} else {
		yyb131 = r.CheckBreak()
	}
	if yyb131 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj131++
	if yyhl131 {
		yyb131 = yyj131 > l
	} else {
		yyb131 = r.CheckBreak()
	}
	if yyb131 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ListMeta = pkg1_unversioned.ListMeta{}
	} else {
		yyv134 := &x.ListMeta
		yym135 := z.DecBinary()
		_ = yym135
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv134) {
		} else {
			z.DecFallback(yyv134, false)
		}
	}
	yyj131++
	if yyhl131 {
		yyb131 = yyj131 > l
	} else {
		yyb131 = r.CheckBreak()
	}
	if yyb131 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Items = nil
	} else {
		yyv136 := &x.Items
		yym137 := z.DecBinary()
		_ = yym137
		if false {
		} else {
			h.decSliceSimple((*[]Simple)(yyv136), d)
		}
	}
	for {
		yyj131++
		if yyhl131 {
			yyb131 = yyj131 > l
		} else {
			yyb131 = r.CheckBreak()
		}
		if yyb131 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj131-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) encSliceSimple(v []Simple, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv138 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yy139 := &yyv138
		yy139.CodecEncodeSelf(e)
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decSliceSimple(v *[]Simple, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv140 := *v
	yyh140, yyl140 := z.DecSliceHelperStart()
	var yyc140 bool
	if yyl140 == 0 {
		if yyv140 == nil {
			yyv140 = []Simple{}
			yyc140 = true
		} else if len(yyv140) != 0 {
			yyv140 = yyv140[:0]
			yyc140 = true
		}
	} else if yyl140 > 0 {
		var yyrr140, yyrl140 int
		var yyrt140 bool
		if yyl140 > cap(yyv140) {

			yyrg140 := len(yyv140) > 0
			yyv2140 := yyv140
			yyrl140, yyrt140 = z.DecInferLen(yyl140, z.DecBasicHandle().MaxInitLen, 216)
			if yyrt140 {
				if yyrl140 <= cap(yyv140) {
					yyv140 = yyv140[:yyrl140]
				} else {
					yyv140 = make([]Simple, yyrl140)
				}
			} else {
				yyv140 = make([]Simple, yyrl140)
			}
			yyc140 = true
			yyrr140 = len(yyv140)
			if yyrg140 {
				copy(yyv140, yyv2140)
			}
		} else if yyl140 != len(yyv140) {
			yyv140 = yyv140[:yyl140]
			yyc140 = true
		}
		yyj140 := 0
		for ; yyj140 < yyrr140; yyj140++ {
			yyh140.ElemContainerState(yyj140)
			if r.TryDecodeAsNil() {
				yyv140[yyj140] = Simple{}
			} else {
				yyv141 := &yyv140[yyj140]
				yyv141.CodecDecodeSelf(d)
			}

		}
		if yyrt140 {
			for ; yyj140 < yyl140; yyj140++ {
				yyv140 = append(yyv140, Simple{})
				yyh140.ElemContainerState(yyj140)
				if r.TryDecodeAsNil() {
					yyv140[yyj140] = Simple{}
				} else {
					yyv142 := &yyv140[yyj140]
					yyv142.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj140 := 0
		for ; !r.CheckBreak(); yyj140++ {

			if yyj140 >= len(yyv140) {
				yyv140 = append(yyv140, Simple{}) // var yyz140 Simple
				yyc140 = true
			}
			yyh140.ElemContainerState(yyj140)
			if yyj140 < len(yyv140) {
				if r.TryDecodeAsNil() {
					yyv140[yyj140] = Simple{}
				} else {
					yyv143 := &yyv140[yyj140]
					yyv143.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj140 < len(yyv140) {
			yyv140 = yyv140[:yyj140]
			yyc140 = true
		} else if yyj140 == 0 && yyv140 == nil {
			yyv140 = []Simple{}
			yyc140 = true
		}
	}
	yyh140.End()
	if yyc140 {
		*v = yyv140
	}
}
