package types

import (
	"fmt"
)

// LLM model to use for custom guardrail evaluation
type CustomGuardrailConfigModel string

const (
	CustomGuardrailConfigModelGemini25FlashLite CustomGuardrailConfigModel = "gemini-2.5-flash-lite"
	CustomGuardrailConfigModelGemini25Flash     CustomGuardrailConfigModel = "gemini-2.5-flash"
	CustomGuardrailConfigModelGemini31FlashLite CustomGuardrailConfigModel = "gemini-3.1-flash-lite"
	CustomGuardrailConfigModelGemini35Flash     CustomGuardrailConfigModel = "gemini-3.5-flash"
	CustomGuardrailConfigModelClaudeHaiku45     CustomGuardrailConfigModel = "claude-haiku-4-5"
	CustomGuardrailConfigModelClaudeSonnet46    CustomGuardrailConfigModel = "claude-sonnet-4-6"
	CustomGuardrailConfigModelGpt54Nano         CustomGuardrailConfigModel = "gpt-5.4-nano"
	CustomGuardrailConfigModelGpt54Mini         CustomGuardrailConfigModel = "gpt-5.4-mini"
)

func NewCustomGuardrailConfigModelFromString(s string) (CustomGuardrailConfigModel, error) {
	switch s {
	case "gemini-2.5-flash-lite":
		return CustomGuardrailConfigModelGemini25FlashLite, nil
	case "gemini-2.5-flash":
		return CustomGuardrailConfigModelGemini25Flash, nil
	case "gemini-3.1-flash-lite":
		return CustomGuardrailConfigModelGemini31FlashLite, nil
	case "gemini-3.5-flash":
		return CustomGuardrailConfigModelGemini35Flash, nil
	case "claude-haiku-4-5":
		return CustomGuardrailConfigModelClaudeHaiku45, nil
	case "claude-sonnet-4-6":
		return CustomGuardrailConfigModelClaudeSonnet46, nil
	case "gpt-5.4-nano":
		return CustomGuardrailConfigModelGpt54Nano, nil
	case "gpt-5.4-mini":
		return CustomGuardrailConfigModelGpt54Mini, nil
	}
	var t CustomGuardrailConfigModel
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CustomGuardrailConfigModel) Ptr() *CustomGuardrailConfigModel {
	return &e
}
