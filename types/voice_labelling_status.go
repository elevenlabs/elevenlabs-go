package types

import (
	"fmt"
)

type VoiceLabellingStatus string

const (
	VoiceLabellingStatusInReview       VoiceLabellingStatus = "in_review"
	VoiceLabellingStatusReviewComplete VoiceLabellingStatus = "review_complete"
)

func NewVoiceLabellingStatusFromString(s string) (VoiceLabellingStatus, error) {
	switch s {
	case "in_review":
		return VoiceLabellingStatusInReview, nil
	case "review_complete":
		return VoiceLabellingStatusReviewComplete, nil
	}
	var t VoiceLabellingStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e VoiceLabellingStatus) Ptr() *VoiceLabellingStatus {
	return &e
}
