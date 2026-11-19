package types

import (
	"fmt"
)

type ConversationFeedbackType string

const (
	ConversationFeedbackTypeThumbs ConversationFeedbackType = "thumbs"
	ConversationFeedbackTypeRating ConversationFeedbackType = "rating"
)

func NewConversationFeedbackTypeFromString(s string) (ConversationFeedbackType, error) {
	switch s {
	case "thumbs":
		return ConversationFeedbackTypeThumbs, nil
	case "rating":
		return ConversationFeedbackTypeRating, nil
	}
	var t ConversationFeedbackType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationFeedbackType) Ptr() *ConversationFeedbackType {
	return &e
}
