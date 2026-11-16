package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Allows the agent to detect when a voicemail system is encountered.
//
// This tool should be invoked by the LLM when it detects that the call has been
// answered by a voicemail system rather than a human. If a voicemail message
// is configured, it will be played; otherwise the call will end immediately.
type VoicemailDetectionToolConfig struct {
	// Optional message to leave on voicemail when detected. If not provided, the call will end immediately when voicemail is detected. Supports dynamic variables (e.g., {{system__time}}, {{system__call_duration_secs}}, {{custom_variable}}).
	VoicemailMessage *string `json:"voicemail_message,omitempty" xml:"voicemail_message,omitempty"`
	extraProperties  map[string]interface{}
}

func (v *VoicemailDetectionToolConfig) GetVoicemailMessage() (value *string) {
	if v == nil {
		return
	}
	return v.VoicemailMessage
}

func (v *VoicemailDetectionToolConfig) SetVoicemailMessage(voicemailMessage *string) {
	v.VoicemailMessage = voicemailMessage
}

func (v *VoicemailDetectionToolConfig) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoicemailDetectionToolConfig) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoicemailDetectionToolConfig to string"
	}
	return string(jsonData)
}

func (v *VoicemailDetectionToolConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v, "system_tool_type")
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoicemailDetectionToolConfig) MarshalJSON() ([]byte, error) {
	type embed VoicemailDetectionToolConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
