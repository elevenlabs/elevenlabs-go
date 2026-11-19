package types

import (
	"fmt"
)

type CampaignSource string

const (
	CampaignSourceUi  CampaignSource = "ui"
	CampaignSourceApi CampaignSource = "api"
)

func NewCampaignSourceFromString(s string) (CampaignSource, error) {
	switch s {
	case "ui":
		return CampaignSourceUi, nil
	case "api":
		return CampaignSourceApi, nil
	}
	var t CampaignSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CampaignSource) Ptr() *CampaignSource {
	return &e
}
