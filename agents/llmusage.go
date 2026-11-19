package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type LlmUsageCalculatorRequestModel struct {
	AgentId string `json:"-" url:"-"`
	// Length of the prompt in characters.
	PromptLength *int `json:"prompt_length,omitempty" xml:"prompt_length,omitempty"`
	// Pages of content in pdf documents OR urls in agent's Knowledge Base.
	NumberOfPages *int `json:"number_of_pages,omitempty" xml:"number_of_pages,omitempty"`
	// Whether RAG is enabled.
	RagEnabled *bool `json:"rag_enabled,omitempty" xml:"rag_enabled,omitempty"`
	// Service tier to price against. Only affects models that offer the tier; others fall back to their standard price.
	ServiceTier *types.LlmServiceTier `json:"service_tier,omitempty" xml:"service_tier,omitempty"`
}

func (l *LlmUsageCalculatorRequestModel) GetAgentId() (value string) {
	if l == nil {
		return
	}
	return l.AgentId
}

func (l *LlmUsageCalculatorRequestModel) GetPromptLength() (value *int) {
	if l == nil {
		return
	}
	return l.PromptLength
}

func (l *LlmUsageCalculatorRequestModel) GetNumberOfPages() (value *int) {
	if l == nil {
		return
	}
	return l.NumberOfPages
}

func (l *LlmUsageCalculatorRequestModel) GetRagEnabled() (value *bool) {
	if l == nil {
		return
	}
	return l.RagEnabled
}

func (l *LlmUsageCalculatorRequestModel) GetServiceTier() (value *types.LlmServiceTier) {
	if l == nil {
		return
	}
	return l.ServiceTier
}

func (l *LlmUsageCalculatorRequestModel) SetAgentId(agentId string) {
	l.AgentId = agentId
}

func (l *LlmUsageCalculatorRequestModel) SetPromptLength(promptLength *int) {
	l.PromptLength = promptLength
}

func (l *LlmUsageCalculatorRequestModel) SetNumberOfPages(numberOfPages *int) {
	l.NumberOfPages = numberOfPages
}

func (l *LlmUsageCalculatorRequestModel) SetRagEnabled(ragEnabled *bool) {
	l.RagEnabled = ragEnabled
}

func (l *LlmUsageCalculatorRequestModel) SetServiceTier(serviceTier *types.LlmServiceTier) {
	l.ServiceTier = serviceTier
}

func (l *LlmUsageCalculatorRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (l LlmUsageCalculatorRequestModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmUsageCalculatorRequestModel to string"
	}
	return string(jsonData)
}

func (l *LlmUsageCalculatorRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, l)
}

func (l *LlmUsageCalculatorRequestModel) MarshalJSON() ([]byte, error) {
	type embed LlmUsageCalculatorRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}

type LlmUsageCalculatorPublicRequestModel struct {
	// Length of the prompt in characters.
	PromptLength int `json:"prompt_length" xml:"prompt_length" required:"true"`
	// Pages of content in PDF documents or URLs in the agent's knowledge base.
	NumberOfPages int `json:"number_of_pages" xml:"number_of_pages" required:"true"`
	// Whether RAG is enabled.
	RagEnabled bool `json:"rag_enabled" xml:"rag_enabled" required:"true"`
	// Service tier to price against. Only affects models that offer the tier; others fall back to their standard price.
	ServiceTier *types.LlmServiceTier `json:"service_tier,omitempty" xml:"service_tier,omitempty"`
}

func (l *LlmUsageCalculatorPublicRequestModel) GetPromptLength() (value int) {
	if l == nil {
		return
	}
	return l.PromptLength
}

func (l *LlmUsageCalculatorPublicRequestModel) GetNumberOfPages() (value int) {
	if l == nil {
		return
	}
	return l.NumberOfPages
}

func (l *LlmUsageCalculatorPublicRequestModel) GetRagEnabled() (value bool) {
	if l == nil {
		return
	}
	return l.RagEnabled
}

func (l *LlmUsageCalculatorPublicRequestModel) GetServiceTier() (value *types.LlmServiceTier) {
	if l == nil {
		return
	}
	return l.ServiceTier
}

func (l *LlmUsageCalculatorPublicRequestModel) SetPromptLength(promptLength int) {
	l.PromptLength = promptLength
}

func (l *LlmUsageCalculatorPublicRequestModel) SetNumberOfPages(numberOfPages int) {
	l.NumberOfPages = numberOfPages
}

func (l *LlmUsageCalculatorPublicRequestModel) SetRagEnabled(ragEnabled bool) {
	l.RagEnabled = ragEnabled
}

func (l *LlmUsageCalculatorPublicRequestModel) SetServiceTier(serviceTier *types.LlmServiceTier) {
	l.ServiceTier = serviceTier
}

func (l *LlmUsageCalculatorPublicRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (l LlmUsageCalculatorPublicRequestModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmUsageCalculatorPublicRequestModel to string"
	}
	return string(jsonData)
}

func (l *LlmUsageCalculatorPublicRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, l)
}

func (l *LlmUsageCalculatorPublicRequestModel) MarshalJSON() ([]byte, error) {
	type embed LlmUsageCalculatorPublicRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
