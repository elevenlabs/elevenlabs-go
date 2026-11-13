package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type PolicyConfigOutput struct {
	// IANA timezone the calling periods are read in.
	Timezone *string      `json:"timezone,omitempty" xml:"timezone,omitempty"`
	Retry    *RetryPolicy `json:"retry,omitempty" xml:"retry,omitempty"`
	// The weekly windows calls may be placed in.
	CallingPeriods []*CallingPeriod `json:"calling_periods,omitempty" xml:"calling_periods,omitempty" minItems:"1"`
	// The agent's data collection items that steer follow-up calls. Each defaults
	// to the built-in item for its purpose; `null` turns it off.
	DataCollection *DataCollectionConfig `json:"data_collection,omitempty" xml:"data_collection,omitempty"`
	// Which `campaign__` dynamic variables follow-up calls get.
	FollowupContext *FollowupContextPolicy `json:"followup_context,omitempty" xml:"followup_context,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PolicyConfigOutput) GetTimezone() (value *string) {
	if p == nil {
		return
	}
	return p.Timezone
}

func (p *PolicyConfigOutput) GetRetry() (value *RetryPolicy) {
	if p == nil {
		return
	}
	return p.Retry
}

func (p *PolicyConfigOutput) GetCallingPeriods() (value []*CallingPeriod) {
	if p == nil {
		return
	}
	return p.CallingPeriods
}

func (p *PolicyConfigOutput) GetDataCollection() (value *DataCollectionConfig) {
	if p == nil {
		return
	}
	return p.DataCollection
}

func (p *PolicyConfigOutput) GetFollowupContext() (value *FollowupContextPolicy) {
	if p == nil {
		return
	}
	return p.FollowupContext
}

func (p *PolicyConfigOutput) SetTimezone(timezone *string) {
	p.Timezone = timezone
}

func (p *PolicyConfigOutput) SetRetry(retry *RetryPolicy) {
	p.Retry = retry
}

func (p *PolicyConfigOutput) SetCallingPeriods(callingPeriods []*CallingPeriod) {
	p.CallingPeriods = callingPeriods
}

func (p *PolicyConfigOutput) SetDataCollection(dataCollection *DataCollectionConfig) {
	p.DataCollection = dataCollection
}

func (p *PolicyConfigOutput) SetFollowupContext(followupContext *FollowupContextPolicy) {
	p.FollowupContext = followupContext
}

func (p *PolicyConfigOutput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PolicyConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PolicyConfigOutput to string"
	}
	return string(jsonData)
}

func (p *PolicyConfigOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler PolicyConfigOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PolicyConfigOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PolicyConfigOutput) MarshalJSON() ([]byte, error) {
	type embed PolicyConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
