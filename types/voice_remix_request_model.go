package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceRemixRequestModel struct {
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
	OutputFormat *AllowedOutputFormats `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// Description of the changes to make to the voice.
	VoiceDescription string `json:"voice_description" xml:"voice_description" required:"true" maxLength:"1000" minLength:"5"`
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
	// The remixing session id.
	RemixingSessionId *string `json:"remixing_session_id,omitempty" xml:"remixing_session_id,omitempty"`
	// The id of the remixing session iteration where these generations should be attached to. If not provided, a new iteration will be created.
	RemixingSessionIterationId *string `json:"remixing_session_iteration_id,omitempty" xml:"remixing_session_iteration_id,omitempty"`
	// Controls the balance of prompt versus reference audio when generating voice samples. 0 means almost no prompt influence, 1 means almost no reference audio influence. Only supported when using the eleven_ttv_v3 model.
	PromptStrength *float64 `json:"prompt_strength,omitempty" xml:"prompt_strength,omitempty" min:"0" max:"1"`
}

func (v *VoiceRemixRequestModel) GetOutputFormat() (value *AllowedOutputFormats) {
	if v == nil {
		return
	}
	return v.OutputFormat
}

func (v *VoiceRemixRequestModel) GetVoiceId() (value string) {
	if v == nil {
		return
	}
	return v.VoiceId
}

func (v *VoiceRemixRequestModel) GetVoiceDescription() (value string) {
	if v == nil {
		return
	}
	return v.VoiceDescription
}

func (v *VoiceRemixRequestModel) GetText() (value *string) {
	if v == nil {
		return
	}
	return v.Text
}

func (v *VoiceRemixRequestModel) GetAutoGenerateText() (value *bool) {
	if v == nil {
		return
	}
	return v.AutoGenerateText
}

func (v *VoiceRemixRequestModel) GetLoudness() (value *float64) {
	if v == nil {
		return
	}
	return v.Loudness
}

func (v *VoiceRemixRequestModel) GetSeed() (value *int) {
	if v == nil {
		return
	}
	return v.Seed
}

func (v *VoiceRemixRequestModel) GetGuidanceScale() (value *float64) {
	if v == nil {
		return
	}
	return v.GuidanceScale
}

func (v *VoiceRemixRequestModel) GetStreamPreviews() (value *bool) {
	if v == nil {
		return
	}
	return v.StreamPreviews
}

func (v *VoiceRemixRequestModel) GetRemixingSessionId() (value *string) {
	if v == nil {
		return
	}
	return v.RemixingSessionId
}

func (v *VoiceRemixRequestModel) GetRemixingSessionIterationId() (value *string) {
	if v == nil {
		return
	}
	return v.RemixingSessionIterationId
}

func (v *VoiceRemixRequestModel) GetPromptStrength() (value *float64) {
	if v == nil {
		return
	}
	return v.PromptStrength
}

func (v *VoiceRemixRequestModel) SetOutputFormat(outputFormat *AllowedOutputFormats) {
	v.OutputFormat = outputFormat
}

func (v *VoiceRemixRequestModel) SetVoiceId(voiceId string) {
	v.VoiceId = voiceId
}

func (v *VoiceRemixRequestModel) SetVoiceDescription(voiceDescription string) {
	v.VoiceDescription = voiceDescription
}

func (v *VoiceRemixRequestModel) SetText(text *string) {
	v.Text = text
}

func (v *VoiceRemixRequestModel) SetAutoGenerateText(autoGenerateText *bool) {
	v.AutoGenerateText = autoGenerateText
}

func (v *VoiceRemixRequestModel) SetLoudness(loudness *float64) {
	v.Loudness = loudness
}

func (v *VoiceRemixRequestModel) SetSeed(seed *int) {
	v.Seed = seed
}

func (v *VoiceRemixRequestModel) SetGuidanceScale(guidanceScale *float64) {
	v.GuidanceScale = guidanceScale
}

func (v *VoiceRemixRequestModel) SetStreamPreviews(streamPreviews *bool) {
	v.StreamPreviews = streamPreviews
}

func (v *VoiceRemixRequestModel) SetRemixingSessionId(remixingSessionId *string) {
	v.RemixingSessionId = remixingSessionId
}

func (v *VoiceRemixRequestModel) SetRemixingSessionIterationId(remixingSessionIterationId *string) {
	v.RemixingSessionIterationId = remixingSessionIterationId
}

func (v *VoiceRemixRequestModel) SetPromptStrength(promptStrength *float64) {
	v.PromptStrength = promptStrength
}

func (v *VoiceRemixRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (v VoiceRemixRequestModel) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceRemixRequestModel to string"
	}
	return string(jsonData)
}

func (v *VoiceRemixRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, v)
}

func (v *VoiceRemixRequestModel) MarshalJSON() ([]byte, error) {
	type embed VoiceRemixRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
