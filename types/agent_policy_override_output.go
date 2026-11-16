package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Each section set here replaces the whole matching section of the policy.
type AgentPolicyOverrideOutput struct {
	Timezone        *string                `json:"timezone,omitempty" xml:"timezone,omitempty"`
	Retry           *RetryPolicy           `json:"retry,omitempty" xml:"retry,omitempty"`
	CallingPeriods  []*CallingPeriod       `json:"calling_periods,omitempty" xml:"calling_periods,omitempty" minItems:"1"`
	DataCollection  *DataCollectionConfig  `json:"data_collection,omitempty" xml:"data_collection,omitempty"`
	FollowupContext *FollowupContextPolicy `json:"followup_context,omitempty" xml:"followup_context,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentPolicyOverrideOutput) GetTimezone() (value *string) {
	if a == nil {
		return
	}
	return a.Timezone
}

func (a *AgentPolicyOverrideOutput) GetRetry() (value *RetryPolicy) {
	if a == nil {
		return
	}
	return a.Retry
}

func (a *AgentPolicyOverrideOutput) GetCallingPeriods() (value []*CallingPeriod) {
	if a == nil {
		return
	}
	return a.CallingPeriods
}

func (a *AgentPolicyOverrideOutput) GetDataCollection() (value *DataCollectionConfig) {
	if a == nil {
		return
	}
	return a.DataCollection
}

func (a *AgentPolicyOverrideOutput) GetFollowupContext() (value *FollowupContextPolicy) {
	if a == nil {
		return
	}
	return a.FollowupContext
}

func (a *AgentPolicyOverrideOutput) SetTimezone(timezone *string) {
	a.Timezone = timezone
}

func (a *AgentPolicyOverrideOutput) SetRetry(retry *RetryPolicy) {
	a.Retry = retry
}

func (a *AgentPolicyOverrideOutput) SetCallingPeriods(callingPeriods []*CallingPeriod) {
	a.CallingPeriods = callingPeriods
}

func (a *AgentPolicyOverrideOutput) SetDataCollection(dataCollection *DataCollectionConfig) {
	a.DataCollection = dataCollection
}

func (a *AgentPolicyOverrideOutput) SetFollowupContext(followupContext *FollowupContextPolicy) {
	a.FollowupContext = followupContext
}

func (a *AgentPolicyOverrideOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentPolicyOverrideOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentPolicyOverrideOutput to string"
	}
	return string(jsonData)
}

func (a *AgentPolicyOverrideOutput) UnmarshalJSON(data []byte) error {
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

func (a *AgentPolicyOverrideOutput) MarshalJSON() ([]byte, error) {
	type embed AgentPolicyOverrideOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
