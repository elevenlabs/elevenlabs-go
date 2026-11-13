package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CampaignConfigOutput struct {
	Policy *PolicyConfigOutput `json:"policy,omitempty" xml:"policy,omitempty"`
	// Overrides of `policy`, keyed by agent ID.
	AgentPolicies map[string]*AgentPolicyOverrideOutput `json:"agent_policies,omitempty"`
	// Overrides for WhatsApp voice campaigns, layered between the workspace policy and an agent's overrides. Sections left out follow WhatsApp voice's built-in defaults.
	WhatsappVoicePolicy *AgentPolicyOverrideOutput `json:"whatsapp_voice_policy,omitempty" xml:"whatsapp_voice_policy,omitempty"`
	extraProperties     map[string]interface{}
}

func (c *CampaignConfigOutput) GetPolicy() (value *PolicyConfigOutput) {
	if c == nil {
		return
	}
	return c.Policy
}

func (c *CampaignConfigOutput) GetAgentPolicies() (value map[string]*AgentPolicyOverrideOutput) {
	if c == nil {
		return
	}
	return c.AgentPolicies
}

func (c *CampaignConfigOutput) GetWhatsappVoicePolicy() (value *AgentPolicyOverrideOutput) {
	if c == nil {
		return
	}
	return c.WhatsappVoicePolicy
}

func (c *CampaignConfigOutput) SetPolicy(policy *PolicyConfigOutput) {
	c.Policy = policy
}

func (c *CampaignConfigOutput) SetAgentPolicies(agentPolicies map[string]*AgentPolicyOverrideOutput) {
	c.AgentPolicies = agentPolicies
}

func (c *CampaignConfigOutput) SetWhatsappVoicePolicy(whatsappVoicePolicy *AgentPolicyOverrideOutput) {
	c.WhatsappVoicePolicy = whatsappVoicePolicy
}

func (c *CampaignConfigOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CampaignConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CampaignConfigOutput to string"
	}
	return string(jsonData)
}

func (c *CampaignConfigOutput) UnmarshalJSON(data []byte) error {
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

func (c *CampaignConfigOutput) MarshalJSON() ([]byte, error) {
	type embed CampaignConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
