package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the ByteDance Seedream 5.0 Pro image model.
//
// ByteDance models are disabled by default and require explicit approval
// before use. Contact support to request access.
type BytedanceSeedream5ProRequest struct {
	// Include to send the generation's result to the workspace's configured flows webhooks once it completes or fails. The webhook payload matches the terminal response of the corresponding GET endpoint.
	Webhook *WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// A text description of the image to generate.
	Prompt string `json:"prompt" xml:"prompt" required:"true"`
	// Up to 10 reference images to edit or draw from.
	Images []*ImageReference `json:"images,omitempty" xml:"images,omitempty" maxItems:"10"`
	// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
	AspectRatio *BytedanceSeedream5ProRequestAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// A seed for reproducible generation: the same seed and inputs give similar output across generations. Omit for random.
	Seed *int `json:"seed,omitempty" xml:"seed,omitempty"`
	// The model to use for the generation.
	modelId string
	// The resolution of the output image.
	Resolution      *BytedanceSeedream5ProRequestResolution `json:"resolution,omitempty" xml:"resolution,omitempty"`
	extraProperties map[string]interface{}
}

func (b *BytedanceSeedream5ProRequest) GetWebhook() (value *WebhookTarget) {
	if b == nil {
		return
	}
	return b.Webhook
}

func (b *BytedanceSeedream5ProRequest) GetPrompt() (value string) {
	if b == nil {
		return
	}
	return b.Prompt
}

func (b *BytedanceSeedream5ProRequest) GetImages() (value []*ImageReference) {
	if b == nil {
		return
	}
	return b.Images
}

func (b *BytedanceSeedream5ProRequest) GetAspectRatio() (value *BytedanceSeedream5ProRequestAspectRatio) {
	if b == nil {
		return
	}
	return b.AspectRatio
}

func (b *BytedanceSeedream5ProRequest) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BytedanceSeedream5ProRequest) GetResolution() (value *BytedanceSeedream5ProRequestResolution) {
	if b == nil {
		return
	}
	return b.Resolution
}

func (b *BytedanceSeedream5ProRequest) SetWebhook(webhook *WebhookTarget) {
	b.Webhook = webhook
}

func (b *BytedanceSeedream5ProRequest) SetPrompt(prompt string) {
	b.Prompt = prompt
}

func (b *BytedanceSeedream5ProRequest) SetImages(images []*ImageReference) {
	b.Images = images
}

func (b *BytedanceSeedream5ProRequest) SetAspectRatio(aspectRatio *BytedanceSeedream5ProRequestAspectRatio) {
	b.AspectRatio = aspectRatio
}

func (b *BytedanceSeedream5ProRequest) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BytedanceSeedream5ProRequest) SetResolution(resolution *BytedanceSeedream5ProRequestResolution) {
	b.Resolution = resolution
}

func (b *BytedanceSeedream5ProRequest) ModelId() string {
	return b.modelId
}

func (b *BytedanceSeedream5ProRequest) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BytedanceSeedream5ProRequest) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BytedanceSeedream5ProRequest to string"
	}
	return string(jsonData)
}

func (b *BytedanceSeedream5ProRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "bytedance-seedream-5-pro" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", b, "bytedance-seedream-5-pro", *literals.ModelId)
	}
	b.modelId = "bytedance-seedream-5-pro"
	extraProperties, err := internal.ExtractExtraProperties(data, *b, "model_id")
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BytedanceSeedream5ProRequest) MarshalJSON() ([]byte, error) {
	type embed BytedanceSeedream5ProRequest
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*b),
		ModelId: "bytedance-seedream-5-pro",
	}
	return json.Marshal(marshaler)
}
