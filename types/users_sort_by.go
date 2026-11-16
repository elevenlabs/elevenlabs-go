package types

import (
	"fmt"
)

type UsersSortBy string

const (
	UsersSortByLastContactUnixSecs   UsersSortBy = "last_contact_unix_secs"
	UsersSortByConversationCount     UsersSortBy = "conversation_count"
	UsersSortByAverageSentimentScore UsersSortBy = "average_sentiment_score"
)

func NewUsersSortByFromString(s string) (UsersSortBy, error) {
	switch s {
	case "last_contact_unix_secs":
		return UsersSortByLastContactUnixSecs, nil
	case "conversation_count":
		return UsersSortByConversationCount, nil
	case "average_sentiment_score":
		return UsersSortByAverageSentimentScore, nil
	}
	var t UsersSortBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e UsersSortBy) Ptr() *UsersSortBy {
	return &e
}
