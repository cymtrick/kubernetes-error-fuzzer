// mkerrors.sh -Wall -Werror -static -I/tmp/include
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build ppc && linux
// +build ppc,linux

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- -Wall -Werror -static -I/tmp/include _const.go

package unix

import "syscall"

const (
	B1000000                         = 0x17
	B115200                          = 0x11
	B1152000                         = 0x18
	B1500000                         = 0x19
	B2000000                         = 0x1a
	B230400                          = 0x12
	B2500000                         = 0x1b
	B3000000                         = 0x1c
	B3500000                         = 0x1d
	B4000000                         = 0x1e
	B460800                          = 0x13
	B500000                          = 0x14
	B57600                           = 0x10
	B576000                          = 0x15
	B921600                          = 0x16
	BLKBSZGET                        = 0x40041270
	BLKBSZSET                        = 0x80041271
	BLKFLSBUF                        = 0x20001261
	BLKFRAGET                        = 0x20001265
	BLKFRASET                        = 0x20001264
	BLKGETSIZE                       = 0x20001260
	BLKGETSIZE64                     = 0x40041272
	BLKPBSZGET                       = 0x2000127b
	BLKRAGET                         = 0x20001263
	BLKRASET                         = 0x20001262
	BLKROGET                         = 0x2000125e
	BLKROSET                         = 0x2000125d
	BLKRRPART                        = 0x2000125f
	BLKSECTGET                       = 0x20001267
	BLKSECTSET                       = 0x20001266
	BLKSSZGET                        = 0x20001268
	BOTHER                           = 0x1f
	BS1                              = 0x8000
	BSDLY                            = 0x8000
	CBAUD                            = 0xff
	CBAUDEX                          = 0x0
	CIBAUD                           = 0xff0000
	CLOCAL                           = 0x8000
	CR1                              = 0x1000
	CR2                              = 0x2000
	CR3                              = 0x3000
	CRDLY                            = 0x3000
	CREAD                            = 0x800
	CS6                              = 0x100
	CS7                              = 0x200
	CS8                              = 0x300
	CSIZE                            = 0x300
	CSTOPB                           = 0x400
	ECCGETLAYOUT                     = 0x41484d11
	ECCGETSTATS                      = 0x40104d12
	ECHOCTL                          = 0x40
	ECHOE                            = 0x2
	ECHOK                            = 0x4
	ECHOKE                           = 0x1
	ECHONL                           = 0x10
	ECHOPRT                          = 0x20
	EFD_CLOEXEC                      = 0x80000
	EFD_NONBLOCK                     = 0x800
	EPOLL_CLOEXEC                    = 0x80000
	EXTPROC                          = 0x10000000
	FF1                              = 0x4000
	FFDLY                            = 0x4000
	FICLONE                          = 0x80049409
	FICLONERANGE                     = 0x8020940d
	FLUSHO                           = 0x800000
	FS_IOC_ENABLE_VERITY             = 0x80806685
	FS_IOC_GETFLAGS                  = 0x40046601
	FS_IOC_GET_ENCRYPTION_NONCE      = 0x4010661b
	FS_IOC_GET_ENCRYPTION_POLICY     = 0x800c6615
	FS_IOC_GET_ENCRYPTION_PWSALT     = 0x80106614
	FS_IOC_SETFLAGS                  = 0x80046602
	FS_IOC_SET_ENCRYPTION_POLICY     = 0x400c6613
	F_GETLK                          = 0xc
	F_GETLK64                        = 0xc
	F_GETOWN                         = 0x9
	F_RDLCK                          = 0x0
	F_SETLK                          = 0xd
	F_SETLK64                        = 0xd
	F_SETLKW                         = 0xe
	F_SETLKW64                       = 0xe
	F_SETOWN                         = 0x8
	F_UNLCK                          = 0x2
	F_WRLCK                          = 0x1
	HIDIOCGRAWINFO                   = 0x40084803
	HIDIOCGRDESC                     = 0x50044802
	HIDIOCGRDESCSIZE                 = 0x40044801
	HUPCL                            = 0x4000
	ICANON                           = 0x100
	IEXTEN                           = 0x400
	IN_CLOEXEC                       = 0x80000
	IN_NONBLOCK                      = 0x800
	IOCTL_VM_SOCKETS_GET_LOCAL_CID   = 0x200007b9
	ISIG                             = 0x80
	IUCLC                            = 0x1000
	IXOFF                            = 0x400
	IXON                             = 0x200
	MAP_ANON                         = 0x20
	MAP_ANONYMOUS                    = 0x20
	MAP_DENYWRITE                    = 0x800
	MAP_EXECUTABLE                   = 0x1000
	MAP_GROWSDOWN                    = 0x100
	MAP_HUGETLB                      = 0x40000
	MAP_LOCKED                       = 0x80
	MAP_NONBLOCK                     = 0x10000
	MAP_NORESERVE                    = 0x40
	MAP_POPULATE                     = 0x8000
	MAP_STACK                        = 0x20000
	MAP_SYNC                         = 0x80000
	MCL_CURRENT                      = 0x2000
	MCL_FUTURE                       = 0x4000
	MCL_ONFAULT                      = 0x8000
	MEMERASE                         = 0x80084d02
	MEMERASE64                       = 0x80104d14
	MEMGETBADBLOCK                   = 0x80084d0b
	MEMGETINFO                       = 0x40204d01
	MEMGETOOBSEL                     = 0x40c84d0a
	MEMGETREGIONCOUNT                = 0x40044d07
	MEMISLOCKED                      = 0x40084d17
	MEMLOCK                          = 0x80084d05
	MEMREADOOB                       = 0xc00c4d04
	MEMSETBADBLOCK                   = 0x80084d0c
	MEMUNLOCK                        = 0x80084d06
	MEMWRITEOOB                      = 0xc00c4d03
	MTDFILEMODE                      = 0x20004d13
	NFDBITS                          = 0x20
	NL2                              = 0x200
	NL3                              = 0x300
	NLDLY                            = 0x300
	NOFLSH                           = 0x80000000
	NS_GET_NSTYPE                    = 0x2000b703
	NS_GET_OWNER_UID                 = 0x2000b704
	NS_GET_PARENT                    = 0x2000b702
	NS_GET_USERNS                    = 0x2000b701
	OLCUC                            = 0x4
	ONLCR                            = 0x2
	OTPERASE                         = 0x800c4d19
	OTPGETREGIONCOUNT                = 0x80044d0e
	OTPGETREGIONINFO                 = 0x800c4d0f
	OTPLOCK                          = 0x400c4d10
	OTPSELECT                        = 0x40044d0d
	O_APPEND                         = 0x400
	O_ASYNC                          = 0x2000
	O_CLOEXEC                        = 0x80000
	O_CREAT                          = 0x40
	O_DIRECT                         = 0x20000
	O_DIRECTORY                      = 0x4000
	O_DSYNC                          = 0x1000
	O_EXCL                           = 0x80
	O_FSYNC                          = 0x101000
	O_LARGEFILE                      = 0x10000
	O_NDELAY                         = 0x800
	O_NOATIME                        = 0x40000
	O_NOCTTY                         = 0x100
	O_NOFOLLOW                       = 0x8000
	O_NONBLOCK                       = 0x800
	O_PATH                           = 0x200000
	O_RSYNC                          = 0x101000
	O_SYNC                           = 0x101000
	O_TMPFILE                        = 0x404000
	O_TRUNC                          = 0x200
	PARENB                           = 0x1000
	PARODD                           = 0x2000
	PENDIN                           = 0x20000000
	PERF_EVENT_IOC_DISABLE           = 0x20002401
	PERF_EVENT_IOC_ENABLE            = 0x20002400
	PERF_EVENT_IOC_ID                = 0x40042407
	PERF_EVENT_IOC_MODIFY_ATTRIBUTES = 0x8004240b
	PERF_EVENT_IOC_PAUSE_OUTPUT      = 0x80042409
	PERF_EVENT_IOC_PERIOD            = 0x80082404
	PERF_EVENT_IOC_QUERY_BPF         = 0xc004240a
	PERF_EVENT_IOC_REFRESH           = 0x20002402
	PERF_EVENT_IOC_RESET             = 0x20002403
	PERF_EVENT_IOC_SET_BPF           = 0x80042408
	PERF_EVENT_IOC_SET_FILTER        = 0x80042406
	PERF_EVENT_IOC_SET_OUTPUT        = 0x20002405
	PPPIOCATTACH                     = 0x8004743d
	PPPIOCATTCHAN                    = 0x80047438
	PPPIOCBRIDGECHAN                 = 0x80047435
	PPPIOCCONNECT                    = 0x8004743a
	PPPIOCDETACH                     = 0x8004743c
	PPPIOCDISCONN                    = 0x20007439
	PPPIOCGASYNCMAP                  = 0x40047458
	PPPIOCGCHAN                      = 0x40047437
	PPPIOCGDEBUG                     = 0x40047441
	PPPIOCGFLAGS                     = 0x4004745a
	PPPIOCGIDLE                      = 0x4008743f
	PPPIOCGIDLE32                    = 0x4008743f
	PPPIOCGIDLE64                    = 0x4010743f
	PPPIOCGL2TPSTATS                 = 0x40487436
	PPPIOCGMRU                       = 0x40047453
	PPPIOCGRASYNCMAP                 = 0x40047455
	PPPIOCGUNIT                      = 0x40047456
	PPPIOCGXASYNCMAP                 = 0x40207450
	PPPIOCSACTIVE                    = 0x80087446
	PPPIOCSASYNCMAP                  = 0x80047457
	PPPIOCSCOMPRESS                  = 0x800c744d
	PPPIOCSDEBUG                     = 0x80047440
	PPPIOCSFLAGS                     = 0x80047459
	PPPIOCSMAXCID                    = 0x80047451
	PPPIOCSMRRU                      = 0x8004743b
	PPPIOCSMRU                       = 0x80047452
	PPPIOCSNPMODE                    = 0x8008744b
	PPPIOCSPASS                      = 0x80087447
	PPPIOCSRASYNCMAP                 = 0x80047454
	PPPIOCSXASYNCMAP                 = 0x8020744f
	PPPIOCUNBRIDGECHAN               = 0x20007434
	PPPIOCXFERUNIT                   = 0x2000744e
	PROT_SAO                         = 0x10
	PR_SET_PTRACER_ANY               = 0xffffffff
	PTRACE_GETEVRREGS                = 0x14
	PTRACE_GETFPREGS                 = 0xe
	PTRACE_GETREGS64                 = 0x16
	PTRACE_GETVRREGS                 = 0x12
	PTRACE_GETVSRREGS                = 0x1b
	PTRACE_GET_DEBUGREG              = 0x19
	PTRACE_SETEVRREGS                = 0x15
	PTRACE_SETFPREGS                 = 0xf
	PTRACE_SETREGS64                 = 0x17
	PTRACE_SETVRREGS                 = 0x13
	PTRACE_SETVSRREGS                = 0x1c
	PTRACE_SET_DEBUGREG              = 0x1a
	PTRACE_SINGLEBLOCK               = 0x100
	PTRACE_SYSEMU                    = 0x1d
	PTRACE_SYSEMU_SINGLESTEP         = 0x1e
	PT_CCR                           = 0x26
	PT_CTR                           = 0x23
	PT_DAR                           = 0x29
	PT_DSCR                          = 0x2c
	PT_DSISR                         = 0x2a
	PT_FPR0                          = 0x30
	PT_FPR31                         = 0x6e
	PT_FPSCR                         = 0x71
	PT_LNK                           = 0x24
	PT_MQ                            = 0x27
	PT_MSR                           = 0x21
	PT_NIP                           = 0x20
	PT_ORIG_R3                       = 0x22
	PT_R0                            = 0x0
	PT_R1                            = 0x1
	PT_R10                           = 0xa
	PT_R11                           = 0xb
	PT_R12                           = 0xc
	PT_R13                           = 0xd
	PT_R14                           = 0xe
	PT_R15                           = 0xf
	PT_R16                           = 0x10
	PT_R17                           = 0x11
	PT_R18                           = 0x12
	PT_R19                           = 0x13
	PT_R2                            = 0x2
	PT_R20                           = 0x14
	PT_R21                           = 0x15
	PT_R22                           = 0x16
	PT_R23                           = 0x17
	PT_R24                           = 0x18
	PT_R25                           = 0x19
	PT_R26                           = 0x1a
	PT_R27                           = 0x1b
	PT_R28                           = 0x1c
	PT_R29                           = 0x1d
	PT_R3                            = 0x3
	PT_R30                           = 0x1e
	PT_R31                           = 0x1f
	PT_R4                            = 0x4
	PT_R5                            = 0x5
	PT_R6                            = 0x6
	PT_R7                            = 0x7
	PT_R8                            = 0x8
	PT_R9                            = 0x9
	PT_REGS_COUNT                    = 0x2c
	PT_RESULT                        = 0x2b
	PT_TRAP                          = 0x28
	PT_XER                           = 0x25
	RLIMIT_AS                        = 0x9
	RLIMIT_MEMLOCK                   = 0x8
	RLIMIT_NOFILE                    = 0x7
	RLIMIT_NPROC                     = 0x6
	RLIMIT_RSS                       = 0x5
	RNDADDENTROPY                    = 0x80085203
	RNDADDTOENTCNT                   = 0x80045201
	RNDCLEARPOOL                     = 0x20005206
	RNDGETENTCNT                     = 0x40045200
	RNDGETPOOL                       = 0x40085202
	RNDRESEEDCRNG                    = 0x20005207
	RNDZAPENTCNT                     = 0x20005204
	RTC_AIE_OFF                      = 0x20007002
	RTC_AIE_ON                       = 0x20007001
	RTC_ALM_READ                     = 0x40247008
	RTC_ALM_SET                      = 0x80247007
	RTC_EPOCH_READ                   = 0x4004700d
	RTC_EPOCH_SET                    = 0x8004700e
	RTC_IRQP_READ                    = 0x4004700b
	RTC_IRQP_SET                     = 0x8004700c
	RTC_PARAM_GET                    = 0x80187013
	RTC_PARAM_SET                    = 0x80187014
	RTC_PIE_OFF                      = 0x20007006
	RTC_PIE_ON                       = 0x20007005
	RTC_PLL_GET                      = 0x401c7011
	RTC_PLL_SET                      = 0x801c7012
	RTC_RD_TIME                      = 0x40247009
	RTC_SET_TIME                     = 0x8024700a
	RTC_UIE_OFF                      = 0x20007004
	RTC_UIE_ON                       = 0x20007003
	RTC_VL_CLR                       = 0x20007014
	RTC_VL_READ                      = 0x40047013
	RTC_WIE_OFF                      = 0x20007010
	RTC_WIE_ON                       = 0x2000700f
	RTC_WKALM_RD                     = 0x40287010
	RTC_WKALM_SET                    = 0x8028700f
	SCM_TIMESTAMPING                 = 0x25
	SCM_TIMESTAMPING_OPT_STATS       = 0x36
	SCM_TIMESTAMPING_PKTINFO         = 0x3a
	SCM_TIMESTAMPNS                  = 0x23
	SCM_TXTIME                       = 0x3d
	SCM_WIFI_STATUS                  = 0x29
	SFD_CLOEXEC                      = 0x80000
	SFD_NONBLOCK                     = 0x800
	SIOCATMARK                       = 0x8905
	SIOCGPGRP                        = 0x8904
	SIOCGSTAMPNS_NEW                 = 0x40108907
	SIOCGSTAMP_NEW                   = 0x40108906
	SIOCINQ                          = 0x4004667f
	SIOCOUTQ                         = 0x40047473
	SIOCSPGRP                        = 0x8902
	SOCK_CLOEXEC                     = 0x80000
	SOCK_DGRAM                       = 0x2
	SOCK_NONBLOCK                    = 0x800
	SOCK_STREAM                      = 0x1
	SOL_SOCKET                       = 0x1
	SO_ACCEPTCONN                    = 0x1e
	SO_ATTACH_BPF                    = 0x32
	SO_ATTACH_REUSEPORT_CBPF         = 0x33
	SO_ATTACH_REUSEPORT_EBPF         = 0x34
	SO_BINDTODEVICE                  = 0x19
	SO_BINDTOIFINDEX                 = 0x3e
	SO_BPF_EXTENSIONS                = 0x30
	SO_BROADCAST                     = 0x6
	SO_BSDCOMPAT                     = 0xe
	SO_BUF_LOCK                      = 0x48
	SO_BUSY_POLL                     = 0x2e
	SO_BUSY_POLL_BUDGET              = 0x46
	SO_CNX_ADVICE                    = 0x35
	SO_COOKIE                        = 0x39
	SO_DETACH_REUSEPORT_BPF          = 0x44
	SO_DOMAIN                        = 0x27
	SO_DONTROUTE                     = 0x5
	SO_ERROR                         = 0x4
	SO_INCOMING_CPU                  = 0x31
	SO_INCOMING_NAPI_ID              = 0x38
	SO_KEEPALIVE                     = 0x9
	SO_LINGER                        = 0xd
	SO_LOCK_FILTER                   = 0x2c
	SO_MARK                          = 0x24
	SO_MAX_PACING_RATE               = 0x2f
	SO_MEMINFO                       = 0x37
	SO_NETNS_COOKIE                  = 0x47
	SO_NOFCS                         = 0x2b
	SO_OOBINLINE                     = 0xa
	SO_PASSCRED                      = 0x14
	SO_PASSSEC                       = 0x22
	SO_PEEK_OFF                      = 0x2a
	SO_PEERCRED                      = 0x15
	SO_PEERGROUPS                    = 0x3b
	SO_PEERSEC                       = 0x1f
	SO_PREFER_BUSY_POLL              = 0x45
	SO_PROTOCOL                      = 0x26
	SO_RCVBUF                        = 0x8
	SO_RCVBUFFORCE                   = 0x21
	SO_RCVLOWAT                      = 0x10
	SO_RCVMARK                       = 0x4b
	SO_RCVTIMEO                      = 0x12
	SO_RCVTIMEO_NEW                  = 0x42
	SO_RCVTIMEO_OLD                  = 0x12
	SO_RESERVE_MEM                   = 0x49
	SO_REUSEADDR                     = 0x2
	SO_REUSEPORT                     = 0xf
	SO_RXQ_OVFL                      = 0x28
	SO_SECURITY_AUTHENTICATION       = 0x16
	SO_SECURITY_ENCRYPTION_NETWORK   = 0x18
	SO_SECURITY_ENCRYPTION_TRANSPORT = 0x17
	SO_SELECT_ERR_QUEUE              = 0x2d
	SO_SNDBUF                        = 0x7
	SO_SNDBUFFORCE                   = 0x20
	SO_SNDLOWAT                      = 0x11
	SO_SNDTIMEO                      = 0x13
	SO_SNDTIMEO_NEW                  = 0x43
	SO_SNDTIMEO_OLD                  = 0x13
	SO_TIMESTAMPING                  = 0x25
	SO_TIMESTAMPING_NEW              = 0x41
	SO_TIMESTAMPING_OLD              = 0x25
	SO_TIMESTAMPNS                   = 0x23
	SO_TIMESTAMPNS_NEW               = 0x40
	SO_TIMESTAMPNS_OLD               = 0x23
	SO_TIMESTAMP_NEW                 = 0x3f
	SO_TXREHASH                      = 0x4a
	SO_TXTIME                        = 0x3d
	SO_TYPE                          = 0x3
	SO_WIFI_STATUS                   = 0x29
	SO_ZEROCOPY                      = 0x3c
	TAB1                             = 0x400
	TAB2                             = 0x800
	TAB3                             = 0xc00
	TABDLY                           = 0xc00
	TCFLSH                           = 0x2000741f
	TCGETA                           = 0x40147417
	TCGETS                           = 0x402c7413
	TCSAFLUSH                        = 0x2
	TCSBRK                           = 0x2000741d
	TCSBRKP                          = 0x5425
	TCSETA                           = 0x80147418
	TCSETAF                          = 0x8014741c
	TCSETAW                          = 0x80147419
	TCSETS                           = 0x802c7414
	TCSETSF                          = 0x802c7416
	TCSETSW                          = 0x802c7415
	TCXONC                           = 0x2000741e
	TFD_CLOEXEC                      = 0x80000
	TFD_NONBLOCK                     = 0x800
	TIOCCBRK                         = 0x5428
	TIOCCONS                         = 0x541d
	TIOCEXCL                         = 0x540c
	TIOCGDEV                         = 0x40045432
	TIOCGETC                         = 0x40067412
	TIOCGETD                         = 0x5424
	TIOCGETP                         = 0x40067408
	TIOCGEXCL                        = 0x40045440
	TIOCGICOUNT                      = 0x545d
	TIOCGISO7816                     = 0x40285442
	TIOCGLCKTRMIOS                   = 0x5456
	TIOCGLTC                         = 0x40067474
	TIOCGPGRP                        = 0x40047477
	TIOCGPKT                         = 0x40045438
	TIOCGPTLCK                       = 0x40045439
	TIOCGPTN                         = 0x40045430
	TIOCGPTPEER                      = 0x20005441
	TIOCGRS485                       = 0x542e
	TIOCGSERIAL                      = 0x541e
	TIOCGSID                         = 0x5429
	TIOCGSOFTCAR                     = 0x5419
	TIOCGWINSZ                       = 0x40087468
	TIOCINQ                          = 0x4004667f
	TIOCLINUX                        = 0x541c
	TIOCMBIC                         = 0x5417
	TIOCMBIS                         = 0x5416
	TIOCMGET                         = 0x5415
	TIOCMIWAIT                       = 0x545c
	TIOCMSET                         = 0x5418
	TIOCM_CAR                        = 0x40
	TIOCM_CD                         = 0x40
	TIOCM_CTS                        = 0x20
	TIOCM_DSR                        = 0x100
	TIOCM_LOOP                       = 0x8000
	TIOCM_OUT1                       = 0x2000
	TIOCM_OUT2                       = 0x4000
	TIOCM_RI                         = 0x80
	TIOCM_RNG                        = 0x80
	TIOCM_SR                         = 0x10
	TIOCM_ST                         = 0x8
	TIOCNOTTY                        = 0x5422
	TIOCNXCL                         = 0x540d
	TIOCOUTQ                         = 0x40047473
	TIOCPKT                          = 0x5420
	TIOCSBRK                         = 0x5427
	TIOCSCTTY                        = 0x540e
	TIOCSERCONFIG                    = 0x5453
	TIOCSERGETLSR                    = 0x5459
	TIOCSERGETMULTI                  = 0x545a
	TIOCSERGSTRUCT                   = 0x5458
	TIOCSERGWILD                     = 0x5454
	TIOCSERSETMULTI                  = 0x545b
	TIOCSERSWILD                     = 0x5455
	TIOCSER_TEMT                     = 0x1
	TIOCSETC                         = 0x80067411
	TIOCSETD                         = 0x5423
	TIOCSETN                         = 0x8006740a
	TIOCSETP                         = 0x80067409
	TIOCSIG                          = 0x80045436
	TIOCSISO7816                     = 0xc0285443
	TIOCSLCKTRMIOS                   = 0x5457
	TIOCSLTC                         = 0x80067475
	TIOCSPGRP                        = 0x80047476
	TIOCSPTLCK                       = 0x80045431
	TIOCSRS485                       = 0x542f
	TIOCSSERIAL                      = 0x541f
	TIOCSSOFTCAR                     = 0x541a
	TIOCSTART                        = 0x2000746e
	TIOCSTI                          = 0x5412
	TIOCSTOP                         = 0x2000746f
	TIOCSWINSZ                       = 0x80087467
	TIOCVHANGUP                      = 0x5437
	TOSTOP                           = 0x400000
	TUNATTACHFILTER                  = 0x800854d5
	TUNDETACHFILTER                  = 0x800854d6
	TUNGETDEVNETNS                   = 0x200054e3
	TUNGETFEATURES                   = 0x400454cf
	TUNGETFILTER                     = 0x400854db
	TUNGETIFF                        = 0x400454d2
	TUNGETSNDBUF                     = 0x400454d3
	TUNGETVNETBE                     = 0x400454df
	TUNGETVNETHDRSZ                  = 0x400454d7
	TUNGETVNETLE                     = 0x400454dd
	TUNSETCARRIER                    = 0x800454e2
	TUNSETDEBUG                      = 0x800454c9
	TUNSETFILTEREBPF                 = 0x400454e1
	TUNSETGROUP                      = 0x800454ce
	TUNSETIFF                        = 0x800454ca
	TUNSETIFINDEX                    = 0x800454da
	TUNSETLINK                       = 0x800454cd
	TUNSETNOCSUM                     = 0x800454c8
	TUNSETOFFLOAD                    = 0x800454d0
	TUNSETOWNER                      = 0x800454cc
	TUNSETPERSIST                    = 0x800454cb
	TUNSETQUEUE                      = 0x800454d9
	TUNSETSNDBUF                     = 0x800454d4
	TUNSETSTEERINGEBPF               = 0x400454e0
	TUNSETTXFILTER                   = 0x800454d1
	TUNSETVNETBE                     = 0x800454de
	TUNSETVNETHDRSZ                  = 0x800454d8
	TUNSETVNETLE                     = 0x800454dc
	UBI_IOCATT                       = 0x80186f40
	UBI_IOCDET                       = 0x80046f41
	UBI_IOCEBCH                      = 0x80044f02
	UBI_IOCEBER                      = 0x80044f01
	UBI_IOCEBISMAP                   = 0x40044f05
	UBI_IOCEBMAP                     = 0x80084f03
	UBI_IOCEBUNMAP                   = 0x80044f04
	UBI_IOCMKVOL                     = 0x80986f00
	UBI_IOCRMVOL                     = 0x80046f01
	UBI_IOCRNVOL                     = 0x91106f03
	UBI_IOCRPEB                      = 0x80046f04
	UBI_IOCRSVOL                     = 0x800c6f02
	UBI_IOCSETVOLPROP                = 0x80104f06
	UBI_IOCSPEB                      = 0x80046f05
	UBI_IOCVOLCRBLK                  = 0x80804f07
	UBI_IOCVOLRMBLK                  = 0x20004f08
	UBI_IOCVOLUP                     = 0x80084f00
	VDISCARD                         = 0x10
	VEOF                             = 0x4
	VEOL                             = 0x6
	VEOL2                            = 0x8
	VMIN                             = 0x5
	VREPRINT                         = 0xb
	VSTART                           = 0xd
	VSTOP                            = 0xe
	VSUSP                            = 0xc
	VSWTC                            = 0x9
	VT1                              = 0x10000
	VTDLY                            = 0x10000
	VTIME                            = 0x7
	VWERASE                          = 0xa
	WDIOC_GETBOOTSTATUS              = 0x40045702
	WDIOC_GETPRETIMEOUT              = 0x40045709
	WDIOC_GETSTATUS                  = 0x40045701
	WDIOC_GETSUPPORT                 = 0x40285700
	WDIOC_GETTEMP                    = 0x40045703
	WDIOC_GETTIMELEFT                = 0x4004570a
	WDIOC_GETTIMEOUT                 = 0x40045707
	WDIOC_KEEPALIVE                  = 0x40045705
	WDIOC_SETOPTIONS                 = 0x40045704
	WORDSIZE                         = 0x20
	XCASE                            = 0x4000
	XTABS                            = 0xc00
	_HIDIOCGRAWNAME                  = 0x40804804
	_HIDIOCGRAWPHYS                  = 0x40404805
	_HIDIOCGRAWUNIQ                  = 0x40404808
)

// Errors
const (
	EADDRINUSE      = syscall.Errno(0x62)
	EADDRNOTAVAIL   = syscall.Errno(0x63)
	EADV            = syscall.Errno(0x44)
	EAFNOSUPPORT    = syscall.Errno(0x61)
	EALREADY        = syscall.Errno(0x72)
	EBADE           = syscall.Errno(0x34)
	EBADFD          = syscall.Errno(0x4d)
	EBADMSG         = syscall.Errno(0x4a)
	EBADR           = syscall.Errno(0x35)
	EBADRQC         = syscall.Errno(0x38)
	EBADSLT         = syscall.Errno(0x39)
	EBFONT          = syscall.Errno(0x3b)
	ECANCELED       = syscall.Errno(0x7d)
	ECHRNG          = syscall.Errno(0x2c)
	ECOMM           = syscall.Errno(0x46)
	ECONNABORTED    = syscall.Errno(0x67)
	ECONNREFUSED    = syscall.Errno(0x6f)
	ECONNRESET      = syscall.Errno(0x68)
	EDEADLK         = syscall.Errno(0x23)
	EDEADLOCK       = syscall.Errno(0x3a)
	EDESTADDRREQ    = syscall.Errno(0x59)
	EDOTDOT         = syscall.Errno(0x49)
	EDQUOT          = syscall.Errno(0x7a)
	EHOSTDOWN       = syscall.Errno(0x70)
	EHOSTUNREACH    = syscall.Errno(0x71)
	EHWPOISON       = syscall.Errno(0x85)
	EIDRM           = syscall.Errno(0x2b)
	EILSEQ          = syscall.Errno(0x54)
	EINPROGRESS     = syscall.Errno(0x73)
	EISCONN         = syscall.Errno(0x6a)
	EISNAM          = syscall.Errno(0x78)
	EKEYEXPIRED     = syscall.Errno(0x7f)
	EKEYREJECTED    = syscall.Errno(0x81)
	EKEYREVOKED     = syscall.Errno(0x80)
	EL2HLT          = syscall.Errno(0x33)
	EL2NSYNC        = syscall.Errno(0x2d)
	EL3HLT          = syscall.Errno(0x2e)
	EL3RST          = syscall.Errno(0x2f)
	ELIBACC         = syscall.Errno(0x4f)
	ELIBBAD         = syscall.Errno(0x50)
	ELIBEXEC        = syscall.Errno(0x53)
	ELIBMAX         = syscall.Errno(0x52)
	ELIBSCN         = syscall.Errno(0x51)
	ELNRNG          = syscall.Errno(0x30)
	ELOOP           = syscall.Errno(0x28)
	EMEDIUMTYPE     = syscall.Errno(0x7c)
	EMSGSIZE        = syscall.Errno(0x5a)
	EMULTIHOP       = syscall.Errno(0x48)
	ENAMETOOLONG    = syscall.Errno(0x24)
	ENAVAIL         = syscall.Errno(0x77)
	ENETDOWN        = syscall.Errno(0x64)
	ENETRESET       = syscall.Errno(0x66)
	ENETUNREACH     = syscall.Errno(0x65)
	ENOANO          = syscall.Errno(0x37)
	ENOBUFS         = syscall.Errno(0x69)
	ENOCSI          = syscall.Errno(0x32)
	ENODATA         = syscall.Errno(0x3d)
	ENOKEY          = syscall.Errno(0x7e)
	ENOLCK          = syscall.Errno(0x25)
	ENOLINK         = syscall.Errno(0x43)
	ENOMEDIUM       = syscall.Errno(0x7b)
	ENOMSG          = syscall.Errno(0x2a)
	ENONET          = syscall.Errno(0x40)
	ENOPKG          = syscall.Errno(0x41)
	ENOPROTOOPT     = syscall.Errno(0x5c)
	ENOSR           = syscall.Errno(0x3f)
	ENOSTR          = syscall.Errno(0x3c)
	ENOSYS          = syscall.Errno(0x26)
	ENOTCONN        = syscall.Errno(0x6b)
	ENOTEMPTY       = syscall.Errno(0x27)
	ENOTNAM         = syscall.Errno(0x76)
	ENOTRECOVERABLE = syscall.Errno(0x83)
	ENOTSOCK        = syscall.Errno(0x58)
	ENOTSUP         = syscall.Errno(0x5f)
	ENOTUNIQ        = syscall.Errno(0x4c)
	EOPNOTSUPP      = syscall.Errno(0x5f)
	EOVERFLOW       = syscall.Errno(0x4b)
	EOWNERDEAD      = syscall.Errno(0x82)
	EPFNOSUPPORT    = syscall.Errno(0x60)
	EPROTO          = syscall.Errno(0x47)
	EPROTONOSUPPORT = syscall.Errno(0x5d)
	EPROTOTYPE      = syscall.Errno(0x5b)
	EREMCHG         = syscall.Errno(0x4e)
	EREMOTE         = syscall.Errno(0x42)
	EREMOTEIO       = syscall.Errno(0x79)
	ERESTART        = syscall.Errno(0x55)
	ERFKILL         = syscall.Errno(0x84)
	ESHUTDOWN       = syscall.Errno(0x6c)
	ESOCKTNOSUPPORT = syscall.Errno(0x5e)
	ESRMNT          = syscall.Errno(0x45)
	ESTALE          = syscall.Errno(0x74)
	ESTRPIPE        = syscall.Errno(0x56)
	ETIME           = syscall.Errno(0x3e)
	ETIMEDOUT       = syscall.Errno(0x6e)
	ETOOMANYREFS    = syscall.Errno(0x6d)
	EUCLEAN         = syscall.Errno(0x75)
	EUNATCH         = syscall.Errno(0x31)
	EUSERS          = syscall.Errno(0x57)
	EXFULL          = syscall.Errno(0x36)
)

// Signals
const (
	SIGBUS    = syscall.Signal(0x7)
	SIGCHLD   = syscall.Signal(0x11)
	SIGCLD    = syscall.Signal(0x11)
	SIGCONT   = syscall.Signal(0x12)
	SIGIO     = syscall.Signal(0x1d)
	SIGPOLL   = syscall.Signal(0x1d)
	SIGPROF   = syscall.Signal(0x1b)
	SIGPWR    = syscall.Signal(0x1e)
	SIGSTKFLT = syscall.Signal(0x10)
	SIGSTOP   = syscall.Signal(0x13)
	SIGSYS    = syscall.Signal(0x1f)
	SIGTSTP   = syscall.Signal(0x14)
	SIGTTIN   = syscall.Signal(0x15)
	SIGTTOU   = syscall.Signal(0x16)
	SIGURG    = syscall.Signal(0x17)
	SIGUSR1   = syscall.Signal(0xa)
	SIGUSR2   = syscall.Signal(0xc)
	SIGVTALRM = syscall.Signal(0x1a)
	SIGWINCH  = syscall.Signal(0x1c)
	SIGXCPU   = syscall.Signal(0x18)
	SIGXFSZ   = syscall.Signal(0x19)
)

// Error table
var errorList = [...]struct {
	num  syscall.Errno
	name string
	desc string
}{
	{1, "EPERM", "operation not permitted"},
	{2, "ENOENT", "no such file or directory"},
	{3, "ESRCH", "no such process"},
	{4, "EINTR", "interrupted system call"},
	{5, "EIO", "input/output error"},
	{6, "ENXIO", "no such device or address"},
	{7, "E2BIG", "argument list too long"},
	{8, "ENOEXEC", "exec format error"},
	{9, "EBADF", "bad file descriptor"},
	{10, "ECHILD", "no child processes"},
	{11, "EAGAIN", "resource temporarily unavailable"},
	{12, "ENOMEM", "cannot allocate memory"},
	{13, "EACCES", "permission denied"},
	{14, "EFAULT", "bad address"},
	{15, "ENOTBLK", "block device required"},
	{16, "EBUSY", "device or resource busy"},
	{17, "EEXIST", "file exists"},
	{18, "EXDEV", "invalid cross-device link"},
	{19, "ENODEV", "no such device"},
	{20, "ENOTDIR", "not a directory"},
	{21, "EISDIR", "is a directory"},
	{22, "EINVAL", "invalid argument"},
	{23, "ENFILE", "too many open files in system"},
	{24, "EMFILE", "too many open files"},
	{25, "ENOTTY", "inappropriate ioctl for device"},
	{26, "ETXTBSY", "text file busy"},
	{27, "EFBIG", "file too large"},
	{28, "ENOSPC", "no space left on device"},
	{29, "ESPIPE", "illegal seek"},
	{30, "EROFS", "read-only file system"},
	{31, "EMLINK", "too many links"},
	{32, "EPIPE", "broken pipe"},
	{33, "EDOM", "numerical argument out of domain"},
	{34, "ERANGE", "numerical result out of range"},
	{35, "EDEADLK", "resource deadlock avoided"},
	{36, "ENAMETOOLONG", "file name too long"},
	{37, "ENOLCK", "no locks available"},
	{38, "ENOSYS", "function not implemented"},
	{39, "ENOTEMPTY", "directory not empty"},
	{40, "ELOOP", "too many levels of symbolic links"},
	{42, "ENOMSG", "no message of desired type"},
	{43, "EIDRM", "identifier removed"},
	{44, "ECHRNG", "channel number out of range"},
	{45, "EL2NSYNC", "level 2 not synchronized"},
	{46, "EL3HLT", "level 3 halted"},
	{47, "EL3RST", "level 3 reset"},
	{48, "ELNRNG", "link number out of range"},
	{49, "EUNATCH", "protocol driver not attached"},
	{50, "ENOCSI", "no CSI structure available"},
	{51, "EL2HLT", "level 2 halted"},
	{52, "EBADE", "invalid exchange"},
	{53, "EBADR", "invalid request descriptor"},
	{54, "EXFULL", "exchange full"},
	{55, "ENOANO", "no anode"},
	{56, "EBADRQC", "invalid request code"},
	{57, "EBADSLT", "invalid slot"},
	{58, "EDEADLOCK", "file locking deadlock error"},
	{59, "EBFONT", "bad font file format"},
	{60, "ENOSTR", "device not a stream"},
	{61, "ENODATA", "no data available"},
	{62, "ETIME", "timer expired"},
	{63, "ENOSR", "out of streams resources"},
	{64, "ENONET", "machine is not on the network"},
	{65, "ENOPKG", "package not installed"},
	{66, "EREMOTE", "object is remote"},
	{67, "ENOLINK", "link has been severed"},
	{68, "EADV", "advertise error"},
	{69, "ESRMNT", "srmount error"},
	{70, "ECOMM", "communication error on send"},
	{71, "EPROTO", "protocol error"},
	{72, "EMULTIHOP", "multihop attempted"},
	{73, "EDOTDOT", "RFS specific error"},
	{74, "EBADMSG", "bad message"},
	{75, "EOVERFLOW", "value too large for defined data type"},
	{76, "ENOTUNIQ", "name not unique on network"},
	{77, "EBADFD", "file descriptor in bad state"},
	{78, "EREMCHG", "remote address changed"},
	{79, "ELIBACC", "can not access a needed shared library"},
	{80, "ELIBBAD", "accessing a corrupted shared library"},
	{81, "ELIBSCN", ".lib section in a.out corrupted"},
	{82, "ELIBMAX", "attempting to link in too many shared libraries"},
	{83, "ELIBEXEC", "cannot exec a shared library directly"},
	{84, "EILSEQ", "invalid or incomplete multibyte or wide character"},
	{85, "ERESTART", "interrupted system call should be restarted"},
	{86, "ESTRPIPE", "streams pipe error"},
	{87, "EUSERS", "too many users"},
	{88, "ENOTSOCK", "socket operation on non-socket"},
	{89, "EDESTADDRREQ", "destination address required"},
	{90, "EMSGSIZE", "message too long"},
	{91, "EPROTOTYPE", "protocol wrong type for socket"},
	{92, "ENOPROTOOPT", "protocol not available"},
	{93, "EPROTONOSUPPORT", "protocol not supported"},
	{94, "ESOCKTNOSUPPORT", "socket type not supported"},
	{95, "ENOTSUP", "operation not supported"},
	{96, "EPFNOSUPPORT", "protocol family not supported"},
	{97, "EAFNOSUPPORT", "address family not supported by protocol"},
	{98, "EADDRINUSE", "address already in use"},
	{99, "EADDRNOTAVAIL", "cannot assign requested address"},
	{100, "ENETDOWN", "network is down"},
	{101, "ENETUNREACH", "network is unreachable"},
	{102, "ENETRESET", "network dropped connection on reset"},
	{103, "ECONNABORTED", "software caused connection abort"},
	{104, "ECONNRESET", "connection reset by peer"},
	{105, "ENOBUFS", "no buffer space available"},
	{106, "EISCONN", "transport endpoint is already connected"},
	{107, "ENOTCONN", "transport endpoint is not connected"},
	{108, "ESHUTDOWN", "cannot send after transport endpoint shutdown"},
	{109, "ETOOMANYREFS", "too many references: cannot splice"},
	{110, "ETIMEDOUT", "connection timed out"},
	{111, "ECONNREFUSED", "connection refused"},
	{112, "EHOSTDOWN", "host is down"},
	{113, "EHOSTUNREACH", "no route to host"},
	{114, "EALREADY", "operation already in progress"},
	{115, "EINPROGRESS", "operation now in progress"},
	{116, "ESTALE", "stale file handle"},
	{117, "EUCLEAN", "structure needs cleaning"},
	{118, "ENOTNAM", "not a XENIX named type file"},
	{119, "ENAVAIL", "no XENIX semaphores available"},
	{120, "EISNAM", "is a named type file"},
	{121, "EREMOTEIO", "remote I/O error"},
	{122, "EDQUOT", "disk quota exceeded"},
	{123, "ENOMEDIUM", "no medium found"},
	{124, "EMEDIUMTYPE", "wrong medium type"},
	{125, "ECANCELED", "operation canceled"},
	{126, "ENOKEY", "required key not available"},
	{127, "EKEYEXPIRED", "key has expired"},
	{128, "EKEYREVOKED", "key has been revoked"},
	{129, "EKEYREJECTED", "key was rejected by service"},
	{130, "EOWNERDEAD", "owner died"},
	{131, "ENOTRECOVERABLE", "state not recoverable"},
	{132, "ERFKILL", "operation not possible due to RF-kill"},
	{133, "EHWPOISON", "memory page has hardware error"},
}

// Signal table
var signalList = [...]struct {
	num  syscall.Signal
	name string
	desc string
}{
	{1, "SIGHUP", "hangup"},
	{2, "SIGINT", "interrupt"},
	{3, "SIGQUIT", "quit"},
	{4, "SIGILL", "illegal instruction"},
	{5, "SIGTRAP", "trace/breakpoint trap"},
	{6, "SIGABRT", "aborted"},
	{7, "SIGBUS", "bus error"},
	{8, "SIGFPE", "floating point exception"},
	{9, "SIGKILL", "killed"},
	{10, "SIGUSR1", "user defined signal 1"},
	{11, "SIGSEGV", "segmentation fault"},
	{12, "SIGUSR2", "user defined signal 2"},
	{13, "SIGPIPE", "broken pipe"},
	{14, "SIGALRM", "alarm clock"},
	{15, "SIGTERM", "terminated"},
	{16, "SIGSTKFLT", "stack fault"},
	{17, "SIGCHLD", "child exited"},
	{18, "SIGCONT", "continued"},
	{19, "SIGSTOP", "stopped (signal)"},
	{20, "SIGTSTP", "stopped"},
	{21, "SIGTTIN", "stopped (tty input)"},
	{22, "SIGTTOU", "stopped (tty output)"},
	{23, "SIGURG", "urgent I/O condition"},
	{24, "SIGXCPU", "CPU time limit exceeded"},
	{25, "SIGXFSZ", "file size limit exceeded"},
	{26, "SIGVTALRM", "virtual timer expired"},
	{27, "SIGPROF", "profiling timer expired"},
	{28, "SIGWINCH", "window changed"},
	{29, "SIGIO", "I/O possible"},
	{30, "SIGPWR", "power failure"},
	{31, "SIGSYS", "bad system call"},
}
