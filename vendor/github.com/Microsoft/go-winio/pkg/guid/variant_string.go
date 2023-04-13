// Code generated by "stringer -type=Variant -trimprefix=Variant -linecomment"; DO NOT EDIT.

package guid

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VariantUnknown-0]
	_ = x[VariantNCS-1]
	_ = x[VariantRFC4122-2]
	_ = x[VariantMicrosoft-3]
	_ = x[VariantFuture-4]
}

const _Variant_name = "UnknownNCSRFC 4122MicrosoftFuture"

var _Variant_index = [...]uint8{0, 7, 10, 18, 27, 33}

func (i Variant) String() string {
	if i >= Variant(len(_Variant_index)-1) {
		return "Variant(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Variant_name[_Variant_index[i]:_Variant_index[i+1]]
}
