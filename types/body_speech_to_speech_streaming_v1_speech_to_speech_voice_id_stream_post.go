package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost struct {
	// ID of the voice to be used. Use the [Get voices](/docs/api-reference/voices/search) endpoint list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
	OutputFormat *StreamSpeechToSpeechRequestOutputFormat `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	// You can turn on latency optimizations at some cost of quality. The best possible final latency varies by model. Possible values:
	// 0 - default mode (no latency optimizations)
	// 1 - normal latency optimizations (about 50% of possible latency improvement of option 3)
	// 2 - strong latency optimizations (about 75% of possible latency improvement of option 3)
	// 3 - max latency optimizations
	// 4 - max latency optimizations, but also with text normalizer turned off for even more latency savings (best latency, but can mispronounce eg numbers and dates).
	//
	// Defaults to None.
	//
	OptimizeStreamingLatency *int `json:"-" url:"optimize_streaming_latency,omitempty" explode:"true" serializationStyle:"form" queryParam:"optimize_streaming_latency"`
	// When enable_logging is set to false zero retention mode will be used for the request. This will mean history features are unavailable for this request, including request stitching. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// The audio file which holds the content and emotion that will control the generated speech.
	Audio io.Reader `json:"audio" xml:"audio" required:"true"`
	// Identifier of the model that will be used, you can query them using GET /v1/models. The model needs to have support for speech to speech, you can check this using the can_do_voice_conversion property.
	ModelId *string `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Voice settings overriding stored settings for the given voice. They are applied only on the given request. Needs to be send as a JSON encoded string.
	VoiceSettings *string `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	// If specified, our system will make a best effort to sample deterministically, such that repeated requests with the same seed and parameters should return the same result. Determinism is not guaranteed. Must be integer between 0 and 4294967295.
	Seed *int `json:"seed,omitempty" xml:"seed,omitempty"`
	// If set, will remove the background noise from your audio input using our audio isolation model. Only applies to Voice Changer.
	RemoveBackgroundNoise *bool `json:"remove_background_noise,omitempty" xml:"remove_background_noise,omitempty"`
	// The format of input audio. Options are 'pcm_s16le_16' or 'other' For `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate, single channel (mono), and little-endian byte order. Latency will be lower than with passing an encoded waveform.
	FileFormat *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormat `json:"file_format,omitempty" xml:"file_format,omitempty"`
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) GetVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.VoiceId
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) GetOutputFormat() (value *StreamSpeechToSpeechRequestOutputFormat) {
	if b == nil {
		return
	}
	return b.OutputFormat
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) GetOptimizeStreamingLatency() (value *int) {
	if b == nil {
		return
	}
	return b.OptimizeStreamingLatency
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) GetEnableLogging() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableLogging
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) GetAudio() (value io.Reader) {
	if b == nil {
		return
	}
	return b.Audio
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) GetModelId() (value *string) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) GetVoiceSettings() (value *string) {
	if b == nil {
		return
	}
	return b.VoiceSettings
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) GetRemoveBackgroundNoise() (value *bool) {
	if b == nil {
		return
	}
	return b.RemoveBackgroundNoise
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) GetFileFormat() (value *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormat) {
	if b == nil {
		return
	}
	return b.FileFormat
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) SetVoiceId(voiceId string) {
	b.VoiceId = voiceId
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) SetOutputFormat(outputFormat *StreamSpeechToSpeechRequestOutputFormat) {
	b.OutputFormat = outputFormat
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) SetOptimizeStreamingLatency(optimizeStreamingLatency *int) {
	b.OptimizeStreamingLatency = optimizeStreamingLatency
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) SetEnableLogging(enableLogging *bool) {
	b.EnableLogging = enableLogging
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) SetAudio(audio io.Reader) {
	b.Audio = audio
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) SetModelId(modelId *string) {
	b.ModelId = modelId
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) SetVoiceSettings(voiceSettings *string) {
	b.VoiceSettings = voiceSettings
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) SetRemoveBackgroundNoise(removeBackgroundNoise *bool) {
	b.RemoveBackgroundNoise = removeBackgroundNoise
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) SetFileFormat(fileFormat *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormat) {
	b.FileFormat = fileFormat
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost to string"
	}
	return string(jsonData)
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost) MarshalJSON() ([]byte, error) {
	type embed BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
