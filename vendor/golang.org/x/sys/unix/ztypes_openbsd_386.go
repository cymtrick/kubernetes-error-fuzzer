// cgo -godefs types_openbsd.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build 386 && openbsd
// +build 386,openbsd

package unix

const (
	SizeofPtr      = 0x4
	SizeofShort    = 0x2
	SizeofInt      = 0x4
	SizeofLong     = 0x4
	SizeofLongLong = 0x8
)

type (
	_C_short     int16
	_C_int       int32
	_C_long      int32
	_C_long_long int64
)

type Timespec struct {
	Sec  int64
	Nsec int32
}

type Timeval struct {
	Sec  int64
	Usec int32
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int32
	Ixrss    int32
	Idrss    int32
	Isrss    int32
	Minflt   int32
	Majflt   int32
	Nswap    int32
	Inblock  int32
	Oublock  int32
	Msgsnd   int32
	Msgrcv   int32
	Nsignals int32
	Nvcsw    int32
	Nivcsw   int32
}

type Rlimit struct {
	Cur uint64
	Max uint64
}

type _Gid_t uint32

type Stat_t struct {
	Mode           uint32
	Dev            int32
	Ino            uint64
	Nlink          uint32
	Uid            uint32
	Gid            uint32
	Rdev           int32
	Atim           Timespec
	Mtim           Timespec
	Ctim           Timespec
	Size           int64
	Blocks         int64
	Blksize        uint32
	Flags          uint32
	Gen            uint32
	X__st_birthtim Timespec
}

type Statfs_t struct {
	F_flags       uint32
	F_bsize       uint32
	F_iosize      uint32
	F_blocks      uint64
	F_bfree       uint64
	F_bavail      int64
	F_files       uint64
	F_ffree       uint64
	F_favail      int64
	F_syncwrites  uint64
	F_syncreads   uint64
	F_asyncwrites uint64
	F_asyncreads  uint64
	F_fsid        Fsid
	F_namemax     uint32
	F_owner       uint32
	F_ctime       uint64
	F_fstypename  [16]int8
	F_mntonname   [90]int8
	F_mntfromname [90]int8
	F_mntfromspec [90]int8
	Pad_cgo_0     [2]byte
	Mount_info    [160]byte
}

type Flock_t struct {
	Start  int64
	Len    int64
	Pid    int32
	Type   int16
	Whence int16
}

type Dirent struct {
	Fileno       uint64
	Off          int64
	Reclen       uint16
	Type         uint8
	Namlen       uint8
	X__d_padding [4]uint8
	Name         [256]int8
}

type Fsid struct {
	Val [2]int32
}

const (
	PathMax = 0x400
)

type RawSockaddrInet4 struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]int8
}

type RawSockaddrInet6 struct {
	Len      uint8
	Family   uint8
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

type RawSockaddrUnix struct {
	Len    uint8
	Family uint8
	Path   [104]int8
}

type RawSockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [24]int8
}

type RawSockaddr struct {
	Len    uint8
	Family uint8
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [92]int8
}

type _Socklen uint32

type Linger struct {
	Onoff  int32
	Linger int32
}

type Iovec struct {
	Base *byte
	Len  uint32
}

type IPMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type IPv6Mreq struct {
	Multiaddr [16]byte /* in6_addr */
	Interface uint32
}

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *Iovec
	Iovlen     uint32
	Control    *byte
	Controllen uint32
	Flags      int32
}

type Cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

type Inet6Pktinfo struct {
	Addr    [16]byte /* in6_addr */
	Ifindex uint32
}

type IPv6MTUInfo struct {
	Addr RawSockaddrInet6
	Mtu  uint32
}

type ICMPv6Filter struct {
	Filt [8]uint32
}

const (
	SizeofSockaddrInet4    = 0x10
	SizeofSockaddrInet6    = 0x1c
	SizeofSockaddrAny      = 0x6c
	SizeofSockaddrUnix     = 0x6a
	SizeofSockaddrDatalink = 0x20
	SizeofLinger           = 0x8
	SizeofIovec            = 0x8
	SizeofIPMreq           = 0x8
	SizeofIPv6Mreq         = 0x14
	SizeofMsghdr           = 0x1c
	SizeofCmsghdr          = 0xc
	SizeofInet6Pktinfo     = 0x14
	SizeofIPv6MTUInfo      = 0x20
	SizeofICMPv6Filter     = 0x20
)

const (
	PTRACE_TRACEME = 0x0
	PTRACE_CONT    = 0x7
	PTRACE_KILL    = 0x8
)

type Kevent_t struct {
	Ident  uint32
	Filter int16
	Flags  uint16
	Fflags uint32
	Data   int64
	Udata  *byte
}

type FdSet struct {
	Bits [32]uint32
}

const (
	SizeofIfMsghdr         = 0xec
	SizeofIfData           = 0xd4
	SizeofIfaMsghdr        = 0x18
	SizeofIfAnnounceMsghdr = 0x1a
	SizeofRtMsghdr         = 0x60
	SizeofRtMetrics        = 0x38
)

type IfMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Hdrlen  uint16
	Index   uint16
	Tableid uint16
	Pad1    uint8
	Pad2    uint8
	Addrs   int32
	Flags   int32
	Xflags  int32
	Data    IfData
}

type IfData struct {
	Type         uint8
	Addrlen      uint8
	Hdrlen       uint8
	Link_state   uint8
	Mtu          uint32
	Metric       uint32
	Pad          uint32
	Baudrate     uint64
	Ipackets     uint64
	Ierrors      uint64
	Opackets     uint64
	Oerrors      uint64
	Collisions   uint64
	Ibytes       uint64
	Obytes       uint64
	Imcasts      uint64
	Omcasts      uint64
	Iqdrops      uint64
	Noproto      uint64
	Capabilities uint32
	Lastchange   Timeval
	Mclpool      [7]Mclpool
}

type IfaMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Hdrlen  uint16
	Index   uint16
	Tableid uint16
	Pad1    uint8
	Pad2    uint8
	Addrs   int32
	Flags   int32
	Metric  int32
}

type IfAnnounceMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Hdrlen  uint16
	Index   uint16
	What    uint16
	Name    [16]int8
}

type RtMsghdr struct {
	Msglen   uint16
	Version  uint8
	Type     uint8
	Hdrlen   uint16
	Index    uint16
	Tableid  uint16
	Priority uint8
	Mpls     uint8
	Addrs    int32
	Flags    int32
	Fmask    int32
	Pid      int32
	Seq      int32
	Errno    int32
	Inits    uint32
	Rmx      RtMetrics
}

type RtMetrics struct {
	Pksent   uint64
	Expire   int64
	Locks    uint32
	Mtu      uint32
	Refcnt   uint32
	Hopcount uint32
	Recvpipe uint32
	Sendpipe uint32
	Ssthresh uint32
	Rtt      uint32
	Rttvar   uint32
	Pad      uint32
}

type Mclpool struct {
	Grown int32
	Alive uint16
	Hwm   uint16
	Cwm   uint16
	Lwm   uint16
}

const (
	SizeofBpfVersion = 0x4
	SizeofBpfStat    = 0x8
	SizeofBpfProgram = 0x8
	SizeofBpfInsn    = 0x8
	SizeofBpfHdr     = 0x14
)

type BpfVersion struct {
	Major uint16
	Minor uint16
}

type BpfStat struct {
	Recv uint32
	Drop uint32
}

type BpfProgram struct {
	Len   uint32
	Insns *BpfInsn
}

type BpfInsn struct {
	Code uint16
	Jt   uint8
	Jf   uint8
	K    uint32
}

type BpfHdr struct {
	Tstamp    BpfTimeval
	Caplen    uint32
	Datalen   uint32
	Hdrlen    uint16
	Pad_cgo_0 [2]byte
}

type BpfTimeval struct {
	Sec  uint32
	Usec uint32
}

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]uint8
	Ispeed int32
	Ospeed int32
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

const (
	AT_FDCWD            = -0x64
	AT_SYMLINK_FOLLOW   = 0x4
	AT_SYMLINK_NOFOLLOW = 0x2
)

type PollFd struct {
	Fd      int32
	Events  int16
	Revents int16
}

const (
	POLLERR    = 0x8
	POLLHUP    = 0x10
	POLLIN     = 0x1
	POLLNVAL   = 0x20
	POLLOUT    = 0x4
	POLLPRI    = 0x2
	POLLRDBAND = 0x80
	POLLRDNORM = 0x40
	POLLWRBAND = 0x100
	POLLWRNORM = 0x4
)

type Sigset_t uint32

type Utsname struct {
	Sysname  [256]byte
	Nodename [256]byte
	Release  [256]byte
	Version  [256]byte
	Machine  [256]byte
}

const SizeofUvmexp = 0x158

type Uvmexp struct {
	Pagesize           int32
	Pagemask           int32
	Pageshift          int32
	Npages             int32
	Free               int32
	Active             int32
	Inactive           int32
	Paging             int32
	Wired              int32
	Zeropages          int32
	Reserve_pagedaemon int32
	Reserve_kernel     int32
	Anonpages          int32
	Vnodepages         int32
	Vtextpages         int32
	Freemin            int32
	Freetarg           int32
	Inactarg           int32
	Wiredmax           int32
	Anonmin            int32
	Vtextmin           int32
	Vnodemin           int32
	Anonminpct         int32
	Vtextminpct        int32
	Vnodeminpct        int32
	Nswapdev           int32
	Swpages            int32
	Swpginuse          int32
	Swpgonly           int32
	Nswget             int32
	Nanon              int32
	Nanonneeded        int32
	Nfreeanon          int32
	Faults             int32
	Traps              int32
	Intrs              int32
	Swtch              int32
	Softs              int32
	Syscalls           int32
	Pageins            int32
	Obsolete_swapins   int32
	Obsolete_swapouts  int32
	Pgswapin           int32
	Pgswapout          int32
	Forks              int32
	Forks_ppwait       int32
	Forks_sharevm      int32
	Pga_zerohit        int32
	Pga_zeromiss       int32
	Zeroaborts         int32
	Fltnoram           int32
	Fltnoanon          int32
	Fltnoamap          int32
	Fltpgwait          int32
	Fltpgrele          int32
	Fltrelck           int32
	Fltrelckok         int32
	Fltanget           int32
	Fltanretry         int32
	Fltamcopy          int32
	Fltnamap           int32
	Fltnomap           int32
	Fltlget            int32
	Fltget             int32
	Flt_anon           int32
	Flt_acow           int32
	Flt_obj            int32
	Flt_prcopy         int32
	Flt_przero         int32
	Pdwoke             int32
	Pdrevs             int32
	Pdswout            int32
	Pdfreed            int32
	Pdscans            int32
	Pdanscan           int32
	Pdobscan           int32
	Pdreact            int32
	Pdbusy             int32
	Pdpageouts         int32
	Pdpending          int32
	Pddeact            int32
	Pdreanon           int32
	Pdrevnode          int32
	Pdrevtext          int32
	Fpswtch            int32
	Kmapent            int32
}

const SizeofClockinfo = 0x14

type Clockinfo struct {
	Hz      int32
	Tick    int32
	Tickadj int32
	Stathz  int32
	Profhz  int32
}
