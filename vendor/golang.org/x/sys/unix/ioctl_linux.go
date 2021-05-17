// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"runtime"
	"unsafe"
)

// IoctlRetInt performs an ioctl operation specified by req on a device
// associated with opened file descriptor fd, and returns a non-negative
// integer that is returned by the ioctl syscall.
func IoctlRetInt(fd int, req uint) (int, error) {
	ret, _, err := Syscall(SYS_IOCTL, uintptr(fd), uintptr(req), 0)
	if err != 0 {
		return 0, err
	}
	return int(ret), nil
}

func IoctlGetUint32(fd int, req uint) (uint32, error) {
	var value uint32
	err := ioctl(fd, req, uintptr(unsafe.Pointer(&value)))
	return value, err
}

func IoctlGetRTCTime(fd int) (*RTCTime, error) {
	var value RTCTime
	err := ioctl(fd, RTC_RD_TIME, uintptr(unsafe.Pointer(&value)))
	return &value, err
}

func IoctlSetRTCTime(fd int, value *RTCTime) error {
	err := ioctl(fd, RTC_SET_TIME, uintptr(unsafe.Pointer(value)))
	runtime.KeepAlive(value)
	return err
}

func IoctlGetRTCWkAlrm(fd int) (*RTCWkAlrm, error) {
	var value RTCWkAlrm
	err := ioctl(fd, RTC_WKALM_RD, uintptr(unsafe.Pointer(&value)))
	return &value, err
}

func IoctlSetRTCWkAlrm(fd int, value *RTCWkAlrm) error {
	err := ioctl(fd, RTC_WKALM_SET, uintptr(unsafe.Pointer(value)))
	runtime.KeepAlive(value)
	return err
}

type ifreqEthtool struct {
	name [IFNAMSIZ]byte
	data unsafe.Pointer
}

// IoctlGetEthtoolDrvinfo fetches ethtool driver information for the network
// device specified by ifname.
func IoctlGetEthtoolDrvinfo(fd int, ifname string) (*EthtoolDrvinfo, error) {
	// Leave room for terminating NULL byte.
	if len(ifname) >= IFNAMSIZ {
		return nil, EINVAL
	}

	value := EthtoolDrvinfo{
		Cmd: ETHTOOL_GDRVINFO,
	}
	ifreq := ifreqEthtool{
		data: unsafe.Pointer(&value),
	}
	copy(ifreq.name[:], ifname)
	err := ioctl(fd, SIOCETHTOOL, uintptr(unsafe.Pointer(&ifreq)))
	runtime.KeepAlive(ifreq)
	return &value, err
}

// IoctlGetWatchdogInfo fetches information about a watchdog device from the
// Linux watchdog API. For more information, see:
// https://www.kernel.org/doc/html/latest/watchdog/watchdog-api.html.
func IoctlGetWatchdogInfo(fd int) (*WatchdogInfo, error) {
	var value WatchdogInfo
	err := ioctl(fd, WDIOC_GETSUPPORT, uintptr(unsafe.Pointer(&value)))
	return &value, err
}

// IoctlWatchdogKeepalive issues a keepalive ioctl to a watchdog device. For
// more information, see:
// https://www.kernel.org/doc/html/latest/watchdog/watchdog-api.html.
func IoctlWatchdogKeepalive(fd int) error {
	return ioctl(fd, WDIOC_KEEPALIVE, 0)
}

// IoctlFileCloneRange performs an FICLONERANGE ioctl operation to clone the
// range of data conveyed in value to the file associated with the file
// descriptor destFd. See the ioctl_ficlonerange(2) man page for details.
func IoctlFileCloneRange(destFd int, value *FileCloneRange) error {
	err := ioctl(destFd, FICLONERANGE, uintptr(unsafe.Pointer(value)))
	runtime.KeepAlive(value)
	return err
}

// IoctlFileClone performs an FICLONE ioctl operation to clone the entire file
// associated with the file description srcFd to the file associated with the
// file descriptor destFd. See the ioctl_ficlone(2) man page for details.
func IoctlFileClone(destFd, srcFd int) error {
	return ioctl(destFd, FICLONE, uintptr(srcFd))
}

type FileDedupeRange struct {
	Src_offset uint64
	Src_length uint64
	Reserved1  uint16
	Reserved2  uint32
	Info       []FileDedupeRangeInfo
}

type FileDedupeRangeInfo struct {
	Dest_fd       int64
	Dest_offset   uint64
	Bytes_deduped uint64
	Status        int32
	Reserved      uint32
}

// IoctlFileDedupeRange performs an FIDEDUPERANGE ioctl operation to share the
// range of data conveyed in value from the file associated with the file
// descriptor srcFd to the value.Info destinations. See the
// ioctl_fideduperange(2) man page for details.
func IoctlFileDedupeRange(srcFd int, value *FileDedupeRange) error {
	buf := make([]byte, SizeofRawFileDedupeRange+
		len(value.Info)*SizeofRawFileDedupeRangeInfo)
	rawrange := (*RawFileDedupeRange)(unsafe.Pointer(&buf[0]))
	rawrange.Src_offset = value.Src_offset
	rawrange.Src_length = value.Src_length
	rawrange.Dest_count = uint16(len(value.Info))
	rawrange.Reserved1 = value.Reserved1
	rawrange.Reserved2 = value.Reserved2

	for i := range value.Info {
		rawinfo := (*RawFileDedupeRangeInfo)(unsafe.Pointer(
			uintptr(unsafe.Pointer(&buf[0])) + uintptr(SizeofRawFileDedupeRange) +
				uintptr(i*SizeofRawFileDedupeRangeInfo)))
		rawinfo.Dest_fd = value.Info[i].Dest_fd
		rawinfo.Dest_offset = value.Info[i].Dest_offset
		rawinfo.Bytes_deduped = value.Info[i].Bytes_deduped
		rawinfo.Status = value.Info[i].Status
		rawinfo.Reserved = value.Info[i].Reserved
	}

	err := ioctl(srcFd, FIDEDUPERANGE, uintptr(unsafe.Pointer(&buf[0])))

	// Output
	for i := range value.Info {
		rawinfo := (*RawFileDedupeRangeInfo)(unsafe.Pointer(
			uintptr(unsafe.Pointer(&buf[0])) + uintptr(SizeofRawFileDedupeRange) +
				uintptr(i*SizeofRawFileDedupeRangeInfo)))
		value.Info[i].Dest_fd = rawinfo.Dest_fd
		value.Info[i].Dest_offset = rawinfo.Dest_offset
		value.Info[i].Bytes_deduped = rawinfo.Bytes_deduped
		value.Info[i].Status = rawinfo.Status
		value.Info[i].Reserved = rawinfo.Reserved
	}

	return err
}

func IoctlHIDGetDesc(fd int, value *HIDRawReportDescriptor) error {
	err := ioctl(fd, HIDIOCGRDESC, uintptr(unsafe.Pointer(value)))
	runtime.KeepAlive(value)
	return err
}

func IoctlHIDGetRawInfo(fd int) (*HIDRawDevInfo, error) {
	var value HIDRawDevInfo
	err := ioctl(fd, HIDIOCGRAWINFO, uintptr(unsafe.Pointer(&value)))
	return &value, err
}

func IoctlHIDGetRawName(fd int) (string, error) {
	var value [_HIDIOCGRAWNAME_LEN]byte
	err := ioctl(fd, _HIDIOCGRAWNAME, uintptr(unsafe.Pointer(&value[0])))
	return ByteSliceToString(value[:]), err
}

func IoctlHIDGetRawPhys(fd int) (string, error) {
	var value [_HIDIOCGRAWPHYS_LEN]byte
	err := ioctl(fd, _HIDIOCGRAWPHYS, uintptr(unsafe.Pointer(&value[0])))
	return ByteSliceToString(value[:]), err
}

func IoctlHIDGetRawUniq(fd int) (string, error) {
	var value [_HIDIOCGRAWUNIQ_LEN]byte
	err := ioctl(fd, _HIDIOCGRAWUNIQ, uintptr(unsafe.Pointer(&value[0])))
	return ByteSliceToString(value[:]), err
}
