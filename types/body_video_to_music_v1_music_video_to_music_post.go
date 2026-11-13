package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyVideoToMusicV1MusicVideoToMusicPost struct {
	// When enable_logging is set to false zero retention mode will be used for the request. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
	OutputFormat *AllowedOutputFormats `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	//
	// One or more video files sent via FormData array (multipart/form-data). They will be combined into one codec in order.
	// A maximum of 10 videos is allowed, where the total size of the combined video is limited to 200MB.
	// In total, the video can be up to 600 seconds long. Note that combining multiple videos may increase the request duration significantly. If possible, combine the videos beforehand.
	//
	Videos []io.Reader `json:"videos" xml:"videos" required:"true" minItems:"1" maxItems:"10"`
	// Optional text description of the music you want. A maximum of 4100 characters is allowed.
	Description *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"4100" minLength:"1"`
	// Optional list of style tags (e.g. ['upbeat', 'cinematic']). A maximum of 10 tags is allowed.
	Tags    []string      `json:"tags,omitempty" xml:"tags,omitempty" maxItems:"10"`
	ModelId *MusicModelId `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// When true, incorporate SFX cues from the video.
	AddSfx *bool `json:"add_sfx,omitempty" xml:"add_sfx,omitempty"`
	// When true, return a ZIP with separate music and SFX audio tracks instead of a single mixed file. Use this for more flexibility to layer the tracks.
	SeparateTracks *bool `json:"separate_tracks,omitempty" xml:"separate_tracks,omitempty"`
	// When true, analyze the video's audio track and use its genres, BPM, instrumentation, mood, and arc as reference characteristics for music generation.
	UseAudioAsReference *bool `json:"use_audio_as_reference,omitempty" xml:"use_audio_as_reference,omitempty"`
	// Whether to sign the generated song with C2PA. Applicable only for mp3 files.
	SignWithC2Pa *bool `json:"sign_with_c2pa,omitempty" xml:"sign_with_c2pa,omitempty"`
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) GetEnableLogging() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableLogging
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) GetOutputFormat() (value *AllowedOutputFormats) {
	if b == nil {
		return
	}
	return b.OutputFormat
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) GetVideos() (value []io.Reader) {
	if b == nil {
		return
	}
	return b.Videos
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) GetDescription() (value *string) {
	if b == nil {
		return
	}
	return b.Description
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) GetTags() (value []string) {
	if b == nil {
		return
	}
	return b.Tags
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) GetModelId() (value *MusicModelId) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) GetAddSfx() (value *bool) {
	if b == nil {
		return
	}
	return b.AddSfx
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) GetSeparateTracks() (value *bool) {
	if b == nil {
		return
	}
	return b.SeparateTracks
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) GetUseAudioAsReference() (value *bool) {
	if b == nil {
		return
	}
	return b.UseAudioAsReference
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) GetSignWithC2Pa() (value *bool) {
	if b == nil {
		return
	}
	return b.SignWithC2Pa
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) SetEnableLogging(enableLogging *bool) {
	b.EnableLogging = enableLogging
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) SetOutputFormat(outputFormat *AllowedOutputFormats) {
	b.OutputFormat = outputFormat
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) SetVideos(videos []io.Reader) {
	b.Videos = videos
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) SetDescription(description *string) {
	b.Description = description
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) SetTags(tags []string) {
	b.Tags = tags
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) SetModelId(modelId *MusicModelId) {
	b.ModelId = modelId
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) SetAddSfx(addSfx *bool) {
	b.AddSfx = addSfx
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) SetSeparateTracks(separateTracks *bool) {
	b.SeparateTracks = separateTracks
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) SetUseAudioAsReference(useAudioAsReference *bool) {
	b.UseAudioAsReference = useAudioAsReference
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) SetSignWithC2Pa(signWithC2Pa *bool) {
	b.SignWithC2Pa = signWithC2Pa
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyVideoToMusicV1MusicVideoToMusicPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyVideoToMusicV1MusicVideoToMusicPost to string"
	}
	return string(jsonData)
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyVideoToMusicV1MusicVideoToMusicPost) MarshalJSON() ([]byte, error) {
	type embed BodyVideoToMusicV1MusicVideoToMusicPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
