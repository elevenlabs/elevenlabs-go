package types

import (
	"fmt"
)

type DirectPublishingReadResponseModelTargetAudience string

const (
	DirectPublishingReadResponseModelTargetAudienceChildren   DirectPublishingReadResponseModelTargetAudience = "children"
	DirectPublishingReadResponseModelTargetAudienceYoungAdult DirectPublishingReadResponseModelTargetAudience = "young adult"
	DirectPublishingReadResponseModelTargetAudienceAdult      DirectPublishingReadResponseModelTargetAudience = "adult"
	DirectPublishingReadResponseModelTargetAudienceAllAges    DirectPublishingReadResponseModelTargetAudience = "all ages"
)

func NewDirectPublishingReadResponseModelTargetAudienceFromString(s string) (DirectPublishingReadResponseModelTargetAudience, error) {
	switch s {
	case "children":
		return DirectPublishingReadResponseModelTargetAudienceChildren, nil
	case "young adult":
		return DirectPublishingReadResponseModelTargetAudienceYoungAdult, nil
	case "adult":
		return DirectPublishingReadResponseModelTargetAudienceAdult, nil
	case "all ages":
		return DirectPublishingReadResponseModelTargetAudienceAllAges, nil
	}
	var t DirectPublishingReadResponseModelTargetAudience
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DirectPublishingReadResponseModelTargetAudience) Ptr() *DirectPublishingReadResponseModelTargetAudience {
	return &e
}
