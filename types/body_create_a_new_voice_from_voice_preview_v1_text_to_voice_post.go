package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost struct {
	// Name to use for the created voice.
	VoiceName string `json:"voice_name" xml:"voice_name" required:"true"`
	// Description to use for the created voice.
	VoiceDescription string `json:"voice_description" xml:"voice_description" required:"true" maxLength:"1000" minLength:"20"`
	// The generated_voice_id to create; obtain it from POST /v1/text-to-voice/design, POST /v1/text-to-voice/:voice_id/remix, or the response headers when generating previews.
	GeneratedVoiceId string `json:"generated_voice_id" xml:"generated_voice_id" required:"true"`
	// Optional, metadata to add to the created voice. Defaults to None.
	Labels map[string]*string `json:"labels,omitempty"`
	// List of voice ids that the user has played but not selected. Used for RLHF.
	PlayedNotSelectedVoiceIds []string `json:"played_not_selected_voice_ids,omitempty" xml:"played_not_selected_voice_ids,omitempty"`
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) GetVoiceName() (value string) {
	if b == nil {
		return
	}
	return b.VoiceName
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) GetVoiceDescription() (value string) {
	if b == nil {
		return
	}
	return b.VoiceDescription
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) GetGeneratedVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.GeneratedVoiceId
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) GetLabels() (value map[string]*string) {
	if b == nil {
		return
	}
	return b.Labels
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) GetPlayedNotSelectedVoiceIds() (value []string) {
	if b == nil {
		return
	}
	return b.PlayedNotSelectedVoiceIds
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) SetVoiceName(voiceName string) {
	b.VoiceName = voiceName
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) SetVoiceDescription(voiceDescription string) {
	b.VoiceDescription = voiceDescription
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) SetGeneratedVoiceId(generatedVoiceId string) {
	b.GeneratedVoiceId = generatedVoiceId
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) SetLabels(labels map[string]*string) {
	b.Labels = labels
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) SetPlayedNotSelectedVoiceIds(playedNotSelectedVoiceIds []string) {
	b.PlayedNotSelectedVoiceIds = playedNotSelectedVoiceIds
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
