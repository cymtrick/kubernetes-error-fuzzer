// Code generated mksyscall_windows.exe DO NOT EDIT

package hns

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modvmcompute = windows.NewLazySystemDLL("vmcompute.dll")

	procHNSCall = modvmcompute.NewProc("HNSCall")
)

func _hnsCall(method string, path string, object string, response **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(method)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(path)
	if hr != nil {
		return
	}
	var _p2 *uint16
	_p2, hr = syscall.UTF16PtrFromString(object)
	if hr != nil {
		return
	}
	return __hnsCall(_p0, _p1, _p2, response)
}

func __hnsCall(method *uint16, path *uint16, object *uint16, response **uint16) (hr error) {
	if hr = procHNSCall.Find(); hr != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procHNSCall.Addr(), 4, uintptr(unsafe.Pointer(method)), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(object)), uintptr(unsafe.Pointer(response)), 0, 0)
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}
