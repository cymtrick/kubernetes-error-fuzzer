// Code generated by "stringer -type Type"; DO NOT EDIT.

package metricdata

import "strconv"

const _Type_name = "TypeGaugeInt64TypeGaugeFloat64TypeGaugeDistributionTypeCumulativeInt64TypeCumulativeFloat64TypeCumulativeDistributionTypeSummary"

var _Type_index = [...]uint8{0, 14, 30, 51, 70, 91, 117, 128}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
