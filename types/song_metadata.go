package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SongMetadata struct {
	// The title of the song
	Title *string `json:"title" xml:"title" required:"true"`
	// The description of the song
	Description *string `json:"description" xml:"description" required:"true"`
	// The genres of the song
	Genres []string `json:"genres" xml:"genres" required:"true"`
	// The languages of the song
	Languages []string `json:"languages" xml:"languages" required:"true"`
	// Whether the song is explicit
	IsExplicit      *bool `json:"is_explicit" xml:"is_explicit" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SongMetadata) GetTitle() (value *string) {
	if s == nil {
		return
	}
	return s.Title
}

func (s *SongMetadata) GetDescription() (value *string) {
	if s == nil {
		return
	}
	return s.Description
}

func (s *SongMetadata) GetGenres() (value []string) {
	if s == nil {
		return
	}
	return s.Genres
}

func (s *SongMetadata) GetLanguages() (value []string) {
	if s == nil {
		return
	}
	return s.Languages
}

func (s *SongMetadata) GetIsExplicit() (value *bool) {
	if s == nil {
		return
	}
	return s.IsExplicit
}

func (s *SongMetadata) SetTitle(title *string) {
	s.Title = title
}

func (s *SongMetadata) SetDescription(description *string) {
	s.Description = description
}

func (s *SongMetadata) SetGenres(genres []string) {
	s.Genres = genres
}

func (s *SongMetadata) SetLanguages(languages []string) {
	s.Languages = languages
}

func (s *SongMetadata) SetIsExplicit(isExplicit *bool) {
	s.IsExplicit = isExplicit
}

func (s *SongMetadata) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SongMetadata) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SongMetadata to string"
	}
	return string(jsonData)
}

func (s *SongMetadata) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SongMetadata) MarshalJSON() ([]byte, error) {
	type embed SongMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
