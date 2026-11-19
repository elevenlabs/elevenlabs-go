package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Each section set here replaces the whole matching section of the policy.
type AgentPolicyOverrideInput struct {
	Timezone        *string                `json:"timezone,omitempty" xml:"timezone,omitempty"`
	Retry           *RetryPolicy           `json:"retry,omitempty" xml:"retry,omitempty"`
	CallingPeriods  []*CallingPeriod       `json:"calling_periods,omitempty" xml:"calling_periods,omitempty" minItems:"1"`
	DataCollection  *DataCollectionConfig  `json:"data_collection,omitempty" xml:"data_collection,omitempty"`
	FollowupContext *FollowupContextPolicy `json:"followup_context,omitempty" xml:"followup_context,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentPolicyOverrideInput) GetTimezone() (value *string) {
	if a == nil {
		return
	}
	return a.Timezone
}

func (a *AgentPolicyOverrideInput) GetRetry() (value *RetryPolicy) {
	if a == nil {
		return
	}
	return a.Retry
}

func (a *AgentPolicyOverrideInput) GetCallingPeriods() (value []*CallingPeriod) {
	if a == nil {
		return
	}
	return a.CallingPeriods
}

func (a *AgentPolicyOverrideInput) GetDataCollection() (value *DataCollectionConfig) {
	if a == nil {
		return
	}
	return a.DataCollection
}

func (a *AgentPolicyOverrideInput) GetFollowupContext() (value *FollowupContextPolicy) {
	if a == nil {
		return
	}
	return a.FollowupContext
}

func (a *AgentPolicyOverrideInput) SetTimezone(timezone *string) {
	a.Timezone = timezone
}

func (a *AgentPolicyOverrideInput) SetRetry(retry *RetryPolicy) {
	a.Retry = retry
}

func (a *AgentPolicyOverrideInput) SetCallingPeriods(callingPeriods []*CallingPeriod) {
	a.CallingPeriods = callingPeriods
}

func (a *AgentPolicyOverrideInput) SetDataCollection(dataCollection *DataCollectionConfig) {
	a.DataCollection = dataCollection
}

func (a *AgentPolicyOverrideInput) SetFollowupContext(followupContext *FollowupContextPolicy) {
	a.FollowupContext = followupContext
}

func (a *AgentPolicyOverrideInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentPolicyOverrideInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentPolicyOverrideInput to string"
	}
	return string(jsonData)
}

func (a *AgentPolicyOverrideInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentPolicyOverrideInput) MarshalJSON() ([]byte, error) {
	type embed AgentPolicyOverrideInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
