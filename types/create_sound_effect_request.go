package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateSoundEffectRequest struct {
	// When enable_logging is set to false zero retention mode will be used for the request. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM and WAV formats with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
	OutputFormat *string `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	// The text that will get converted into a sound effect.
	Text string `json:"text" xml:"text" required:"true"`
	// Whether to create a sound effect that loops smoothly.
	Loop *bool `json:"loop,omitempty" xml:"loop,omitempty"`
	// The duration of the sound which will be generated in seconds. Must be at least 0.5 and at most 30. If set to None we will guess the optimal duration using the prompt. Defaults to None.
	DurationSeconds *float64 `json:"duration_seconds,omitempty" xml:"duration_seconds,omitempty"`
	// A higher prompt influence makes your generation follow the prompt more closely while also making generations less variable. Must be a value between 0 and 1. Defaults to 0.3.
	PromptInfluence *float64    `json:"prompt_influence,omitempty" xml:"prompt_influence,omitempty"`
	ModelId         *SfxModelId `json:"model_id,omitempty" xml:"model_id,omitempty"`
}

func (c *CreateSoundEffectRequest) GetEnableLogging() (value *bool) {
	if c == nil {
		return
	}
	return c.EnableLogging
}

func (c *CreateSoundEffectRequest) GetOutputFormat() (value *string) {
	if c == nil {
		return
	}
	return c.OutputFormat
}

func (c *CreateSoundEffectRequest) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *CreateSoundEffectRequest) GetLoop() (value *bool) {
	if c == nil {
		return
	}
	return c.Loop
}

func (c *CreateSoundEffectRequest) GetDurationSeconds() (value *float64) {
	if c == nil {
		return
	}
	return c.DurationSeconds
}

func (c *CreateSoundEffectRequest) GetPromptInfluence() (value *float64) {
	if c == nil {
		return
	}
	return c.PromptInfluence
}

func (c *CreateSoundEffectRequest) GetModelId() (value *SfxModelId) {
	if c == nil {
		return
	}
	return c.ModelId
}

func (c *CreateSoundEffectRequest) SetEnableLogging(enableLogging *bool) {
	c.EnableLogging = enableLogging
}

func (c *CreateSoundEffectRequest) SetOutputFormat(outputFormat *string) {
	c.OutputFormat = outputFormat
}

func (c *CreateSoundEffectRequest) SetText(text string) {
	c.Text = text
}

func (c *CreateSoundEffectRequest) SetLoop(loop *bool) {
	c.Loop = loop
}

func (c *CreateSoundEffectRequest) SetDurationSeconds(durationSeconds *float64) {
	c.DurationSeconds = durationSeconds
}

func (c *CreateSoundEffectRequest) SetPromptInfluence(promptInfluence *float64) {
	c.PromptInfluence = promptInfluence
}

func (c *CreateSoundEffectRequest) SetModelId(modelId *SfxModelId) {
	c.ModelId = modelId
}

func (c *CreateSoundEffectRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreateSoundEffectRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateSoundEffectRequest to string"
	}
	return string(jsonData)
}

func (c *CreateSoundEffectRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

func (c *CreateSoundEffectRequest) MarshalJSON() ([]byte, error) {
	type embed CreateSoundEffectRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
