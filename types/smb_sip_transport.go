package types

import (
	"fmt"
)

// Signalling transport for outbound SIP INVITEs.
//
// Values match the Agents platform “SIPTrunkTransportEnum“ so the SMB adapter
// maps by value. “AUTO“ is omitted: the Agents UI hides it too, since it can
// pick UDP, which is known to misbehave (CAI-299). Reception defaults to TLS.
type SmbSipTransport string

const (
	SmbSipTransportUdp SmbSipTransport = "udp"
	SmbSipTransportTcp SmbSipTransport = "tcp"
	SmbSipTransportTls SmbSipTransport = "tls"
)

func NewSmbSipTransportFromString(s string) (SmbSipTransport, error) {
	switch s {
	case "udp":
		return SmbSipTransportUdp, nil
	case "tcp":
		return SmbSipTransportTcp, nil
	case "tls":
		return SmbSipTransportTls, nil
	}
	var t SmbSipTransport
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SmbSipTransport) Ptr() *SmbSipTransport {
	return &e
}
