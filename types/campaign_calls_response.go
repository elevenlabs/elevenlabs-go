package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CampaignCallsResponse struct {
	Calls           []*CallView `json:"calls" xml:"calls" required:"true"`
	NextCursor      *string     `json:"next_cursor" xml:"next_cursor" required:"true"`
	HasMore         bool        `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CampaignCallsResponse) GetCalls() (value []*CallView) {
	if c == nil {
		return
	}
	return c.Calls
}

func (c *CampaignCallsResponse) GetNextCursor() (value *string) {
	if c == nil {
		return
	}
	return c.NextCursor
}

func (c *CampaignCallsResponse) GetHasMore() (value bool) {
	if c == nil {
		return
	}
	return c.HasMore
}

func (c *CampaignCallsResponse) SetCalls(calls []*CallView) {
	c.Calls = calls
}

func (c *CampaignCallsResponse) SetNextCursor(nextCursor *string) {
	c.NextCursor = nextCursor
}

func (c *CampaignCallsResponse) SetHasMore(hasMore bool) {
	c.HasMore = hasMore
}

func (c *CampaignCallsResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CampaignCallsResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CampaignCallsResponse to string"
	}
	return string(jsonData)
}

func (c *CampaignCallsResponse) UnmarshalJSON(data []byte) error {
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

func (c *CampaignCallsResponse) MarshalJSON() ([]byte, error) {
	type embed CampaignCallsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
