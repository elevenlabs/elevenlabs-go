package types

import (
	"fmt"
)

type MessageSearchSortBy string

const (
	MessageSearchSortBySearchScore MessageSearchSortBy = "search_score"
	MessageSearchSortByCreatedAt   MessageSearchSortBy = "created_at"
)

func NewMessageSearchSortByFromString(s string) (MessageSearchSortBy, error) {
	switch s {
	case "search_score":
		return MessageSearchSortBySearchScore, nil
	case "created_at":
		return MessageSearchSortByCreatedAt, nil
	}
	var t MessageSearchSortBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MessageSearchSortBy) Ptr() *MessageSearchSortBy {
	return &e
}
