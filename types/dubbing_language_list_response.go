package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingLanguageListResponse struct {
	// The page of language targets for the project.
	Languages []*DubbingLanguageResponse `json:"languages" xml:"languages" required:"true"`
	// Opaque cursor to pass back as `cursor` for the next page, or null when there are no more results.
	NextCursor      *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubbingLanguageListResponse) GetLanguages() (value []*DubbingLanguageResponse) {
	if d == nil {
		return
	}
	return d.Languages
}

func (d *DubbingLanguageListResponse) GetNextCursor() (value *string) {
	if d == nil {
		return
	}
	return d.NextCursor
}

func (d *DubbingLanguageListResponse) SetLanguages(languages []*DubbingLanguageResponse) {
	d.Languages = languages
}

func (d *DubbingLanguageListResponse) SetNextCursor(nextCursor *string) {
	d.NextCursor = nextCursor
}

func (d *DubbingLanguageListResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingLanguageListResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingLanguageListResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingLanguageListResponse) UnmarshalJSON(data []byte) error {
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

func (d *DubbingLanguageListResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingLanguageListResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
