package types

import (
	"fmt"
)

type TransferTypeEnum string

const (
	TransferTypeEnumBlind      TransferTypeEnum = "blind"
	TransferTypeEnumConference TransferTypeEnum = "conference"
	TransferTypeEnumSipRefer   TransferTypeEnum = "sip_refer"
)

func NewTransferTypeEnumFromString(s string) (TransferTypeEnum, error) {
	switch s {
	case "blind":
		return TransferTypeEnumBlind, nil
	case "conference":
		return TransferTypeEnumConference, nil
	case "sip_refer":
		return TransferTypeEnumSipRefer, nil
	}
	var t TransferTypeEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TransferTypeEnum) Ptr() *TransferTypeEnum {
	return &e
}
