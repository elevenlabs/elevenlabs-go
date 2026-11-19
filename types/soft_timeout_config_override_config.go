package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SoftTimeoutConfigOverrideConfig struct {
	// Whether to allow overriding the message field.
	Message *bool `json:"message,omitempty" xml:"message,omitempty"`
	// Whether to allow overriding the additional_soft_timeout_messages field.
	AdditionalSoftTimeoutMessages *bool `json:"additional_soft_timeout_messages,omitempty" xml:"additional_soft_timeout_messages,omitempty"`
	extraProperties               map[string]interface{}
}

func (s *SoftTimeoutConfigOverrideConfig) GetMessage() (value *bool) {
	if s == nil {
		return
	}
	return s.Message
}

func (s *SoftTimeoutConfigOverrideConfig) GetAdditionalSoftTimeoutMessages() (value *bool) {
	if s == nil {
		return
	}
	return s.AdditionalSoftTimeoutMessages
}

func (s *SoftTimeoutConfigOverrideConfig) SetMessage(message *bool) {
	s.Message = message
}

func (s *SoftTimeoutConfigOverrideConfig) SetAdditionalSoftTimeoutMessages(additionalSoftTimeoutMessages *bool) {
	s.AdditionalSoftTimeoutMessages = additionalSoftTimeoutMessages
}

func (s *SoftTimeoutConfigOverrideConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SoftTimeoutConfigOverrideConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SoftTimeoutConfigOverrideConfig to string"
	}
	return string(jsonData)
}

func (s *SoftTimeoutConfigOverrideConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler SoftTimeoutConfigOverrideConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SoftTimeoutConfigOverrideConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SoftTimeoutConfigOverrideConfig) MarshalJSON() ([]byte, error) {
	type embed SoftTimeoutConfigOverrideConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
