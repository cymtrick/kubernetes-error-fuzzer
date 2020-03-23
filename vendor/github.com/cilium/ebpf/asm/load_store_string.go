// Code generated by "stringer -output load_store_string.go -type=Mode,Size"; DO NOT EDIT.

package asm

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidMode-255]
	_ = x[ImmMode-0]
	_ = x[AbsMode-32]
	_ = x[IndMode-64]
	_ = x[MemMode-96]
	_ = x[XAddMode-192]
}

const (
	_Mode_name_0 = "ImmMode"
	_Mode_name_1 = "AbsMode"
	_Mode_name_2 = "IndMode"
	_Mode_name_3 = "MemMode"
	_Mode_name_4 = "XAddMode"
	_Mode_name_5 = "InvalidMode"
)

func (i Mode) String() string {
	switch {
	case i == 0:
		return _Mode_name_0
	case i == 32:
		return _Mode_name_1
	case i == 64:
		return _Mode_name_2
	case i == 96:
		return _Mode_name_3
	case i == 192:
		return _Mode_name_4
	case i == 255:
		return _Mode_name_5
	default:
		return "Mode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidSize-255]
	_ = x[DWord-24]
	_ = x[Word-0]
	_ = x[Half-8]
	_ = x[Byte-16]
}

const (
	_Size_name_0 = "Word"
	_Size_name_1 = "Half"
	_Size_name_2 = "Byte"
	_Size_name_3 = "DWord"
	_Size_name_4 = "InvalidSize"
)

func (i Size) String() string {
	switch {
	case i == 0:
		return _Size_name_0
	case i == 8:
		return _Size_name_1
	case i == 16:
		return _Size_name_2
	case i == 24:
		return _Size_name_3
	case i == 255:
		return _Size_name_4
	default:
		return "Size(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
