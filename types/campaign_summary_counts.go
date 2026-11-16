package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CampaignSummaryCounts struct {
	Leads int `json:"leads" xml:"leads" required:"true"`
	// Leads that won't be called again.
	TerminalLeads int `json:"terminal_leads" xml:"terminal_leads" required:"true"`
	// Attention items not yet resolved.
	OpenAttention int `json:"open_attention" xml:"open_attention" required:"true"`
	// How many messages reached each stage; set when getting one message campaign.
	Delivery        *RecipientDeliveryCounts `json:"delivery,omitempty" xml:"delivery,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CampaignSummaryCounts) GetLeads() (value int) {
	if c == nil {
		return
	}
	return c.Leads
}

func (c *CampaignSummaryCounts) GetTerminalLeads() (value int) {
	if c == nil {
		return
	}
	return c.TerminalLeads
}

func (c *CampaignSummaryCounts) GetOpenAttention() (value int) {
	if c == nil {
		return
	}
	return c.OpenAttention
}

func (c *CampaignSummaryCounts) GetDelivery() (value *RecipientDeliveryCounts) {
	if c == nil {
		return
	}
	return c.Delivery
}

func (c *CampaignSummaryCounts) SetLeads(leads int) {
	c.Leads = leads
}

func (c *CampaignSummaryCounts) SetTerminalLeads(terminalLeads int) {
	c.TerminalLeads = terminalLeads
}

func (c *CampaignSummaryCounts) SetOpenAttention(openAttention int) {
	c.OpenAttention = openAttention
}

func (c *CampaignSummaryCounts) SetDelivery(delivery *RecipientDeliveryCounts) {
	c.Delivery = delivery
}

func (c *CampaignSummaryCounts) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CampaignSummaryCounts) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CampaignSummaryCounts to string"
	}
	return string(jsonData)
}

func (c *CampaignSummaryCounts) UnmarshalJSON(data []byte) error {
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

func (c *CampaignSummaryCounts) MarshalJSON() ([]byte, error) {
	type embed CampaignSummaryCounts
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
