package types

import (
	"fmt"
)

// Type of knowledge source.
type KnowledgeSourceType string

const (
	KnowledgeSourceTypeFile KnowledgeSourceType = "file"
	KnowledgeSourceTypeUrl  KnowledgeSourceType = "url"
)

func NewKnowledgeSourceTypeFromString(s string) (KnowledgeSourceType, error) {
	switch s {
	case "file":
		return KnowledgeSourceTypeFile, nil
	case "url":
		return KnowledgeSourceTypeUrl, nil
	}
	var t KnowledgeSourceType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e KnowledgeSourceType) Ptr() *KnowledgeSourceType {
	return &e
}
