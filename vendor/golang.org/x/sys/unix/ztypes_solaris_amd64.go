// cgo -godefs types_solaris.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build amd64 && solaris
// +build amd64,solaris

package unix

const (
	SizeofPtr      = 0x8
	SizeofShort    = 0x2
	SizeofInt      = 0x4
	SizeofLong     = 0x8
	SizeofLongLong = 0x8
	PathMax        = 0x400
	MaxHostNameLen = 0x100
)

type (
	_C_short     int16
	_C_int       int32
	_C_long      int64
	_C_long_long int64
)

type Timespec struct {
	Sec  int64
	Nsec int64
}

type Timeval struct {
	Sec  int64
	Usec int64
}

type Timeval32 struct {
	Sec  int32
	Usec int32
}

type Tms struct {
	Utime  int64
	Stime  int64
	Cutime int64
	Cstime int64
}

type Utimbuf struct {
	Actime  int64
	Modtime int64
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int64
	Ixrss    int64
	Idrss    int64
	Isrss    int64
	Minflt   int64
	Majflt   int64
	Nswap    int64
	Inblock  int64
	Oublock  int64
	Msgsnd   int64
	Msgrcv   int64
	Nsignals int64
	Nvcsw    int64
	Nivcsw   int64
}

type Rlimit struct {
	Cur uint64
	Max uint64
}

type _Gid_t uint32

type Stat_t struct {
	Dev     uint64
	Ino     uint64
	Mode    uint32
	Nlink   uint32
	Uid     uint32
	Gid     uint32
	Rdev    uint64
	Size    int64
	Atim    Timespec
	Mtim    Timespec
	Ctim    Timespec
	Blksize int32
	Blocks  int64
	Fstype  [16]int8
}

type Flock_t struct {
	Type   int16
	Whence int16
	Start  int64
	Len    int64
	Sysid  int32
	Pid    int32
	Pad    [4]int64
}

type Dirent struct {
	Ino    uint64
	Off    int64
	Reclen uint16
	Name   [1]int8
	_      [5]byte
}

type _Fsblkcnt_t uint64

type Statvfs_t struct {
	Bsize    uint64
	Frsize   uint64
	Blocks   uint64
	Bfree    uint64
	Bavail   uint64
	Files    uint64
	Ffree    uint64
	Favail   uint64
	Fsid     uint64
	Basetype [16]int8
	Flag     uint64
	Namemax  uint64
	Fstr     [32]int8
}

type RawSockaddrInet4 struct {
	Family uint16
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]int8
}

type RawSockaddrInet6 struct {
	Family   uint16
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
	_        uint32
}

type RawSockaddrUnix struct {
	Family uint16
	Path   [108]int8
}

type RawSockaddrDatalink struct {
	Family uint16
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [244]int8
}

type RawSockaddr struct {
	Family uint16
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [236]int8
}

type _Socklen uint32

type Linger struct {
	Onoff  int32
	Linger int32
}

type Iovec struct {
	Base *byte
	Len  uint64
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
	Name         *byte
	Namelen      uint32
	Iov          *Iovec
	Iovlen       int32
	Accrights    *int8
	Accrightslen int32
	_            [4]byte
}

type Cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

type Inet4Pktinfo struct {
	Ifindex  uint32
	Spec_dst [4]byte /* in_addr */
	Addr     [4]byte /* in_addr */
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
	SizeofSockaddrInet6    = 0x20
	SizeofSockaddrAny      = 0xfc
	SizeofSockaddrUnix     = 0x6e
	SizeofSockaddrDatalink = 0xfc
	SizeofLinger           = 0x8
	SizeofIovec            = 0x10
	SizeofIPMreq           = 0x8
	SizeofIPv6Mreq         = 0x14
	SizeofMsghdr           = 0x30
	SizeofCmsghdr          = 0xc
	SizeofInet4Pktinfo     = 0xc
	SizeofInet6Pktinfo     = 0x14
	SizeofIPv6MTUInfo      = 0x24
	SizeofICMPv6Filter     = 0x20
)

type FdSet struct {
	Bits [1024]int64
}

type Utsname struct {
	Sysname  [257]byte
	Nodename [257]byte
	Release  [257]byte
	Version  [257]byte
	Machine  [257]byte
}

type Ustat_t struct {
	Tfree  int64
	Tinode uint64
	Fname  [6]int8
	Fpack  [6]int8
	_      [4]byte
}

const (
	AT_FDCWD            = 0xffd19553
	AT_SYMLINK_NOFOLLOW = 0x1000
	AT_SYMLINK_FOLLOW   = 0x2000
	AT_REMOVEDIR        = 0x1
	AT_EACCESS          = 0x4
)

const (
	SizeofIfMsghdr  = 0x54
	SizeofIfData    = 0x44
	SizeofIfaMsghdr = 0x14
	SizeofRtMsghdr  = 0x4c
	SizeofRtMetrics = 0x28
)

type IfMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	Data    IfData
}

type IfData struct {
	Type       uint8
	Addrlen    uint8
	Hdrlen     uint8
	Mtu        uint32
	Metric     uint32
	Baudrate   uint32
	Ipackets   uint32
	Ierrors    uint32
	Opackets   uint32
	Oerrors    uint32
	Collisions uint32
	Ibytes     uint32
	Obytes     uint32
	Imcasts    uint32
	Omcasts    uint32
	Iqdrops    uint32
	Noproto    uint32
	Lastchange Timeval32
}

type IfaMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	Metric  int32
}

type RtMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Index   uint16
	Flags   int32
	Addrs   int32
	Pid     int32
	Seq     int32
	Errno   int32
	Use     int32
	Inits   uint32
	Rmx     RtMetrics
}

type RtMetrics struct {
	Locks    uint32
	Mtu      uint32
	Hopcount uint32
	Expire   uint32
	Recvpipe uint32
	Sendpipe uint32
	Ssthresh uint32
	Rtt      uint32
	Rttvar   uint32
	Pksent   uint32
}

const (
	SizeofBpfVersion = 0x4
	SizeofBpfStat    = 0x80
	SizeofBpfProgram = 0x10
	SizeofBpfInsn    = 0x8
	SizeofBpfHdr     = 0x14
)

type BpfVersion struct {
	Major uint16
	Minor uint16
}

type BpfStat struct {
	Recv uint64
	Drop uint64
	Capt uint64
	_    [13]uint64
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

type BpfTimeval struct {
	Sec  int32
	Usec int32
}

type BpfHdr struct {
	Tstamp  BpfTimeval
	Caplen  uint32
	Datalen uint32
	Hdrlen  uint16
	_       [2]byte
}

type Termios struct {
	Iflag uint32
	Oflag uint32
	Cflag uint32
	Lflag uint32
	Cc    [19]uint8
	_     [1]byte
}

type Termio struct {
	Iflag uint16
	Oflag uint16
	Cflag uint16
	Lflag uint16
	Line  int8
	Cc    [8]uint8
	_     [1]byte
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

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

type fileObj struct {
	Atim Timespec
	Mtim Timespec
	Ctim Timespec
	Pad  [3]uint64
	Name *int8
}

type portEvent struct {
	Events int32
	Source uint16
	Pad    uint16
	Object uint64
	User   *byte
}

const (
	PORT_SOURCE_AIO    = 0x1
	PORT_SOURCE_TIMER  = 0x2
	PORT_SOURCE_USER   = 0x3
	PORT_SOURCE_FD     = 0x4
	PORT_SOURCE_ALERT  = 0x5
	PORT_SOURCE_MQ     = 0x6
	PORT_SOURCE_FILE   = 0x7
	PORT_ALERT_SET     = 0x1
	PORT_ALERT_UPDATE  = 0x2
	PORT_ALERT_INVALID = 0x3
	FILE_ACCESS        = 0x1
	FILE_MODIFIED      = 0x2
	FILE_ATTRIB        = 0x4
	FILE_TRUNC         = 0x100000
	FILE_NOFOLLOW      = 0x10000000
	FILE_DELETE        = 0x10
	FILE_RENAME_TO     = 0x20
	FILE_RENAME_FROM   = 0x40
	UNMOUNTED          = 0x20000000
	MOUNTEDOVER        = 0x40000000
	FILE_EXCEPTION     = 0x60000070
)

const (
	TUNNEWPPA = 0x540001
	TUNSETPPA = 0x540002

	I_STR     = 0x5308
	I_POP     = 0x5303
	I_PUSH    = 0x5302
	I_LINK    = 0x530c
	I_UNLINK  = 0x530d
	I_PLINK   = 0x5316
	I_PUNLINK = 0x5317

	IF_UNITSEL = -0x7ffb8cca
)

type strbuf struct {
	Maxlen int32
	Len    int32
	Buf    *int8
}

type Strioctl struct {
	Cmd    int32
	Timout int32
	Len    int32
	Dp     *int8
}

type Lifreq struct {
	Name   [32]int8
	Lifru1 [4]byte
	Type   uint32
	Lifru  [336]byte
}
