package types

import (
	"fmt"
)

// Valid Twilio region IDs.
type TwilioRegionId string

const (
	TwilioRegionIdUs1 TwilioRegionId = "us1"
	TwilioRegionIdIe1 TwilioRegionId = "ie1"
	TwilioRegionIdAu1 TwilioRegionId = "au1"
)

func NewTwilioRegionIdFromString(s string) (TwilioRegionId, error) {
	switch s {
	case "us1":
		return TwilioRegionIdUs1, nil
	case "ie1":
		return TwilioRegionIdIe1, nil
	case "au1":
		return TwilioRegionIdAu1, nil
	}
	var t TwilioRegionId
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TwilioRegionId) Ptr() *TwilioRegionId {
	return &e
}
