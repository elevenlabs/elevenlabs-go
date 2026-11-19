package types

import (
	"fmt"
)

// Request-time processing tier for a model.
//
// STANDARD: default price and throughput.
// FAST: pays a per-token price premium for higher throughput / lower latency.
// See OpenAI's fast mode: https://openai.com/api-fast-mode/ and, for Gemini, Vertex
// Priority PayGo: https://docs.cloud.google.com/gemini-enterprise-agent-platform/models/priority-paygo
type LlmServiceTier string

const (
	LlmServiceTierStandard LlmServiceTier = "standard"
	LlmServiceTierFast     LlmServiceTier = "fast"
)

func NewLlmServiceTierFromString(s string) (LlmServiceTier, error) {
	switch s {
	case "standard":
		return LlmServiceTierStandard, nil
	case "fast":
		return LlmServiceTierFast, nil
	}
	var t LlmServiceTier
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e LlmServiceTier) Ptr() *LlmServiceTier {
	return &e
}
