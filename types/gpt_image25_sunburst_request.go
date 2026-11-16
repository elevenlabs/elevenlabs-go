package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the OpenAI GPT Image 2.5 Sunburst model.
type GptImage25SunburstRequest struct {
	// Include to send the generation's result to the workspace's configured flows webhooks once it completes or fails. The webhook payload matches the terminal response of the corresponding GET endpoint.
	Webhook *WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// A text description of the image to generate.
	Prompt string `json:"prompt" xml:"prompt" required:"true"`
	// Up to 10 reference images to edit or draw from.
	Images []*ImageReference `json:"images,omitempty" xml:"images,omitempty" maxItems:"10"`
	// An image whose fully transparent areas mark where the first reference image may be edited; requires `images`.
	Mask *ImageReference `json:"mask,omitempty" xml:"mask,omitempty"`
	// The quality of the output image.
	Quality *GptImage25SunburstRequestQuality `json:"quality,omitempty" xml:"quality,omitempty"`
	// The model to use for the generation.
	modelId string
	// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
	AspectRatio *GptImage25SunburstRequestAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// The resolution of the output image.
	Resolution      *GptImage25SunburstRequestResolution `json:"resolution,omitempty" xml:"resolution,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GptImage25SunburstRequest) GetWebhook() (value *WebhookTarget) {
	if g == nil {
		return
	}
	return g.Webhook
}

func (g *GptImage25SunburstRequest) GetPrompt() (value string) {
	if g == nil {
		return
	}
	return g.Prompt
}

func (g *GptImage25SunburstRequest) GetImages() (value []*ImageReference) {
	if g == nil {
		return
	}
	return g.Images
}

func (g *GptImage25SunburstRequest) GetMask() (value *ImageReference) {
	if g == nil {
		return
	}
	return g.Mask
}

func (g *GptImage25SunburstRequest) GetQuality() (value *GptImage25SunburstRequestQuality) {
	if g == nil {
		return
	}
	return g.Quality
}

func (g *GptImage25SunburstRequest) GetAspectRatio() (value *GptImage25SunburstRequestAspectRatio) {
	if g == nil {
		return
	}
	return g.AspectRatio
}

func (g *GptImage25SunburstRequest) GetResolution() (value *GptImage25SunburstRequestResolution) {
	if g == nil {
		return
	}
	return g.Resolution
}

func (g *GptImage25SunburstRequest) SetWebhook(webhook *WebhookTarget) {
	g.Webhook = webhook
}

func (g *GptImage25SunburstRequest) SetPrompt(prompt string) {
	g.Prompt = prompt
}

func (g *GptImage25SunburstRequest) SetImages(images []*ImageReference) {
	g.Images = images
}

func (g *GptImage25SunburstRequest) SetMask(mask *ImageReference) {
	g.Mask = mask
}

func (g *GptImage25SunburstRequest) SetQuality(quality *GptImage25SunburstRequestQuality) {
	g.Quality = quality
}

func (g *GptImage25SunburstRequest) SetAspectRatio(aspectRatio *GptImage25SunburstRequestAspectRatio) {
	g.AspectRatio = aspectRatio
}

func (g *GptImage25SunburstRequest) SetResolution(resolution *GptImage25SunburstRequestResolution) {
	g.Resolution = resolution
}

func (g *GptImage25SunburstRequest) ModelId() string {
	return g.modelId
}

func (g *GptImage25SunburstRequest) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GptImage25SunburstRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GptImage25SunburstRequest to string"
	}
	return string(jsonData)
}

func (g *GptImage25SunburstRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "gpt-image-2.5-sunburst" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", g, "gpt-image-2.5-sunburst", *literals.ModelId)
	}
	g.modelId = "gpt-image-2.5-sunburst"
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "model_id")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GptImage25SunburstRequest) MarshalJSON() ([]byte, error) {
	type embed GptImage25SunburstRequest
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*g),
		ModelId: "gpt-image-2.5-sunburst",
	}
	return json.Marshal(marshaler)
}
