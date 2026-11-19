package types

import (
	"fmt"
)

type GetConversationResponseModelStatus string

const (
	GetConversationResponseModelStatusInitiated  GetConversationResponseModelStatus = "initiated"
	GetConversationResponseModelStatusInProgress GetConversationResponseModelStatus = "in-progress"
	GetConversationResponseModelStatusProcessing GetConversationResponseModelStatus = "processing"
	GetConversationResponseModelStatusDone       GetConversationResponseModelStatus = "done"
	GetConversationResponseModelStatusFailed     GetConversationResponseModelStatus = "failed"
)

func NewGetConversationResponseModelStatusFromString(s string) (GetConversationResponseModelStatus, error) {
	switch s {
	case "initiated":
		return GetConversationResponseModelStatusInitiated, nil
	case "in-progress":
		return GetConversationResponseModelStatusInProgress, nil
	case "processing":
		return GetConversationResponseModelStatusProcessing, nil
	case "done":
		return GetConversationResponseModelStatusDone, nil
	case "failed":
		return GetConversationResponseModelStatusFailed, nil
	}
	var t GetConversationResponseModelStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GetConversationResponseModelStatus) Ptr() *GetConversationResponseModelStatus {
	return &e
}
