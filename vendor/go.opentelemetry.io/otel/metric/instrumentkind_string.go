// Code generated by "stringer -type=InstrumentKind"; DO NOT EDIT.

package metric

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ValueRecorderInstrumentKind-0]
	_ = x[ValueObserverInstrumentKind-1]
	_ = x[CounterInstrumentKind-2]
	_ = x[UpDownCounterInstrumentKind-3]
	_ = x[SumObserverInstrumentKind-4]
	_ = x[UpDownSumObserverInstrumentKind-5]
}

const _InstrumentKind_name = "ValueRecorderInstrumentKindValueObserverInstrumentKindCounterInstrumentKindUpDownCounterInstrumentKindSumObserverInstrumentKindUpDownSumObserverInstrumentKind"

var _InstrumentKind_index = [...]uint8{0, 27, 54, 75, 102, 127, 158}

func (i InstrumentKind) String() string {
	if i < 0 || i >= InstrumentKind(len(_InstrumentKind_index)-1) {
		return "InstrumentKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _InstrumentKind_name[_InstrumentKind_index[i]:_InstrumentKind_index[i+1]]
}
