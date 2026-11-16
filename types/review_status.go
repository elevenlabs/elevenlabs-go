package types

import (
	"fmt"
)

// The review status of the voice.
type ReviewStatus string

const (
	ReviewStatusNotRequested          ReviewStatus = "not_requested"
	ReviewStatusPending               ReviewStatus = "pending"
	ReviewStatusDeclined              ReviewStatus = "declined"
	ReviewStatusDeclinedResubmittable ReviewStatus = "declined_resubmittable"
	ReviewStatusAllowed               ReviewStatus = "allowed"
	ReviewStatusAllowedWithChanges    ReviewStatus = "allowed_with_changes"
)

func NewReviewStatusFromString(s string) (ReviewStatus, error) {
	switch s {
	case "not_requested":
		return ReviewStatusNotRequested, nil
	case "pending":
		return ReviewStatusPending, nil
	case "declined":
		return ReviewStatusDeclined, nil
	case "declined_resubmittable":
		return ReviewStatusDeclinedResubmittable, nil
	case "allowed":
		return ReviewStatusAllowed, nil
	case "allowed_with_changes":
		return ReviewStatusAllowedWithChanges, nil
	}
	var t ReviewStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReviewStatus) Ptr() *ReviewStatus {
	return &e
}
