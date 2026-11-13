package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CampaignPageResponse struct {
	Campaigns       []*CampaignResponse `json:"campaigns" xml:"campaigns" required:"true"`
	NextCursor      *string             `json:"next_cursor" xml:"next_cursor" required:"true"`
	HasMore         bool                `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CampaignPageResponse) GetCampaigns() (value []*CampaignResponse) {
	if c == nil {
		return
	}
	return c.Campaigns
}

func (c *CampaignPageResponse) GetNextCursor() (value *string) {
	if c == nil {
		return
	}
	return c.NextCursor
}

func (c *CampaignPageResponse) GetHasMore() (value bool) {
	if c == nil {
		return
	}
	return c.HasMore
}

func (c *CampaignPageResponse) SetCampaigns(campaigns []*CampaignResponse) {
	c.Campaigns = campaigns
}

func (c *CampaignPageResponse) SetNextCursor(nextCursor *string) {
	c.NextCursor = nextCursor
}

func (c *CampaignPageResponse) SetHasMore(hasMore bool) {
	c.HasMore = hasMore
}

func (c *CampaignPageResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CampaignPageResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CampaignPageResponse to string"
	}
	return string(jsonData)
}

func (c *CampaignPageResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CampaignPageResponse) MarshalJSON() ([]byte, error) {
	type embed CampaignPageResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
