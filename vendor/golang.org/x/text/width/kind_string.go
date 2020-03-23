// Code generated by "stringer -type=Kind"; DO NOT EDIT.

package width

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Neutral-0]
	_ = x[EastAsianAmbiguous-1]
	_ = x[EastAsianWide-2]
	_ = x[EastAsianNarrow-3]
	_ = x[EastAsianFullwidth-4]
	_ = x[EastAsianHalfwidth-5]
}

const _Kind_name = "NeutralEastAsianAmbiguousEastAsianWideEastAsianNarrowEastAsianFullwidthEastAsianHalfwidth"

var _Kind_index = [...]uint8{0, 7, 25, 38, 53, 71, 89}

func (i Kind) String() string {
	if i < 0 || i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
