package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the Google Gemini 2.5 Flash image model.
type Gemini25FlashImageRequest struct {
	// Include to send the generation's result to the workspace's configured flows webhooks once it completes or fails. The webhook payload matches the terminal response of the corresponding GET endpoint.
	Webhook *WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// A text description of the image to generate.
	Prompt string `json:"prompt" xml:"prompt" required:"true"`
	// The model to use for the generation.
	modelId string
	// Up to 5 reference images to edit or draw from.
	Images []*ImageReference `json:"images,omitempty" xml:"images,omitempty" maxItems:"5"`
	// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
	AspectRatio     *Gemini25FlashImageRequestAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	extraProperties map[string]interface{}
}

func (g *Gemini25FlashImageRequest) GetWebhook() (value *WebhookTarget) {
	if g == nil {
		return
	}
	return g.Webhook
}

func (g *Gemini25FlashImageRequest) GetPrompt() (value string) {
	if g == nil {
		return
	}
	return g.Prompt
}

func (g *Gemini25FlashImageRequest) GetImages() (value []*ImageReference) {
	if g == nil {
		return
	}
	return g.Images
}

func (g *Gemini25FlashImageRequest) GetAspectRatio() (value *Gemini25FlashImageRequestAspectRatio) {
	if g == nil {
		return
	}
	return g.AspectRatio
}

func (g *Gemini25FlashImageRequest) SetWebhook(webhook *WebhookTarget) {
	g.Webhook = webhook
}

func (g *Gemini25FlashImageRequest) SetPrompt(prompt string) {
	g.Prompt = prompt
}

func (g *Gemini25FlashImageRequest) SetImages(images []*ImageReference) {
	g.Images = images
}

func (g *Gemini25FlashImageRequest) SetAspectRatio(aspectRatio *Gemini25FlashImageRequestAspectRatio) {
	g.AspectRatio = aspectRatio
}

func (g *Gemini25FlashImageRequest) ModelId() string {
	return g.modelId
}

func (g *Gemini25FlashImageRequest) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g Gemini25FlashImageRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: Gemini25FlashImageRequest to string"
	}
	return string(jsonData)
}

func (g *Gemini25FlashImageRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "gemini-2.5-flash-image" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", g, "gemini-2.5-flash-image", *literals.ModelId)
	}
	g.modelId = "gemini-2.5-flash-image"
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "model_id")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *Gemini25FlashImageRequest) MarshalJSON() ([]byte, error) {
	type embed Gemini25FlashImageRequest
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*g),
		ModelId: "gemini-2.5-flash-image",
	}
	return json.Marshal(marshaler)
}
