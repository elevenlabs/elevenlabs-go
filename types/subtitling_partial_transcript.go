package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitlingPartialTranscript struct {
	messageType     string
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingPartialTranscript) GetText() (value string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SubtitlingPartialTranscript) SetText(text string) {
	s.Text = text
}

func (s *SubtitlingPartialTranscript) MessageType() string {
	return s.messageType
}

func (s *SubtitlingPartialTranscript) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingPartialTranscript) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingPartialTranscript to string"
	}
	return string(jsonData)
}

func (s *SubtitlingPartialTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingPartialTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingPartialTranscript(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "partial_transcript" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "partial_transcript", *literals.MessageType)
	}
	s.messageType = "partial_transcript"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingPartialTranscript) MarshalJSON() ([]byte, error) {
	type embed SubtitlingPartialTranscript
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "partial_transcript",
	}
	return json.Marshal(marshaler)
}
