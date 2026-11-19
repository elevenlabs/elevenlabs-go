package types

import (
	"fmt"
)

type BucketingStatus string

const (
	BucketingStatusPending   BucketingStatus = "pending"
	BucketingStatusCompleted BucketingStatus = "completed"
	BucketingStatusFailed    BucketingStatus = "failed"
)

func NewBucketingStatusFromString(s string) (BucketingStatus, error) {
	switch s {
	case "pending":
		return BucketingStatusPending, nil
	case "completed":
		return BucketingStatusCompleted, nil
	case "failed":
		return BucketingStatusFailed, nil
	}
	var t BucketingStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BucketingStatus) Ptr() *BucketingStatus {
	return &e
}
