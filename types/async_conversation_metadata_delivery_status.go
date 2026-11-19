package types

import (
	"fmt"
)

type AsyncConversationMetadataDeliveryStatus string

const (
	AsyncConversationMetadataDeliveryStatusPending AsyncConversationMetadataDeliveryStatus = "pending"
	AsyncConversationMetadataDeliveryStatusSuccess AsyncConversationMetadataDeliveryStatus = "success"
	AsyncConversationMetadataDeliveryStatusFailed  AsyncConversationMetadataDeliveryStatus = "failed"
)

func NewAsyncConversationMetadataDeliveryStatusFromString(s string) (AsyncConversationMetadataDeliveryStatus, error) {
	switch s {
	case "pending":
		return AsyncConversationMetadataDeliveryStatusPending, nil
	case "success":
		return AsyncConversationMetadataDeliveryStatusSuccess, nil
	case "failed":
		return AsyncConversationMetadataDeliveryStatusFailed, nil
	}
	var t AsyncConversationMetadataDeliveryStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AsyncConversationMetadataDeliveryStatus) Ptr() *AsyncConversationMetadataDeliveryStatus {
	return &e
}
