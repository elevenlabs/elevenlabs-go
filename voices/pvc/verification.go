package pvc

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// Verification documents
	Files []io.Reader `json:"files" xml:"files" required:"true"`
	// Extra text to be used in the manual verification process.
	ExtraText *string `json:"extra_text,omitempty" xml:"extra_text,omitempty"`
}

func (b *BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost) GetVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.VoiceId
}

func (b *BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost) GetFiles() (value []io.Reader) {
	if b == nil {
		return
	}
	return b.Files
}

func (b *BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost) GetExtraText() (value *string) {
	if b == nil {
		return
	}
	return b.ExtraText
}

func (b *BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost) SetVoiceId(voiceId string) {
	b.VoiceId = voiceId
}

func (b *BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost) SetFiles(files []io.Reader) {
	b.Files = files
}

func (b *BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost) SetExtraText(extraText *string) {
	b.ExtraText = extraText
}

func (b *BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost to string"
	}
	return string(jsonData)
}

func (b *BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost) MarshalJSON() ([]byte, error) {
	type embed BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
