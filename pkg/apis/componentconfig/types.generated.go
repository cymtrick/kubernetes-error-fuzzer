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

package componentconfig

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg1_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	"reflect"
	"runtime"
)

const (
	codecSelferC_UTF81234         = 1
	codecSelferC_RAW1234          = 0
	codecSelferValueTypeArray1234 = 10
	codecSelferValueTypeMap1234   = 9
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 4 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			4, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg1_unversioned.TypeMeta
		_ = v0
	}
}

func (x *KubeProxyConfiguration) CodecEncodeSelf(e *codec1978.Encoder) {
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
			var yyq2 [18]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(18)
			} else {
				var yynn2 int = 16
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
			}
			if yyr2 || yy2arr2 {
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
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
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
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				yym10 := z.EncBinary()
				_ = yym10
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.BindAddress))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("bindAddress"))
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.BindAddress))
				}
			}
			if yyr2 || yy2arr2 {
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else {
					r.EncodeBool(bool(x.CleanupIPTables))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("cleanupIPTables"))
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else {
					r.EncodeBool(bool(x.CleanupIPTables))
				}
			}
			if yyr2 || yy2arr2 {
				yym16 := z.EncBinary()
				_ = yym16
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.HealthzBindAddress))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("healthzBindAddress"))
				yym17 := z.EncBinary()
				_ = yym17
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.HealthzBindAddress))
				}
			}
			if yyr2 || yy2arr2 {
				yym19 := z.EncBinary()
				_ = yym19
				if false {
				} else {
					r.EncodeInt(int64(x.HealthzPort))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("healthzPort"))
				yym20 := z.EncBinary()
				_ = yym20
				if false {
				} else {
					r.EncodeInt(int64(x.HealthzPort))
				}
			}
			if yyr2 || yy2arr2 {
				yym22 := z.EncBinary()
				_ = yym22
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.HostnameOverride))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("hostnameOverride"))
				yym23 := z.EncBinary()
				_ = yym23
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.HostnameOverride))
				}
			}
			if yyr2 || yy2arr2 {
				yym25 := z.EncBinary()
				_ = yym25
				if false {
				} else {
					r.EncodeInt(int64(x.IPTablesSyncePeriodSeconds))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("iptablesSyncPeriodSeconds"))
				yym26 := z.EncBinary()
				_ = yym26
				if false {
				} else {
					r.EncodeInt(int64(x.IPTablesSyncePeriodSeconds))
				}
			}
			if yyr2 || yy2arr2 {
				yym28 := z.EncBinary()
				_ = yym28
				if false {
				} else {
					r.EncodeInt(int64(x.KubeAPIBurst))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("kubeAPIBurst"))
				yym29 := z.EncBinary()
				_ = yym29
				if false {
				} else {
					r.EncodeInt(int64(x.KubeAPIBurst))
				}
			}
			if yyr2 || yy2arr2 {
				yym31 := z.EncBinary()
				_ = yym31
				if false {
				} else {
					r.EncodeInt(int64(x.KubeAPIQPS))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("kubeAPIQPS"))
				yym32 := z.EncBinary()
				_ = yym32
				if false {
				} else {
					r.EncodeInt(int64(x.KubeAPIQPS))
				}
			}
			if yyr2 || yy2arr2 {
				yym34 := z.EncBinary()
				_ = yym34
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.KubeconfigPath))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("kubeconfigPath"))
				yym35 := z.EncBinary()
				_ = yym35
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.KubeconfigPath))
				}
			}
			if yyr2 || yy2arr2 {
				yym37 := z.EncBinary()
				_ = yym37
				if false {
				} else {
					r.EncodeBool(bool(x.MasqueradeAll))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("masqueradeAll"))
				yym38 := z.EncBinary()
				_ = yym38
				if false {
				} else {
					r.EncodeBool(bool(x.MasqueradeAll))
				}
			}
			if yyr2 || yy2arr2 {
				yym40 := z.EncBinary()
				_ = yym40
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Master))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("master"))
				yym41 := z.EncBinary()
				_ = yym41
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.Master))
				}
			}
			if yyr2 || yy2arr2 {
				if x.OOMScoreAdj == nil {
					r.EncodeNil()
				} else {
					yy43 := *x.OOMScoreAdj
					yym44 := z.EncBinary()
					_ = yym44
					if false {
					} else {
						r.EncodeInt(int64(yy43))
					}
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("oomScoreAdj"))
				if x.OOMScoreAdj == nil {
					r.EncodeNil()
				} else {
					yy45 := *x.OOMScoreAdj
					yym46 := z.EncBinary()
					_ = yym46
					if false {
					} else {
						r.EncodeInt(int64(yy45))
					}
				}
			}
			if yyr2 || yy2arr2 {
				x.Mode.CodecEncodeSelf(e)
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("mode"))
				x.Mode.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				yym49 := z.EncBinary()
				_ = yym49
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.PortRange))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("portRange"))
				yym50 := z.EncBinary()
				_ = yym50
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.PortRange))
				}
			}
			if yyr2 || yy2arr2 {
				yym52 := z.EncBinary()
				_ = yym52
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.ResourceContainer))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("resourceContainer"))
				yym53 := z.EncBinary()
				_ = yym53
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.ResourceContainer))
				}
			}
			if yyr2 || yy2arr2 {
				yym55 := z.EncBinary()
				_ = yym55
				if false {
				} else {
					r.EncodeInt(int64(x.UDPTimeoutMilliseconds))
				}
			} else {
				r.EncodeString(codecSelferC_UTF81234, string("udpTimeoutMilliseconds"))
				yym56 := z.EncBinary()
				_ = yym56
				if false {
				} else {
					r.EncodeInt(int64(x.UDPTimeoutMilliseconds))
				}
			}
			if yysep2 {
				r.EncodeEnd()
			}
		}
	}
}

func (x *KubeProxyConfiguration) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym57 := z.DecBinary()
	_ = yym57
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		if r.IsContainerType(codecSelferValueTypeMap1234) {
			yyl58 := r.ReadMapStart()
			if yyl58 == 0 {
				r.ReadEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl58, d)
			}
		} else if r.IsContainerType(codecSelferValueTypeArray1234) {
			yyl58 := r.ReadArrayStart()
			if yyl58 == 0 {
				r.ReadEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl58, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *KubeProxyConfiguration) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
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
		yys59Slc = r.DecodeBytes(yys59Slc, true, true)
		yys59 := string(yys59Slc)
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
		case "bindAddress":
			if r.TryDecodeAsNil() {
				x.BindAddress = ""
			} else {
				x.BindAddress = string(r.DecodeString())
			}
		case "cleanupIPTables":
			if r.TryDecodeAsNil() {
				x.CleanupIPTables = false
			} else {
				x.CleanupIPTables = bool(r.DecodeBool())
			}
		case "healthzBindAddress":
			if r.TryDecodeAsNil() {
				x.HealthzBindAddress = ""
			} else {
				x.HealthzBindAddress = string(r.DecodeString())
			}
		case "healthzPort":
			if r.TryDecodeAsNil() {
				x.HealthzPort = 0
			} else {
				x.HealthzPort = int(r.DecodeInt(codecSelferBitsize1234))
			}
		case "hostnameOverride":
			if r.TryDecodeAsNil() {
				x.HostnameOverride = ""
			} else {
				x.HostnameOverride = string(r.DecodeString())
			}
		case "iptablesSyncPeriodSeconds":
			if r.TryDecodeAsNil() {
				x.IPTablesSyncePeriodSeconds = 0
			} else {
				x.IPTablesSyncePeriodSeconds = int(r.DecodeInt(codecSelferBitsize1234))
			}
		case "kubeAPIBurst":
			if r.TryDecodeAsNil() {
				x.KubeAPIBurst = 0
			} else {
				x.KubeAPIBurst = int(r.DecodeInt(codecSelferBitsize1234))
			}
		case "kubeAPIQPS":
			if r.TryDecodeAsNil() {
				x.KubeAPIQPS = 0
			} else {
				x.KubeAPIQPS = int(r.DecodeInt(codecSelferBitsize1234))
			}
		case "kubeconfigPath":
			if r.TryDecodeAsNil() {
				x.KubeconfigPath = ""
			} else {
				x.KubeconfigPath = string(r.DecodeString())
			}
		case "masqueradeAll":
			if r.TryDecodeAsNil() {
				x.MasqueradeAll = false
			} else {
				x.MasqueradeAll = bool(r.DecodeBool())
			}
		case "master":
			if r.TryDecodeAsNil() {
				x.Master = ""
			} else {
				x.Master = string(r.DecodeString())
			}
		case "oomScoreAdj":
			if r.TryDecodeAsNil() {
				if x.OOMScoreAdj != nil {
					x.OOMScoreAdj = nil
				}
			} else {
				if x.OOMScoreAdj == nil {
					x.OOMScoreAdj = new(int)
				}
				yym74 := z.DecBinary()
				_ = yym74
				if false {
				} else {
					*((*int)(x.OOMScoreAdj)) = int(r.DecodeInt(codecSelferBitsize1234))
				}
			}
		case "mode":
			if r.TryDecodeAsNil() {
				x.Mode = ""
			} else {
				x.Mode = ProxyMode(r.DecodeString())
			}
		case "portRange":
			if r.TryDecodeAsNil() {
				x.PortRange = ""
			} else {
				x.PortRange = string(r.DecodeString())
			}
		case "resourceContainer":
			if r.TryDecodeAsNil() {
				x.ResourceContainer = ""
			} else {
				x.ResourceContainer = string(r.DecodeString())
			}
		case "udpTimeoutMilliseconds":
			if r.TryDecodeAsNil() {
				x.UDPTimeoutMilliseconds = 0
			} else {
				x.UDPTimeoutMilliseconds = int(r.DecodeInt(codecSelferBitsize1234))
			}
		default:
			z.DecStructFieldNotFound(-1, yys59)
		} // end switch yys59
	} // end for yyj59
	if !yyhl59 {
		r.ReadEnd()
	}
}

func (x *KubeProxyConfiguration) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj79 int
	var yyb79 bool
	var yyhl79 bool = l >= 0
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.BindAddress = ""
	} else {
		x.BindAddress = string(r.DecodeString())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.CleanupIPTables = false
	} else {
		x.CleanupIPTables = bool(r.DecodeBool())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.HealthzBindAddress = ""
	} else {
		x.HealthzBindAddress = string(r.DecodeString())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.HealthzPort = 0
	} else {
		x.HealthzPort = int(r.DecodeInt(codecSelferBitsize1234))
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.HostnameOverride = ""
	} else {
		x.HostnameOverride = string(r.DecodeString())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.IPTablesSyncePeriodSeconds = 0
	} else {
		x.IPTablesSyncePeriodSeconds = int(r.DecodeInt(codecSelferBitsize1234))
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.KubeAPIBurst = 0
	} else {
		x.KubeAPIBurst = int(r.DecodeInt(codecSelferBitsize1234))
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.KubeAPIQPS = 0
	} else {
		x.KubeAPIQPS = int(r.DecodeInt(codecSelferBitsize1234))
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.KubeconfigPath = ""
	} else {
		x.KubeconfigPath = string(r.DecodeString())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.MasqueradeAll = false
	} else {
		x.MasqueradeAll = bool(r.DecodeBool())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Master = ""
	} else {
		x.Master = string(r.DecodeString())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		if x.OOMScoreAdj != nil {
			x.OOMScoreAdj = nil
		}
	} else {
		if x.OOMScoreAdj == nil {
			x.OOMScoreAdj = new(int)
		}
		yym94 := z.DecBinary()
		_ = yym94
		if false {
		} else {
			*((*int)(x.OOMScoreAdj)) = int(r.DecodeInt(codecSelferBitsize1234))
		}
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Mode = ""
	} else {
		x.Mode = ProxyMode(r.DecodeString())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.PortRange = ""
	} else {
		x.PortRange = string(r.DecodeString())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.ResourceContainer = ""
	} else {
		x.ResourceContainer = string(r.DecodeString())
	}
	yyj79++
	if yyhl79 {
		yyb79 = yyj79 > l
	} else {
		yyb79 = r.CheckBreak()
	}
	if yyb79 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.UDPTimeoutMilliseconds = 0
	} else {
		x.UDPTimeoutMilliseconds = int(r.DecodeInt(codecSelferBitsize1234))
	}
	for {
		yyj79++
		if yyhl79 {
			yyb79 = yyj79 > l
		} else {
			yyb79 = r.CheckBreak()
		}
		if yyb79 {
			break
		}
		z.DecStructFieldNotFound(yyj79-1, "")
	}
	r.ReadEnd()
}

func (x ProxyMode) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	yym99 := z.EncBinary()
	_ = yym99
	if false {
	} else if z.HasExtensions() && z.EncExt(x) {
	} else {
		r.EncodeString(codecSelferC_UTF81234, string(x))
	}
}

func (x *ProxyMode) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym100 := z.DecBinary()
	_ = yym100
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		*((*string)(x)) = r.DecodeString()
	}
}
