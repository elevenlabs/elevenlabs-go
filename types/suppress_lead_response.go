package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SuppressLeadResponse struct {
	CampaignId string `json:"campaign_id" xml:"campaign_id" required:"true"`
	LeadId     string `json:"lead_id" xml:"lead_id" required:"true"`
	// Whether a call already handed off for dialing may still go through.
	HasCallsInFlight bool `json:"has_calls_in_flight" xml:"has_calls_in_flight" required:"true"`
	extraProperties  map[string]interface{}
}

func (s *SuppressLeadResponse) GetCampaignId() (value string) {
	if s == nil {
		return
	}
	return s.CampaignId
}

func (s *SuppressLeadResponse) GetLeadId() (value string) {
	if s == nil {
		return
	}
	return s.LeadId
}

func (s *SuppressLeadResponse) GetHasCallsInFlight() (value bool) {
	if s == nil {
		return
	}
	return s.HasCallsInFlight
}

func (s *SuppressLeadResponse) SetCampaignId(campaignId string) {
	s.CampaignId = campaignId
}

func (s *SuppressLeadResponse) SetLeadId(leadId string) {
	s.LeadId = leadId
}

func (s *SuppressLeadResponse) SetHasCallsInFlight(hasCallsInFlight bool) {
	s.HasCallsInFlight = hasCallsInFlight
}

func (s *SuppressLeadResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SuppressLeadResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SuppressLeadResponse to string"
	}
	return string(jsonData)
}

func (s *SuppressLeadResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SuppressLeadResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SuppressLeadResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SuppressLeadResponse) MarshalJSON() ([]byte, error) {
	type embed SuppressLeadResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
