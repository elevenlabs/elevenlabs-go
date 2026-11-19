package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CampaignResponse struct {
	CampaignId string `json:"campaign_id" xml:"campaign_id" required:"true"`
	// The campaign's name.
	CallName      string                         `json:"call_name" xml:"call_name" required:"true"`
	AgentId       string                         `json:"agent_id" xml:"agent_id" required:"true"`
	ChannelConfig *CampaignResponseChannelConfig `json:"channel_config" xml:"channel_config" required:"true"`
	Status        CampaignStatus                 `json:"status" xml:"status" required:"true"`
	// The policy the campaign took when it was created, laid out like a submitted `policy`.
	Policy  *CampaignPolicySnapshot `json:"policy" xml:"policy" required:"true"`
	Variant CampaignVariant         `json:"variant" xml:"variant" required:"true"`
	// Whether the campaign was created in the app or through the API.
	Source        *CampaignSource `json:"source" xml:"source" required:"true"`
	CreatedAtUnix int             `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	// When the campaign starts calling; unset to start once created.
	ScheduledTimeUnix *int `json:"scheduled_time_unix" xml:"scheduled_time_unix" required:"true"`
	UpdatedAtUnix     int  `json:"updated_at_unix" xml:"updated_at_unix" required:"true"`
	// The agent branch the campaign calls with.
	BranchId *string `json:"branch_id" xml:"branch_id" required:"true"`
	// The environment the agent's environment variables resolve in.
	Environment *string `json:"environment" xml:"environment" required:"true"`
	// The most calls the campaign places at once.
	TargetConcurrencyLimit *int                   `json:"target_concurrency_limit" xml:"target_concurrency_limit" required:"true"`
	Channel                OutboundChannel        `json:"channel" xml:"channel" required:"true"`
	Mode                   OutboundMode           `json:"mode" xml:"mode" required:"true"`
	Counts                 *CampaignSummaryCounts `json:"counts" xml:"counts" required:"true"`
	AgentName              string                 `json:"agent_name" xml:"agent_name" required:"true"`
	extraProperties        map[string]interface{}
}

func (c *CampaignResponse) GetCampaignId() (value string) {
	if c == nil {
		return
	}
	return c.CampaignId
}

func (c *CampaignResponse) GetCallName() (value string) {
	if c == nil {
		return
	}
	return c.CallName
}

func (c *CampaignResponse) GetAgentId() (value string) {
	if c == nil {
		return
	}
	return c.AgentId
}

func (c *CampaignResponse) GetChannelConfig() (value *CampaignResponseChannelConfig) {
	if c == nil {
		return
	}
	return c.ChannelConfig
}

func (c *CampaignResponse) GetStatus() (value CampaignStatus) {
	if c == nil {
		return
	}
	return c.Status
}

func (c *CampaignResponse) GetPolicy() (value *CampaignPolicySnapshot) {
	if c == nil {
		return
	}
	return c.Policy
}

func (c *CampaignResponse) GetVariant() (value CampaignVariant) {
	if c == nil {
		return
	}
	return c.Variant
}

func (c *CampaignResponse) GetSource() (value *CampaignSource) {
	if c == nil {
		return
	}
	return c.Source
}

func (c *CampaignResponse) GetCreatedAtUnix() (value int) {
	if c == nil {
		return
	}
	return c.CreatedAtUnix
}

func (c *CampaignResponse) GetScheduledTimeUnix() (value *int) {
	if c == nil {
		return
	}
	return c.ScheduledTimeUnix
}

func (c *CampaignResponse) GetUpdatedAtUnix() (value int) {
	if c == nil {
		return
	}
	return c.UpdatedAtUnix
}

func (c *CampaignResponse) GetBranchId() (value *string) {
	if c == nil {
		return
	}
	return c.BranchId
}

func (c *CampaignResponse) GetEnvironment() (value *string) {
	if c == nil {
		return
	}
	return c.Environment
}

func (c *CampaignResponse) GetTargetConcurrencyLimit() (value *int) {
	if c == nil {
		return
	}
	return c.TargetConcurrencyLimit
}

func (c *CampaignResponse) GetChannel() (value OutboundChannel) {
	if c == nil {
		return
	}
	return c.Channel
}

func (c *CampaignResponse) GetMode() (value OutboundMode) {
	if c == nil {
		return
	}
	return c.Mode
}

func (c *CampaignResponse) GetCounts() (value *CampaignSummaryCounts) {
	if c == nil {
		return
	}
	return c.Counts
}

func (c *CampaignResponse) GetAgentName() (value string) {
	if c == nil {
		return
	}
	return c.AgentName
}

func (c *CampaignResponse) SetCampaignId(campaignId string) {
	c.CampaignId = campaignId
}

func (c *CampaignResponse) SetCallName(callName string) {
	c.CallName = callName
}

func (c *CampaignResponse) SetAgentId(agentId string) {
	c.AgentId = agentId
}

func (c *CampaignResponse) SetChannelConfig(channelConfig *CampaignResponseChannelConfig) {
	c.ChannelConfig = channelConfig
}

func (c *CampaignResponse) SetStatus(status CampaignStatus) {
	c.Status = status
}

func (c *CampaignResponse) SetPolicy(policy *CampaignPolicySnapshot) {
	c.Policy = policy
}

func (c *CampaignResponse) SetVariant(variant CampaignVariant) {
	c.Variant = variant
}

func (c *CampaignResponse) SetSource(source *CampaignSource) {
	c.Source = source
}

func (c *CampaignResponse) SetCreatedAtUnix(createdAtUnix int) {
	c.CreatedAtUnix = createdAtUnix
}

func (c *CampaignResponse) SetScheduledTimeUnix(scheduledTimeUnix *int) {
	c.ScheduledTimeUnix = scheduledTimeUnix
}

func (c *CampaignResponse) SetUpdatedAtUnix(updatedAtUnix int) {
	c.UpdatedAtUnix = updatedAtUnix
}

func (c *CampaignResponse) SetBranchId(branchId *string) {
	c.BranchId = branchId
}

func (c *CampaignResponse) SetEnvironment(environment *string) {
	c.Environment = environment
}

func (c *CampaignResponse) SetTargetConcurrencyLimit(targetConcurrencyLimit *int) {
	c.TargetConcurrencyLimit = targetConcurrencyLimit
}

func (c *CampaignResponse) SetChannel(channel OutboundChannel) {
	c.Channel = channel
}

func (c *CampaignResponse) SetMode(mode OutboundMode) {
	c.Mode = mode
}

func (c *CampaignResponse) SetCounts(counts *CampaignSummaryCounts) {
	c.Counts = counts
}

func (c *CampaignResponse) SetAgentName(agentName string) {
	c.AgentName = agentName
}

func (c *CampaignResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CampaignResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CampaignResponse to string"
	}
	return string(jsonData)
}

func (c *CampaignResponse) UnmarshalJSON(data []byte) error {
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

func (c *CampaignResponse) MarshalJSON() ([]byte, error) {
	type embed CampaignResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
