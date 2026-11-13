package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoicePreviewResponseModel struct {
	// The base64 encoded audio of the preview.
	AudioBase64 string `json:"audio_base_64" xml:"audio_base_64" required:"true"`
	// The ID of the generated voice. Use it to create a voice from the preview.
	GeneratedVoiceId string `json:"generated_voice_id" xml:"generated_voice_id" required:"true"`
	// The media type of the preview.
	MediaType string `json:"media_type" xml:"media_type" required:"true"`
	// The duration of the preview in seconds.
	DurationSecs float64 `json:"duration_secs" xml:"duration_secs" required:"true"`
	// The language of the preview.
	Language        *string `json:"language" xml:"language" required:"true"`
	extraProperties map[string]interface{}
}

func (v *VoicePreviewResponseModel) GetAudioBase64() (value string) {
	if v == nil {
		return
	}
	return v.AudioBase64
}

func (v *VoicePreviewResponseModel) GetGeneratedVoiceId() (value string) {
	if v == nil {
		return
	}
	return v.GeneratedVoiceId
}

func (v *VoicePreviewResponseModel) GetMediaType() (value string) {
	if v == nil {
		return
	}
	return v.MediaType
}

func (v *VoicePreviewResponseModel) GetDurationSecs() (value float64) {
	if v == nil {
		return
	}
	return v.DurationSecs
}

func (v *VoicePreviewResponseModel) GetLanguage() (value *string) {
	if v == nil {
		return
	}
	return v.Language
}

func (v *VoicePreviewResponseModel) SetAudioBase64(audioBase64 string) {
	v.AudioBase64 = audioBase64
}

func (v *VoicePreviewResponseModel) SetGeneratedVoiceId(generatedVoiceId string) {
	v.GeneratedVoiceId = generatedVoiceId
}

func (v *VoicePreviewResponseModel) SetMediaType(mediaType string) {
	v.MediaType = mediaType
}

func (v *VoicePreviewResponseModel) SetDurationSecs(durationSecs float64) {
	v.DurationSecs = durationSecs
}

func (v *VoicePreviewResponseModel) SetLanguage(language *string) {
	v.Language = language
}

func (v *VoicePreviewResponseModel) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoicePreviewResponseModel) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoicePreviewResponseModel to string"
	}
	return string(jsonData)
}

func (v *VoicePreviewResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoicePreviewResponseModel) MarshalJSON() ([]byte, error) {
	type embed VoicePreviewResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
