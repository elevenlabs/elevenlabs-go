package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EditProjectResponseModel struct {
	Project         *ProjectResponse `json:"project" xml:"project" required:"true"`
	extraProperties map[string]interface{}
}

func (e *EditProjectResponseModel) GetProject() (value *ProjectResponse) {
	if e == nil {
		return
	}
	return e.Project
}

func (e *EditProjectResponseModel) SetProject(project *ProjectResponse) {
	e.Project = project
}

func (e *EditProjectResponseModel) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EditProjectResponseModel) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditProjectResponseModel to string"
	}
	return string(jsonData)
}

func (e *EditProjectResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EditProjectResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EditProjectResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EditProjectResponseModel) MarshalJSON() ([]byte, error) {
	type embed EditProjectResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
