package verification

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// Audio recording of the user
	Recording io.Reader `json:"recording" xml:"recording" required:"true"`
}

func (b *BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost) GetVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.VoiceId
}

func (b *BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost) GetRecording() (value io.Reader) {
	if b == nil {
		return
	}
	return b.Recording
}

func (b *BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost) SetVoiceId(voiceId string) {
	b.VoiceId = voiceId
}

func (b *BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost) SetRecording(recording io.Reader) {
	b.Recording = recording
}

func (b *BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost to string"
	}
	return string(jsonData)
}

func (b *BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost(tmp)
	return nil
}

func (b *BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost) MarshalJSON() ([]byte, error) {
	type embed BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
