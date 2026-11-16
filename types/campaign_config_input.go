package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CampaignConfigInput struct {
	Policy *PolicyConfigInput `json:"policy,omitempty" xml:"policy,omitempty"`
	// Overrides of `policy`, keyed by agent ID.
	AgentPolicies map[string]*AgentPolicyOverrideInput `json:"agent_policies,omitempty"`
	// Overrides for WhatsApp voice campaigns, layered between the workspace policy and an agent's overrides. Sections left out follow WhatsApp voice's built-in defaults.
	WhatsappVoicePolicy *AgentPolicyOverrideInput `json:"whatsapp_voice_policy,omitempty" xml:"whatsapp_voice_policy,omitempty"`
	extraProperties     map[string]interface{}
}

func (c *CampaignConfigInput) GetPolicy() (value *PolicyConfigInput) {
	if c == nil {
		return
	}
	return c.Policy
}

func (c *CampaignConfigInput) GetAgentPolicies() (value map[string]*AgentPolicyOverrideInput) {
	if c == nil {
		return
	}
	return c.AgentPolicies
}

func (c *CampaignConfigInput) GetWhatsappVoicePolicy() (value *AgentPolicyOverrideInput) {
	if c == nil {
		return
	}
	return c.WhatsappVoicePolicy
}

func (c *CampaignConfigInput) SetPolicy(policy *PolicyConfigInput) {
	c.Policy = policy
}

func (c *CampaignConfigInput) SetAgentPolicies(agentPolicies map[string]*AgentPolicyOverrideInput) {
	c.AgentPolicies = agentPolicies
}

func (c *CampaignConfigInput) SetWhatsappVoicePolicy(whatsappVoicePolicy *AgentPolicyOverrideInput) {
	c.WhatsappVoicePolicy = whatsappVoicePolicy
}

func (c *CampaignConfigInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CampaignConfigInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CampaignConfigInput to string"
	}
	return string(jsonData)
}

func (c *CampaignConfigInput) UnmarshalJSON(data []byte) error {
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

func (c *CampaignConfigInput) MarshalJSON() ([]byte, error) {
	type embed CampaignConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
