package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PhoneVoiceConfigInput struct {
	// How long to ring the recipient before giving up, in seconds. Note that this will also be limited by the provider's own constraints.
	RingingTimeoutSecs *int `json:"ringing_timeout_secs,omitempty" xml:"ringing_timeout_secs,omitempty" min:"1" max:"999"`
	// Whether to record the call using Twilio call recording. Ignored for non-Twilio providers. Recordings are stored in your Twilio account.
	TwilioCallRecordingEnabled *bool `json:"twilio_call_recording_enabled,omitempty" xml:"twilio_call_recording_enabled,omitempty"`
	// Configuration for Twilio's carrier-level answering machine detection (AMD). Omit or set to null to disable it. Ignored for non-Twilio providers and for inbound calls. The resulting verdict is delivered as its own `answering_machine_detection` webhook event, which requires that event to be enabled on the workspace or agent webhook settings; it is not part of the conversation or the post-call webhook. Detection runs asynchronously so it never delays the start of the conversation, and the verdict can arrive at any point during the call -- with `detect_message_end`, even after it has ended. Twilio bills separately for AMD.
	TwilioMachineDetection *TwilioMachineDetectionConfig `json:"twilio_machine_detection,omitempty" xml:"twilio_machine_detection,omitempty"`
	Mode                   *string                       `json:"mode,omitempty" xml:"mode,omitempty"`
	extraProperties        map[string]interface{}
}

func (p *PhoneVoiceConfigInput) GetRingingTimeoutSecs() (value *int) {
	if p == nil {
		return
	}
	return p.RingingTimeoutSecs
}

func (p *PhoneVoiceConfigInput) GetTwilioCallRecordingEnabled() (value *bool) {
	if p == nil {
		return
	}
	return p.TwilioCallRecordingEnabled
}

func (p *PhoneVoiceConfigInput) GetTwilioMachineDetection() (value *TwilioMachineDetectionConfig) {
	if p == nil {
		return
	}
	return p.TwilioMachineDetection
}

func (p *PhoneVoiceConfigInput) GetMode() (value *string) {
	if p == nil {
		return
	}
	return p.Mode
}

func (p *PhoneVoiceConfigInput) SetRingingTimeoutSecs(ringingTimeoutSecs *int) {
	p.RingingTimeoutSecs = ringingTimeoutSecs
}

func (p *PhoneVoiceConfigInput) SetTwilioCallRecordingEnabled(twilioCallRecordingEnabled *bool) {
	p.TwilioCallRecordingEnabled = twilioCallRecordingEnabled
}

func (p *PhoneVoiceConfigInput) SetTwilioMachineDetection(twilioMachineDetection *TwilioMachineDetectionConfig) {
	p.TwilioMachineDetection = twilioMachineDetection
}

func (p *PhoneVoiceConfigInput) SetMode(mode *string) {
	p.Mode = mode
}

func (p *PhoneVoiceConfigInput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PhoneVoiceConfigInput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PhoneVoiceConfigInput to string"
	}
	return string(jsonData)
}

func (p *PhoneVoiceConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PhoneVoiceConfigInput) MarshalJSON() ([]byte, error) {
	type embed PhoneVoiceConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
