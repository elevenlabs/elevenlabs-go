package conversations

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type ConversationFeedbackRequestModel struct {
	// The id of the conversation you're taking the action on.
	ConversationId string `json:"-" url:"-"`
	// Either 'like' or 'dislike' to indicate the feedback for the conversation.
	Feedback *types.UserFeedbackScore `json:"feedback,omitempty" xml:"feedback,omitempty"`
}

func (c *ConversationFeedbackRequestModel) GetConversationId() (value string) {
	if c == nil {
		return
	}
	return c.ConversationId
}

func (c *ConversationFeedbackRequestModel) GetFeedback() (value *types.UserFeedbackScore) {
	if c == nil {
		return
	}
	return c.Feedback
}

func (c *ConversationFeedbackRequestModel) SetConversationId(conversationId string) {
	c.ConversationId = conversationId
}

func (c *ConversationFeedbackRequestModel) SetFeedback(feedback *types.UserFeedbackScore) {
	c.Feedback = feedback
}

func (c *ConversationFeedbackRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c ConversationFeedbackRequestModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationFeedbackRequestModel to string"
	}
	return string(jsonData)
}

func (c *ConversationFeedbackRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

func (c *ConversationFeedbackRequestModel) MarshalJSON() ([]byte, error) {
	type embed ConversationFeedbackRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
