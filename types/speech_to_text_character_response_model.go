package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeechToTextCharacterResponseModel struct {
	// The character that was transcribed.
	Text string `json:"text" xml:"text" required:"true"`
	// The start time of the character in seconds.
	Start *float64 `json:"start,omitempty" xml:"start,omitempty"`
	// The end time of the character in seconds.
	End             *float64 `json:"end,omitempty" xml:"end,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SpeechToTextCharacterResponseModel) GetText() (value string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SpeechToTextCharacterResponseModel) GetStart() (value *float64) {
	if s == nil {
		return
	}
	return s.Start
}

func (s *SpeechToTextCharacterResponseModel) GetEnd() (value *float64) {
	if s == nil {
		return
	}
	return s.End
}

func (s *SpeechToTextCharacterResponseModel) SetText(text string) {
	s.Text = text
}

func (s *SpeechToTextCharacterResponseModel) SetStart(start *float64) {
	s.Start = start
}

func (s *SpeechToTextCharacterResponseModel) SetEnd(end *float64) {
	s.End = end
}

func (s *SpeechToTextCharacterResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeechToTextCharacterResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeechToTextCharacterResponseModel to string"
	}
	return string(jsonData)
}

func (s *SpeechToTextCharacterResponseModel) UnmarshalJSON(data []byte) error {
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

func (s *SpeechToTextCharacterResponseModel) MarshalJSON() ([]byte, error) {
	type embed SpeechToTextCharacterResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
