package types

import (
	"fmt"
)

type TopicSortBy string

const (
	TopicSortByConversations TopicSortBy = "conversations"
	TopicSortBySentiment     TopicSortBy = "sentiment"
	TopicSortBySuccessRate   TopicSortBy = "success_rate"
	TopicSortByFrustration   TopicSortBy = "frustration"
)

func NewTopicSortByFromString(s string) (TopicSortBy, error) {
	switch s {
	case "conversations":
		return TopicSortByConversations, nil
	case "sentiment":
		return TopicSortBySentiment, nil
	case "success_rate":
		return TopicSortBySuccessRate, nil
	case "frustration":
		return TopicSortByFrustration, nil
	}
	var t TopicSortBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TopicSortBy) Ptr() *TopicSortBy {
	return &e
}
