package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type StopAllCampaignsResponse struct {
	StoppedCampaignIds []string `json:"stopped_campaign_ids" xml:"stopped_campaign_ids" required:"true"`
	// Campaigns whose calls already placed are still finishing.
	StoppingCampaignIds []string `json:"stopping_campaign_ids" xml:"stopping_campaign_ids" required:"true"`
	extraProperties     map[string]interface{}
}

func (s *StopAllCampaignsResponse) GetStoppedCampaignIds() (value []string) {
	if s == nil {
		return
	}
	return s.StoppedCampaignIds
}

func (s *StopAllCampaignsResponse) GetStoppingCampaignIds() (value []string) {
	if s == nil {
		return
	}
	return s.StoppingCampaignIds
}

func (s *StopAllCampaignsResponse) SetStoppedCampaignIds(stoppedCampaignIds []string) {
	s.StoppedCampaignIds = stoppedCampaignIds
}

func (s *StopAllCampaignsResponse) SetStoppingCampaignIds(stoppingCampaignIds []string) {
	s.StoppingCampaignIds = stoppingCampaignIds
}

func (s *StopAllCampaignsResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StopAllCampaignsResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StopAllCampaignsResponse to string"
	}
	return string(jsonData)
}

func (s *StopAllCampaignsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler StopAllCampaignsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = StopAllCampaignsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StopAllCampaignsResponse) MarshalJSON() ([]byte, error) {
	type embed StopAllCampaignsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
