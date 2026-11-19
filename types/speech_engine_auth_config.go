package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SpeechEngineAuthConfig struct {
	// If true, starting a conversation with the speech engine requires a signed URL or token from an authorized caller, rather than only its ID
	EnableAuth      *bool `json:"enable_auth,omitempty" xml:"enable_auth,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SpeechEngineAuthConfig) GetEnableAuth() (value *bool) {
	if s == nil {
		return
	}
	return s.EnableAuth
}

func (s *SpeechEngineAuthConfig) SetEnableAuth(enableAuth *bool) {
	s.EnableAuth = enableAuth
}

func (s *SpeechEngineAuthConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeechEngineAuthConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeechEngineAuthConfig to string"
	}
	return string(jsonData)
}

func (s *SpeechEngineAuthConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler SpeechEngineAuthConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SpeechEngineAuthConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeechEngineAuthConfig) MarshalJSON() ([]byte, error) {
	type embed SpeechEngineAuthConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
