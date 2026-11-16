package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetProjectsResponse struct {
	// A list of projects with their metadata.
	Projects        []*ProjectResponse `json:"projects" xml:"projects" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetProjectsResponse) GetProjects() (value []*ProjectResponse) {
	if g == nil {
		return
	}
	return g.Projects
}

func (g *GetProjectsResponse) SetProjects(projects []*ProjectResponse) {
	g.Projects = projects
}

func (g *GetProjectsResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetProjectsResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetProjectsResponse to string"
	}
	return string(jsonData)
}

func (g *GetProjectsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetProjectsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetProjectsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetProjectsResponse) MarshalJSON() ([]byte, error) {
	type embed GetProjectsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
