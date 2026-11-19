package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type RetryTriggerAction struct {
	// Custom feedback to inject into the agent when retrying after guardrail trigger.
	Feedback        *string `json:"feedback,omitempty" xml:"feedback,omitempty"`
	extraProperties map[string]interface{}
}

func (r *RetryTriggerAction) GetFeedback() (value *string) {
	if r == nil {
		return
	}
	return r.Feedback
}

func (r *RetryTriggerAction) SetFeedback(feedback *string) {
	r.Feedback = feedback
}

func (r *RetryTriggerAction) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RetryTriggerAction) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RetryTriggerAction to string"
	}
	return string(jsonData)
}

func (r *RetryTriggerAction) UnmarshalJSON(data []byte) error {
	type unmarshaler RetryTriggerAction
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RetryTriggerAction(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r, "type")
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RetryTriggerAction) MarshalJSON() ([]byte, error) {
	type embed RetryTriggerAction
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
