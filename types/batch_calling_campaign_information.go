package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BatchCallingCampaignInformation struct {
	CampaignId       string                    `json:"campaign_id" xml:"campaign_id" required:"true"`
	CampaignLeadId   string                    `json:"campaign_lead_id" xml:"campaign_lead_id" required:"true"`
	ColdAttempt      *int                      `json:"cold_attempt,omitempty" xml:"cold_attempt,omitempty"`
	ConversationType *CampaignConversationType `json:"conversation_type,omitempty" xml:"conversation_type,omitempty"`
	extraProperties  map[string]interface{}
}

func (b *BatchCallingCampaignInformation) GetCampaignId() (value string) {
	if b == nil {
		return
	}
	return b.CampaignId
}

func (b *BatchCallingCampaignInformation) GetCampaignLeadId() (value string) {
	if b == nil {
		return
	}
	return b.CampaignLeadId
}

func (b *BatchCallingCampaignInformation) GetColdAttempt() (value *int) {
	if b == nil {
		return
	}
	return b.ColdAttempt
}

func (b *BatchCallingCampaignInformation) GetConversationType() (value *CampaignConversationType) {
	if b == nil {
		return
	}
	return b.ConversationType
}

func (b *BatchCallingCampaignInformation) SetCampaignId(campaignId string) {
	b.CampaignId = campaignId
}

func (b *BatchCallingCampaignInformation) SetCampaignLeadId(campaignLeadId string) {
	b.CampaignLeadId = campaignLeadId
}

func (b *BatchCallingCampaignInformation) SetColdAttempt(coldAttempt *int) {
	b.ColdAttempt = coldAttempt
}

func (b *BatchCallingCampaignInformation) SetConversationType(conversationType *CampaignConversationType) {
	b.ConversationType = conversationType
}

func (b *BatchCallingCampaignInformation) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BatchCallingCampaignInformation) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BatchCallingCampaignInformation to string"
	}
	return string(jsonData)
}

func (b *BatchCallingCampaignInformation) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BatchCallingCampaignInformation) MarshalJSON() ([]byte, error) {
	type embed BatchCallingCampaignInformation
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
