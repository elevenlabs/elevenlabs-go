package types

import (
	"fmt"
)

// Buckets that make up the non-LLM “platform_charge“ of a conversation.
type PlatformCategory string

const (
	PlatformCategoryVoice       PlatformCategory = "voice"
	PlatformCategorySilence     PlatformCategory = "silence"
	PlatformCategoryBurst       PlatformCategory = "burst"
	PlatformCategoryAsr         PlatformCategory = "asr"
	PlatformCategoryTextMessage PlatformCategory = "text_message"
	PlatformCategoryReception   PlatformCategory = "reception"
	PlatformCategoryV3Addon     PlatformCategory = "v3_addon"
	PlatformCategoryV4Addon     PlatformCategory = "v4_addon"
	PlatformCategoryRagAddon    PlatformCategory = "rag_addon"
	PlatformCategoryOther       PlatformCategory = "other"
)

func NewPlatformCategoryFromString(s string) (PlatformCategory, error) {
	switch s {
	case "voice":
		return PlatformCategoryVoice, nil
	case "silence":
		return PlatformCategorySilence, nil
	case "burst":
		return PlatformCategoryBurst, nil
	case "asr":
		return PlatformCategoryAsr, nil
	case "text_message":
		return PlatformCategoryTextMessage, nil
	case "reception":
		return PlatformCategoryReception, nil
	case "v3_addon":
		return PlatformCategoryV3Addon, nil
	case "v4_addon":
		return PlatformCategoryV4Addon, nil
	case "rag_addon":
		return PlatformCategoryRagAddon, nil
	case "other":
		return PlatformCategoryOther, nil
	}
	var t PlatformCategory
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e PlatformCategory) Ptr() *PlatformCategory {
	return &e
}
