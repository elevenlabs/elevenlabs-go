package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeWarning struct {
	MessageType     *string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	Warning         string  `json:"warning" xml:"warning" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeWarning) GetMessageType() (value *string) {
	if s == nil {
		return
	}
	return s.MessageType
}

func (s *ScribeWarning) GetWarning() (value string) {
	if s == nil {
		return
	}
	return s.Warning
}

func (s *ScribeWarning) SetMessageType(messageType *string) {
	s.MessageType = messageType
}

func (s *ScribeWarning) SetWarning(warning string) {
	s.Warning = warning
}

func (s *ScribeWarning) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeWarning) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeWarning to string"
	}
	return string(jsonData)
}

func (s *ScribeWarning) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeWarning
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeWarning(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeWarning) MarshalJSON() ([]byte, error) {
	type embed ScribeWarning
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
