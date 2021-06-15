// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs_linux.go

package ipv4

const (
	sysIP_RECVTTL = 0xc

	sizeofKernelSockaddrStorage = 0x80
	sizeofSockaddrInet          = 0x10
	sizeofInetPktinfo           = 0xc
	sizeofSockExtendedErr       = 0x10

	sizeofIPMreq         = 0x8
	sizeofIPMreqn        = 0xc
	sizeofIPMreqSource   = 0xc
	sizeofGroupReq       = 0x84
	sizeofGroupSourceReq = 0x104

	sizeofICMPFilter = 0x4
)

type kernelSockaddrStorage struct {
	Family  uint16
	X__data [126]int8
}

type sockaddrInet struct {
	Family uint16
	Port   uint16
	Addr   [4]byte /* in_addr */
	X__pad [8]uint8
}

type inetPktinfo struct {
	Ifindex  int32
	Spec_dst [4]byte /* in_addr */
	Addr     [4]byte /* in_addr */
}

type sockExtendedErr struct {
	Errno  uint32
	Origin uint8
	Type   uint8
	Code   uint8
	Pad    uint8
	Info   uint32
	Data   uint32
}

type ipMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type ipMreqn struct {
	Multiaddr [4]byte /* in_addr */
	Address   [4]byte /* in_addr */
	Ifindex   int32
}

type ipMreqSource struct {
	Multiaddr  uint32
	Interface  uint32
	Sourceaddr uint32
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

type icmpFilter struct {
	Data uint32
}
