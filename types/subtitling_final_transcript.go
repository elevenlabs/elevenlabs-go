package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitlingFinalTranscript struct {
	messageType     string
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingFinalTranscript) GetText() (value string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SubtitlingFinalTranscript) SetText(text string) {
	s.Text = text
}

func (s *SubtitlingFinalTranscript) MessageType() string {
	return s.messageType
}

func (s *SubtitlingFinalTranscript) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingFinalTranscript) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingFinalTranscript to string"
	}
	return string(jsonData)
}

func (s *SubtitlingFinalTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingFinalTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingFinalTranscript(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "final_transcript" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "final_transcript", *literals.MessageType)
	}
	s.messageType = "final_transcript"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingFinalTranscript) MarshalJSON() ([]byte, error) {
	type embed SubtitlingFinalTranscript
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "final_transcript",
	}
	return json.Marshal(marshaler)
}
