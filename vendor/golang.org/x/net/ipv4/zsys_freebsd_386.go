// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs_freebsd.go

package ipv4

const (
	sysIP_RECVDSTADDR = 0x7
	sysIP_RECVIF      = 0x14
	sysIP_RECVTTL     = 0x41

	sizeofSockaddrStorage = 0x80
	sizeofSockaddrInet    = 0x10

	sizeofIPMreq         = 0x8
	sizeofIPMreqn        = 0xc
	sizeofIPMreqSource   = 0xc
	sizeofGroupReq       = 0x84
	sizeofGroupSourceReq = 0x104
)

type sockaddrStorage struct {
	Len         uint8
	Family      uint8
	X__ss_pad1  [6]int8
	X__ss_align int64
	X__ss_pad2  [112]int8
}

type sockaddrInet struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]int8
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
	Multiaddr  [4]byte /* in_addr */
	Sourceaddr [4]byte /* in_addr */
	Interface  [4]byte /* in_addr */
}

type groupReq struct {
	Interface uint32
	Group     sockaddrStorage
}

type groupSourceReq struct {
	Interface uint32
	Group     sockaddrStorage
	Source    sockaddrStorage
}
