package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SoftTimeoutConfigOverride struct {
	// Message to show when the first soft timeout is reached while waiting for LLM response. Supports dynamic variables (e.g., {{system__time}}, {{custom_variable}}).
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Extra static filler messages for subsequent soft timeouts in the same LLM generation. The first timeout uses `message`. If fewer messages are configured than `max_soft_timeouts_per_generation`, the last configured message is repeated; otherwise a built-in filler is used.
	AdditionalSoftTimeoutMessages []string `json:"additional_soft_timeout_messages,omitempty" xml:"additional_soft_timeout_messages,omitempty"`
	extraProperties               map[string]interface{}
}

func (s *SoftTimeoutConfigOverride) GetMessage() (value *string) {
	if s == nil {
		return
	}
	return s.Message
}

func (s *SoftTimeoutConfigOverride) GetAdditionalSoftTimeoutMessages() (value []string) {
	if s == nil {
		return
	}
	return s.AdditionalSoftTimeoutMessages
}

func (s *SoftTimeoutConfigOverride) SetMessage(message *string) {
	s.Message = message
}

func (s *SoftTimeoutConfigOverride) SetAdditionalSoftTimeoutMessages(additionalSoftTimeoutMessages []string) {
	s.AdditionalSoftTimeoutMessages = additionalSoftTimeoutMessages
}

func (s *SoftTimeoutConfigOverride) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SoftTimeoutConfigOverride) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SoftTimeoutConfigOverride to string"
	}
	return string(jsonData)
}

func (s *SoftTimeoutConfigOverride) UnmarshalJSON(data []byte) error {
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

func (s *SoftTimeoutConfigOverride) MarshalJSON() ([]byte, error) {
	type embed SoftTimeoutConfigOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
