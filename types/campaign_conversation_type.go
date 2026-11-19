package types

import (
	"fmt"
)

type CampaignConversationType string

const (
	CampaignConversationTypeInitial  CampaignConversationType = "initial"
	CampaignConversationTypeRetry    CampaignConversationType = "retry"
	CampaignConversationTypeCallback CampaignConversationType = "callback"
	CampaignConversationTypeRedirect CampaignConversationType = "redirect"
)

func NewCampaignConversationTypeFromString(s string) (CampaignConversationType, error) {
	switch s {
	case "initial":
		return CampaignConversationTypeInitial, nil
	case "retry":
		return CampaignConversationTypeRetry, nil
	case "callback":
		return CampaignConversationTypeCallback, nil
	case "redirect":
		return CampaignConversationTypeRedirect, nil
	}
	var t CampaignConversationType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CampaignConversationType) Ptr() *CampaignConversationType {
	return &e
}
