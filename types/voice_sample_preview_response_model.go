package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceSamplePreviewResponseModel struct {
	// The base64 encoded audio.
	AudioBase64 string `json:"audio_base_64" xml:"audio_base_64" required:"true"`
	// The ID of the voice.
	VoiceId string `json:"voice_id" xml:"voice_id" required:"true"`
	// The ID of the sample.
	SampleId string `json:"sample_id" xml:"sample_id" required:"true"`
	// The media type of the audio.
	MediaType string `json:"media_type" xml:"media_type" required:"true"`
	// The duration of the audio in seconds.
	DurationSecs    *float64 `json:"duration_secs,omitempty" xml:"duration_secs,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VoiceSamplePreviewResponseModel) GetAudioBase64() (value string) {
	if v == nil {
		return
	}
	return v.AudioBase64
}

func (v *VoiceSamplePreviewResponseModel) GetVoiceId() (value string) {
	if v == nil {
		return
	}
	return v.VoiceId
}

func (v *VoiceSamplePreviewResponseModel) GetSampleId() (value string) {
	if v == nil {
		return
	}
	return v.SampleId
}

func (v *VoiceSamplePreviewResponseModel) GetMediaType() (value string) {
	if v == nil {
		return
	}
	return v.MediaType
}

func (v *VoiceSamplePreviewResponseModel) GetDurationSecs() (value *float64) {
	if v == nil {
		return
	}
	return v.DurationSecs
}

func (v *VoiceSamplePreviewResponseModel) SetAudioBase64(audioBase64 string) {
	v.AudioBase64 = audioBase64
}

func (v *VoiceSamplePreviewResponseModel) SetVoiceId(voiceId string) {
	v.VoiceId = voiceId
}

func (v *VoiceSamplePreviewResponseModel) SetSampleId(sampleId string) {
	v.SampleId = sampleId
}

func (v *VoiceSamplePreviewResponseModel) SetMediaType(mediaType string) {
	v.MediaType = mediaType
}

func (v *VoiceSamplePreviewResponseModel) SetDurationSecs(durationSecs *float64) {
	v.DurationSecs = durationSecs
}

func (v *VoiceSamplePreviewResponseModel) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceSamplePreviewResponseModel) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceSamplePreviewResponseModel to string"
	}
	return string(jsonData)
}

func (v *VoiceSamplePreviewResponseModel) UnmarshalJSON(data []byte) error {
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

func (v *VoiceSamplePreviewResponseModel) MarshalJSON() ([]byte, error) {
	type embed VoiceSamplePreviewResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
