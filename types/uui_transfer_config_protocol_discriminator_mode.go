package types

import (
	"fmt"
)

// How to attach protocol_discriminator. 'prefix' prepends the octet to the hex payload (User-to-User=XX<hex>;encoding=hex). 'pd_parameter' sends it as a separate parameter (User-to-User=<hex>;pd=XX;encoding=hex). Ignored when protocol_discriminator is unset.
type UuiTransferConfigProtocolDiscriminatorMode string

const (
	UuiTransferConfigProtocolDiscriminatorModePrefix      UuiTransferConfigProtocolDiscriminatorMode = "prefix"
	UuiTransferConfigProtocolDiscriminatorModePdParameter UuiTransferConfigProtocolDiscriminatorMode = "pd_parameter"
)

func NewUuiTransferConfigProtocolDiscriminatorModeFromString(s string) (UuiTransferConfigProtocolDiscriminatorMode, error) {
	switch s {
	case "prefix":
		return UuiTransferConfigProtocolDiscriminatorModePrefix, nil
	case "pd_parameter":
		return UuiTransferConfigProtocolDiscriminatorModePdParameter, nil
	}
	var t UuiTransferConfigProtocolDiscriminatorMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e UuiTransferConfigProtocolDiscriminatorMode) Ptr() *UuiTransferConfigProtocolDiscriminatorMode {
	return &e
}
