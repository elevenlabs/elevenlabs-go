package types

import (
	"fmt"
)

type KnowledgeBaseSortBy string

const (
	KnowledgeBaseSortByName      KnowledgeBaseSortBy = "name"
	KnowledgeBaseSortByCreatedAt KnowledgeBaseSortBy = "created_at"
	KnowledgeBaseSortByUpdatedAt KnowledgeBaseSortBy = "updated_at"
	KnowledgeBaseSortBySize      KnowledgeBaseSortBy = "size"
)

func NewKnowledgeBaseSortByFromString(s string) (KnowledgeBaseSortBy, error) {
	switch s {
	case "name":
		return KnowledgeBaseSortByName, nil
	case "created_at":
		return KnowledgeBaseSortByCreatedAt, nil
	case "updated_at":
		return KnowledgeBaseSortByUpdatedAt, nil
	case "size":
		return KnowledgeBaseSortBySize, nil
	}
	var t KnowledgeBaseSortBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e KnowledgeBaseSortBy) Ptr() *KnowledgeBaseSortBy {
	return &e
}
