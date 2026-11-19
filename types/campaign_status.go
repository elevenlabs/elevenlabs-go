package types

import (
	"fmt"
)

type CampaignStatus string

const (
	CampaignStatusActive    CampaignStatus = "active"
	CampaignStatusPaused    CampaignStatus = "paused"
	CampaignStatusStopping  CampaignStatus = "stopping"
	CampaignStatusStopped   CampaignStatus = "stopped"
	CampaignStatusCompleted CampaignStatus = "completed"
	CampaignStatusFailed    CampaignStatus = "failed"
	CampaignStatusAttention CampaignStatus = "attention"
)

func NewCampaignStatusFromString(s string) (CampaignStatus, error) {
	switch s {
	case "active":
		return CampaignStatusActive, nil
	case "paused":
		return CampaignStatusPaused, nil
	case "stopping":
		return CampaignStatusStopping, nil
	case "stopped":
		return CampaignStatusStopped, nil
	case "completed":
		return CampaignStatusCompleted, nil
	case "failed":
		return CampaignStatusFailed, nil
	case "attention":
		return CampaignStatusAttention, nil
	}
	var t CampaignStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CampaignStatus) Ptr() *CampaignStatus {
	return &e
}
