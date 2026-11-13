package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceSample struct {
	// The ID of the sample.
	SampleId *string `json:"sample_id,omitempty" xml:"sample_id,omitempty"`
	// The name of the sample file.
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// The MIME type of the sample file.
	MimeType *string `json:"mime_type,omitempty" xml:"mime_type,omitempty"`
	// The size of the sample file in bytes.
	SizeBytes *int `json:"size_bytes,omitempty" xml:"size_bytes,omitempty"`
	// The hash of the sample file.
	Hash                    *string                         `json:"hash,omitempty" xml:"hash,omitempty"`
	DurationSecs            *float64                        `json:"duration_secs,omitempty" xml:"duration_secs,omitempty"`
	RemoveBackgroundNoise   *bool                           `json:"remove_background_noise,omitempty" xml:"remove_background_noise,omitempty"`
	HasIsolatedAudio        *bool                           `json:"has_isolated_audio,omitempty" xml:"has_isolated_audio,omitempty"`
	HasIsolatedAudioPreview *bool                           `json:"has_isolated_audio_preview,omitempty" xml:"has_isolated_audio_preview,omitempty"`
	SpeakerSeparation       *SpeakerSeparationResponseModel `json:"speaker_separation,omitempty" xml:"speaker_separation,omitempty"`
	TrimStart               *int                            `json:"trim_start,omitempty" xml:"trim_start,omitempty"`
	TrimEnd                 *int                            `json:"trim_end,omitempty" xml:"trim_end,omitempty"`
	extraProperties         map[string]interface{}
}

func (v *VoiceSample) GetSampleId() (value *string) {
	if v == nil {
		return
	}
	return v.SampleId
}

func (v *VoiceSample) GetFileName() (value *string) {
	if v == nil {
		return
	}
	return v.FileName
}

func (v *VoiceSample) GetMimeType() (value *string) {
	if v == nil {
		return
	}
	return v.MimeType
}

func (v *VoiceSample) GetSizeBytes() (value *int) {
	if v == nil {
		return
	}
	return v.SizeBytes
}

func (v *VoiceSample) GetHash() (value *string) {
	if v == nil {
		return
	}
	return v.Hash
}

func (v *VoiceSample) GetDurationSecs() (value *float64) {
	if v == nil {
		return
	}
	return v.DurationSecs
}

func (v *VoiceSample) GetRemoveBackgroundNoise() (value *bool) {
	if v == nil {
		return
	}
	return v.RemoveBackgroundNoise
}

func (v *VoiceSample) GetHasIsolatedAudio() (value *bool) {
	if v == nil {
		return
	}
	return v.HasIsolatedAudio
}

func (v *VoiceSample) GetHasIsolatedAudioPreview() (value *bool) {
	if v == nil {
		return
	}
	return v.HasIsolatedAudioPreview
}

func (v *VoiceSample) GetSpeakerSeparation() (value *SpeakerSeparationResponseModel) {
	if v == nil {
		return
	}
	return v.SpeakerSeparation
}

func (v *VoiceSample) GetTrimStart() (value *int) {
	if v == nil {
		return
	}
	return v.TrimStart
}

func (v *VoiceSample) GetTrimEnd() (value *int) {
	if v == nil {
		return
	}
	return v.TrimEnd
}

func (v *VoiceSample) SetSampleId(sampleId *string) {
	v.SampleId = sampleId
}

func (v *VoiceSample) SetFileName(fileName *string) {
	v.FileName = fileName
}

func (v *VoiceSample) SetMimeType(mimeType *string) {
	v.MimeType = mimeType
}

func (v *VoiceSample) SetSizeBytes(sizeBytes *int) {
	v.SizeBytes = sizeBytes
}

func (v *VoiceSample) SetHash(hash *string) {
	v.Hash = hash
}

func (v *VoiceSample) SetDurationSecs(durationSecs *float64) {
	v.DurationSecs = durationSecs
}

func (v *VoiceSample) SetRemoveBackgroundNoise(removeBackgroundNoise *bool) {
	v.RemoveBackgroundNoise = removeBackgroundNoise
}

func (v *VoiceSample) SetHasIsolatedAudio(hasIsolatedAudio *bool) {
	v.HasIsolatedAudio = hasIsolatedAudio
}

func (v *VoiceSample) SetHasIsolatedAudioPreview(hasIsolatedAudioPreview *bool) {
	v.HasIsolatedAudioPreview = hasIsolatedAudioPreview
}

func (v *VoiceSample) SetSpeakerSeparation(speakerSeparation *SpeakerSeparationResponseModel) {
	v.SpeakerSeparation = speakerSeparation
}

func (v *VoiceSample) SetTrimStart(trimStart *int) {
	v.TrimStart = trimStart
}

func (v *VoiceSample) SetTrimEnd(trimEnd *int) {
	v.TrimEnd = trimEnd
}

func (v *VoiceSample) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceSample) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceSample to string"
	}
	return string(jsonData)
}

func (v *VoiceSample) UnmarshalJSON(data []byte) error {
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

func (v *VoiceSample) MarshalJSON() ([]byte, error) {
	type embed VoiceSample
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
