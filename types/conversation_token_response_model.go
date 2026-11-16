package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationTokenResponseModel struct {
	// The ID of the agent
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// The token for the agent
	ConversationToken string `json:"conversation_token" xml:"conversation_token" required:"true"`
	// The expiration time of the token in unix seconds
	ExpirationTimeUnixSecs *int `json:"expiration_time_unix_secs,omitempty" xml:"expiration_time_unix_secs,omitempty"`
	// The ID of the conversation
	ConversationId *string                  `json:"conversation_id,omitempty" xml:"conversation_id,omitempty"`
	Purpose        ConversationTokenPurpose `json:"purpose" xml:"purpose" required:"true"`
	// The user ID of the entity who requested the token
	TokenRequesterUserId *string `json:"token_requester_user_id,omitempty" xml:"token_requester_user_id,omitempty"`
	extraProperties      map[string]interface{}
}

func (c *ConversationTokenResponseModel) GetAgentId() (value string) {
	if c == nil {
		return
	}
	return c.AgentId
}

func (c *ConversationTokenResponseModel) GetConversationToken() (value string) {
	if c == nil {
		return
	}
	return c.ConversationToken
}

func (c *ConversationTokenResponseModel) GetExpirationTimeUnixSecs() (value *int) {
	if c == nil {
		return
	}
	return c.ExpirationTimeUnixSecs
}

func (c *ConversationTokenResponseModel) GetConversationId() (value *string) {
	if c == nil {
		return
	}
	return c.ConversationId
}

func (c *ConversationTokenResponseModel) GetPurpose() (value ConversationTokenPurpose) {
	if c == nil {
		return
	}
	return c.Purpose
}

func (c *ConversationTokenResponseModel) GetTokenRequesterUserId() (value *string) {
	if c == nil {
		return
	}
	return c.TokenRequesterUserId
}

func (c *ConversationTokenResponseModel) SetAgentId(agentId string) {
	c.AgentId = agentId
}

func (c *ConversationTokenResponseModel) SetConversationToken(conversationToken string) {
	c.ConversationToken = conversationToken
}

func (c *ConversationTokenResponseModel) SetExpirationTimeUnixSecs(expirationTimeUnixSecs *int) {
	c.ExpirationTimeUnixSecs = expirationTimeUnixSecs
}

func (c *ConversationTokenResponseModel) SetConversationId(conversationId *string) {
	c.ConversationId = conversationId
}

func (c *ConversationTokenResponseModel) SetPurpose(purpose ConversationTokenPurpose) {
	c.Purpose = purpose
}

func (c *ConversationTokenResponseModel) SetTokenRequesterUserId(tokenRequesterUserId *string) {
	c.TokenRequesterUserId = tokenRequesterUserId
}

func (c *ConversationTokenResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationTokenResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationTokenResponseModel to string"
	}
	return string(jsonData)
}

func (c *ConversationTokenResponseModel) UnmarshalJSON(data []byte) error {
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

func (c *ConversationTokenResponseModel) MarshalJSON() ([]byte, error) {
	type embed ConversationTokenResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
