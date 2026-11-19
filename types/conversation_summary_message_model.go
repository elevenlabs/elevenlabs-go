package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationSummaryMessageModel struct {
	Role            ConversationSummaryMessageModelRole `json:"role" xml:"role" required:"true"`
	Message         string                              `json:"message" xml:"message" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationSummaryMessageModel) GetRole() (value ConversationSummaryMessageModelRole) {
	if c == nil {
		return
	}
	return c.Role
}

func (c *ConversationSummaryMessageModel) GetMessage() (value string) {
	if c == nil {
		return
	}
	return c.Message
}

func (c *ConversationSummaryMessageModel) SetRole(role ConversationSummaryMessageModelRole) {
	c.Role = role
}

func (c *ConversationSummaryMessageModel) SetMessage(message string) {
	c.Message = message
}

func (c *ConversationSummaryMessageModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationSummaryMessageModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationSummaryMessageModel to string"
	}
	return string(jsonData)
}

func (c *ConversationSummaryMessageModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationSummaryMessageModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationSummaryMessageModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationSummaryMessageModel) MarshalJSON() ([]byte, error) {
	type embed ConversationSummaryMessageModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
