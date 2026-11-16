package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingProjectListResponse struct {
	// The page of dubbing projects the caller can access.
	Projects []*DubbingProjectResponse `json:"projects" xml:"projects" required:"true"`
	// Opaque cursor to pass back as `cursor` for the next page, or null when there are no more results.
	NextCursor      *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubbingProjectListResponse) GetProjects() (value []*DubbingProjectResponse) {
	if d == nil {
		return
	}
	return d.Projects
}

func (d *DubbingProjectListResponse) GetNextCursor() (value *string) {
	if d == nil {
		return
	}
	return d.NextCursor
}

func (d *DubbingProjectListResponse) SetProjects(projects []*DubbingProjectResponse) {
	d.Projects = projects
}

func (d *DubbingProjectListResponse) SetNextCursor(nextCursor *string) {
	d.NextCursor = nextCursor
}

func (d *DubbingProjectListResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingProjectListResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingProjectListResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingProjectListResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingProjectListResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingProjectListResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
