package types

import (
	"fmt"
)

type KnowledgeBaseToolStatus string

const (
	KnowledgeBaseToolStatusSuccess             KnowledgeBaseToolStatus = "success"
	KnowledgeBaseToolStatusNoMatchingDocuments KnowledgeBaseToolStatus = "no_matching_documents"
	KnowledgeBaseToolStatusNoResults           KnowledgeBaseToolStatus = "no_results"
)

func NewKnowledgeBaseToolStatusFromString(s string) (KnowledgeBaseToolStatus, error) {
	switch s {
	case "success":
		return KnowledgeBaseToolStatusSuccess, nil
	case "no_matching_documents":
		return KnowledgeBaseToolStatusNoMatchingDocuments, nil
	case "no_results":
		return KnowledgeBaseToolStatusNoResults, nil
	}
	var t KnowledgeBaseToolStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e KnowledgeBaseToolStatus) Ptr() *KnowledgeBaseToolStatus {
	return &e
}
