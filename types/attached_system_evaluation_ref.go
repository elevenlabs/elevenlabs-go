package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AttachedSystemEvaluationRef struct {
	source         string
	AnalysisItemId SystemEvaluationId `json:"analysis_item_id" xml:"analysis_item_id" required:"true"`
	Scope          *AnalysisScope     `json:"scope,omitempty" xml:"scope,omitempty"`
	// Optional relative weight for aggregate scoring.
	Weight          *float64 `json:"weight,omitempty" xml:"weight,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AttachedSystemEvaluationRef) GetAnalysisItemId() (value SystemEvaluationId) {
	if a == nil {
		return
	}
	return a.AnalysisItemId
}

func (a *AttachedSystemEvaluationRef) GetScope() (value *AnalysisScope) {
	if a == nil {
		return
	}
	return a.Scope
}

func (a *AttachedSystemEvaluationRef) GetWeight() (value *float64) {
	if a == nil {
		return
	}
	return a.Weight
}

func (a *AttachedSystemEvaluationRef) SetAnalysisItemId(analysisItemId SystemEvaluationId) {
	a.AnalysisItemId = analysisItemId
}

func (a *AttachedSystemEvaluationRef) SetScope(scope *AnalysisScope) {
	a.Scope = scope
}

func (a *AttachedSystemEvaluationRef) SetWeight(weight *float64) {
	a.Weight = weight
}

func (a *AttachedSystemEvaluationRef) Source() string {
	return a.source
}

func (a *AttachedSystemEvaluationRef) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AttachedSystemEvaluationRef) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachedSystemEvaluationRef to string"
	}
	return string(jsonData)
}

func (a *AttachedSystemEvaluationRef) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	var literals struct {
		Source *string `json:"source"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Source != nil && *literals.Source != "system" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "system", *literals.Source)
	}
	a.source = "system"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "source")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AttachedSystemEvaluationRef) MarshalJSON() ([]byte, error) {
	type embed AttachedSystemEvaluationRef
	var marshaler = struct {
		embed
		Source string `json:"source"`
	}{
		embed:  embed(*a),
		Source: "system",
	}
	return json.Marshal(marshaler)
}
