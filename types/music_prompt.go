package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Composition plan for the `music_v1` model. Using this field with any other model will result in an error.
type MusicPrompt struct {
	// The styles and musical directions that should be present in the entire song. Use English language for best result.
	PositiveGlobalStyles []string `json:"positive_global_styles" xml:"positive_global_styles" required:"true" maxItems:"50"`
	// The styles and musical directions that should not be present in the entire song. Use English language for best result.
	NegativeGlobalStyles []string `json:"negative_global_styles" xml:"negative_global_styles" required:"true" maxItems:"50"`
	// The sections of the song.
	Sections        []*SongSection `json:"sections" xml:"sections" required:"true" maxItems:"30"`
	extraProperties map[string]interface{}
}

func (m *MusicPrompt) GetPositiveGlobalStyles() (value []string) {
	if m == nil {
		return
	}
	return m.PositiveGlobalStyles
}

func (m *MusicPrompt) GetNegativeGlobalStyles() (value []string) {
	if m == nil {
		return
	}
	return m.NegativeGlobalStyles
}

func (m *MusicPrompt) GetSections() (value []*SongSection) {
	if m == nil {
		return
	}
	return m.Sections
}

func (m *MusicPrompt) SetPositiveGlobalStyles(positiveGlobalStyles []string) {
	m.PositiveGlobalStyles = positiveGlobalStyles
}

func (m *MusicPrompt) SetNegativeGlobalStyles(negativeGlobalStyles []string) {
	m.NegativeGlobalStyles = negativeGlobalStyles
}

func (m *MusicPrompt) SetSections(sections []*SongSection) {
	m.Sections = sections
}

func (m *MusicPrompt) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MusicPrompt) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MusicPrompt to string"
	}
	return string(jsonData)
}

func (m *MusicPrompt) UnmarshalJSON(data []byte) error {
	type unmarshaler MusicPrompt
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MusicPrompt(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MusicPrompt) MarshalJSON() ([]byte, error) {
	type embed MusicPrompt
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
