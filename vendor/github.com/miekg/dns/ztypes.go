// Code generated by "go run types_generate.go"; DO NOT EDIT.

package dns

import (
	"encoding/base64"
	"net"
)

// TypeToRR is a map of constructors for each RR type.
var TypeToRR = map[uint16]func() RR{
	TypeA:          func() RR { return new(A) },
	TypeAAAA:       func() RR { return new(AAAA) },
	TypeAFSDB:      func() RR { return new(AFSDB) },
	TypeANY:        func() RR { return new(ANY) },
	TypeAVC:        func() RR { return new(AVC) },
	TypeCAA:        func() RR { return new(CAA) },
	TypeCDNSKEY:    func() RR { return new(CDNSKEY) },
	TypeCDS:        func() RR { return new(CDS) },
	TypeCERT:       func() RR { return new(CERT) },
	TypeCNAME:      func() RR { return new(CNAME) },
	TypeCSYNC:      func() RR { return new(CSYNC) },
	TypeDHCID:      func() RR { return new(DHCID) },
	TypeDLV:        func() RR { return new(DLV) },
	TypeDNAME:      func() RR { return new(DNAME) },
	TypeDNSKEY:     func() RR { return new(DNSKEY) },
	TypeDS:         func() RR { return new(DS) },
	TypeEID:        func() RR { return new(EID) },
	TypeEUI48:      func() RR { return new(EUI48) },
	TypeEUI64:      func() RR { return new(EUI64) },
	TypeGID:        func() RR { return new(GID) },
	TypeGPOS:       func() RR { return new(GPOS) },
	TypeHINFO:      func() RR { return new(HINFO) },
	TypeHIP:        func() RR { return new(HIP) },
	TypeKEY:        func() RR { return new(KEY) },
	TypeKX:         func() RR { return new(KX) },
	TypeL32:        func() RR { return new(L32) },
	TypeL64:        func() RR { return new(L64) },
	TypeLOC:        func() RR { return new(LOC) },
	TypeLP:         func() RR { return new(LP) },
	TypeMB:         func() RR { return new(MB) },
	TypeMD:         func() RR { return new(MD) },
	TypeMF:         func() RR { return new(MF) },
	TypeMG:         func() RR { return new(MG) },
	TypeMINFO:      func() RR { return new(MINFO) },
	TypeMR:         func() RR { return new(MR) },
	TypeMX:         func() RR { return new(MX) },
	TypeNAPTR:      func() RR { return new(NAPTR) },
	TypeNID:        func() RR { return new(NID) },
	TypeNIMLOC:     func() RR { return new(NIMLOC) },
	TypeNINFO:      func() RR { return new(NINFO) },
	TypeNS:         func() RR { return new(NS) },
	TypeNSAPPTR:    func() RR { return new(NSAPPTR) },
	TypeNSEC:       func() RR { return new(NSEC) },
	TypeNSEC3:      func() RR { return new(NSEC3) },
	TypeNSEC3PARAM: func() RR { return new(NSEC3PARAM) },
	TypeNULL:       func() RR { return new(NULL) },
	TypeOPENPGPKEY: func() RR { return new(OPENPGPKEY) },
	TypeOPT:        func() RR { return new(OPT) },
	TypePTR:        func() RR { return new(PTR) },
	TypePX:         func() RR { return new(PX) },
	TypeRKEY:       func() RR { return new(RKEY) },
	TypeRP:         func() RR { return new(RP) },
	TypeRRSIG:      func() RR { return new(RRSIG) },
	TypeRT:         func() RR { return new(RT) },
	TypeSIG:        func() RR { return new(SIG) },
	TypeSMIMEA:     func() RR { return new(SMIMEA) },
	TypeSOA:        func() RR { return new(SOA) },
	TypeSPF:        func() RR { return new(SPF) },
	TypeSRV:        func() RR { return new(SRV) },
	TypeSSHFP:      func() RR { return new(SSHFP) },
	TypeTA:         func() RR { return new(TA) },
	TypeTALINK:     func() RR { return new(TALINK) },
	TypeTKEY:       func() RR { return new(TKEY) },
	TypeTLSA:       func() RR { return new(TLSA) },
	TypeTSIG:       func() RR { return new(TSIG) },
	TypeTXT:        func() RR { return new(TXT) },
	TypeUID:        func() RR { return new(UID) },
	TypeUINFO:      func() RR { return new(UINFO) },
	TypeURI:        func() RR { return new(URI) },
	TypeX25:        func() RR { return new(X25) },
}

// TypeToString is a map of strings for each RR type.
var TypeToString = map[uint16]string{
	TypeA:          "A",
	TypeAAAA:       "AAAA",
	TypeAFSDB:      "AFSDB",
	TypeANY:        "ANY",
	TypeATMA:       "ATMA",
	TypeAVC:        "AVC",
	TypeAXFR:       "AXFR",
	TypeCAA:        "CAA",
	TypeCDNSKEY:    "CDNSKEY",
	TypeCDS:        "CDS",
	TypeCERT:       "CERT",
	TypeCNAME:      "CNAME",
	TypeCSYNC:      "CSYNC",
	TypeDHCID:      "DHCID",
	TypeDLV:        "DLV",
	TypeDNAME:      "DNAME",
	TypeDNSKEY:     "DNSKEY",
	TypeDS:         "DS",
	TypeEID:        "EID",
	TypeEUI48:      "EUI48",
	TypeEUI64:      "EUI64",
	TypeGID:        "GID",
	TypeGPOS:       "GPOS",
	TypeHINFO:      "HINFO",
	TypeHIP:        "HIP",
	TypeISDN:       "ISDN",
	TypeIXFR:       "IXFR",
	TypeKEY:        "KEY",
	TypeKX:         "KX",
	TypeL32:        "L32",
	TypeL64:        "L64",
	TypeLOC:        "LOC",
	TypeLP:         "LP",
	TypeMAILA:      "MAILA",
	TypeMAILB:      "MAILB",
	TypeMB:         "MB",
	TypeMD:         "MD",
	TypeMF:         "MF",
	TypeMG:         "MG",
	TypeMINFO:      "MINFO",
	TypeMR:         "MR",
	TypeMX:         "MX",
	TypeNAPTR:      "NAPTR",
	TypeNID:        "NID",
	TypeNIMLOC:     "NIMLOC",
	TypeNINFO:      "NINFO",
	TypeNS:         "NS",
	TypeNSEC:       "NSEC",
	TypeNSEC3:      "NSEC3",
	TypeNSEC3PARAM: "NSEC3PARAM",
	TypeNULL:       "NULL",
	TypeNXT:        "NXT",
	TypeNone:       "None",
	TypeOPENPGPKEY: "OPENPGPKEY",
	TypeOPT:        "OPT",
	TypePTR:        "PTR",
	TypePX:         "PX",
	TypeRKEY:       "RKEY",
	TypeRP:         "RP",
	TypeRRSIG:      "RRSIG",
	TypeRT:         "RT",
	TypeReserved:   "Reserved",
	TypeSIG:        "SIG",
	TypeSMIMEA:     "SMIMEA",
	TypeSOA:        "SOA",
	TypeSPF:        "SPF",
	TypeSRV:        "SRV",
	TypeSSHFP:      "SSHFP",
	TypeTA:         "TA",
	TypeTALINK:     "TALINK",
	TypeTKEY:       "TKEY",
	TypeTLSA:       "TLSA",
	TypeTSIG:       "TSIG",
	TypeTXT:        "TXT",
	TypeUID:        "UID",
	TypeUINFO:      "UINFO",
	TypeUNSPEC:     "UNSPEC",
	TypeURI:        "URI",
	TypeX25:        "X25",
	TypeNSAPPTR:    "NSAP-PTR",
}

func (rr *A) Header() *RR_Header          { return &rr.Hdr }
func (rr *AAAA) Header() *RR_Header       { return &rr.Hdr }
func (rr *AFSDB) Header() *RR_Header      { return &rr.Hdr }
func (rr *ANY) Header() *RR_Header        { return &rr.Hdr }
func (rr *AVC) Header() *RR_Header        { return &rr.Hdr }
func (rr *CAA) Header() *RR_Header        { return &rr.Hdr }
func (rr *CDNSKEY) Header() *RR_Header    { return &rr.Hdr }
func (rr *CDS) Header() *RR_Header        { return &rr.Hdr }
func (rr *CERT) Header() *RR_Header       { return &rr.Hdr }
func (rr *CNAME) Header() *RR_Header      { return &rr.Hdr }
func (rr *CSYNC) Header() *RR_Header      { return &rr.Hdr }
func (rr *DHCID) Header() *RR_Header      { return &rr.Hdr }
func (rr *DLV) Header() *RR_Header        { return &rr.Hdr }
func (rr *DNAME) Header() *RR_Header      { return &rr.Hdr }
func (rr *DNSKEY) Header() *RR_Header     { return &rr.Hdr }
func (rr *DS) Header() *RR_Header         { return &rr.Hdr }
func (rr *EID) Header() *RR_Header        { return &rr.Hdr }
func (rr *EUI48) Header() *RR_Header      { return &rr.Hdr }
func (rr *EUI64) Header() *RR_Header      { return &rr.Hdr }
func (rr *GID) Header() *RR_Header        { return &rr.Hdr }
func (rr *GPOS) Header() *RR_Header       { return &rr.Hdr }
func (rr *HINFO) Header() *RR_Header      { return &rr.Hdr }
func (rr *HIP) Header() *RR_Header        { return &rr.Hdr }
func (rr *KEY) Header() *RR_Header        { return &rr.Hdr }
func (rr *KX) Header() *RR_Header         { return &rr.Hdr }
func (rr *L32) Header() *RR_Header        { return &rr.Hdr }
func (rr *L64) Header() *RR_Header        { return &rr.Hdr }
func (rr *LOC) Header() *RR_Header        { return &rr.Hdr }
func (rr *LP) Header() *RR_Header         { return &rr.Hdr }
func (rr *MB) Header() *RR_Header         { return &rr.Hdr }
func (rr *MD) Header() *RR_Header         { return &rr.Hdr }
func (rr *MF) Header() *RR_Header         { return &rr.Hdr }
func (rr *MG) Header() *RR_Header         { return &rr.Hdr }
func (rr *MINFO) Header() *RR_Header      { return &rr.Hdr }
func (rr *MR) Header() *RR_Header         { return &rr.Hdr }
func (rr *MX) Header() *RR_Header         { return &rr.Hdr }
func (rr *NAPTR) Header() *RR_Header      { return &rr.Hdr }
func (rr *NID) Header() *RR_Header        { return &rr.Hdr }
func (rr *NIMLOC) Header() *RR_Header     { return &rr.Hdr }
func (rr *NINFO) Header() *RR_Header      { return &rr.Hdr }
func (rr *NS) Header() *RR_Header         { return &rr.Hdr }
func (rr *NSAPPTR) Header() *RR_Header    { return &rr.Hdr }
func (rr *NSEC) Header() *RR_Header       { return &rr.Hdr }
func (rr *NSEC3) Header() *RR_Header      { return &rr.Hdr }
func (rr *NSEC3PARAM) Header() *RR_Header { return &rr.Hdr }
func (rr *NULL) Header() *RR_Header       { return &rr.Hdr }
func (rr *OPENPGPKEY) Header() *RR_Header { return &rr.Hdr }
func (rr *OPT) Header() *RR_Header        { return &rr.Hdr }
func (rr *PTR) Header() *RR_Header        { return &rr.Hdr }
func (rr *PX) Header() *RR_Header         { return &rr.Hdr }
func (rr *RFC3597) Header() *RR_Header    { return &rr.Hdr }
func (rr *RKEY) Header() *RR_Header       { return &rr.Hdr }
func (rr *RP) Header() *RR_Header         { return &rr.Hdr }
func (rr *RRSIG) Header() *RR_Header      { return &rr.Hdr }
func (rr *RT) Header() *RR_Header         { return &rr.Hdr }
func (rr *SIG) Header() *RR_Header        { return &rr.Hdr }
func (rr *SMIMEA) Header() *RR_Header     { return &rr.Hdr }
func (rr *SOA) Header() *RR_Header        { return &rr.Hdr }
func (rr *SPF) Header() *RR_Header        { return &rr.Hdr }
func (rr *SRV) Header() *RR_Header        { return &rr.Hdr }
func (rr *SSHFP) Header() *RR_Header      { return &rr.Hdr }
func (rr *TA) Header() *RR_Header         { return &rr.Hdr }
func (rr *TALINK) Header() *RR_Header     { return &rr.Hdr }
func (rr *TKEY) Header() *RR_Header       { return &rr.Hdr }
func (rr *TLSA) Header() *RR_Header       { return &rr.Hdr }
func (rr *TSIG) Header() *RR_Header       { return &rr.Hdr }
func (rr *TXT) Header() *RR_Header        { return &rr.Hdr }
func (rr *UID) Header() *RR_Header        { return &rr.Hdr }
func (rr *UINFO) Header() *RR_Header      { return &rr.Hdr }
func (rr *URI) Header() *RR_Header        { return &rr.Hdr }
func (rr *X25) Header() *RR_Header        { return &rr.Hdr }

// len() functions
func (rr *A) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += net.IPv4len // A
	return l
}
func (rr *AAAA) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += net.IPv6len // AAAA
	return l
}
func (rr *AFSDB) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Subtype
	l += domainNameLen(rr.Hostname, off+l, compression, false)
	return l
}
func (rr *ANY) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	return l
}
func (rr *AVC) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	for _, x := range rr.Txt {
		l += len(x) + 1
	}
	return l
}
func (rr *CAA) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l++ // Flag
	l += len(rr.Tag) + 1
	l += len(rr.Value)
	return l
}
func (rr *CERT) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Type
	l += 2 // KeyTag
	l++    // Algorithm
	l += base64.StdEncoding.DecodedLen(len(rr.Certificate))
	return l
}
func (rr *CNAME) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Target, off+l, compression, true)
	return l
}
func (rr *DHCID) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += base64.StdEncoding.DecodedLen(len(rr.Digest))
	return l
}
func (rr *DNAME) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Target, off+l, compression, false)
	return l
}
func (rr *DNSKEY) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Flags
	l++    // Protocol
	l++    // Algorithm
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	return l
}
func (rr *DS) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // KeyTag
	l++    // Algorithm
	l++    // DigestType
	l += len(rr.Digest)/2 + 1
	return l
}
func (rr *EID) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += len(rr.Endpoint)/2 + 1
	return l
}
func (rr *EUI48) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 6 // Address
	return l
}
func (rr *EUI64) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 8 // Address
	return l
}
func (rr *GID) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 4 // Gid
	return l
}
func (rr *GPOS) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += len(rr.Longitude) + 1
	l += len(rr.Latitude) + 1
	l += len(rr.Altitude) + 1
	return l
}
func (rr *HINFO) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += len(rr.Cpu) + 1
	l += len(rr.Os) + 1
	return l
}
func (rr *HIP) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l++    // HitLength
	l++    // PublicKeyAlgorithm
	l += 2 // PublicKeyLength
	l += len(rr.Hit) / 2
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	for _, x := range rr.RendezvousServers {
		l += domainNameLen(x, off+l, compression, false)
	}
	return l
}
func (rr *KX) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Preference
	l += domainNameLen(rr.Exchanger, off+l, compression, false)
	return l
}
func (rr *L32) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2           // Preference
	l += net.IPv4len // Locator32
	return l
}
func (rr *L64) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Preference
	l += 8 // Locator64
	return l
}
func (rr *LOC) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l++    // Version
	l++    // Size
	l++    // HorizPre
	l++    // VertPre
	l += 4 // Latitude
	l += 4 // Longitude
	l += 4 // Altitude
	return l
}
func (rr *LP) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Preference
	l += domainNameLen(rr.Fqdn, off+l, compression, false)
	return l
}
func (rr *MB) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Mb, off+l, compression, true)
	return l
}
func (rr *MD) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Md, off+l, compression, true)
	return l
}
func (rr *MF) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Mf, off+l, compression, true)
	return l
}
func (rr *MG) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Mg, off+l, compression, true)
	return l
}
func (rr *MINFO) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Rmail, off+l, compression, true)
	l += domainNameLen(rr.Email, off+l, compression, true)
	return l
}
func (rr *MR) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Mr, off+l, compression, true)
	return l
}
func (rr *MX) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Preference
	l += domainNameLen(rr.Mx, off+l, compression, true)
	return l
}
func (rr *NAPTR) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Order
	l += 2 // Preference
	l += len(rr.Flags) + 1
	l += len(rr.Service) + 1
	l += len(rr.Regexp) + 1
	l += domainNameLen(rr.Replacement, off+l, compression, false)
	return l
}
func (rr *NID) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Preference
	l += 8 // NodeID
	return l
}
func (rr *NIMLOC) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += len(rr.Locator)/2 + 1
	return l
}
func (rr *NINFO) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	for _, x := range rr.ZSData {
		l += len(x) + 1
	}
	return l
}
func (rr *NS) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Ns, off+l, compression, true)
	return l
}
func (rr *NSAPPTR) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Ptr, off+l, compression, false)
	return l
}
func (rr *NSEC3PARAM) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l++    // Hash
	l++    // Flags
	l += 2 // Iterations
	l++    // SaltLength
	l += len(rr.Salt) / 2
	return l
}
func (rr *NULL) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += len(rr.Data)
	return l
}
func (rr *OPENPGPKEY) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	return l
}
func (rr *PTR) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Ptr, off+l, compression, true)
	return l
}
func (rr *PX) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Preference
	l += domainNameLen(rr.Map822, off+l, compression, false)
	l += domainNameLen(rr.Mapx400, off+l, compression, false)
	return l
}
func (rr *RFC3597) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += len(rr.Rdata)/2 + 1
	return l
}
func (rr *RKEY) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Flags
	l++    // Protocol
	l++    // Algorithm
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	return l
}
func (rr *RP) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Mbox, off+l, compression, false)
	l += domainNameLen(rr.Txt, off+l, compression, false)
	return l
}
func (rr *RRSIG) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // TypeCovered
	l++    // Algorithm
	l++    // Labels
	l += 4 // OrigTtl
	l += 4 // Expiration
	l += 4 // Inception
	l += 2 // KeyTag
	l += domainNameLen(rr.SignerName, off+l, compression, false)
	l += base64.StdEncoding.DecodedLen(len(rr.Signature))
	return l
}
func (rr *RT) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Preference
	l += domainNameLen(rr.Host, off+l, compression, false)
	return l
}
func (rr *SMIMEA) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l++ // Usage
	l++ // Selector
	l++ // MatchingType
	l += len(rr.Certificate)/2 + 1
	return l
}
func (rr *SOA) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Ns, off+l, compression, true)
	l += domainNameLen(rr.Mbox, off+l, compression, true)
	l += 4 // Serial
	l += 4 // Refresh
	l += 4 // Retry
	l += 4 // Expire
	l += 4 // Minttl
	return l
}
func (rr *SPF) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	for _, x := range rr.Txt {
		l += len(x) + 1
	}
	return l
}
func (rr *SRV) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Priority
	l += 2 // Weight
	l += 2 // Port
	l += domainNameLen(rr.Target, off+l, compression, false)
	return l
}
func (rr *SSHFP) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l++ // Algorithm
	l++ // Type
	l += len(rr.FingerPrint)/2 + 1
	return l
}
func (rr *TA) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // KeyTag
	l++    // Algorithm
	l++    // DigestType
	l += len(rr.Digest)/2 + 1
	return l
}
func (rr *TALINK) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.PreviousName, off+l, compression, false)
	l += domainNameLen(rr.NextName, off+l, compression, false)
	return l
}
func (rr *TKEY) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Algorithm, off+l, compression, false)
	l += 4 // Inception
	l += 4 // Expiration
	l += 2 // Mode
	l += 2 // Error
	l += 2 // KeySize
	l += len(rr.Key) / 2
	l += 2 // OtherLen
	l += len(rr.OtherData) / 2
	return l
}
func (rr *TLSA) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l++ // Usage
	l++ // Selector
	l++ // MatchingType
	l += len(rr.Certificate)/2 + 1
	return l
}
func (rr *TSIG) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += domainNameLen(rr.Algorithm, off+l, compression, false)
	l += 6 // TimeSigned
	l += 2 // Fudge
	l += 2 // MACSize
	l += len(rr.MAC) / 2
	l += 2 // OrigId
	l += 2 // Error
	l += 2 // OtherLen
	l += len(rr.OtherData) / 2
	return l
}
func (rr *TXT) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	for _, x := range rr.Txt {
		l += len(x) + 1
	}
	return l
}
func (rr *UID) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 4 // Uid
	return l
}
func (rr *UINFO) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += len(rr.Uinfo) + 1
	return l
}
func (rr *URI) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += 2 // Priority
	l += 2 // Weight
	l += len(rr.Target)
	return l
}
func (rr *X25) len(off int, compression map[string]struct{}) int {
	l := rr.Hdr.len(off, compression)
	l += len(rr.PSDNAddress) + 1
	return l
}

// copy() functions
func (rr *A) copy() RR {
	return &A{rr.Hdr, copyIP(rr.A)}
}
func (rr *AAAA) copy() RR {
	return &AAAA{rr.Hdr, copyIP(rr.AAAA)}
}
func (rr *AFSDB) copy() RR {
	return &AFSDB{rr.Hdr, rr.Subtype, rr.Hostname}
}
func (rr *ANY) copy() RR {
	return &ANY{rr.Hdr}
}
func (rr *AVC) copy() RR {
	Txt := make([]string, len(rr.Txt))
	copy(Txt, rr.Txt)
	return &AVC{rr.Hdr, Txt}
}
func (rr *CAA) copy() RR {
	return &CAA{rr.Hdr, rr.Flag, rr.Tag, rr.Value}
}
func (rr *CERT) copy() RR {
	return &CERT{rr.Hdr, rr.Type, rr.KeyTag, rr.Algorithm, rr.Certificate}
}
func (rr *CNAME) copy() RR {
	return &CNAME{rr.Hdr, rr.Target}
}
func (rr *CSYNC) copy() RR {
	TypeBitMap := make([]uint16, len(rr.TypeBitMap))
	copy(TypeBitMap, rr.TypeBitMap)
	return &CSYNC{rr.Hdr, rr.Serial, rr.Flags, TypeBitMap}
}
func (rr *DHCID) copy() RR {
	return &DHCID{rr.Hdr, rr.Digest}
}
func (rr *DNAME) copy() RR {
	return &DNAME{rr.Hdr, rr.Target}
}
func (rr *DNSKEY) copy() RR {
	return &DNSKEY{rr.Hdr, rr.Flags, rr.Protocol, rr.Algorithm, rr.PublicKey}
}
func (rr *DS) copy() RR {
	return &DS{rr.Hdr, rr.KeyTag, rr.Algorithm, rr.DigestType, rr.Digest}
}
func (rr *EID) copy() RR {
	return &EID{rr.Hdr, rr.Endpoint}
}
func (rr *EUI48) copy() RR {
	return &EUI48{rr.Hdr, rr.Address}
}
func (rr *EUI64) copy() RR {
	return &EUI64{rr.Hdr, rr.Address}
}
func (rr *GID) copy() RR {
	return &GID{rr.Hdr, rr.Gid}
}
func (rr *GPOS) copy() RR {
	return &GPOS{rr.Hdr, rr.Longitude, rr.Latitude, rr.Altitude}
}
func (rr *HINFO) copy() RR {
	return &HINFO{rr.Hdr, rr.Cpu, rr.Os}
}
func (rr *HIP) copy() RR {
	RendezvousServers := make([]string, len(rr.RendezvousServers))
	copy(RendezvousServers, rr.RendezvousServers)
	return &HIP{rr.Hdr, rr.HitLength, rr.PublicKeyAlgorithm, rr.PublicKeyLength, rr.Hit, rr.PublicKey, RendezvousServers}
}
func (rr *KX) copy() RR {
	return &KX{rr.Hdr, rr.Preference, rr.Exchanger}
}
func (rr *L32) copy() RR {
	return &L32{rr.Hdr, rr.Preference, copyIP(rr.Locator32)}
}
func (rr *L64) copy() RR {
	return &L64{rr.Hdr, rr.Preference, rr.Locator64}
}
func (rr *LOC) copy() RR {
	return &LOC{rr.Hdr, rr.Version, rr.Size, rr.HorizPre, rr.VertPre, rr.Latitude, rr.Longitude, rr.Altitude}
}
func (rr *LP) copy() RR {
	return &LP{rr.Hdr, rr.Preference, rr.Fqdn}
}
func (rr *MB) copy() RR {
	return &MB{rr.Hdr, rr.Mb}
}
func (rr *MD) copy() RR {
	return &MD{rr.Hdr, rr.Md}
}
func (rr *MF) copy() RR {
	return &MF{rr.Hdr, rr.Mf}
}
func (rr *MG) copy() RR {
	return &MG{rr.Hdr, rr.Mg}
}
func (rr *MINFO) copy() RR {
	return &MINFO{rr.Hdr, rr.Rmail, rr.Email}
}
func (rr *MR) copy() RR {
	return &MR{rr.Hdr, rr.Mr}
}
func (rr *MX) copy() RR {
	return &MX{rr.Hdr, rr.Preference, rr.Mx}
}
func (rr *NAPTR) copy() RR {
	return &NAPTR{rr.Hdr, rr.Order, rr.Preference, rr.Flags, rr.Service, rr.Regexp, rr.Replacement}
}
func (rr *NID) copy() RR {
	return &NID{rr.Hdr, rr.Preference, rr.NodeID}
}
func (rr *NIMLOC) copy() RR {
	return &NIMLOC{rr.Hdr, rr.Locator}
}
func (rr *NINFO) copy() RR {
	ZSData := make([]string, len(rr.ZSData))
	copy(ZSData, rr.ZSData)
	return &NINFO{rr.Hdr, ZSData}
}
func (rr *NS) copy() RR {
	return &NS{rr.Hdr, rr.Ns}
}
func (rr *NSAPPTR) copy() RR {
	return &NSAPPTR{rr.Hdr, rr.Ptr}
}
func (rr *NSEC) copy() RR {
	TypeBitMap := make([]uint16, len(rr.TypeBitMap))
	copy(TypeBitMap, rr.TypeBitMap)
	return &NSEC{rr.Hdr, rr.NextDomain, TypeBitMap}
}
func (rr *NSEC3) copy() RR {
	TypeBitMap := make([]uint16, len(rr.TypeBitMap))
	copy(TypeBitMap, rr.TypeBitMap)
	return &NSEC3{rr.Hdr, rr.Hash, rr.Flags, rr.Iterations, rr.SaltLength, rr.Salt, rr.HashLength, rr.NextDomain, TypeBitMap}
}
func (rr *NSEC3PARAM) copy() RR {
	return &NSEC3PARAM{rr.Hdr, rr.Hash, rr.Flags, rr.Iterations, rr.SaltLength, rr.Salt}
}
func (rr *NULL) copy() RR {
	return &NULL{rr.Hdr, rr.Data}
}
func (rr *OPENPGPKEY) copy() RR {
	return &OPENPGPKEY{rr.Hdr, rr.PublicKey}
}
func (rr *OPT) copy() RR {
	Option := make([]EDNS0, len(rr.Option))
	for i, e := range rr.Option {
		Option[i] = e.copy()
	}
	return &OPT{rr.Hdr, Option}
}
func (rr *PTR) copy() RR {
	return &PTR{rr.Hdr, rr.Ptr}
}
func (rr *PX) copy() RR {
	return &PX{rr.Hdr, rr.Preference, rr.Map822, rr.Mapx400}
}
func (rr *RFC3597) copy() RR {
	return &RFC3597{rr.Hdr, rr.Rdata}
}
func (rr *RKEY) copy() RR {
	return &RKEY{rr.Hdr, rr.Flags, rr.Protocol, rr.Algorithm, rr.PublicKey}
}
func (rr *RP) copy() RR {
	return &RP{rr.Hdr, rr.Mbox, rr.Txt}
}
func (rr *RRSIG) copy() RR {
	return &RRSIG{rr.Hdr, rr.TypeCovered, rr.Algorithm, rr.Labels, rr.OrigTtl, rr.Expiration, rr.Inception, rr.KeyTag, rr.SignerName, rr.Signature}
}
func (rr *RT) copy() RR {
	return &RT{rr.Hdr, rr.Preference, rr.Host}
}
func (rr *SMIMEA) copy() RR {
	return &SMIMEA{rr.Hdr, rr.Usage, rr.Selector, rr.MatchingType, rr.Certificate}
}
func (rr *SOA) copy() RR {
	return &SOA{rr.Hdr, rr.Ns, rr.Mbox, rr.Serial, rr.Refresh, rr.Retry, rr.Expire, rr.Minttl}
}
func (rr *SPF) copy() RR {
	Txt := make([]string, len(rr.Txt))
	copy(Txt, rr.Txt)
	return &SPF{rr.Hdr, Txt}
}
func (rr *SRV) copy() RR {
	return &SRV{rr.Hdr, rr.Priority, rr.Weight, rr.Port, rr.Target}
}
func (rr *SSHFP) copy() RR {
	return &SSHFP{rr.Hdr, rr.Algorithm, rr.Type, rr.FingerPrint}
}
func (rr *TA) copy() RR {
	return &TA{rr.Hdr, rr.KeyTag, rr.Algorithm, rr.DigestType, rr.Digest}
}
func (rr *TALINK) copy() RR {
	return &TALINK{rr.Hdr, rr.PreviousName, rr.NextName}
}
func (rr *TKEY) copy() RR {
	return &TKEY{rr.Hdr, rr.Algorithm, rr.Inception, rr.Expiration, rr.Mode, rr.Error, rr.KeySize, rr.Key, rr.OtherLen, rr.OtherData}
}
func (rr *TLSA) copy() RR {
	return &TLSA{rr.Hdr, rr.Usage, rr.Selector, rr.MatchingType, rr.Certificate}
}
func (rr *TSIG) copy() RR {
	return &TSIG{rr.Hdr, rr.Algorithm, rr.TimeSigned, rr.Fudge, rr.MACSize, rr.MAC, rr.OrigId, rr.Error, rr.OtherLen, rr.OtherData}
}
func (rr *TXT) copy() RR {
	Txt := make([]string, len(rr.Txt))
	copy(Txt, rr.Txt)
	return &TXT{rr.Hdr, Txt}
}
func (rr *UID) copy() RR {
	return &UID{rr.Hdr, rr.Uid}
}
func (rr *UINFO) copy() RR {
	return &UINFO{rr.Hdr, rr.Uinfo}
}
func (rr *URI) copy() RR {
	return &URI{rr.Hdr, rr.Priority, rr.Weight, rr.Target}
}
func (rr *X25) copy() RR {
	return &X25{rr.Hdr, rr.PSDNAddress}
}
