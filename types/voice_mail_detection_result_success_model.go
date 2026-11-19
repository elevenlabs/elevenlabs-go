package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceMailDetectionResultSuccessModel struct {
	Status           *string `json:"status,omitempty" xml:"status,omitempty"`
	VoicemailMessage *string `json:"voicemail_message,omitempty" xml:"voicemail_message,omitempty"`
	Reason           *string `json:"reason,omitempty" xml:"reason,omitempty"`
	extraProperties  map[string]interface{}
}

func (v *VoiceMailDetectionResultSuccessModel) GetStatus() (value *string) {
	if v == nil {
		return
	}
	return v.Status
}

func (v *VoiceMailDetectionResultSuccessModel) GetVoicemailMessage() (value *string) {
	if v == nil {
		return
	}
	return v.VoicemailMessage
}

func (v *VoiceMailDetectionResultSuccessModel) GetReason() (value *string) {
	if v == nil {
		return
	}
	return v.Reason
}

func (v *VoiceMailDetectionResultSuccessModel) SetStatus(status *string) {
	v.Status = status
}

func (v *VoiceMailDetectionResultSuccessModel) SetVoicemailMessage(voicemailMessage *string) {
	v.VoicemailMessage = voicemailMessage
}

func (v *VoiceMailDetectionResultSuccessModel) SetReason(reason *string) {
	v.Reason = reason
}

func (v *VoiceMailDetectionResultSuccessModel) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceMailDetectionResultSuccessModel) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceMailDetectionResultSuccessModel to string"
	}
	return string(jsonData)
}

func (v *VoiceMailDetectionResultSuccessModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v, "result_type")
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoiceMailDetectionResultSuccessModel) MarshalJSON() ([]byte, error) {
	type embed VoiceMailDetectionResultSuccessModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
