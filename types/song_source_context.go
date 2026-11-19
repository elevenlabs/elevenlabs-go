package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SongSourceContext struct {
	SongId             string         `json:"song_id" xml:"song_id" required:"true"`
	ChatId             *string        `json:"chat_id,omitempty" xml:"chat_id,omitempty"`
	Title              *string        `json:"title,omitempty" xml:"title,omitempty"`
	Description        *string        `json:"description,omitempty" xml:"description,omitempty"`
	Genres             []string       `json:"genres,omitempty" xml:"genres,omitempty"`
	Languages          []string       `json:"languages,omitempty" xml:"languages,omitempty"`
	IsExplicit         *bool          `json:"is_explicit,omitempty" xml:"is_explicit,omitempty"`
	Bpm                *int           `json:"bpm,omitempty" xml:"bpm,omitempty"`
	GenerationSettings map[string]any `json:"generation_settings,omitempty"`
	extraProperties    map[string]interface{}
}

func (s *SongSourceContext) GetSongId() (value string) {
	if s == nil {
		return
	}
	return s.SongId
}

func (s *SongSourceContext) GetChatId() (value *string) {
	if s == nil {
		return
	}
	return s.ChatId
}

func (s *SongSourceContext) GetTitle() (value *string) {
	if s == nil {
		return
	}
	return s.Title
}

func (s *SongSourceContext) GetDescription() (value *string) {
	if s == nil {
		return
	}
	return s.Description
}

func (s *SongSourceContext) GetGenres() (value []string) {
	if s == nil {
		return
	}
	return s.Genres
}

func (s *SongSourceContext) GetLanguages() (value []string) {
	if s == nil {
		return
	}
	return s.Languages
}

func (s *SongSourceContext) GetIsExplicit() (value *bool) {
	if s == nil {
		return
	}
	return s.IsExplicit
}

func (s *SongSourceContext) GetBpm() (value *int) {
	if s == nil {
		return
	}
	return s.Bpm
}

func (s *SongSourceContext) GetGenerationSettings() (value map[string]any) {
	if s == nil {
		return
	}
	return s.GenerationSettings
}

func (s *SongSourceContext) SetSongId(songId string) {
	s.SongId = songId
}

func (s *SongSourceContext) SetChatId(chatId *string) {
	s.ChatId = chatId
}

func (s *SongSourceContext) SetTitle(title *string) {
	s.Title = title
}

func (s *SongSourceContext) SetDescription(description *string) {
	s.Description = description
}

func (s *SongSourceContext) SetGenres(genres []string) {
	s.Genres = genres
}

func (s *SongSourceContext) SetLanguages(languages []string) {
	s.Languages = languages
}

func (s *SongSourceContext) SetIsExplicit(isExplicit *bool) {
	s.IsExplicit = isExplicit
}

func (s *SongSourceContext) SetBpm(bpm *int) {
	s.Bpm = bpm
}

func (s *SongSourceContext) SetGenerationSettings(generationSettings map[string]any) {
	s.GenerationSettings = generationSettings
}

func (s *SongSourceContext) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SongSourceContext) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SongSourceContext to string"
	}
	return string(jsonData)
}

func (s *SongSourceContext) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "source_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SongSourceContext) MarshalJSON() ([]byte, error) {
	type embed SongSourceContext
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
