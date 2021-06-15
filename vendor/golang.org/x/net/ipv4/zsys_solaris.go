// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs_solaris.go

package ipv4

const (
	sysIP_RECVDSTADDR = 0x7
	sysIP_RECVIF      = 0x9
	sysIP_RECVTTL     = 0xb

	sizeofSockaddrStorage = 0x100
	sizeofSockaddrInet    = 0x10
	sizeofInetPktinfo     = 0xc

	sizeofIPMreq         = 0x8
	sizeofIPMreqSource   = 0xc
	sizeofGroupReq       = 0x104
	sizeofGroupSourceReq = 0x204
)

type sockaddrStorage struct {
	Family     uint16
	X_ss_pad1  [6]int8
	X_ss_align float64
	X_ss_pad2  [240]int8
}

type sockaddrInet struct {
	Family uint16
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]int8
}

type inetPktinfo struct {
	Ifindex  uint32
	Spec_dst [4]byte /* in_addr */
	Addr     [4]byte /* in_addr */
}

type ipMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type ipMreqSource struct {
	Multiaddr  [4]byte /* in_addr */
	Sourceaddr [4]byte /* in_addr */
	Interface  [4]byte /* in_addr */
}

type groupReq struct {
	Interface uint32
	Pad_cgo_0 [256]byte
}

type groupSourceReq struct {
	Interface uint32
	Pad_cgo_0 [256]byte
	Pad_cgo_1 [256]byte
}
