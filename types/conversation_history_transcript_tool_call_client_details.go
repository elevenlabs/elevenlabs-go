package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptToolCallClientDetails struct {
	Parameters      string `json:"parameters" xml:"parameters" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationHistoryTranscriptToolCallClientDetails) GetParameters() (value string) {
	if c == nil {
		return
	}
	return c.Parameters
}

func (c *ConversationHistoryTranscriptToolCallClientDetails) SetParameters(parameters string) {
	c.Parameters = parameters
}

func (c *ConversationHistoryTranscriptToolCallClientDetails) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptToolCallClientDetails) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptToolCallClientDetails to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptToolCallClientDetails) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationHistoryTranscriptToolCallClientDetails
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationHistoryTranscriptToolCallClientDetails(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptToolCallClientDetails) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptToolCallClientDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
