package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateAgentConversationResponse struct {
	ConversationId  string `json:"conversation_id" xml:"conversation_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateAgentConversationResponse) GetConversationId() (value string) {
	if c == nil {
		return
	}
	return c.ConversationId
}

func (c *CreateAgentConversationResponse) SetConversationId(conversationId string) {
	c.ConversationId = conversationId
}

func (c *CreateAgentConversationResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateAgentConversationResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateAgentConversationResponse to string"
	}
	return string(jsonData)
}

func (c *CreateAgentConversationResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateAgentConversationResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateAgentConversationResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateAgentConversationResponse) MarshalJSON() ([]byte, error) {
	type embed CreateAgentConversationResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
