// mksyscall_solaris.pl syscall_solaris.go syscall_solaris_amd64.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

// +build amd64,solaris

package unix

import (
	"syscall"
	"unsafe"
)

//go:cgo_import_dynamic libc_getgroups getgroups "libc.so"
//go:cgo_import_dynamic libc_setgroups setgroups "libc.so"
//go:cgo_import_dynamic libc_fcntl fcntl "libc.so"
//go:cgo_import_dynamic libsocket_accept accept "libsocket.so"
//go:cgo_import_dynamic libsocket_sendmsg sendmsg "libsocket.so"
//go:cgo_import_dynamic libc_access access "libc.so"
//go:cgo_import_dynamic libc_adjtime adjtime "libc.so"
//go:cgo_import_dynamic libc_chdir chdir "libc.so"
//go:cgo_import_dynamic libc_chmod chmod "libc.so"
//go:cgo_import_dynamic libc_chown chown "libc.so"
//go:cgo_import_dynamic libc_chroot chroot "libc.so"
//go:cgo_import_dynamic libc_close close "libc.so"
//go:cgo_import_dynamic libc_dup dup "libc.so"
//go:cgo_import_dynamic libc_exit exit "libc.so"
//go:cgo_import_dynamic libc_fchdir fchdir "libc.so"
//go:cgo_import_dynamic libc_fchmod fchmod "libc.so"
//go:cgo_import_dynamic libc_fchown fchown "libc.so"
//go:cgo_import_dynamic libc_fpathconf fpathconf "libc.so"
//go:cgo_import_dynamic libc_fstat fstat "libc.so"
//go:cgo_import_dynamic libc_getdents getdents "libc.so"
//go:cgo_import_dynamic libc_getgid getgid "libc.so"
//go:cgo_import_dynamic libc_getpid getpid "libc.so"
//go:cgo_import_dynamic libc_geteuid geteuid "libc.so"
//go:cgo_import_dynamic libc_getegid getegid "libc.so"
//go:cgo_import_dynamic libc_getppid getppid "libc.so"
//go:cgo_import_dynamic libc_getpriority getpriority "libc.so"
//go:cgo_import_dynamic libc_getrlimit getrlimit "libc.so"
//go:cgo_import_dynamic libc_gettimeofday gettimeofday "libc.so"
//go:cgo_import_dynamic libc_getuid getuid "libc.so"
//go:cgo_import_dynamic libc_kill kill "libc.so"
//go:cgo_import_dynamic libc_lchown lchown "libc.so"
//go:cgo_import_dynamic libc_link link "libc.so"
//go:cgo_import_dynamic libsocket_listen listen "libsocket.so"
//go:cgo_import_dynamic libc_lstat lstat "libc.so"
//go:cgo_import_dynamic libc_madvise madvise "libc.so"
//go:cgo_import_dynamic libc_mkdir mkdir "libc.so"
//go:cgo_import_dynamic libc_mknod mknod "libc.so"
//go:cgo_import_dynamic libc_nanosleep nanosleep "libc.so"
//go:cgo_import_dynamic libc_open open "libc.so"
//go:cgo_import_dynamic libc_pathconf pathconf "libc.so"
//go:cgo_import_dynamic libc_pread pread "libc.so"
//go:cgo_import_dynamic libc_pwrite pwrite "libc.so"
//go:cgo_import_dynamic libc_read read "libc.so"
//go:cgo_import_dynamic libc_readlink readlink "libc.so"
//go:cgo_import_dynamic libc_rename rename "libc.so"
//go:cgo_import_dynamic libc_rmdir rmdir "libc.so"
//go:cgo_import_dynamic libc_lseek lseek "libc.so"
//go:cgo_import_dynamic libc_setegid setegid "libc.so"
//go:cgo_import_dynamic libc_seteuid seteuid "libc.so"
//go:cgo_import_dynamic libc_setgid setgid "libc.so"
//go:cgo_import_dynamic libc_setpgid setpgid "libc.so"
//go:cgo_import_dynamic libc_setpriority setpriority "libc.so"
//go:cgo_import_dynamic libc_setregid setregid "libc.so"
//go:cgo_import_dynamic libc_setreuid setreuid "libc.so"
//go:cgo_import_dynamic libc_setrlimit setrlimit "libc.so"
//go:cgo_import_dynamic libc_setsid setsid "libc.so"
//go:cgo_import_dynamic libc_setuid setuid "libc.so"
//go:cgo_import_dynamic libsocket_shutdown shutdown "libsocket.so"
//go:cgo_import_dynamic libc_stat stat "libc.so"
//go:cgo_import_dynamic libc_symlink symlink "libc.so"
//go:cgo_import_dynamic libc_sync sync "libc.so"
//go:cgo_import_dynamic libc_truncate truncate "libc.so"
//go:cgo_import_dynamic libc_fsync fsync "libc.so"
//go:cgo_import_dynamic libc_ftruncate ftruncate "libc.so"
//go:cgo_import_dynamic libc_umask umask "libc.so"
//go:cgo_import_dynamic libc_unlink unlink "libc.so"
//go:cgo_import_dynamic libc_utimes utimes "libc.so"
//go:cgo_import_dynamic libsocket_bind bind "libsocket.so"
//go:cgo_import_dynamic libsocket_connect connect "libsocket.so"
//go:cgo_import_dynamic libc_mmap mmap "libc.so"
//go:cgo_import_dynamic libc_munmap munmap "libc.so"
//go:cgo_import_dynamic libsocket_sendto sendto "libsocket.so"
//go:cgo_import_dynamic libsocket_socket socket "libsocket.so"
//go:cgo_import_dynamic libsocket_socketpair socketpair "libsocket.so"
//go:cgo_import_dynamic libc_write write "libc.so"
//go:cgo_import_dynamic libsocket_getsockopt getsockopt "libsocket.so"
//go:cgo_import_dynamic libsocket_getpeername getpeername "libsocket.so"
//go:cgo_import_dynamic libsocket_getsockname getsockname "libsocket.so"
//go:cgo_import_dynamic libsocket_setsockopt setsockopt "libsocket.so"
//go:cgo_import_dynamic libsocket_recvfrom recvfrom "libsocket.so"
//go:cgo_import_dynamic libsocket_recvmsg recvmsg "libsocket.so"

//go:linkname procgetgroups libc_getgroups
//go:linkname procsetgroups libc_setgroups
//go:linkname procfcntl libc_fcntl
//go:linkname procaccept libsocket_accept
//go:linkname procsendmsg libsocket_sendmsg
//go:linkname procAccess libc_access
//go:linkname procAdjtime libc_adjtime
//go:linkname procChdir libc_chdir
//go:linkname procChmod libc_chmod
//go:linkname procChown libc_chown
//go:linkname procChroot libc_chroot
//go:linkname procClose libc_close
//go:linkname procDup libc_dup
//go:linkname procExit libc_exit
//go:linkname procFchdir libc_fchdir
//go:linkname procFchmod libc_fchmod
//go:linkname procFchown libc_fchown
//go:linkname procFpathconf libc_fpathconf
//go:linkname procFstat libc_fstat
//go:linkname procGetdents libc_getdents
//go:linkname procGetgid libc_getgid
//go:linkname procGetpid libc_getpid
//go:linkname procGeteuid libc_geteuid
//go:linkname procGetegid libc_getegid
//go:linkname procGetppid libc_getppid
//go:linkname procGetpriority libc_getpriority
//go:linkname procGetrlimit libc_getrlimit
//go:linkname procGettimeofday libc_gettimeofday
//go:linkname procGetuid libc_getuid
//go:linkname procKill libc_kill
//go:linkname procLchown libc_lchown
//go:linkname procLink libc_link
//go:linkname proclisten libsocket_listen
//go:linkname procLstat libc_lstat
//go:linkname procMadvise libc_madvise
//go:linkname procMkdir libc_mkdir
//go:linkname procMknod libc_mknod
//go:linkname procNanosleep libc_nanosleep
//go:linkname procOpen libc_open
//go:linkname procPathconf libc_pathconf
//go:linkname procPread libc_pread
//go:linkname procPwrite libc_pwrite
//go:linkname procread libc_read
//go:linkname procReadlink libc_readlink
//go:linkname procRename libc_rename
//go:linkname procRmdir libc_rmdir
//go:linkname proclseek libc_lseek
//go:linkname procSetegid libc_setegid
//go:linkname procSeteuid libc_seteuid
//go:linkname procSetgid libc_setgid
//go:linkname procSetpgid libc_setpgid
//go:linkname procSetpriority libc_setpriority
//go:linkname procSetregid libc_setregid
//go:linkname procSetreuid libc_setreuid
//go:linkname procSetrlimit libc_setrlimit
//go:linkname procSetsid libc_setsid
//go:linkname procSetuid libc_setuid
//go:linkname procshutdown libsocket_shutdown
//go:linkname procStat libc_stat
//go:linkname procSymlink libc_symlink
//go:linkname procSync libc_sync
//go:linkname procTruncate libc_truncate
//go:linkname procFsync libc_fsync
//go:linkname procFtruncate libc_ftruncate
//go:linkname procUmask libc_umask
//go:linkname procUnlink libc_unlink
//go:linkname procUtimes libc_utimes
//go:linkname procbind libsocket_bind
//go:linkname procconnect libsocket_connect
//go:linkname procmmap libc_mmap
//go:linkname procmunmap libc_munmap
//go:linkname procsendto libsocket_sendto
//go:linkname procsocket libsocket_socket
//go:linkname procsocketpair libsocket_socketpair
//go:linkname procwrite libc_write
//go:linkname procgetsockopt libsocket_getsockopt
//go:linkname procgetpeername libsocket_getpeername
//go:linkname procgetsockname libsocket_getsockname
//go:linkname procsetsockopt libsocket_setsockopt
//go:linkname procrecvfrom libsocket_recvfrom
//go:linkname procrecvmsg libsocket_recvmsg

var (
	procgetgroups,
	procsetgroups,
	procfcntl,
	procaccept,
	procsendmsg,
	procAccess,
	procAdjtime,
	procChdir,
	procChmod,
	procChown,
	procChroot,
	procClose,
	procDup,
	procExit,
	procFchdir,
	procFchmod,
	procFchown,
	procFpathconf,
	procFstat,
	procGetdents,
	procGetgid,
	procGetpid,
	procGeteuid,
	procGetegid,
	procGetppid,
	procGetpriority,
	procGetrlimit,
	procGettimeofday,
	procGetuid,
	procKill,
	procLchown,
	procLink,
	proclisten,
	procLstat,
	procMadvise,
	procMkdir,
	procMknod,
	procNanosleep,
	procOpen,
	procPathconf,
	procPread,
	procPwrite,
	procread,
	procReadlink,
	procRename,
	procRmdir,
	proclseek,
	procSetegid,
	procSeteuid,
	procSetgid,
	procSetpgid,
	procSetpriority,
	procSetregid,
	procSetreuid,
	procSetrlimit,
	procSetsid,
	procSetuid,
	procshutdown,
	procStat,
	procSymlink,
	procSync,
	procTruncate,
	procFsync,
	procFtruncate,
	procUmask,
	procUnlink,
	procUtimes,
	procbind,
	procconnect,
	procmmap,
	procmunmap,
	procsendto,
	procsocket,
	procsocketpair,
	procwrite,
	procgetsockopt,
	procgetpeername,
	procgetsockname,
	procsetsockopt,
	procrecvfrom,
	procrecvmsg syscallFunc
)

func getgroups(ngid int, gid *_Gid_t) (n int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procgetgroups)), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func setgroups(ngid int, gid *_Gid_t) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procsetgroups)), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procfcntl)), 3, uintptr(fd), uintptr(cmd), uintptr(arg), 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procaccept)), 3, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procsendmsg)), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Access(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procAccess)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func Adjtime(delta *Timeval, olddelta *Timeval) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procAdjtime)), 2, uintptr(unsafe.Pointer(delta)), uintptr(unsafe.Pointer(olddelta)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Chdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChdir)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func Chmod(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChmod)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func Chown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChown)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func Chroot(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChroot)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func Close(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procClose)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Dup(fd int) (nfd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procDup)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	nfd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Exit(code int) {
	sysvicall6(uintptr(unsafe.Pointer(&procExit)), 1, uintptr(code), 0, 0, 0, 0, 0)
	return
}

func Fchdir(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchdir)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fchmod(fd int, mode uint32) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchmod)), 2, uintptr(fd), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fchown(fd int, uid int, gid int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchown)), 3, uintptr(fd), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fpathconf(fd int, name int) (val int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFpathconf)), 2, uintptr(fd), uintptr(name), 0, 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fstat(fd int, stat *Stat_t) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFstat)), 2, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getdents(fd int, buf []byte, basep *uintptr) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procGetdents)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(unsafe.Pointer(basep)), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getgid() (gid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&procGetgid)), 0, 0, 0, 0, 0, 0, 0)
	gid = int(r0)
	return
}

func Getpid() (pid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&procGetpid)), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	return
}

func Geteuid() (euid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procGeteuid)), 0, 0, 0, 0, 0, 0, 0)
	euid = int(r0)
	return
}

func Getegid() (egid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procGetegid)), 0, 0, 0, 0, 0, 0, 0)
	egid = int(r0)
	return
}

func Getppid() (ppid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procGetppid)), 0, 0, 0, 0, 0, 0, 0)
	ppid = int(r0)
	return
}

func Getpriority(which int, who int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procGetpriority)), 2, uintptr(which), uintptr(who), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGetrlimit)), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Gettimeofday(tv *Timeval) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGettimeofday)), 1, uintptr(unsafe.Pointer(tv)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getuid() (uid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&procGetuid)), 0, 0, 0, 0, 0, 0, 0)
	uid = int(r0)
	return
}

func Kill(pid int, signum syscall.Signal) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procKill)), 2, uintptr(pid), uintptr(signum), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Lchown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procLchown)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

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
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procLink)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	if e1 != 0 {
		err = e1
	}
	return
}

func Listen(s int, backlog int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proclisten)), 2, uintptr(s), uintptr(backlog), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Lstat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procLstat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func Madvise(b []byte, advice int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMadvise)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), uintptr(advice), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mkdir(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMkdir)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func Mknod(path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMknod)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev), 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func Nanosleep(time *Timespec, leftover *Timespec) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procNanosleep)), 2, uintptr(unsafe.Pointer(time)), uintptr(unsafe.Pointer(leftover)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Open(path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procOpen)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm), 0, 0, 0)
	use(unsafe.Pointer(_p0))
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Pathconf(path string, name int) (val int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procPathconf)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(name), 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Pread(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procPread)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procPwrite)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func read(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procread)), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Readlink(path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	if len(buf) > 0 {
		_p1 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procReadlink)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(len(buf)), 0, 0, 0)
	use(unsafe.Pointer(_p0))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

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
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procRename)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	if e1 != 0 {
		err = e1
	}
	return
}

func Rmdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procRmdir)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func Seek(fd int, offset int64, whence int) (newoffset int64, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proclseek)), 3, uintptr(fd), uintptr(offset), uintptr(whence), 0, 0, 0)
	newoffset = int64(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setegid(egid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetegid)), 1, uintptr(egid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Seteuid(euid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSeteuid)), 1, uintptr(euid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setgid(gid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetgid)), 1, uintptr(gid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setpgid(pid int, pgid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetpgid)), 2, uintptr(pid), uintptr(pgid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setpriority(which int, who int, prio int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSetpriority)), 3, uintptr(which), uintptr(who), uintptr(prio), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setregid(rgid int, egid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetregid)), 2, uintptr(rgid), uintptr(egid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setreuid(ruid int, euid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetreuid)), 2, uintptr(ruid), uintptr(euid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetrlimit)), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setsid() (pid int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetsid)), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setuid(uid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetuid)), 1, uintptr(uid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Shutdown(s int, how int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procshutdown)), 2, uintptr(s), uintptr(how), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Stat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procStat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

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
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSymlink)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	use(unsafe.Pointer(_p1))
	if e1 != 0 {
		err = e1
	}
	return
}

func Sync() (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSync)), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Truncate(path string, length int64) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procTruncate)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(length), 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func Fsync(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFsync)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Ftruncate(fd int, length int64) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFtruncate)), 2, uintptr(fd), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Umask(newmask int) (oldmask int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procUmask)), 1, uintptr(newmask), 0, 0, 0, 0, 0)
	oldmask = int(r0)
	return
}

func Unlink(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procUnlink)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func Utimes(path string, times *[2]Timeval) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procUtimes)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)), 0, 0, 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = e1
	}
	return
}

func bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procbind)), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procconnect)), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func mmap(addr uintptr, length uintptr, prot int, flag int, fd int, pos int64) (ret uintptr, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procmmap)), 6, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flag), uintptr(fd), uintptr(pos))
	ret = uintptr(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func munmap(addr uintptr, length uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procmunmap)), 2, uintptr(addr), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procsendto)), 6, uintptr(s), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(addrlen))
	if e1 != 0 {
		err = e1
	}
	return
}

func socket(domain int, typ int, proto int) (fd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procsocket)), 3, uintptr(domain), uintptr(typ), uintptr(proto), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func socketpair(domain int, typ int, proto int, fd *[2]int32) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procsocketpair)), 4, uintptr(domain), uintptr(typ), uintptr(proto), uintptr(unsafe.Pointer(fd)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func write(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procwrite)), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procgetsockopt)), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procgetpeername)), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procgetsockname)), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procsetsockopt)), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procrecvfrom)), 6, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procrecvmsg)), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
