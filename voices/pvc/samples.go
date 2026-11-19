package pvc

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// Audio files used to create the voice.
	Files []io.Reader `json:"files" xml:"files" required:"true"`
	// If set will remove background noise for voice samples using our audio isolation model. If the samples do not include background noise, it can make the quality worse.
	RemoveBackgroundNoise *bool `json:"remove_background_noise,omitempty" xml:"remove_background_noise,omitempty"`
}

func (b *BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost) GetVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.VoiceId
}

func (b *BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost) GetFiles() (value []io.Reader) {
	if b == nil {
		return
	}
	return b.Files
}

func (b *BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost) GetRemoveBackgroundNoise() (value *bool) {
	if b == nil {
		return
	}
	return b.RemoveBackgroundNoise
}

func (b *BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost) SetVoiceId(voiceId string) {
	b.VoiceId = voiceId
}

func (b *BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost) SetFiles(files []io.Reader) {
	b.Files = files
}

func (b *BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost) SetRemoveBackgroundNoise(removeBackgroundNoise *bool) {
	b.RemoveBackgroundNoise = removeBackgroundNoise
}

func (b *BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost to string"
	}
	return string(jsonData)
}

func (b *BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost(tmp)
	return nil
}

func (b *BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost) MarshalJSON() ([]byte, error) {
	type embed BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost struct {
	// Sample ID to be used
	SampleId string `json:"-" url:"-"`
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// If set will remove background noise for voice samples using our audio isolation model. If the samples do not include background noise, it can make the quality worse.
	RemoveBackgroundNoise *bool `json:"remove_background_noise,omitempty" xml:"remove_background_noise,omitempty"`
	// Speaker IDs to be used for PVC training. Make sure you send all the speaker IDs you want to use for PVC training in one request because the last request will override the previous ones.
	SelectedSpeakerIds []string `json:"selected_speaker_ids,omitempty" xml:"selected_speaker_ids,omitempty"`
	// The start time of the audio to be used for PVC training. Time should be in milliseconds
	TrimStartTime *int `json:"trim_start_time,omitempty" xml:"trim_start_time,omitempty"`
	// The end time of the audio to be used for PVC training. Time should be in milliseconds
	TrimEndTime *int `json:"trim_end_time,omitempty" xml:"trim_end_time,omitempty"`
	// The name of the audio file to be used for PVC training.
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) GetSampleId() (value string) {
	if b == nil {
		return
	}
	return b.SampleId
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) GetVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.VoiceId
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) GetRemoveBackgroundNoise() (value *bool) {
	if b == nil {
		return
	}
	return b.RemoveBackgroundNoise
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) GetSelectedSpeakerIds() (value []string) {
	if b == nil {
		return
	}
	return b.SelectedSpeakerIds
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) GetTrimStartTime() (value *int) {
	if b == nil {
		return
	}
	return b.TrimStartTime
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) GetTrimEndTime() (value *int) {
	if b == nil {
		return
	}
	return b.TrimEndTime
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) GetFileName() (value *string) {
	if b == nil {
		return
	}
	return b.FileName
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) SetSampleId(sampleId string) {
	b.SampleId = sampleId
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) SetVoiceId(voiceId string) {
	b.VoiceId = voiceId
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) SetRemoveBackgroundNoise(removeBackgroundNoise *bool) {
	b.RemoveBackgroundNoise = removeBackgroundNoise
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) SetSelectedSpeakerIds(selectedSpeakerIds []string) {
	b.SelectedSpeakerIds = selectedSpeakerIds
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) SetTrimStartTime(trimStartTime *int) {
	b.TrimStartTime = trimStartTime
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) SetTrimEndTime(trimEndTime *int) {
	b.TrimEndTime = trimEndTime
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) SetFileName(fileName *string) {
	b.FileName = fileName
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost to string"
	}
	return string(jsonData)
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost) MarshalJSON() ([]byte, error) {
	type embed BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
