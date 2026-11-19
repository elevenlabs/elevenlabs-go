package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProjectCreationMetaResponseModel struct {
	// The progress of the project creation.
	CreationProgress float64 `json:"creation_progress" xml:"creation_progress" required:"true"`
	// The status of the project creation action.
	Status          ProjectCreationMetaResponseModelStatus `json:"status" xml:"status" required:"true"`
	Type            ProjectCreationMetaType                `json:"type" xml:"type" required:"true"`
	extraProperties map[string]interface{}
}

func (p *ProjectCreationMetaResponseModel) GetCreationProgress() (value float64) {
	if p == nil {
		return
	}
	return p.CreationProgress
}

func (p *ProjectCreationMetaResponseModel) GetStatus() (value ProjectCreationMetaResponseModelStatus) {
	if p == nil {
		return
	}
	return p.Status
}

func (p *ProjectCreationMetaResponseModel) GetType() (value ProjectCreationMetaType) {
	if p == nil {
		return
	}
	return p.Type
}

func (p *ProjectCreationMetaResponseModel) SetCreationProgress(creationProgress float64) {
	p.CreationProgress = creationProgress
}

func (p *ProjectCreationMetaResponseModel) SetStatus(status ProjectCreationMetaResponseModelStatus) {
	p.Status = status
}

func (p *ProjectCreationMetaResponseModel) SetType(type_ ProjectCreationMetaType) {
	p.Type = type_
}

func (p *ProjectCreationMetaResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectCreationMetaResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectCreationMetaResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProjectCreationMetaResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias ProjectCreationMetaResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = ProjectCreationMetaResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProjectCreationMetaResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProjectCreationMetaResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
