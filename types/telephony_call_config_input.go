package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TelephonyCallConfigInput struct {
	// How long to ring the recipient before giving up, in seconds. Note that this will also be limited by the provider's own constraints.
	RingingTimeoutSecs *int `json:"ringing_timeout_secs,omitempty" xml:"ringing_timeout_secs,omitempty" min:"1" max:"999"`
	// Whether to record the call using Twilio call recording. Ignored for non-Twilio providers. Recordings are stored in your Twilio account.
	TwilioCallRecordingEnabled *bool `json:"twilio_call_recording_enabled,omitempty" xml:"twilio_call_recording_enabled,omitempty"`
	// Configuration for Twilio's carrier-level answering machine detection (AMD). Omit or set to null to disable it. Ignored for non-Twilio providers and for inbound calls. The resulting verdict is delivered as its own `answering_machine_detection` webhook event, which requires that event to be enabled on the workspace or agent webhook settings; it is not part of the conversation or the post-call webhook. Detection runs asynchronously so it never delays the start of the conversation, and the verdict can arrive at any point during the call -- with `detect_message_end`, even after it has ended. Twilio bills separately for AMD.
	TwilioMachineDetection *TwilioMachineDetectionConfig `json:"twilio_machine_detection,omitempty" xml:"twilio_machine_detection,omitempty"`
	extraProperties        map[string]interface{}
}

func (t *TelephonyCallConfigInput) GetRingingTimeoutSecs() (value *int) {
	if t == nil {
		return
	}
	return t.RingingTimeoutSecs
}

func (t *TelephonyCallConfigInput) GetTwilioCallRecordingEnabled() (value *bool) {
	if t == nil {
		return
	}
	return t.TwilioCallRecordingEnabled
}

func (t *TelephonyCallConfigInput) GetTwilioMachineDetection() (value *TwilioMachineDetectionConfig) {
	if t == nil {
		return
	}
	return t.TwilioMachineDetection
}

func (t *TelephonyCallConfigInput) SetRingingTimeoutSecs(ringingTimeoutSecs *int) {
	t.RingingTimeoutSecs = ringingTimeoutSecs
}

func (t *TelephonyCallConfigInput) SetTwilioCallRecordingEnabled(twilioCallRecordingEnabled *bool) {
	t.TwilioCallRecordingEnabled = twilioCallRecordingEnabled
}

func (t *TelephonyCallConfigInput) SetTwilioMachineDetection(twilioMachineDetection *TwilioMachineDetectionConfig) {
	t.TwilioMachineDetection = twilioMachineDetection
}

func (t *TelephonyCallConfigInput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TelephonyCallConfigInput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TelephonyCallConfigInput to string"
	}
	return string(jsonData)
}

func (t *TelephonyCallConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TelephonyCallConfigInput) MarshalJSON() ([]byte, error) {
	type embed TelephonyCallConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
