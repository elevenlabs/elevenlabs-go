package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CampaignProgressResponse struct {
	Leads         *LeadsProgress        `json:"leads" xml:"leads" required:"true"`
	Conversations *ConversationProgress `json:"conversations" xml:"conversations" required:"true"`
	// When the next call is due; unset when none is scheduled.
	NextCallTimeUnix *int `json:"next_call_time_unix" xml:"next_call_time_unix" required:"true"`
	extraProperties  map[string]interface{}
}

func (c *CampaignProgressResponse) GetLeads() (value *LeadsProgress) {
	if c == nil {
		return
	}
	return c.Leads
}

func (c *CampaignProgressResponse) GetConversations() (value *ConversationProgress) {
	if c == nil {
		return
	}
	return c.Conversations
}

func (c *CampaignProgressResponse) GetNextCallTimeUnix() (value *int) {
	if c == nil {
		return
	}
	return c.NextCallTimeUnix
}

func (c *CampaignProgressResponse) SetLeads(leads *LeadsProgress) {
	c.Leads = leads
}

func (c *CampaignProgressResponse) SetConversations(conversations *ConversationProgress) {
	c.Conversations = conversations
}

func (c *CampaignProgressResponse) SetNextCallTimeUnix(nextCallTimeUnix *int) {
	c.NextCallTimeUnix = nextCallTimeUnix
}

func (c *CampaignProgressResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CampaignProgressResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CampaignProgressResponse to string"
	}
	return string(jsonData)
}

func (c *CampaignProgressResponse) UnmarshalJSON(data []byte) error {
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

func (c *CampaignProgressResponse) MarshalJSON() ([]byte, error) {
	type embed CampaignProgressResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
