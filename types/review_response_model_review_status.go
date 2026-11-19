package types

import (
	"fmt"
)

type ReviewResponseModelReviewStatus string

const (
	ReviewResponseModelReviewStatusApproved      ReviewResponseModelReviewStatus = "approved"
	ReviewResponseModelReviewStatusEditsRequired ReviewResponseModelReviewStatus = "edits_required"
	ReviewResponseModelReviewStatusRejected      ReviewResponseModelReviewStatus = "rejected"
)

func NewReviewResponseModelReviewStatusFromString(s string) (ReviewResponseModelReviewStatus, error) {
	switch s {
	case "approved":
		return ReviewResponseModelReviewStatusApproved, nil
	case "edits_required":
		return ReviewResponseModelReviewStatusEditsRequired, nil
	case "rejected":
		return ReviewResponseModelReviewStatusRejected, nil
	}
	var t ReviewResponseModelReviewStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReviewResponseModelReviewStatus) Ptr() *ReviewResponseModelReviewStatus {
	return &e
}
