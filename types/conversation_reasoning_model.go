package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationReasoningModel struct {
	Summary         *string `json:"summary,omitempty" xml:"summary,omitempty"`
	ProviderRedact  *bool   `json:"provider_redact,omitempty" xml:"provider_redact,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationReasoningModel) GetSummary() (value *string) {
	if c == nil {
		return
	}
	return c.Summary
}

func (c *ConversationReasoningModel) GetProviderRedact() (value *bool) {
	if c == nil {
		return
	}
	return c.ProviderRedact
}

func (c *ConversationReasoningModel) SetSummary(summary *string) {
	c.Summary = summary
}

func (c *ConversationReasoningModel) SetProviderRedact(providerRedact *bool) {
	c.ProviderRedact = providerRedact
}

func (c *ConversationReasoningModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationReasoningModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationReasoningModel to string"
	}
	return string(jsonData)
}

func (c *ConversationReasoningModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationReasoningModel) MarshalJSON() ([]byte, error) {
	type embed ConversationReasoningModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
