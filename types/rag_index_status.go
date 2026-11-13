package types

import (
	"fmt"
)

type RagIndexStatus string

const (
	RagIndexStatusNew               RagIndexStatus = "new"
	RagIndexStatusCreated           RagIndexStatus = "created"
	RagIndexStatusProcessing        RagIndexStatus = "processing"
	RagIndexStatusFailed            RagIndexStatus = "failed"
	RagIndexStatusSucceeded         RagIndexStatus = "succeeded"
	RagIndexStatusRagLimitExceeded  RagIndexStatus = "rag_limit_exceeded"
	RagIndexStatusDocumentTooSmall  RagIndexStatus = "document_too_small"
	RagIndexStatusCannotIndexFolder RagIndexStatus = "cannot_index_folder"
)

func NewRagIndexStatusFromString(s string) (RagIndexStatus, error) {
	switch s {
	case "new":
		return RagIndexStatusNew, nil
	case "created":
		return RagIndexStatusCreated, nil
	case "processing":
		return RagIndexStatusProcessing, nil
	case "failed":
		return RagIndexStatusFailed, nil
	case "succeeded":
		return RagIndexStatusSucceeded, nil
	case "rag_limit_exceeded":
		return RagIndexStatusRagLimitExceeded, nil
	case "document_too_small":
		return RagIndexStatusDocumentTooSmall, nil
	case "cannot_index_folder":
		return RagIndexStatusCannotIndexFolder, nil
	}
	var t RagIndexStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e RagIndexStatus) Ptr() *RagIndexStatus {
	return &e
}
