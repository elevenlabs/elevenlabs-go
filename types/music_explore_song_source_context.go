package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MusicExploreSongSourceContext struct {
	MusicExploreSongId string  `json:"music_explore_song_id" xml:"music_explore_song_id" required:"true"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	Bpm                *int    `json:"bpm,omitempty" xml:"bpm,omitempty"`
	Vocals             *string `json:"vocals,omitempty" xml:"vocals,omitempty"`
	Lyrics             *string `json:"lyrics,omitempty" xml:"lyrics,omitempty"`
	extraProperties    map[string]interface{}
}

func (m *MusicExploreSongSourceContext) GetMusicExploreSongId() (value string) {
	if m == nil {
		return
	}
	return m.MusicExploreSongId
}

func (m *MusicExploreSongSourceContext) GetTitle() (value *string) {
	if m == nil {
		return
	}
	return m.Title
}

func (m *MusicExploreSongSourceContext) GetDescription() (value *string) {
	if m == nil {
		return
	}
	return m.Description
}

func (m *MusicExploreSongSourceContext) GetBpm() (value *int) {
	if m == nil {
		return
	}
	return m.Bpm
}

func (m *MusicExploreSongSourceContext) GetVocals() (value *string) {
	if m == nil {
		return
	}
	return m.Vocals
}

func (m *MusicExploreSongSourceContext) GetLyrics() (value *string) {
	if m == nil {
		return
	}
	return m.Lyrics
}

func (m *MusicExploreSongSourceContext) SetMusicExploreSongId(musicExploreSongId string) {
	m.MusicExploreSongId = musicExploreSongId
}

func (m *MusicExploreSongSourceContext) SetTitle(title *string) {
	m.Title = title
}

func (m *MusicExploreSongSourceContext) SetDescription(description *string) {
	m.Description = description
}

func (m *MusicExploreSongSourceContext) SetBpm(bpm *int) {
	m.Bpm = bpm
}

func (m *MusicExploreSongSourceContext) SetVocals(vocals *string) {
	m.Vocals = vocals
}

func (m *MusicExploreSongSourceContext) SetLyrics(lyrics *string) {
	m.Lyrics = lyrics
}

func (m *MusicExploreSongSourceContext) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MusicExploreSongSourceContext) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MusicExploreSongSourceContext to string"
	}
	return string(jsonData)
}

func (m *MusicExploreSongSourceContext) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "source_type")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MusicExploreSongSourceContext) MarshalJSON() ([]byte, error) {
	type embed MusicExploreSongSourceContext
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
