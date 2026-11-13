package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryFeedbackCommonModel struct {
	Type            *ConversationFeedbackType `json:"type,omitempty" xml:"type,omitempty"`
	OverallScore    *UserFeedbackScore        `json:"overall_score,omitempty" xml:"overall_score,omitempty"`
	Likes           *int                      `json:"likes,omitempty" xml:"likes,omitempty"`
	Dislikes        *int                      `json:"dislikes,omitempty" xml:"dislikes,omitempty"`
	Rating          *int                      `json:"rating,omitempty" xml:"rating,omitempty"`
	Comment         *string                   `json:"comment,omitempty" xml:"comment,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationHistoryFeedbackCommonModel) GetType() (value *ConversationFeedbackType) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationHistoryFeedbackCommonModel) GetOverallScore() (value *UserFeedbackScore) {
	if c == nil {
		return
	}
	return c.OverallScore
}

func (c *ConversationHistoryFeedbackCommonModel) GetLikes() (value *int) {
	if c == nil {
		return
	}
	return c.Likes
}

func (c *ConversationHistoryFeedbackCommonModel) GetDislikes() (value *int) {
	if c == nil {
		return
	}
	return c.Dislikes
}

func (c *ConversationHistoryFeedbackCommonModel) GetRating() (value *int) {
	if c == nil {
		return
	}
	return c.Rating
}

func (c *ConversationHistoryFeedbackCommonModel) GetComment() (value *string) {
	if c == nil {
		return
	}
	return c.Comment
}

func (c *ConversationHistoryFeedbackCommonModel) SetType(type_ *ConversationFeedbackType) {
	c.Type = type_
}

func (c *ConversationHistoryFeedbackCommonModel) SetOverallScore(overallScore *UserFeedbackScore) {
	c.OverallScore = overallScore
}

func (c *ConversationHistoryFeedbackCommonModel) SetLikes(likes *int) {
	c.Likes = likes
}

func (c *ConversationHistoryFeedbackCommonModel) SetDislikes(dislikes *int) {
	c.Dislikes = dislikes
}

func (c *ConversationHistoryFeedbackCommonModel) SetRating(rating *int) {
	c.Rating = rating
}

func (c *ConversationHistoryFeedbackCommonModel) SetComment(comment *string) {
	c.Comment = comment
}

func (c *ConversationHistoryFeedbackCommonModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryFeedbackCommonModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryFeedbackCommonModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryFeedbackCommonModel) UnmarshalJSON(data []byte) error {
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

func (c *ConversationHistoryFeedbackCommonModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryFeedbackCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
