package types

import (
	"fmt"
)

type DocumentUsageModeEnum string

const (
	DocumentUsageModeEnumPrompt DocumentUsageModeEnum = "prompt"
	DocumentUsageModeEnumAuto   DocumentUsageModeEnum = "auto"
)

func NewDocumentUsageModeEnumFromString(s string) (DocumentUsageModeEnum, error) {
	switch s {
	case "prompt":
		return DocumentUsageModeEnumPrompt, nil
	case "auto":
		return DocumentUsageModeEnumAuto, nil
	}
	var t DocumentUsageModeEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DocumentUsageModeEnum) Ptr() *DocumentUsageModeEnum {
	return &e
}
