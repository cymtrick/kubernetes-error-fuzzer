package kubelet

import (
	"C"
)
import (
	"testing"
)

//export TestFuzzPodMutator
func FuzzPodMutator(data *C.char, size C.size_t, t *testing.T) {
	TestHandleMemExceeded(t)
}

func TestHandleMemExceeded(t *testing.T) {
	panic("unimplemented")
}

//go build -o libgofuzzer.so -buildmode=c-shared FUZZ_TARGET.go
