package types

import (
	"fmt"
)

type CampaignVariant string

const (
	CampaignVariantSimple   CampaignVariant = "simple"
	CampaignVariantAdvanced CampaignVariant = "advanced"
)

func NewCampaignVariantFromString(s string) (CampaignVariant, error) {
	switch s {
	case "simple":
		return CampaignVariantSimple, nil
	case "advanced":
		return CampaignVariantAdvanced, nil
	}
	var t CampaignVariant
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CampaignVariant) Ptr() *CampaignVariant {
	return &e
}
