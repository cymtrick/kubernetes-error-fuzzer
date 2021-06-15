// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs_linux.go

package ipv6

const (
	sysIPV6_NEXTHOP  = 0x9
	sysIPV6_PKTINFO  = 0x32
	sysIPV6_HOPLIMIT = 0x34
	sysIPV6_PATHMTU  = 0x3d
	sysIPV6_TCLASS   = 0x43

	sizeofKernelSockaddrStorage = 0x80
	sizeofSockaddrInet6         = 0x1c
	sizeofInet6Pktinfo          = 0x14
	sizeofIPv6Mtuinfo           = 0x20
	sizeofIPv6FlowlabelReq      = 0x20

	sizeofIPv6Mreq       = 0x14
	sizeofGroupReq       = 0x84
	sizeofGroupSourceReq = 0x104

	sizeofICMPv6Filter = 0x20
)

type kernelSockaddrStorage struct {
	Family  uint16
	X__data [126]uint8
}

type sockaddrInet6 struct {
	Family   uint16
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

type inet6Pktinfo struct {
	Addr    [16]byte /* in6_addr */
	Ifindex int32
}

type ipv6Mtuinfo struct {
	Addr sockaddrInet6
	Mtu  uint32
}

type ipv6FlowlabelReq struct {
	Dst        [16]byte /* in6_addr */
	Label      uint32
	Action     uint8
	Share      uint8
	Flags      uint16
	Expires    uint16
	Linger     uint16
	X__flr_pad uint32
}

type ipv6Mreq struct {
	Multiaddr [16]byte /* in6_addr */
	Ifindex   int32
}

type groupReq struct {
	Interface uint32
	Group     kernelSockaddrStorage
}

type groupSourceReq struct {
	Interface uint32
	Group     kernelSockaddrStorage
	Source    kernelSockaddrStorage
}

type icmpv6Filter struct {
	Data [8]uint32
}
