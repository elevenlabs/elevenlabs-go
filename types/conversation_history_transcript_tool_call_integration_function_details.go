package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationHistoryTranscriptToolCallIntegrationFunctionDetails struct {
	ResolvedParameters map[string]string `json:"resolved_parameters" required:"true"`
	extraProperties    map[string]interface{}
}

func (c *ConversationHistoryTranscriptToolCallIntegrationFunctionDetails) GetResolvedParameters() (value map[string]string) {
	if c == nil {
		return
	}
	return c.ResolvedParameters
}

func (c *ConversationHistoryTranscriptToolCallIntegrationFunctionDetails) SetResolvedParameters(resolvedParameters map[string]string) {
	c.ResolvedParameters = resolvedParameters
}

func (c *ConversationHistoryTranscriptToolCallIntegrationFunctionDetails) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptToolCallIntegrationFunctionDetails) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptToolCallIntegrationFunctionDetails to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptToolCallIntegrationFunctionDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationHistoryTranscriptToolCallIntegrationFunctionDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationHistoryTranscriptToolCallIntegrationFunctionDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptToolCallIntegrationFunctionDetails) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptToolCallIntegrationFunctionDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
