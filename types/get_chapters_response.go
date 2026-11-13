package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetChaptersResponse struct {
	Chapters        []*ChapterResponse `json:"chapters" xml:"chapters" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetChaptersResponse) GetChapters() (value []*ChapterResponse) {
	if g == nil {
		return
	}
	return g.Chapters
}

func (g *GetChaptersResponse) SetChapters(chapters []*ChapterResponse) {
	g.Chapters = chapters
}

func (g *GetChaptersResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetChaptersResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetChaptersResponse to string"
	}
	return string(jsonData)
}

func (g *GetChaptersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetChaptersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetChaptersResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetChaptersResponse) MarshalJSON() ([]byte, error) {
	type embed GetChaptersResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
