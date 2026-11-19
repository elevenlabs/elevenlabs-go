package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Settings to evaluate an agent's performance.
// Agents are evaluated against a set of criteria, with success being defined as meeting some combination of those criteria.
type EvaluationSettingsInput struct {
	// Individual criteria that the agent should be evaluated against. Each criterion is stored as an analysis item attached to the agent (see analysis_items).
	Criteria        []*PromptEvaluationCriteria `json:"criteria,omitempty" xml:"criteria,omitempty"`
	extraProperties map[string]interface{}
}

func (e *EvaluationSettingsInput) GetCriteria() (value []*PromptEvaluationCriteria) {
	if e == nil {
		return
	}
	return e.Criteria
}

func (e *EvaluationSettingsInput) SetCriteria(criteria []*PromptEvaluationCriteria) {
	e.Criteria = criteria
}

func (e *EvaluationSettingsInput) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EvaluationSettingsInput) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EvaluationSettingsInput to string"
	}
	return string(jsonData)
}

func (e *EvaluationSettingsInput) UnmarshalJSON(data []byte) error {
	type unmarshaler EvaluationSettingsInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EvaluationSettingsInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EvaluationSettingsInput) MarshalJSON() ([]byte, error) {
	type embed EvaluationSettingsInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
