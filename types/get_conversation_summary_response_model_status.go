package types

import (
	"fmt"
)

type GetConversationSummaryResponseModelStatus string

const (
	GetConversationSummaryResponseModelStatusInitiated  GetConversationSummaryResponseModelStatus = "initiated"
	GetConversationSummaryResponseModelStatusInProgress GetConversationSummaryResponseModelStatus = "in-progress"
	GetConversationSummaryResponseModelStatusProcessing GetConversationSummaryResponseModelStatus = "processing"
	GetConversationSummaryResponseModelStatusDone       GetConversationSummaryResponseModelStatus = "done"
	GetConversationSummaryResponseModelStatusFailed     GetConversationSummaryResponseModelStatus = "failed"
)

func NewGetConversationSummaryResponseModelStatusFromString(s string) (GetConversationSummaryResponseModelStatus, error) {
	switch s {
	case "initiated":
		return GetConversationSummaryResponseModelStatusInitiated, nil
	case "in-progress":
		return GetConversationSummaryResponseModelStatusInProgress, nil
	case "processing":
		return GetConversationSummaryResponseModelStatusProcessing, nil
	case "done":
		return GetConversationSummaryResponseModelStatusDone, nil
	case "failed":
		return GetConversationSummaryResponseModelStatusFailed, nil
	}
	var t GetConversationSummaryResponseModelStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GetConversationSummaryResponseModelStatus) Ptr() *GetConversationSummaryResponseModelStatus {
	return &e
}
