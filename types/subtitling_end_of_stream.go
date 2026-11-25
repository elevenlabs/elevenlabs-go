package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitlingEndOfStream struct {
	messageType     string
	extraProperties map[string]interface{}
}

func (s *SubtitlingEndOfStream) MessageType() string {
	return s.messageType
}

func (s *SubtitlingEndOfStream) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingEndOfStream) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingEndOfStream to string"
	}
	return string(jsonData)
}

func (s *SubtitlingEndOfStream) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingEndOfStream
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingEndOfStream(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "end_of_stream" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "end_of_stream", *literals.MessageType)
	}
	s.messageType = "end_of_stream"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingEndOfStream) MarshalJSON() ([]byte, error) {
	type embed SubtitlingEndOfStream
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "end_of_stream",
	}
	return json.Marshal(marshaler)
}
