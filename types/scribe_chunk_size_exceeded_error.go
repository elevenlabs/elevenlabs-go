package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeChunkSizeExceededError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeChunkSizeExceededError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeChunkSizeExceededError) SetError(error string) {
	s.Error = error
}

func (s *ScribeChunkSizeExceededError) MessageType() string {
	return s.messageType
}

func (s *ScribeChunkSizeExceededError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeChunkSizeExceededError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeChunkSizeExceededError to string"
	}
	return string(jsonData)
}

func (s *ScribeChunkSizeExceededError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeChunkSizeExceededError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeChunkSizeExceededError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "chunk_size_exceeded" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "chunk_size_exceeded", *literals.MessageType)
	}
	s.messageType = "chunk_size_exceeded"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeChunkSizeExceededError) MarshalJSON() ([]byte, error) {
	type embed ScribeChunkSizeExceededError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "chunk_size_exceeded",
	}
	return json.Marshal(marshaler)
}
