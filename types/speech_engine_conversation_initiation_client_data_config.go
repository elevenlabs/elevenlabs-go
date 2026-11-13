package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SpeechEngineConversationInitiationClientDataConfig struct {
	// Whether the first message can be overridden by the client
	FirstMessage    *bool `json:"first_message,omitempty" xml:"first_message,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SpeechEngineConversationInitiationClientDataConfig) GetFirstMessage() (value *bool) {
	if s == nil {
		return
	}
	return s.FirstMessage
}

func (s *SpeechEngineConversationInitiationClientDataConfig) SetFirstMessage(firstMessage *bool) {
	s.FirstMessage = firstMessage
}

func (s *SpeechEngineConversationInitiationClientDataConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeechEngineConversationInitiationClientDataConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeechEngineConversationInitiationClientDataConfig to string"
	}
	return string(jsonData)
}

func (s *SpeechEngineConversationInitiationClientDataConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler SpeechEngineConversationInitiationClientDataConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SpeechEngineConversationInitiationClientDataConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeechEngineConversationInitiationClientDataConfig) MarshalJSON() ([]byte, error) {
	type embed SpeechEngineConversationInitiationClientDataConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
