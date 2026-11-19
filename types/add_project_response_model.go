package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AddProjectResponseModel struct {
	Project         *ProjectResponse `json:"project" xml:"project" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AddProjectResponseModel) GetProject() (value *ProjectResponse) {
	if a == nil {
		return
	}
	return a.Project
}

func (a *AddProjectResponseModel) SetProject(project *ProjectResponse) {
	a.Project = project
}

func (a *AddProjectResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AddProjectResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddProjectResponseModel to string"
	}
	return string(jsonData)
}

func (a *AddProjectResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddProjectResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddProjectResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AddProjectResponseModel) MarshalJSON() ([]byte, error) {
	type embed AddProjectResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
