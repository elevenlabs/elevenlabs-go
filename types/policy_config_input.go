package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type PolicyConfigInput struct {
	// IANA timezone the calling periods are read in.
	Timezone *string      `json:"timezone,omitempty" xml:"timezone,omitempty"`
	Retry    *RetryPolicy `json:"retry,omitempty" xml:"retry,omitempty"`
	// The weekly windows calls may be placed in.
	CallingPeriods []*CallingPeriod `json:"calling_periods,omitempty" xml:"calling_periods,omitempty" minItems:"1"`
	// Local dates in `timezone` whose calling windows replace `calling_periods`. Date ranges may not overlap.
	DateOverrides []*DateOverride `json:"date_overrides,omitempty" xml:"date_overrides,omitempty" maxItems:"100"`
	// The agent's data collection items that steer follow-up calls. Each defaults
	// to the built-in item for its purpose; `null` turns it off.
	DataCollection *DataCollectionConfig `json:"data_collection,omitempty" xml:"data_collection,omitempty"`
	// Which `campaign__` dynamic variables follow-up calls get.
	FollowupContext *FollowupContextPolicy `json:"followup_context,omitempty" xml:"followup_context,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PolicyConfigInput) GetTimezone() (value *string) {
	if p == nil {
		return
	}
	return p.Timezone
}

func (p *PolicyConfigInput) GetRetry() (value *RetryPolicy) {
	if p == nil {
		return
	}
	return p.Retry
}

func (p *PolicyConfigInput) GetCallingPeriods() (value []*CallingPeriod) {
	if p == nil {
		return
	}
	return p.CallingPeriods
}

func (p *PolicyConfigInput) GetDateOverrides() (value []*DateOverride) {
	if p == nil {
		return
	}
	return p.DateOverrides
}

func (p *PolicyConfigInput) GetDataCollection() (value *DataCollectionConfig) {
	if p == nil {
		return
	}
	return p.DataCollection
}

func (p *PolicyConfigInput) GetFollowupContext() (value *FollowupContextPolicy) {
	if p == nil {
		return
	}
	return p.FollowupContext
}

func (p *PolicyConfigInput) SetTimezone(timezone *string) {
	p.Timezone = timezone
}

func (p *PolicyConfigInput) SetRetry(retry *RetryPolicy) {
	p.Retry = retry
}

func (p *PolicyConfigInput) SetCallingPeriods(callingPeriods []*CallingPeriod) {
	p.CallingPeriods = callingPeriods
}

func (p *PolicyConfigInput) SetDateOverrides(dateOverrides []*DateOverride) {
	p.DateOverrides = dateOverrides
}

func (p *PolicyConfigInput) SetDataCollection(dataCollection *DataCollectionConfig) {
	p.DataCollection = dataCollection
}

func (p *PolicyConfigInput) SetFollowupContext(followupContext *FollowupContextPolicy) {
	p.FollowupContext = followupContext
}

func (p *PolicyConfigInput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PolicyConfigInput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PolicyConfigInput to string"
	}
	return string(jsonData)
}

func (p *PolicyConfigInput) UnmarshalJSON(data []byte) error {
	type unmarshaler PolicyConfigInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PolicyConfigInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PolicyConfigInput) MarshalJSON() ([]byte, error) {
	type embed PolicyConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
