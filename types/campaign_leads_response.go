package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CampaignLeadsResponse struct {
	Leads           []*LeadView `json:"leads" xml:"leads" required:"true"`
	NextCursor      *string     `json:"next_cursor" xml:"next_cursor" required:"true"`
	HasMore         bool        `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CampaignLeadsResponse) GetLeads() (value []*LeadView) {
	if c == nil {
		return
	}
	return c.Leads
}

func (c *CampaignLeadsResponse) GetNextCursor() (value *string) {
	if c == nil {
		return
	}
	return c.NextCursor
}

func (c *CampaignLeadsResponse) GetHasMore() (value bool) {
	if c == nil {
		return
	}
	return c.HasMore
}

func (c *CampaignLeadsResponse) SetLeads(leads []*LeadView) {
	c.Leads = leads
}

func (c *CampaignLeadsResponse) SetNextCursor(nextCursor *string) {
	c.NextCursor = nextCursor
}

func (c *CampaignLeadsResponse) SetHasMore(hasMore bool) {
	c.HasMore = hasMore
}

func (c *CampaignLeadsResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CampaignLeadsResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CampaignLeadsResponse to string"
	}
	return string(jsonData)
}

func (c *CampaignLeadsResponse) UnmarshalJSON(data []byte) error {
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

func (c *CampaignLeadsResponse) MarshalJSON() ([]byte, error) {
	type embed CampaignLeadsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
