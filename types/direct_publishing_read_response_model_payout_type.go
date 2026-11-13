package types

import (
	"fmt"
)

type DirectPublishingReadResponseModelPayoutType string

const (
	DirectPublishingReadResponseModelPayoutTypeNone            DirectPublishingReadResponseModelPayoutType = "none"
	DirectPublishingReadResponseModelPayoutTypeEngagementBased DirectPublishingReadResponseModelPayoutType = "engagement_based"
	DirectPublishingReadResponseModelPayoutTypeFixedPayout     DirectPublishingReadResponseModelPayoutType = "fixed_payout"
)

func NewDirectPublishingReadResponseModelPayoutTypeFromString(s string) (DirectPublishingReadResponseModelPayoutType, error) {
	switch s {
	case "none":
		return DirectPublishingReadResponseModelPayoutTypeNone, nil
	case "engagement_based":
		return DirectPublishingReadResponseModelPayoutTypeEngagementBased, nil
	case "fixed_payout":
		return DirectPublishingReadResponseModelPayoutTypeFixedPayout, nil
	}
	var t DirectPublishingReadResponseModelPayoutType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DirectPublishingReadResponseModelPayoutType) Ptr() *DirectPublishingReadResponseModelPayoutType {
	return &e
}
