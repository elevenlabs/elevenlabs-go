package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AttachedSystemDataCollectionRef struct {
	source          string
	AnalysisItemId  SystemDataCollectionId `json:"analysis_item_id" xml:"analysis_item_id" required:"true"`
	Scope           *AnalysisScope         `json:"scope,omitempty" xml:"scope,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AttachedSystemDataCollectionRef) GetAnalysisItemId() (value SystemDataCollectionId) {
	if a == nil {
		return
	}
	return a.AnalysisItemId
}

func (a *AttachedSystemDataCollectionRef) GetScope() (value *AnalysisScope) {
	if a == nil {
		return
	}
	return a.Scope
}

func (a *AttachedSystemDataCollectionRef) SetAnalysisItemId(analysisItemId SystemDataCollectionId) {
	a.AnalysisItemId = analysisItemId
}

func (a *AttachedSystemDataCollectionRef) SetScope(scope *AnalysisScope) {
	a.Scope = scope
}

func (a *AttachedSystemDataCollectionRef) Source() string {
	return a.source
}

func (a *AttachedSystemDataCollectionRef) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AttachedSystemDataCollectionRef) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachedSystemDataCollectionRef to string"
	}
	return string(jsonData)
}

func (a *AttachedSystemDataCollectionRef) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AttachedSystemDataCollectionRef
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AttachedSystemDataCollectionRef(tmp)
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

func (a *AttachedSystemDataCollectionRef) MarshalJSON() ([]byte, error) {
	type embed AttachedSystemDataCollectionRef
	var marshaler = struct {
		embed
		Source string `json:"source"`
	}{
		embed:  embed(*a),
		Source: "system",
	}
	return json.Marshal(marshaler)
}
