package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the Creatify Aurora lipsync video model.
type CreatifyAuroraRequest struct {
	// Include to send the generation's result to the workspace's configured flows webhooks once it completes or fails. The webhook payload matches the terminal response of the corresponding GET endpoint.
	Webhook *WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// The model to use for the generation.
	modelId string
	Image   *ImageReference `json:"image" xml:"image" required:"true"`
	Audio   *AudioReference `json:"audio" xml:"audio" required:"true"`
	// The resolution of the output video.
	Resolution *CreatifyAuroraRequestResolution `json:"resolution,omitempty" xml:"resolution,omitempty"`
	// How strongly the generation adheres to the input image. Omit to use the model's default.
	GuidanceScale *float64 `json:"guidance_scale,omitempty" xml:"guidance_scale,omitempty" min:"0" max:"5"`
	// How strongly the lip movements adhere to the audio. Omit to use the model's default.
	AudioGuidanceScale *float64 `json:"audio_guidance_scale,omitempty" xml:"audio_guidance_scale,omitempty" min:"0" max:"5"`
	extraProperties    map[string]interface{}
}

func (c *CreatifyAuroraRequest) GetWebhook() (value *WebhookTarget) {
	if c == nil {
		return
	}
	return c.Webhook
}

func (c *CreatifyAuroraRequest) GetImage() (value *ImageReference) {
	if c == nil {
		return
	}
	return c.Image
}

func (c *CreatifyAuroraRequest) GetAudio() (value *AudioReference) {
	if c == nil {
		return
	}
	return c.Audio
}

func (c *CreatifyAuroraRequest) GetResolution() (value *CreatifyAuroraRequestResolution) {
	if c == nil {
		return
	}
	return c.Resolution
}

func (c *CreatifyAuroraRequest) GetGuidanceScale() (value *float64) {
	if c == nil {
		return
	}
	return c.GuidanceScale
}

func (c *CreatifyAuroraRequest) GetAudioGuidanceScale() (value *float64) {
	if c == nil {
		return
	}
	return c.AudioGuidanceScale
}

func (c *CreatifyAuroraRequest) SetWebhook(webhook *WebhookTarget) {
	c.Webhook = webhook
}

func (c *CreatifyAuroraRequest) SetImage(image *ImageReference) {
	c.Image = image
}

func (c *CreatifyAuroraRequest) SetAudio(audio *AudioReference) {
	c.Audio = audio
}

func (c *CreatifyAuroraRequest) SetResolution(resolution *CreatifyAuroraRequestResolution) {
	c.Resolution = resolution
}

func (c *CreatifyAuroraRequest) SetGuidanceScale(guidanceScale *float64) {
	c.GuidanceScale = guidanceScale
}

func (c *CreatifyAuroraRequest) SetAudioGuidanceScale(audioGuidanceScale *float64) {
	c.AudioGuidanceScale = audioGuidanceScale
}

func (c *CreatifyAuroraRequest) ModelId() string {
	return c.modelId
}

func (c *CreatifyAuroraRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreatifyAuroraRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatifyAuroraRequest to string"
	}
	return string(jsonData)
}

func (c *CreatifyAuroraRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "creatify-aurora" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "creatify-aurora", *literals.ModelId)
	}
	c.modelId = "creatify-aurora"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "model_id")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreatifyAuroraRequest) MarshalJSON() ([]byte, error) {
	type embed CreatifyAuroraRequest
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*c),
		ModelId: "creatify-aurora",
	}
	return json.Marshal(marshaler)
}
