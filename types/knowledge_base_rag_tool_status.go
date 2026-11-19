package types

import (
	"fmt"
)

type KnowledgeBaseRagToolStatus string

const (
	KnowledgeBaseRagToolStatusSuccess     KnowledgeBaseRagToolStatus = "success"
	KnowledgeBaseRagToolStatusNoDocuments KnowledgeBaseRagToolStatus = "no_documents"
	KnowledgeBaseRagToolStatusNoResults   KnowledgeBaseRagToolStatus = "no_results"
)

func NewKnowledgeBaseRagToolStatusFromString(s string) (KnowledgeBaseRagToolStatus, error) {
	switch s {
	case "success":
		return KnowledgeBaseRagToolStatusSuccess, nil
	case "no_documents":
		return KnowledgeBaseRagToolStatusNoDocuments, nil
	case "no_results":
		return KnowledgeBaseRagToolStatusNoResults, nil
	}
	var t KnowledgeBaseRagToolStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e KnowledgeBaseRagToolStatus) Ptr() *KnowledgeBaseRagToolStatus {
	return &e
}
