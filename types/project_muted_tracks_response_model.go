package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ProjectMutedTracksResponseModel struct {
	// List of chapter IDs that have muted tracks.
	ChapterIds      []string `json:"chapter_ids" xml:"chapter_ids" required:"true"`
	extraProperties map[string]interface{}
}

func (p *ProjectMutedTracksResponseModel) GetChapterIds() (value []string) {
	if p == nil {
		return
	}
	return p.ChapterIds
}

func (p *ProjectMutedTracksResponseModel) SetChapterIds(chapterIds []string) {
	p.ChapterIds = chapterIds
}

func (p *ProjectMutedTracksResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectMutedTracksResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectMutedTracksResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProjectMutedTracksResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler ProjectMutedTracksResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ProjectMutedTracksResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProjectMutedTracksResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProjectMutedTracksResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
