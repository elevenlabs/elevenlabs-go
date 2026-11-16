package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SuggestedAudioTag struct {
	// Audio tag to use (for best performance, 1-2 words, e.g., 'happy', 'excited')
	Tag string `json:"tag" xml:"tag" required:"true" maxLength:"30" minLength:"1"`
	// Optional description of when to use this tag
	Description     *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"200"`
	extraProperties map[string]interface{}
}

func (s *SuggestedAudioTag) GetTag() (value string) {
	if s == nil {
		return
	}
	return s.Tag
}

func (s *SuggestedAudioTag) GetDescription() (value *string) {
	if s == nil {
		return
	}
	return s.Description
}

func (s *SuggestedAudioTag) SetTag(tag string) {
	s.Tag = tag
}

func (s *SuggestedAudioTag) SetDescription(description *string) {
	s.Description = description
}

func (s *SuggestedAudioTag) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SuggestedAudioTag) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SuggestedAudioTag to string"
	}
	return string(jsonData)
}

func (s *SuggestedAudioTag) UnmarshalJSON(data []byte) error {
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

func (s *SuggestedAudioTag) MarshalJSON() ([]byte, error) {
	type embed SuggestedAudioTag
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
