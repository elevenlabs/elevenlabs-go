package types

import (
	"fmt"
)

type KnowledgeBaseDocumentType string

const (
	KnowledgeBaseDocumentTypeFile   KnowledgeBaseDocumentType = "file"
	KnowledgeBaseDocumentTypeUrl    KnowledgeBaseDocumentType = "url"
	KnowledgeBaseDocumentTypeText   KnowledgeBaseDocumentType = "text"
	KnowledgeBaseDocumentTypeFolder KnowledgeBaseDocumentType = "folder"
)

func NewKnowledgeBaseDocumentTypeFromString(s string) (KnowledgeBaseDocumentType, error) {
	switch s {
	case "file":
		return KnowledgeBaseDocumentTypeFile, nil
	case "url":
		return KnowledgeBaseDocumentTypeUrl, nil
	case "text":
		return KnowledgeBaseDocumentTypeText, nil
	case "folder":
		return KnowledgeBaseDocumentTypeFolder, nil
	}
	var t KnowledgeBaseDocumentType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e KnowledgeBaseDocumentType) Ptr() *KnowledgeBaseDocumentType {
	return &e
}
