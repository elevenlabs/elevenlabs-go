package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PodcastProjectResponseModel struct {
	Project         *ProjectResponse `json:"project" xml:"project" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PodcastProjectResponseModel) GetProject() (value *ProjectResponse) {
	if p == nil {
		return
	}
	return p.Project
}

func (p *PodcastProjectResponseModel) SetProject(project *ProjectResponse) {
	p.Project = project
}

func (p *PodcastProjectResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PodcastProjectResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PodcastProjectResponseModel to string"
	}
	return string(jsonData)
}

func (p *PodcastProjectResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PodcastProjectResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PodcastProjectResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PodcastProjectResponseModel) MarshalJSON() ([]byte, error) {
	type embed PodcastProjectResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
