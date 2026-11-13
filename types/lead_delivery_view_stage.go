package types

import (
	"fmt"
)

// The furthest the message got, or `failed` if sending it failed.
type LeadDeliveryViewStage string

const (
	LeadDeliveryViewStageQueued    LeadDeliveryViewStage = "queued"
	LeadDeliveryViewStageSent      LeadDeliveryViewStage = "sent"
	LeadDeliveryViewStageDelivered LeadDeliveryViewStage = "delivered"
	LeadDeliveryViewStageRead      LeadDeliveryViewStage = "read"
	LeadDeliveryViewStageReplied   LeadDeliveryViewStage = "replied"
	LeadDeliveryViewStageFailed    LeadDeliveryViewStage = "failed"
)

func NewLeadDeliveryViewStageFromString(s string) (LeadDeliveryViewStage, error) {
	switch s {
	case "queued":
		return LeadDeliveryViewStageQueued, nil
	case "sent":
		return LeadDeliveryViewStageSent, nil
	case "delivered":
		return LeadDeliveryViewStageDelivered, nil
	case "read":
		return LeadDeliveryViewStageRead, nil
	case "replied":
		return LeadDeliveryViewStageReplied, nil
	case "failed":
		return LeadDeliveryViewStageFailed, nil
	}
	var t LeadDeliveryViewStage
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e LeadDeliveryViewStage) Ptr() *LeadDeliveryViewStage {
	return &e
}
