package types

import (
	"fmt"
)

type BanReasonType string

const (
	BanReasonTypeSafety BanReasonType = "safety"
	BanReasonTypeManual BanReasonType = "manual"
)

func NewBanReasonTypeFromString(s string) (BanReasonType, error) {
	switch s {
	case "safety":
		return BanReasonTypeSafety, nil
	case "manual":
		return BanReasonTypeManual, nil
	}
	var t BanReasonType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BanReasonType) Ptr() *BanReasonType {
	return &e
}
