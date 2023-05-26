// Code generated by "stringer -type RoundingMode"; DO NOT EDIT.

package number

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ToNearestEven-0]
	_ = x[ToNearestZero-1]
	_ = x[ToNearestAway-2]
	_ = x[ToPositiveInf-3]
	_ = x[ToNegativeInf-4]
	_ = x[ToZero-5]
	_ = x[AwayFromZero-6]
	_ = x[numModes-7]
}

const _RoundingMode_name = "ToNearestEvenToNearestZeroToNearestAwayToPositiveInfToNegativeInfToZeroAwayFromZeronumModes"

var _RoundingMode_index = [...]uint8{0, 13, 26, 39, 52, 65, 71, 83, 91}

func (i RoundingMode) String() string {
	if i >= RoundingMode(len(_RoundingMode_index)-1) {
		return "RoundingMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RoundingMode_name[_RoundingMode_index[i]:_RoundingMode_index[i+1]]
}
