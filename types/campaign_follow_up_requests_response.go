package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CampaignFollowUpRequestsResponse struct {
	Callbacks       *FollowUpRequestCounts `json:"callbacks" xml:"callbacks" required:"true"`
	Redirects       *FollowUpRequestCounts `json:"redirects" xml:"redirects" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CampaignFollowUpRequestsResponse) GetCallbacks() (value *FollowUpRequestCounts) {
	if c == nil {
		return
	}
	return c.Callbacks
}

func (c *CampaignFollowUpRequestsResponse) GetRedirects() (value *FollowUpRequestCounts) {
	if c == nil {
		return
	}
	return c.Redirects
}

func (c *CampaignFollowUpRequestsResponse) SetCallbacks(callbacks *FollowUpRequestCounts) {
	c.Callbacks = callbacks
}

func (c *CampaignFollowUpRequestsResponse) SetRedirects(redirects *FollowUpRequestCounts) {
	c.Redirects = redirects
}

func (c *CampaignFollowUpRequestsResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CampaignFollowUpRequestsResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CampaignFollowUpRequestsResponse to string"
	}
	return string(jsonData)
}

func (c *CampaignFollowUpRequestsResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CampaignFollowUpRequestsResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CampaignFollowUpRequestsResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CampaignFollowUpRequestsResponse) MarshalJSON() ([]byte, error) {
	type embed CampaignFollowUpRequestsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
