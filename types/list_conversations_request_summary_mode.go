package types

import (
	"fmt"
)

// Whether to include transcript summaries in the response.
type ListConversationsRequestSummaryMode string

const (
	ListConversationsRequestSummaryModeExclude ListConversationsRequestSummaryMode = "exclude"
	ListConversationsRequestSummaryModeInclude ListConversationsRequestSummaryMode = "include"
)

func NewListConversationsRequestSummaryModeFromString(s string) (ListConversationsRequestSummaryMode, error) {
	switch s {
	case "exclude":
		return ListConversationsRequestSummaryModeExclude, nil
	case "include":
		return ListConversationsRequestSummaryModeInclude, nil
	}
	var t ListConversationsRequestSummaryMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListConversationsRequestSummaryMode) Ptr() *ListConversationsRequestSummaryMode {
	return &e
}
