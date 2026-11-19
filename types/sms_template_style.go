package types

import (
	"fmt"
)

// Preset tone for an SMS message body.
type SmsTemplateStyle string

const (
	SmsTemplateStyleFriendly     SmsTemplateStyle = "friendly"
	SmsTemplateStyleProfessional SmsTemplateStyle = "professional"
	SmsTemplateStyleConcise      SmsTemplateStyle = "concise"
)

func NewSmsTemplateStyleFromString(s string) (SmsTemplateStyle, error) {
	switch s {
	case "friendly":
		return SmsTemplateStyleFriendly, nil
	case "professional":
		return SmsTemplateStyleProfessional, nil
	case "concise":
		return SmsTemplateStyleConcise, nil
	}
	var t SmsTemplateStyle
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SmsTemplateStyle) Ptr() *SmsTemplateStyle {
	return &e
}
