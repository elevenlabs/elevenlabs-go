package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubmitCampaignResponse struct {
	Campaign *CampaignResponse `json:"campaign" xml:"campaign" required:"true"`
	// False when the request repeated an earlier one and returned its campaign.
	Created bool `json:"created" xml:"created" required:"true"`
	// Whether calling starts right away rather than at the next calling window.
	WithinCallingWindow bool `json:"within_calling_window" xml:"within_calling_window" required:"true"`
	extraProperties     map[string]interface{}
}

func (s *SubmitCampaignResponse) GetCampaign() (value *CampaignResponse) {
	if s == nil {
		return
	}
	return s.Campaign
}

func (s *SubmitCampaignResponse) GetCreated() (value bool) {
	if s == nil {
		return
	}
	return s.Created
}

func (s *SubmitCampaignResponse) GetWithinCallingWindow() (value bool) {
	if s == nil {
		return
	}
	return s.WithinCallingWindow
}

func (s *SubmitCampaignResponse) SetCampaign(campaign *CampaignResponse) {
	s.Campaign = campaign
}

func (s *SubmitCampaignResponse) SetCreated(created bool) {
	s.Created = created
}

func (s *SubmitCampaignResponse) SetWithinCallingWindow(withinCallingWindow bool) {
	s.WithinCallingWindow = withinCallingWindow
}

func (s *SubmitCampaignResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubmitCampaignResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubmitCampaignResponse to string"
	}
	return string(jsonData)
}

func (s *SubmitCampaignResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubmitCampaignResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubmitCampaignResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubmitCampaignResponse) MarshalJSON() ([]byte, error) {
	type embed SubmitCampaignResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
