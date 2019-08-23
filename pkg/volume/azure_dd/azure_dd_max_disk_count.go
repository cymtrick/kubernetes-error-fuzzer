// +build !providerless

/*
Copyright 2019 The Kubernetes Authors.

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

package azure_dd

// about how to get all VM size list,
// refer to https://github.com/kubernetes/kubernetes/issues/77461#issuecomment-492488756
var maxDataDiskCountMap = map[string]int64{
	"BASIC_A0":               1,
	"BASIC_A1":               2,
	"BASIC_A2":               4,
	"BASIC_A3":               8,
	"BASIC_A4":               16,
	"STANDARD_A0":            1,
	"STANDARD_A10":           32,
	"STANDARD_A11":           64,
	"STANDARD_A1":            2,
	"STANDARD_A1_V2":         2,
	"STANDARD_A2":            4,
	"STANDARD_A2M_V2":        4,
	"STANDARD_A2_V2":         4,
	"STANDARD_A3":            8,
	"STANDARD_A4":            16,
	"STANDARD_A4M_V2":        8,
	"STANDARD_A4_V2":         8,
	"STANDARD_A5":            4,
	"STANDARD_A6":            8,
	"STANDARD_A7":            16,
	"STANDARD_A8":            32,
	"STANDARD_A8M_V2":        16,
	"STANDARD_A8_V2":         16,
	"STANDARD_A9":            64,
	"STANDARD_B1LS":          2,
	"STANDARD_B1MS":          2,
	"STANDARD_B1S":           2,
	"STANDARD_B2MS":          4,
	"STANDARD_B2S":           4,
	"STANDARD_B4MS":          8,
	"STANDARD_B8MS":          16,
	"STANDARD_D11":           8,
	"STANDARD_D11_V2":        8,
	"STANDARD_D11_V2_PROMO":  8,
	"STANDARD_D12":           16,
	"STANDARD_D12_V2":        16,
	"STANDARD_D12_V2_PROMO":  16,
	"STANDARD_D13":           32,
	"STANDARD_D13_V2":        32,
	"STANDARD_D13_V2_PROMO":  32,
	"STANDARD_D1":            4,
	"STANDARD_D14":           64,
	"STANDARD_D14_V2":        64,
	"STANDARD_D14_V2_PROMO":  64,
	"STANDARD_D15_V2":        64,
	"STANDARD_D16S_V3":       32,
	"STANDARD_D16_V3":        32,
	"STANDARD_D1_V2":         4,
	"STANDARD_D2":            8,
	"STANDARD_D2S_V3":        4,
	"STANDARD_D2_V2":         8,
	"STANDARD_D2_V2_PROMO":   8,
	"STANDARD_D2_V3":         4,
	"STANDARD_D3":            16,
	"STANDARD_D32S_V3":       32,
	"STANDARD_D32_V3":        32,
	"STANDARD_D3_V2":         16,
	"STANDARD_D3_V2_PROMO":   16,
	"STANDARD_D4":            32,
	"STANDARD_D4S_V3":        8,
	"STANDARD_D4_V2":         32,
	"STANDARD_D4_V2_PROMO":   32,
	"STANDARD_D4_V3":         8,
	"STANDARD_D5_V2":         64,
	"STANDARD_D5_V2_PROMO":   64,
	"STANDARD_D64S_V3":       32,
	"STANDARD_D64_V3":        32,
	"STANDARD_D8S_V3":        16,
	"STANDARD_D8_V3":         16,
	"STANDARD_DC2S":          2,
	"STANDARD_DC4S":          4,
	"STANDARD_DS11-1_V2":     8,
	"STANDARD_DS11":          8,
	"STANDARD_DS11_V2":       8,
	"STANDARD_DS11_V2_PROMO": 8,
	"STANDARD_DS12":          16,
	"STANDARD_DS12-1_V2":     16,
	"STANDARD_DS12-2_V2":     16,
	"STANDARD_DS12_V2":       16,
	"STANDARD_DS12_V2_PROMO": 16,
	"STANDARD_DS13-2_V2":     32,
	"STANDARD_DS13":          32,
	"STANDARD_DS13-4_V2":     32,
	"STANDARD_DS13_V2":       32,
	"STANDARD_DS13_V2_PROMO": 32,
	"STANDARD_DS1":           4,
	"STANDARD_DS14-4_V2":     64,
	"STANDARD_DS14":          64,
	"STANDARD_DS14-8_V2":     64,
	"STANDARD_DS14_V2":       64,
	"STANDARD_DS14_V2_PROMO": 64,
	"STANDARD_DS15_V2":       64,
	"STANDARD_DS1_V2":        4,
	"STANDARD_DS2":           8,
	"STANDARD_DS2_V2":        8,
	"STANDARD_DS2_V2_PROMO":  8,
	"STANDARD_DS3":           16,
	"STANDARD_DS3_V2":        16,
	"STANDARD_DS3_V2_PROMO":  16,
	"STANDARD_DS4":           32,
	"STANDARD_DS4_V2":        32,
	"STANDARD_DS4_V2_PROMO":  32,
	"STANDARD_DS5_V2":        64,
	"STANDARD_DS5_V2_PROMO":  64,
	"STANDARD_E16-4S_V3":     32,
	"STANDARD_E16-8S_V3":     32,
	"STANDARD_E16S_V3":       32,
	"STANDARD_E16_V3":        32,
	"STANDARD_E20S_V3":       32,
	"STANDARD_E20_V3":        32,
	"STANDARD_E2S_V3":        4,
	"STANDARD_E2_V3":         4,
	"STANDARD_E32-16S_V3":    32,
	"STANDARD_E32-8S_V3":     32,
	"STANDARD_E32S_V3":       32,
	"STANDARD_E32_V3":        32,
	"STANDARD_E4-2S_V3":      8,
	"STANDARD_E4S_V3":        8,
	"STANDARD_E4_V3":         8,
	"STANDARD_E64-16S_V3":    32,
	"STANDARD_E64-32S_V3":    32,
	"STANDARD_E64IS_V3":      32,
	"STANDARD_E64I_V3":       32,
	"STANDARD_E64S_V3":       32,
	"STANDARD_E64_V3":        32,
	"STANDARD_E8-2S_V3":      16,
	"STANDARD_E8-4S_V3":      16,
	"STANDARD_E8S_V3":        16,
	"STANDARD_E8_V3":         16,
	"STANDARD_F1":            4,
	"STANDARD_F16":           64,
	"STANDARD_F16S":          64,
	"STANDARD_F16S_V2":       32,
	"STANDARD_F1S":           4,
	"STANDARD_F2":            8,
	"STANDARD_F2S":           8,
	"STANDARD_F2S_V2":        4,
	"STANDARD_F32S_V2":       32,
	"STANDARD_F4":            16,
	"STANDARD_F4S":           16,
	"STANDARD_F4S_V2":        8,
	"STANDARD_F64S_V2":       32,
	"STANDARD_F72S_V2":       32,
	"STANDARD_F8":            32,
	"STANDARD_F8S":           32,
	"STANDARD_F8S_V2":        16,
	"STANDARD_G1":            8,
	"STANDARD_G2":            16,
	"STANDARD_G3":            32,
	"STANDARD_G4":            64,
	"STANDARD_G5":            64,
	"STANDARD_GS1":           8,
	"STANDARD_GS2":           16,
	"STANDARD_GS3":           32,
	"STANDARD_GS4-4":         64,
	"STANDARD_GS4":           64,
	"STANDARD_GS4-8":         64,
	"STANDARD_GS5-16":        64,
	"STANDARD_GS5":           64,
	"STANDARD_GS5-8":         64,
	"STANDARD_H16":           64,
	"STANDARD_H16M":          64,
	"STANDARD_H16M_PROMO":    64,
	"STANDARD_H16MR":         64,
	"STANDARD_H16MR_PROMO":   64,
	"STANDARD_H16_PROMO":     64,
	"STANDARD_H16R":          64,
	"STANDARD_H16R_PROMO":    64,
	"STANDARD_H8":            32,
	"STANDARD_H8M":           32,
	"STANDARD_H8M_PROMO":     32,
	"STANDARD_H8_PROMO":      32,
	"STANDARD_HB60RS":        4,
	"STANDARD_HC44RS":        4,
	"STANDARD_L16S":          64,
	"STANDARD_L16S_V2":       32,
	"STANDARD_L32S":          64,
	"STANDARD_L32S_V2":       32,
	"STANDARD_L4S":           16,
	"STANDARD_L64S_V2":       32,
	"STANDARD_L80S_V2":       32,
	"STANDARD_L8S":           32,
	"STANDARD_L8S_V2":        16,
	"STANDARD_M128-32MS":     64,
	"STANDARD_M128":          64,
	"STANDARD_M128-64MS":     64,
	"STANDARD_M128M":         64,
	"STANDARD_M128MS":        64,
	"STANDARD_M128S":         64,
	"STANDARD_M16-4MS":       16,
	"STANDARD_M16-8MS":       16,
	"STANDARD_M16MS":         16,
	"STANDARD_M208MS_V2":     64,
	"STANDARD_M208S_V2":      64,
	"STANDARD_M32-16MS":      32,
	"STANDARD_M32-8MS":       32,
	"STANDARD_M32LS":         32,
	"STANDARD_M32MS":         32,
	"STANDARD_M32TS":         32,
	"STANDARD_M64-16MS":      64,
	"STANDARD_M64-32MS":      64,
	"STANDARD_M64":           64,
	"STANDARD_M64LS":         64,
	"STANDARD_M64M":          64,
	"STANDARD_M64MS":         64,
	"STANDARD_M64S":          64,
	"STANDARD_M8-2MS":        8,
	"STANDARD_M8-4MS":        8,
	"STANDARD_M8MS":          8,
	"STANDARD_NC12":          48,
	"STANDARD_NC12_PROMO":    48,
	"STANDARD_NC12S_V2":      24,
	"STANDARD_NC12S_V3":      24,
	"STANDARD_NC24":          64,
	"STANDARD_NC24_PROMO":    64,
	"STANDARD_NC24R":         64,
	"STANDARD_NC24R_PROMO":   64,
	"STANDARD_NC24RS_V2":     32,
	"STANDARD_NC24RS_V3":     32,
	"STANDARD_NC24S_V2":      32,
	"STANDARD_NC24S_V3":      32,
	"STANDARD_NC6":           24,
	"STANDARD_NC6_PROMO":     24,
	"STANDARD_NC6S_V2":       12,
	"STANDARD_NC6S_V3":       12,
	"STANDARD_ND12S":         24,
	"STANDARD_ND24RS":        32,
	"STANDARD_ND24S":         32,
	"STANDARD_ND6S":          12,
	"STANDARD_NV12":          48,
	"STANDARD_NV12_PROMO":    48,
	"STANDARD_NV12S_V2":      24,
	"STANDARD_NV12S_V3":      12,
	"STANDARD_NV24":          64,
	"STANDARD_NV24_PROMO":    64,
	"STANDARD_NV24S_V2":      32,
	"STANDARD_NV24S_V3":      24,
	"STANDARD_NV48S_V3":      32,
	"STANDARD_NV6":           24,
	"STANDARD_NV6_PROMO":     24,
	"STANDARD_NV6S_V2":       12,
	"STANDARD_PB6S":          12,
}
