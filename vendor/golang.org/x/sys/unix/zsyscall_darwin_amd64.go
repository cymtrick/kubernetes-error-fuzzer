// mksyscall.pl -tags darwin,amd64 syscall_bsd.go syscall_darwin.go syscall_darwin_amd64.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build darwin,amd64

package unix

import (
	"syscall"
	"unsafe"
)

var _ syscall.Errno

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getgroups(ngid int, gid *_Gid_t) (n int, err error) {
	r0, _, e1 := RawSyscall(SYS_GETGROUPS, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func setgroups(ngid int, gid *_Gid_t) (err error) {
	_, _, e1 := RawSyscall(SYS_SETGROUPS, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func wait4(pid int, wstatus *_C_int, options int, rusage *Rusage) (wpid int, err error) {
	r0, _, e1 := Syscall6(SYS_WAIT4, uintptr(pid), uintptr(unsafe.Pointer(wstatus)), uintptr(options), uintptr(unsafe.Pointer(rusage)), 0, 0)
	wpid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	r0, _, e1 := Syscall(SYS_ACCEPT, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := Syscall(SYS_BIND, uintptr(s), uintptr(addr), uintptr(addrlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := Syscall(SYS_CONNECT, uintptr(s), uintptr(addr), uintptr(addrlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func socket(domain int, typ int, proto int) (fd int, err error) {
	r0, _, e1 := RawSyscall(SYS_SOCKET, uintptr(domain), uintptr(typ), uintptr(proto))
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	_, _, e1 := Syscall6(SYS_GETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
	_, _, e1 := Syscall6(SYS_SETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := RawSyscall(SYS_GETPEERNAME, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := RawSyscall(SYS_GETSOCKNAME, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Shutdown(s int, how int) (err error) {
	_, _, e1 := Syscall(SYS_SHUTDOWN, uintptr(s), uintptr(how), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func socketpair(domain int, typ int, proto int, fd *[2]int32) (err error) {
	_, _, e1 := RawSyscall6(SYS_SOCKETPAIR, uintptr(domain), uintptr(typ), uintptr(proto), uintptr(unsafe.Pointer(fd)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(SYS_RECVFROM, uintptr(fd), uintptr(_p0), uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
	var _p0 unsafe.Pointer
	if len(buf) > 0 {
		_p0 = unsafe.Pointer(&buf[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall6(SYS_SENDTO, uintptr(s), uintptr(_p0), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(addrlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := Syscall(SYS_RECVMSG, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := Syscall(SYS_SENDMSG, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func kevent(kq int, change unsafe.Pointer, nchange int, event unsafe.Pointer, nevent int, timeout *Timespec) (n int, err error) {
	r0, _, e1 := Syscall6(SYS_KEVENT, uintptr(kq), uintptr(change), uintptr(nchange), uintptr(event), uintptr(nevent), uintptr(unsafe.Pointer(timeout)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func sysctl(mib []_C_int, old *byte, oldlen *uintptr, new *byte, newlen uintptr) (err error) {
	var _p0 unsafe.Pointer
	if len(mib) > 0 {
		_p0 = unsafe.Pointer(&mib[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall6(SYS___SYSCTL, uintptr(_p0), uintptr(len(mib)), uintptr(unsafe.Pointer(old)), uintptr(unsafe.Pointer(oldlen)), uintptr(unsafe.Pointer(new)), uintptr(newlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func utimes(path string, timeval *[2]Timeval) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_UTIMES, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(timeval)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func futimes(fd int, timeval *[2]Timeval) (err error) {
	_, _, e1 := Syscall(SYS_FUTIMES, uintptr(fd), uintptr(unsafe.Pointer(timeval)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, _, e1 := Syscall(SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg))
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Madvise(b []byte, behav int) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall(SYS_MADVISE, uintptr(_p0), uintptr(len(b)), uintptr(behav))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mlock(b []byte) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall(SYS_MLOCK, uintptr(_p0), uintptr(len(b)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mlockall(flags int) (err error) {
	_, _, e1 := Syscall(SYS_MLOCKALL, uintptr(flags), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mprotect(b []byte, prot int) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall(SYS_MPROTECT, uintptr(_p0), uintptr(len(b)), uintptr(prot))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Msync(b []byte, flags int) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall(SYS_MSYNC, uintptr(_p0), uintptr(len(b)), uintptr(flags))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Munlock(b []byte) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := Syscall(SYS_MUNLOCK, uintptr(_p0), uintptr(len(b)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Munlockall() (err error) {
	_, _, e1 := Syscall(SYS_MUNLOCKALL, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func ptrace(request int, pid int, addr uintptr, data uintptr) (err error) {
	_, _, e1 := Syscall6(SYS_PTRACE, uintptr(request), uintptr(pid), uintptr(addr), uintptr(data), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func pipe() (r int, w int, err error) {
	r0, r1, e1 := RawSyscall(SYS_PIPE, 0, 0, 0)
	r = int(r0)
	w = int(r1)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func kill(pid int, signum int, posix int) (err error) {
	_, _, e1 := Syscall(SYS_KILL, uintptr(pid), uintptr(signum), uintptr(posix))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func ioctl(fd int, req uint, arg uintptr) (err error) {
	_, _, e1 := Syscall(SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(arg))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Access(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_ACCESS, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Adjtime(delta *Timeval, olddelta *Timeval) (err error) {
	_, _, e1 := Syscall(SYS_ADJTIME, uintptr(unsafe.Pointer(delta)), uintptr(unsafe.Pointer(olddelta)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_CHDIR, uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chflags(path string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_CHFLAGS, uintptr(unsafe.Pointer(_p0)), uintptr(flags), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chmod(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_CHMOD, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_CHOWN, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chroot(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_CHROOT, uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Close(fd int) (err error) {
	_, _, e1 := Syscall(SYS_CLOSE, uintptr(fd), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Dup(fd int) (nfd int, err error) {
	r0, _, e1 := Syscall(SYS_DUP, uintptr(fd), 0, 0)
	nfd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Dup2(from int, to int) (err error) {
	_, _, e1 := Syscall(SYS_DUP2, uintptr(from), uintptr(to), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Exchangedata(path1 string, path2 string, options int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path1)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(path2)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_EXCHANGEDATA, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(options))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Exit(code int) {
	Syscall(SYS_EXIT, uintptr(code), 0, 0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Faccessat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_FACCESSAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(flags), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchdir(fd int) (err error) {
	_, _, e1 := Syscall(SYS_FCHDIR, uintptr(fd), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchflags(fd int, flags int) (err error) {
	_, _, e1 := Syscall(SYS_FCHFLAGS, uintptr(fd), uintptr(flags), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchmod(fd int, mode uint32) (err error) {
	_, _, e1 := Syscall(SYS_FCHMOD, uintptr(fd), uintptr(mode), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_FCHMODAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(flags), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchown(fd int, uid int, gid int) (err error) {
	_, _, e1 := Syscall(SYS_FCHOWN, uintptr(fd), uintptr(uid), uintptr(gid))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_FCHOWNAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), uintptr(flags), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Flock(fd int, how int) (err error) {
	_, _, e1 := Syscall(SYS_FLOCK, uintptr(fd), uintptr(how), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fpathconf(fd int, name int) (val int, err error) {
	r0, _, e1 := Syscall(SYS_FPATHCONF, uintptr(fd), uintptr(name), 0)
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fstat(fd int, stat *Stat_t) (err error) {
	_, _, e1 := Syscall(SYS_FSTAT64, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fstatfs(fd int, stat *Statfs_t) (err error) {
	_, _, e1 := Syscall(SYS_FSTATFS64, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fsync(fd int) (err error) {
	_, _, e1 := Syscall(SYS_FSYNC, uintptr(fd), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Ftruncate(fd int, length int64) (err error) {
	_, _, e1 := Syscall(SYS_FTRUNCATE, uintptr(fd), uintptr(length), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getdirentries(fd int, buf []byte, basep *uintptr) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(buf) > 0 {
		_p0 = unsafe.Pointer(&buf[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(SYS_GETDIRENTRIES64, uintptr(fd), uintptr(_p0), uintptr(len(buf)), uintptr(unsafe.Pointer(basep)), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getdtablesize() (size int) {
	r0, _, _ := Syscall(SYS_GETDTABLESIZE, 0, 0, 0)
	size = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getegid() (egid int) {
	r0, _, _ := RawSyscall(SYS_GETEGID, 0, 0, 0)
	egid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Geteuid() (uid int) {
	r0, _, _ := RawSyscall(SYS_GETEUID, 0, 0, 0)
	uid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getgid() (gid int) {
	r0, _, _ := RawSyscall(SYS_GETGID, 0, 0, 0)
	gid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpgid(pid int) (pgid int, err error) {
	r0, _, e1 := RawSyscall(SYS_GETPGID, uintptr(pid), 0, 0)
	pgid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpgrp() (pgrp int) {
	r0, _, _ := RawSyscall(SYS_GETPGRP, 0, 0, 0)
	pgrp = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpid() (pid int) {
	r0, _, _ := RawSyscall(SYS_GETPID, 0, 0, 0)
	pid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getppid() (ppid int) {
	r0, _, _ := RawSyscall(SYS_GETPPID, 0, 0, 0)
	ppid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpriority(which int, who int) (prio int, err error) {
	r0, _, e1 := Syscall(SYS_GETPRIORITY, uintptr(which), uintptr(who), 0)
	prio = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := RawSyscall(SYS_GETRLIMIT, uintptr(which), uintptr(unsafe.Pointer(lim)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getrusage(who int, rusage *Rusage) (err error) {
	_, _, e1 := RawSyscall(SYS_GETRUSAGE, uintptr(who), uintptr(unsafe.Pointer(rusage)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getsid(pid int) (sid int, err error) {
	r0, _, e1 := RawSyscall(SYS_GETSID, uintptr(pid), 0, 0)
	sid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getuid() (uid int) {
	r0, _, _ := RawSyscall(SYS_GETUID, 0, 0, 0)
	uid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Issetugid() (tainted bool) {
	r0, _, _ := RawSyscall(SYS_ISSETUGID, 0, 0, 0)
	tainted = bool(r0 != 0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Kqueue() (fd int, err error) {
	r0, _, e1 := Syscall(SYS_KQUEUE, 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Lchown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_LCHOWN, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Link(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_LINK, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Linkat(pathfd int, path string, linkfd int, link string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_LINKAT, uintptr(pathfd), uintptr(unsafe.Pointer(_p0)), uintptr(linkfd), uintptr(unsafe.Pointer(_p1)), uintptr(flags), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Listen(s int, backlog int) (err error) {
	_, _, e1 := Syscall(SYS_LISTEN, uintptr(s), uintptr(backlog), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Lstat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_LSTAT64, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mkdir(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_MKDIR, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mkdirat(dirfd int, path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_MKDIRAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mkfifo(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_MKFIFO, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mknod(path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_MKNOD, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Open(path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := Syscall(SYS_OPEN, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm))
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Openat(dirfd int, path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := Syscall6(SYS_OPENAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm), 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pathconf(path string, name int) (val int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := Syscall(SYS_PATHCONF, uintptr(unsafe.Pointer(_p0)), uintptr(name), 0)
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pread(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(SYS_PREAD, uintptr(fd), uintptr(_p0), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(SYS_PWRITE, uintptr(fd), uintptr(_p0), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func read(fd int, p []byte) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(_p0), uintptr(len(p)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Readlink(path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 unsafe.Pointer
	if len(buf) > 0 {
		_p1 = unsafe.Pointer(&buf[0])
	} else {
		_p1 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(SYS_READLINK, uintptr(unsafe.Pointer(_p0)), uintptr(_p1), uintptr(len(buf)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Readlinkat(dirfd int, path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 unsafe.Pointer
	if len(buf) > 0 {
		_p1 = unsafe.Pointer(&buf[0])
	} else {
		_p1 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(SYS_READLINKAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(_p1), uintptr(len(buf)), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Rename(from string, to string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(from)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(to)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_RENAME, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Renameat(fromfd int, from string, tofd int, to string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(from)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(to)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_RENAMEAT, uintptr(fromfd), uintptr(unsafe.Pointer(_p0)), uintptr(tofd), uintptr(unsafe.Pointer(_p1)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Revoke(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_REVOKE, uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Rmdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_RMDIR, uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Seek(fd int, offset int64, whence int) (newoffset int64, err error) {
	r0, _, e1 := Syscall(SYS_LSEEK, uintptr(fd), uintptr(offset), uintptr(whence))
	newoffset = int64(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Select(n int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (err error) {
	_, _, e1 := Syscall6(SYS_SELECT, uintptr(n), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(timeout)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setegid(egid int) (err error) {
	_, _, e1 := Syscall(SYS_SETEGID, uintptr(egid), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Seteuid(euid int) (err error) {
	_, _, e1 := RawSyscall(SYS_SETEUID, uintptr(euid), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setgid(gid int) (err error) {
	_, _, e1 := RawSyscall(SYS_SETGID, uintptr(gid), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setlogin(name string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(name)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_SETLOGIN, uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setpgid(pid int, pgid int) (err error) {
	_, _, e1 := RawSyscall(SYS_SETPGID, uintptr(pid), uintptr(pgid), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setpriority(which int, who int, prio int) (err error) {
	_, _, e1 := Syscall(SYS_SETPRIORITY, uintptr(which), uintptr(who), uintptr(prio))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setprivexec(flag int) (err error) {
	_, _, e1 := Syscall(SYS_SETPRIVEXEC, uintptr(flag), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setregid(rgid int, egid int) (err error) {
	_, _, e1 := RawSyscall(SYS_SETREGID, uintptr(rgid), uintptr(egid), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setreuid(ruid int, euid int) (err error) {
	_, _, e1 := RawSyscall(SYS_SETREUID, uintptr(ruid), uintptr(euid), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := RawSyscall(SYS_SETRLIMIT, uintptr(which), uintptr(unsafe.Pointer(lim)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setsid() (pid int, err error) {
	r0, _, e1 := RawSyscall(SYS_SETSID, 0, 0, 0)
	pid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Settimeofday(tp *Timeval) (err error) {
	_, _, e1 := RawSyscall(SYS_SETTIMEOFDAY, uintptr(unsafe.Pointer(tp)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setuid(uid int) (err error) {
	_, _, e1 := RawSyscall(SYS_SETUID, uintptr(uid), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Stat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_STAT64, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Statfs(path string, stat *Statfs_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_STATFS64, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Symlink(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_SYMLINK, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Symlinkat(oldpath string, newdirfd int, newpath string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(oldpath)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(newpath)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_SYMLINKAT, uintptr(unsafe.Pointer(_p0)), uintptr(newdirfd), uintptr(unsafe.Pointer(_p1)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Sync() (err error) {
	_, _, e1 := Syscall(SYS_SYNC, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Truncate(path string, length int64) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_TRUNCATE, uintptr(unsafe.Pointer(_p0)), uintptr(length), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Umask(newmask int) (oldmask int) {
	r0, _, _ := Syscall(SYS_UMASK, uintptr(newmask), 0, 0)
	oldmask = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Undelete(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_UNDELETE, uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Unlink(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_UNLINK, uintptr(unsafe.Pointer(_p0)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Unlinkat(dirfd int, path string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_UNLINKAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(flags))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Unmount(path string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := Syscall(SYS_UNMOUNT, uintptr(unsafe.Pointer(_p0)), uintptr(flags), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func write(fd int, p []byte) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(_p0), uintptr(len(p)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func mmap(addr uintptr, length uintptr, prot int, flag int, fd int, pos int64) (ret uintptr, err error) {
	r0, _, e1 := Syscall6(SYS_MMAP, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flag), uintptr(fd), uintptr(pos))
	ret = uintptr(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func munmap(addr uintptr, length uintptr) (err error) {
	_, _, e1 := Syscall(SYS_MUNMAP, uintptr(addr), uintptr(length), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func readlen(fd int, buf *byte, nbuf int) (n int, err error) {
	r0, _, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func writelen(fd int, buf *byte, nbuf int) (n int, err error) {
	r0, _, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func gettimeofday(tp *Timeval) (sec int64, usec int32, err error) {
	r0, r1, e1 := RawSyscall(SYS_GETTIMEOFDAY, uintptr(unsafe.Pointer(tp)), 0, 0)
	sec = int64(r0)
	usec = int32(r1)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
