package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitlingStatus struct {
	messageType     string
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingStatus) GetStatus() (value string) {
	if s == nil {
		return
	}
	return s.Status
}

func (s *SubtitlingStatus) SetStatus(status string) {
	s.Status = status
}

func (s *SubtitlingStatus) MessageType() string {
	return s.messageType
}

func (s *SubtitlingStatus) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingStatus) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingStatus to string"
	}
	return string(jsonData)
}

func (s *SubtitlingStatus) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingStatus
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingStatus(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "status" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "status", *literals.MessageType)
	}
	s.messageType = "status"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingStatus) MarshalJSON() ([]byte, error) {
	type embed SubtitlingStatus
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "status",
	}
	return json.Marshal(marshaler)
}
