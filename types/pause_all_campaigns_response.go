package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type PauseAllCampaignsResponse struct {
	PausedCampaignIds []string `json:"paused_campaign_ids" xml:"paused_campaign_ids" required:"true"`
	extraProperties   map[string]interface{}
}

func (p *PauseAllCampaignsResponse) GetPausedCampaignIds() (value []string) {
	if p == nil {
		return
	}
	return p.PausedCampaignIds
}

func (p *PauseAllCampaignsResponse) SetPausedCampaignIds(pausedCampaignIds []string) {
	p.PausedCampaignIds = pausedCampaignIds
}

func (p *PauseAllCampaignsResponse) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PauseAllCampaignsResponse) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PauseAllCampaignsResponse to string"
	}
	return string(jsonData)
}

func (p *PauseAllCampaignsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PauseAllCampaignsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PauseAllCampaignsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PauseAllCampaignsResponse) MarshalJSON() ([]byte, error) {
	type embed PauseAllCampaignsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
