package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationUserResponseModel struct {
	UserId                      string                       `json:"user_id" xml:"user_id" required:"true"`
	LastContactUnixSecs         int                          `json:"last_contact_unix_secs" xml:"last_contact_unix_secs" required:"true"`
	FirstContactUnixSecs        int                          `json:"first_contact_unix_secs" xml:"first_contact_unix_secs" required:"true"`
	ConversationCount           int                          `json:"conversation_count" xml:"conversation_count" required:"true"`
	LastContactAgentId          *string                      `json:"last_contact_agent_id,omitempty" xml:"last_contact_agent_id,omitempty"`
	LastContactConversationId   string                       `json:"last_contact_conversation_id" xml:"last_contact_conversation_id" required:"true"`
	LastContactAgentName        *string                      `json:"last_contact_agent_name,omitempty" xml:"last_contact_agent_name,omitempty"`
	Sentiment                   *SentimentAggregate          `json:"sentiment" xml:"sentiment" required:"true"`
	MostFrustratedConversations []*FrustratedConversationRef `json:"most_frustrated_conversations,omitempty" xml:"most_frustrated_conversations,omitempty"`
	extraProperties             map[string]interface{}
}

func (c *ConversationUserResponseModel) GetUserId() (value string) {
	if c == nil {
		return
	}
	return c.UserId
}

func (c *ConversationUserResponseModel) GetLastContactUnixSecs() (value int) {
	if c == nil {
		return
	}
	return c.LastContactUnixSecs
}

func (c *ConversationUserResponseModel) GetFirstContactUnixSecs() (value int) {
	if c == nil {
		return
	}
	return c.FirstContactUnixSecs
}

func (c *ConversationUserResponseModel) GetConversationCount() (value int) {
	if c == nil {
		return
	}
	return c.ConversationCount
}

func (c *ConversationUserResponseModel) GetLastContactAgentId() (value *string) {
	if c == nil {
		return
	}
	return c.LastContactAgentId
}

func (c *ConversationUserResponseModel) GetLastContactConversationId() (value string) {
	if c == nil {
		return
	}
	return c.LastContactConversationId
}

func (c *ConversationUserResponseModel) GetLastContactAgentName() (value *string) {
	if c == nil {
		return
	}
	return c.LastContactAgentName
}

func (c *ConversationUserResponseModel) GetSentiment() (value *SentimentAggregate) {
	if c == nil {
		return
	}
	return c.Sentiment
}

func (c *ConversationUserResponseModel) GetMostFrustratedConversations() (value []*FrustratedConversationRef) {
	if c == nil {
		return
	}
	return c.MostFrustratedConversations
}

func (c *ConversationUserResponseModel) SetUserId(userId string) {
	c.UserId = userId
}

func (c *ConversationUserResponseModel) SetLastContactUnixSecs(lastContactUnixSecs int) {
	c.LastContactUnixSecs = lastContactUnixSecs
}

func (c *ConversationUserResponseModel) SetFirstContactUnixSecs(firstContactUnixSecs int) {
	c.FirstContactUnixSecs = firstContactUnixSecs
}

func (c *ConversationUserResponseModel) SetConversationCount(conversationCount int) {
	c.ConversationCount = conversationCount
}

func (c *ConversationUserResponseModel) SetLastContactAgentId(lastContactAgentId *string) {
	c.LastContactAgentId = lastContactAgentId
}

func (c *ConversationUserResponseModel) SetLastContactConversationId(lastContactConversationId string) {
	c.LastContactConversationId = lastContactConversationId
}

func (c *ConversationUserResponseModel) SetLastContactAgentName(lastContactAgentName *string) {
	c.LastContactAgentName = lastContactAgentName
}

func (c *ConversationUserResponseModel) SetSentiment(sentiment *SentimentAggregate) {
	c.Sentiment = sentiment
}

func (c *ConversationUserResponseModel) SetMostFrustratedConversations(mostFrustratedConversations []*FrustratedConversationRef) {
	c.MostFrustratedConversations = mostFrustratedConversations
}

func (c *ConversationUserResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationUserResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationUserResponseModel to string"
	}
	return string(jsonData)
}

func (c *ConversationUserResponseModel) UnmarshalJSON(data []byte) error {
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

func (c *ConversationUserResponseModel) MarshalJSON() ([]byte, error) {
	type embed ConversationUserResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
