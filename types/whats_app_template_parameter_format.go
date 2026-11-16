package types

import (
	"fmt"
)

type WhatsAppTemplateParameterFormat string

const (
	WhatsAppTemplateParameterFormatPositional WhatsAppTemplateParameterFormat = "POSITIONAL"
	WhatsAppTemplateParameterFormatNamed      WhatsAppTemplateParameterFormat = "NAMED"
)

func NewWhatsAppTemplateParameterFormatFromString(s string) (WhatsAppTemplateParameterFormat, error) {
	switch s {
	case "POSITIONAL":
		return WhatsAppTemplateParameterFormatPositional, nil
	case "NAMED":
		return WhatsAppTemplateParameterFormatNamed, nil
	}
	var t WhatsAppTemplateParameterFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WhatsAppTemplateParameterFormat) Ptr() *WhatsAppTemplateParameterFormat {
	return &e
}
