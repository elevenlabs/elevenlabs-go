package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceDesignRequestModel struct {
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
	OutputFormat *AllowedOutputFormats `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	// Description to use for the created voice.
	VoiceDescription string `json:"voice_description" xml:"voice_description" required:"true" maxLength:"1000" minLength:"20"`
	// Model to use for the voice generation. Possible values: eleven_multilingual_ttv_v2, eleven_ttv_v3.
	ModelId *VoiceDesignRequestModelModelId `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Text to generate, text length has to be between 100 and 1000.
	Text *string `json:"text,omitempty" xml:"text,omitempty" maxLength:"1000" minLength:"100"`
	// Whether to automatically generate a text suitable for the voice description.
	AutoGenerateText *bool `json:"auto_generate_text,omitempty" xml:"auto_generate_text,omitempty"`
	// Controls the volume level of the generated voice. -1 is quietest, 1 is loudest, 0 corresponds to roughly -24 LUFS.
	Loudness *float64 `json:"loudness,omitempty" xml:"loudness,omitempty" min:"-1" max:"1"`
	// Random number that controls the voice generation. Same seed with same inputs produces same voice.
	Seed *int `json:"seed,omitempty" xml:"seed,omitempty" min:"0" max:"2147483647"`
	// Controls how closely the AI follows the prompt. Lower numbers give the AI more freedom to be creative, while higher numbers force it to stick more to the prompt. High numbers can cause voice to sound artificial or robotic. We recommend to use longer, more detailed prompts at lower Guidance Scale.
	GuidanceScale *float64 `json:"guidance_scale,omitempty" xml:"guidance_scale,omitempty" min:"0" max:"100"`
	// Determines whether the Text to Voice previews should be included in the response. If true, only the generated IDs will be returned which can then be streamed via the /v1/text-to-voice/:generated_voice_id/stream endpoint.
	StreamPreviews *bool `json:"stream_previews,omitempty" xml:"stream_previews,omitempty"`
	// Whether to enhance the voice description using AI to add more detail and improve voice generation quality. When enabled, the system will automatically expand simple prompts into more detailed voice descriptions. Defaults to False
	ShouldEnhance *bool `json:"should_enhance,omitempty" xml:"should_enhance,omitempty"`
	// The remixing session id.
	RemixingSessionId *string `json:"remixing_session_id,omitempty" xml:"remixing_session_id,omitempty"`
	// The id of the remixing session iteration where these generations should be attached to. If not provided, a new iteration will be created.
	RemixingSessionIterationId *string `json:"remixing_session_iteration_id,omitempty" xml:"remixing_session_iteration_id,omitempty"`
	// Higher quality results in better voice output but less variety.
	Quality *float64 `json:"quality,omitempty" xml:"quality,omitempty" min:"-1" max:"1"`
	// Reference audio to use for the voice generation. The audio should be base64 encoded. Only supported when using the  eleven_ttv_v3 model.
	ReferenceAudioBase64 *string `json:"reference_audio_base64,omitempty" xml:"reference_audio_base64,omitempty"`
	// Controls the balance of prompt versus reference audio when generating voice samples. 0 means almost no prompt influence, 1 means almost no reference audio influence. Only supported when using the eleven_ttv_v3 model.
	PromptStrength *float64 `json:"prompt_strength,omitempty" xml:"prompt_strength,omitempty" min:"0" max:"1"`
}

func (v *VoiceDesignRequestModel) GetOutputFormat() (value *AllowedOutputFormats) {
	if v == nil {
		return
	}
	return v.OutputFormat
}

func (v *VoiceDesignRequestModel) GetVoiceDescription() (value string) {
	if v == nil {
		return
	}
	return v.VoiceDescription
}

func (v *VoiceDesignRequestModel) GetModelId() (value *VoiceDesignRequestModelModelId) {
	if v == nil {
		return
	}
	return v.ModelId
}

func (v *VoiceDesignRequestModel) GetText() (value *string) {
	if v == nil {
		return
	}
	return v.Text
}

func (v *VoiceDesignRequestModel) GetAutoGenerateText() (value *bool) {
	if v == nil {
		return
	}
	return v.AutoGenerateText
}

func (v *VoiceDesignRequestModel) GetLoudness() (value *float64) {
	if v == nil {
		return
	}
	return v.Loudness
}

func (v *VoiceDesignRequestModel) GetSeed() (value *int) {
	if v == nil {
		return
	}
	return v.Seed
}

func (v *VoiceDesignRequestModel) GetGuidanceScale() (value *float64) {
	if v == nil {
		return
	}
	return v.GuidanceScale
}

func (v *VoiceDesignRequestModel) GetStreamPreviews() (value *bool) {
	if v == nil {
		return
	}
	return v.StreamPreviews
}

func (v *VoiceDesignRequestModel) GetShouldEnhance() (value *bool) {
	if v == nil {
		return
	}
	return v.ShouldEnhance
}

func (v *VoiceDesignRequestModel) GetRemixingSessionId() (value *string) {
	if v == nil {
		return
	}
	return v.RemixingSessionId
}

func (v *VoiceDesignRequestModel) GetRemixingSessionIterationId() (value *string) {
	if v == nil {
		return
	}
	return v.RemixingSessionIterationId
}

func (v *VoiceDesignRequestModel) GetQuality() (value *float64) {
	if v == nil {
		return
	}
	return v.Quality
}

func (v *VoiceDesignRequestModel) GetReferenceAudioBase64() (value *string) {
	if v == nil {
		return
	}
	return v.ReferenceAudioBase64
}

func (v *VoiceDesignRequestModel) GetPromptStrength() (value *float64) {
	if v == nil {
		return
	}
	return v.PromptStrength
}

func (v *VoiceDesignRequestModel) SetOutputFormat(outputFormat *AllowedOutputFormats) {
	v.OutputFormat = outputFormat
}

func (v *VoiceDesignRequestModel) SetVoiceDescription(voiceDescription string) {
	v.VoiceDescription = voiceDescription
}

func (v *VoiceDesignRequestModel) SetModelId(modelId *VoiceDesignRequestModelModelId) {
	v.ModelId = modelId
}

func (v *VoiceDesignRequestModel) SetText(text *string) {
	v.Text = text
}

func (v *VoiceDesignRequestModel) SetAutoGenerateText(autoGenerateText *bool) {
	v.AutoGenerateText = autoGenerateText
}

func (v *VoiceDesignRequestModel) SetLoudness(loudness *float64) {
	v.Loudness = loudness
}

func (v *VoiceDesignRequestModel) SetSeed(seed *int) {
	v.Seed = seed
}

func (v *VoiceDesignRequestModel) SetGuidanceScale(guidanceScale *float64) {
	v.GuidanceScale = guidanceScale
}

func (v *VoiceDesignRequestModel) SetStreamPreviews(streamPreviews *bool) {
	v.StreamPreviews = streamPreviews
}

func (v *VoiceDesignRequestModel) SetShouldEnhance(shouldEnhance *bool) {
	v.ShouldEnhance = shouldEnhance
}

func (v *VoiceDesignRequestModel) SetRemixingSessionId(remixingSessionId *string) {
	v.RemixingSessionId = remixingSessionId
}

func (v *VoiceDesignRequestModel) SetRemixingSessionIterationId(remixingSessionIterationId *string) {
	v.RemixingSessionIterationId = remixingSessionIterationId
}

func (v *VoiceDesignRequestModel) SetQuality(quality *float64) {
	v.Quality = quality
}

func (v *VoiceDesignRequestModel) SetReferenceAudioBase64(referenceAudioBase64 *string) {
	v.ReferenceAudioBase64 = referenceAudioBase64
}

func (v *VoiceDesignRequestModel) SetPromptStrength(promptStrength *float64) {
	v.PromptStrength = promptStrength
}

func (v *VoiceDesignRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (v VoiceDesignRequestModel) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceDesignRequestModel to string"
	}
	return string(jsonData)
}

func (v *VoiceDesignRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, v)
}

func (v *VoiceDesignRequestModel) MarshalJSON() ([]byte, error) {
	type embed VoiceDesignRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
