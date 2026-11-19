package types

import (
	"fmt"
)

// Processing status for knowledge sources.
//
// - `processing`: the file or site is being read.
// - `ready`: done; the results are available.
// - `error`: it failed and can be retried.
//
// See `progress` for detail.
type KnowledgeSourceStatus string

const (
	KnowledgeSourceStatusProcessing KnowledgeSourceStatus = "processing"
	KnowledgeSourceStatusReady      KnowledgeSourceStatus = "ready"
	KnowledgeSourceStatusError      KnowledgeSourceStatus = "error"
)

func NewKnowledgeSourceStatusFromString(s string) (KnowledgeSourceStatus, error) {
	switch s {
	case "processing":
		return KnowledgeSourceStatusProcessing, nil
	case "ready":
		return KnowledgeSourceStatusReady, nil
	case "error":
		return KnowledgeSourceStatusError, nil
	}
	var t KnowledgeSourceStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e KnowledgeSourceStatus) Ptr() *KnowledgeSourceStatus {
	return &e
}
