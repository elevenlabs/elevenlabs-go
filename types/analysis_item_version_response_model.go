package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AnalysisItemVersionResponseModel struct {
	Id                string                                      `json:"id" xml:"id" required:"true"`
	VersionNumber     int                                         `json:"version_number" xml:"version_number" required:"true"`
	CreatedAtUnixSecs int                                         `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	Definition        *AnalysisItemVersionResponseModelDefinition `json:"definition" xml:"definition" required:"true"`
	extraProperties   map[string]interface{}
}

func (a *AnalysisItemVersionResponseModel) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AnalysisItemVersionResponseModel) GetVersionNumber() (value int) {
	if a == nil {
		return
	}
	return a.VersionNumber
}

func (a *AnalysisItemVersionResponseModel) GetCreatedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAtUnixSecs
}

func (a *AnalysisItemVersionResponseModel) GetDefinition() (value *AnalysisItemVersionResponseModelDefinition) {
	if a == nil {
		return
	}
	return a.Definition
}

func (a *AnalysisItemVersionResponseModel) SetId(id string) {
	a.Id = id
}

func (a *AnalysisItemVersionResponseModel) SetVersionNumber(versionNumber int) {
	a.VersionNumber = versionNumber
}

func (a *AnalysisItemVersionResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	a.CreatedAtUnixSecs = createdAtUnixSecs
}

func (a *AnalysisItemVersionResponseModel) SetDefinition(definition *AnalysisItemVersionResponseModelDefinition) {
	a.Definition = definition
}

func (a *AnalysisItemVersionResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnalysisItemVersionResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnalysisItemVersionResponseModel to string"
	}
	return string(jsonData)
}

func (a *AnalysisItemVersionResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AnalysisItemVersionResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AnalysisItemVersionResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AnalysisItemVersionResponseModel) MarshalJSON() ([]byte, error) {
	type embed AnalysisItemVersionResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
