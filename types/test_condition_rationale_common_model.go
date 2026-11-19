package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Structured rationale for test condition results containing individual failure/success reasons.
type TestConditionRationaleCommonModel struct {
	// List of individual parameter evaluation messages or reasons
	Messages []string `json:"messages,omitempty" xml:"messages,omitempty"`
	// High-level summary of the evaluation result
	Summary         *string `json:"summary,omitempty" xml:"summary,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TestConditionRationaleCommonModel) GetMessages() (value []string) {
	if t == nil {
		return
	}
	return t.Messages
}

func (t *TestConditionRationaleCommonModel) GetSummary() (value *string) {
	if t == nil {
		return
	}
	return t.Summary
}

func (t *TestConditionRationaleCommonModel) SetMessages(messages []string) {
	t.Messages = messages
}

func (t *TestConditionRationaleCommonModel) SetSummary(summary *string) {
	t.Summary = summary
}

func (t *TestConditionRationaleCommonModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TestConditionRationaleCommonModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TestConditionRationaleCommonModel to string"
	}
	return string(jsonData)
}

func (t *TestConditionRationaleCommonModel) UnmarshalJSON(data []byte) error {
	type unmarshaler TestConditionRationaleCommonModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TestConditionRationaleCommonModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TestConditionRationaleCommonModel) MarshalJSON() ([]byte, error) {
	type embed TestConditionRationaleCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
