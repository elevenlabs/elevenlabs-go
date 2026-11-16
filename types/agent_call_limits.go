package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentCallLimits struct {
	// The maximum number of concurrent conversations. -1 indicates that there is no maximum
	AgentConcurrencyLimit *int `json:"agent_concurrency_limit,omitempty" xml:"agent_concurrency_limit,omitempty"`
	// The maximum number of conversations per day
	DailyLimit *int `json:"daily_limit,omitempty" xml:"daily_limit,omitempty"`
	// Whether to enable bursting. If true, exceeding workspace concurrency limit will be allowed up to 3 times the limit. Calls will be charged at double rate when exceeding the limit.
	BurstingEnabled *bool `json:"bursting_enabled,omitempty" xml:"bursting_enabled,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentCallLimits) GetAgentConcurrencyLimit() (value *int) {
	if a == nil {
		return
	}
	return a.AgentConcurrencyLimit
}

func (a *AgentCallLimits) GetDailyLimit() (value *int) {
	if a == nil {
		return
	}
	return a.DailyLimit
}

func (a *AgentCallLimits) GetBurstingEnabled() (value *bool) {
	if a == nil {
		return
	}
	return a.BurstingEnabled
}

func (a *AgentCallLimits) SetAgentConcurrencyLimit(agentConcurrencyLimit *int) {
	a.AgentConcurrencyLimit = agentConcurrencyLimit
}

func (a *AgentCallLimits) SetDailyLimit(dailyLimit *int) {
	a.DailyLimit = dailyLimit
}

func (a *AgentCallLimits) SetBurstingEnabled(burstingEnabled *bool) {
	a.BurstingEnabled = burstingEnabled
}

func (a *AgentCallLimits) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentCallLimits) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentCallLimits to string"
	}
	return string(jsonData)
}

func (a *AgentCallLimits) UnmarshalJSON(data []byte) error {
	type unmarshaler AgentCallLimits
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AgentCallLimits(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentCallLimits) MarshalJSON() ([]byte, error) {
	type embed AgentCallLimits
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
