package types

import (
	"fmt"
)

type ConversationSummaryResponseModelStatus string

const (
	ConversationSummaryResponseModelStatusInitiated  ConversationSummaryResponseModelStatus = "initiated"
	ConversationSummaryResponseModelStatusInProgress ConversationSummaryResponseModelStatus = "in-progress"
	ConversationSummaryResponseModelStatusProcessing ConversationSummaryResponseModelStatus = "processing"
	ConversationSummaryResponseModelStatusDone       ConversationSummaryResponseModelStatus = "done"
	ConversationSummaryResponseModelStatusFailed     ConversationSummaryResponseModelStatus = "failed"
)

func NewConversationSummaryResponseModelStatusFromString(s string) (ConversationSummaryResponseModelStatus, error) {
	switch s {
	case "initiated":
		return ConversationSummaryResponseModelStatusInitiated, nil
	case "in-progress":
		return ConversationSummaryResponseModelStatusInProgress, nil
	case "processing":
		return ConversationSummaryResponseModelStatusProcessing, nil
	case "done":
		return ConversationSummaryResponseModelStatusDone, nil
	case "failed":
		return ConversationSummaryResponseModelStatusFailed, nil
	}
	var t ConversationSummaryResponseModelStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationSummaryResponseModelStatus) Ptr() *ConversationSummaryResponseModelStatus {
	return &e
}
