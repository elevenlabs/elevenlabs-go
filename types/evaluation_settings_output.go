package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Settings to evaluate an agent's performance.
// Agents are evaluated against a set of criteria, with success being defined as meeting some combination of those criteria.
type EvaluationSettingsOutput struct {
	// Individual criteria that the agent should be evaluated against. Each criterion is stored as an analysis item attached to the agent (see analysis_items).
	Criteria        []*PromptEvaluationCriteria `json:"criteria,omitempty" xml:"criteria,omitempty"`
	extraProperties map[string]interface{}
}

func (e *EvaluationSettingsOutput) GetCriteria() (value []*PromptEvaluationCriteria) {
	if e == nil {
		return
	}
	return e.Criteria
}

func (e *EvaluationSettingsOutput) SetCriteria(criteria []*PromptEvaluationCriteria) {
	e.Criteria = criteria
}

func (e *EvaluationSettingsOutput) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EvaluationSettingsOutput) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EvaluationSettingsOutput to string"
	}
	return string(jsonData)
}

func (e *EvaluationSettingsOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler EvaluationSettingsOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EvaluationSettingsOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EvaluationSettingsOutput) MarshalJSON() ([]byte, error) {
	type embed EvaluationSettingsOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
