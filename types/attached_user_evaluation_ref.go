package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AttachedUserEvaluationRef struct {
	source string
	// Id of the referenced user evaluation item.
	AnalysisItemId string `json:"analysis_item_id" xml:"analysis_item_id" required:"true"`
	// Primary item version whose result feeds scoring. 'latest' tracks the item's latest published version.
	VersionId *string `json:"version_id,omitempty" xml:"version_id,omitempty"`
	// Extra item versions to also run for comparison (A/B). These are executed and stored but excluded from scoring; the primary version_id is the one that scores.
	AdditionalVersionIds []string       `json:"additional_version_ids,omitempty" xml:"additional_version_ids,omitempty"`
	Scope                *AnalysisScope `json:"scope,omitempty" xml:"scope,omitempty"`
	// Optional relative weight for aggregate scoring.
	Weight          *float64 `json:"weight,omitempty" xml:"weight,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AttachedUserEvaluationRef) GetAnalysisItemId() (value string) {
	if a == nil {
		return
	}
	return a.AnalysisItemId
}

func (a *AttachedUserEvaluationRef) GetVersionId() (value *string) {
	if a == nil {
		return
	}
	return a.VersionId
}

func (a *AttachedUserEvaluationRef) GetAdditionalVersionIds() (value []string) {
	if a == nil {
		return
	}
	return a.AdditionalVersionIds
}

func (a *AttachedUserEvaluationRef) GetScope() (value *AnalysisScope) {
	if a == nil {
		return
	}
	return a.Scope
}

func (a *AttachedUserEvaluationRef) GetWeight() (value *float64) {
	if a == nil {
		return
	}
	return a.Weight
}

func (a *AttachedUserEvaluationRef) SetAnalysisItemId(analysisItemId string) {
	a.AnalysisItemId = analysisItemId
}

func (a *AttachedUserEvaluationRef) SetVersionId(versionId *string) {
	a.VersionId = versionId
}

func (a *AttachedUserEvaluationRef) SetAdditionalVersionIds(additionalVersionIds []string) {
	a.AdditionalVersionIds = additionalVersionIds
}

func (a *AttachedUserEvaluationRef) SetScope(scope *AnalysisScope) {
	a.Scope = scope
}

func (a *AttachedUserEvaluationRef) SetWeight(weight *float64) {
	a.Weight = weight
}

func (a *AttachedUserEvaluationRef) Source() string {
	return a.source
}

func (a *AttachedUserEvaluationRef) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AttachedUserEvaluationRef) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachedUserEvaluationRef to string"
	}
	return string(jsonData)
}

func (a *AttachedUserEvaluationRef) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	var literals struct {
		Source *string `json:"source"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Source != nil && *literals.Source != "user" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "user", *literals.Source)
	}
	a.source = "user"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "source")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AttachedUserEvaluationRef) MarshalJSON() ([]byte, error) {
	type embed AttachedUserEvaluationRef
	var marshaler = struct {
		embed
		Source string `json:"source"`
	}{
		embed:  embed(*a),
		Source: "user",
	}
	return json.Marshal(marshaler)
}
