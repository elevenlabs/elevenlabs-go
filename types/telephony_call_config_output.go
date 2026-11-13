package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TelephonyCallConfigOutput struct {
	// How long to ring the recipient before giving up, in seconds. Note that this will also be limited by the provider's own constraints.
	RingingTimeoutSecs *int `json:"ringing_timeout_secs,omitempty" xml:"ringing_timeout_secs,omitempty" min:"1" max:"999"`
	// Whether to record the call using Twilio call recording. Ignored for non-Twilio providers. Recordings are stored in your Twilio account.
	TwilioCallRecordingEnabled *bool `json:"twilio_call_recording_enabled,omitempty" xml:"twilio_call_recording_enabled,omitempty"`
	// Configuration for Twilio's carrier-level answering machine detection (AMD). Omit or set to null to disable it. Ignored for non-Twilio providers and for inbound calls. The resulting verdict is delivered as its own `answering_machine_detection` webhook event, which requires that event to be enabled on the workspace or agent webhook settings; it is not part of the conversation or the post-call webhook. Detection runs asynchronously so it never delays the start of the conversation, and the verdict can arrive at any point during the call -- with `detect_message_end`, even after it has ended. Twilio bills separately for AMD.
	TwilioMachineDetection *TwilioMachineDetectionConfig `json:"twilio_machine_detection,omitempty" xml:"twilio_machine_detection,omitempty"`
	extraProperties        map[string]interface{}
}

func (t *TelephonyCallConfigOutput) GetRingingTimeoutSecs() (value *int) {
	if t == nil {
		return
	}
	return t.RingingTimeoutSecs
}

func (t *TelephonyCallConfigOutput) GetTwilioCallRecordingEnabled() (value *bool) {
	if t == nil {
		return
	}
	return t.TwilioCallRecordingEnabled
}

func (t *TelephonyCallConfigOutput) GetTwilioMachineDetection() (value *TwilioMachineDetectionConfig) {
	if t == nil {
		return
	}
	return t.TwilioMachineDetection
}

func (t *TelephonyCallConfigOutput) SetRingingTimeoutSecs(ringingTimeoutSecs *int) {
	t.RingingTimeoutSecs = ringingTimeoutSecs
}

func (t *TelephonyCallConfigOutput) SetTwilioCallRecordingEnabled(twilioCallRecordingEnabled *bool) {
	t.TwilioCallRecordingEnabled = twilioCallRecordingEnabled
}

func (t *TelephonyCallConfigOutput) SetTwilioMachineDetection(twilioMachineDetection *TwilioMachineDetectionConfig) {
	t.TwilioMachineDetection = twilioMachineDetection
}

func (t *TelephonyCallConfigOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TelephonyCallConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TelephonyCallConfigOutput to string"
	}
	return string(jsonData)
}

func (t *TelephonyCallConfigOutput) UnmarshalJSON(data []byte) error {
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

func (t *TelephonyCallConfigOutput) MarshalJSON() ([]byte, error) {
	type embed TelephonyCallConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
