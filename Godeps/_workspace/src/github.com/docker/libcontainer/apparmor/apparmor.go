// +build apparmor,linux

package apparmor

// #cgo LDFLAGS: -lapparmor
// #include <sys/apparmor.h>
// #include <stdlib.h>
import "C"
import (
	"io/ioutil"
	"os"
	"unsafe"
)

func IsEnabled() bool {
	if _, err := os.Stat("/sys/kernel/security/apparmor"); err == nil && os.Getenv("container") == "" {
		if _, err = os.Stat("/sbin/apparmor_parser"); err == nil {
			buf, err := ioutil.ReadFile("/sys/module/apparmor/parameters/enabled")
			return err == nil && len(buf) > 1 && buf[0] == 'Y'
		}
	}
	return false
}

func ApplyProfile(name string) error {
	if name == "" {
		return nil
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	if _, err := C.aa_change_onexec(cName); err != nil {
		return err
	}
	return nil
}
