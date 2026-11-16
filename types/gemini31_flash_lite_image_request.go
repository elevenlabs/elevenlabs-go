package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the Google Gemini 3.1 Flash Lite image model.
type Gemini31FlashLiteImageRequest struct {
	// Include to send the generation's result to the workspace's configured flows webhooks once it completes or fails. The webhook payload matches the terminal response of the corresponding GET endpoint.
	Webhook *WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// A text description of the image to generate.
	Prompt string `json:"prompt" xml:"prompt" required:"true"`
	// The model to use for the generation.
	modelId string
	// Up to 14 reference images to edit or draw from.
	Images []*ImageReference `json:"images,omitempty" xml:"images,omitempty" maxItems:"14"`
	// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
	AspectRatio *Gemini31FlashLiteImageRequestAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// The resolution of the output image.
	Resolution      *string `json:"resolution,omitempty" xml:"resolution,omitempty"`
	extraProperties map[string]interface{}
}

func (g *Gemini31FlashLiteImageRequest) GetWebhook() (value *WebhookTarget) {
	if g == nil {
		return
	}
	return g.Webhook
}

func (g *Gemini31FlashLiteImageRequest) GetPrompt() (value string) {
	if g == nil {
		return
	}
	return g.Prompt
}

func (g *Gemini31FlashLiteImageRequest) GetImages() (value []*ImageReference) {
	if g == nil {
		return
	}
	return g.Images
}

func (g *Gemini31FlashLiteImageRequest) GetAspectRatio() (value *Gemini31FlashLiteImageRequestAspectRatio) {
	if g == nil {
		return
	}
	return g.AspectRatio
}

func (g *Gemini31FlashLiteImageRequest) GetResolution() (value *string) {
	if g == nil {
		return
	}
	return g.Resolution
}

func (g *Gemini31FlashLiteImageRequest) SetWebhook(webhook *WebhookTarget) {
	g.Webhook = webhook
}

func (g *Gemini31FlashLiteImageRequest) SetPrompt(prompt string) {
	g.Prompt = prompt
}

func (g *Gemini31FlashLiteImageRequest) SetImages(images []*ImageReference) {
	g.Images = images
}

func (g *Gemini31FlashLiteImageRequest) SetAspectRatio(aspectRatio *Gemini31FlashLiteImageRequestAspectRatio) {
	g.AspectRatio = aspectRatio
}

func (g *Gemini31FlashLiteImageRequest) SetResolution(resolution *string) {
	g.Resolution = resolution
}

func (g *Gemini31FlashLiteImageRequest) ModelId() string {
	return g.modelId
}

func (g *Gemini31FlashLiteImageRequest) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g Gemini31FlashLiteImageRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: Gemini31FlashLiteImageRequest to string"
	}
	return string(jsonData)
}

func (g *Gemini31FlashLiteImageRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "gemini-3.1-flash-lite-image" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", g, "gemini-3.1-flash-lite-image", *literals.ModelId)
	}
	g.modelId = "gemini-3.1-flash-lite-image"
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "model_id")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *Gemini31FlashLiteImageRequest) MarshalJSON() ([]byte, error) {
	type embed Gemini31FlashLiteImageRequest
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*g),
		ModelId: "gemini-3.1-flash-lite-image",
	}
	return json.Marshal(marshaler)
}
