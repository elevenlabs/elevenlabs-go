package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AttachedUserDataCollectionRef struct {
	source string
	// Id of the referenced user data-collection item.
	AnalysisItemId string `json:"analysis_item_id" xml:"analysis_item_id" required:"true"`
	// Pinned item version. 'latest' tracks the item's latest published version.
	VersionId       *string        `json:"version_id,omitempty" xml:"version_id,omitempty"`
	Scope           *AnalysisScope `json:"scope,omitempty" xml:"scope,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AttachedUserDataCollectionRef) GetAnalysisItemId() (value string) {
	if a == nil {
		return
	}
	return a.AnalysisItemId
}

func (a *AttachedUserDataCollectionRef) GetVersionId() (value *string) {
	if a == nil {
		return
	}
	return a.VersionId
}

func (a *AttachedUserDataCollectionRef) GetScope() (value *AnalysisScope) {
	if a == nil {
		return
	}
	return a.Scope
}

func (a *AttachedUserDataCollectionRef) SetAnalysisItemId(analysisItemId string) {
	a.AnalysisItemId = analysisItemId
}

func (a *AttachedUserDataCollectionRef) SetVersionId(versionId *string) {
	a.VersionId = versionId
}

func (a *AttachedUserDataCollectionRef) SetScope(scope *AnalysisScope) {
	a.Scope = scope
}

func (a *AttachedUserDataCollectionRef) Source() string {
	return a.source
}

func (a *AttachedUserDataCollectionRef) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AttachedUserDataCollectionRef) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachedUserDataCollectionRef to string"
	}
	return string(jsonData)
}

func (a *AttachedUserDataCollectionRef) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AttachedUserDataCollectionRef
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AttachedUserDataCollectionRef(tmp)
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

func (a *AttachedUserDataCollectionRef) MarshalJSON() ([]byte, error) {
	type embed AttachedUserDataCollectionRef
	var marshaler = struct {
		embed
		Source string `json:"source"`
	}{
		embed:  embed(*a),
		Source: "user",
	}
	return json.Marshal(marshaler)
}
