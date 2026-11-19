package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the OpenAI GPT Image 2.5 Flare model.
type GptImage25FlareRequest struct {
	// Include to send the generation's result to the workspace's configured flows webhooks once it completes or fails. The webhook payload matches the terminal response of the corresponding GET endpoint.
	Webhook *WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// A text description of the image to generate.
	Prompt string `json:"prompt" xml:"prompt" required:"true"`
	// Up to 10 reference images to edit or draw from.
	Images []*ImageReference `json:"images,omitempty" xml:"images,omitempty" maxItems:"10"`
	// An image whose fully transparent areas mark where the first reference image may be edited; requires `images`.
	Mask *ImageReference `json:"mask,omitempty" xml:"mask,omitempty"`
	// The quality of the output image.
	Quality *GptImage25FlareRequestQuality `json:"quality,omitempty" xml:"quality,omitempty"`
	// The model to use for the generation.
	modelId string
	// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
	AspectRatio *GptImage25FlareRequestAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// The resolution of the output image.
	Resolution      *GptImage25FlareRequestResolution `json:"resolution,omitempty" xml:"resolution,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GptImage25FlareRequest) GetWebhook() (value *WebhookTarget) {
	if g == nil {
		return
	}
	return g.Webhook
}

func (g *GptImage25FlareRequest) GetPrompt() (value string) {
	if g == nil {
		return
	}
	return g.Prompt
}

func (g *GptImage25FlareRequest) GetImages() (value []*ImageReference) {
	if g == nil {
		return
	}
	return g.Images
}

func (g *GptImage25FlareRequest) GetMask() (value *ImageReference) {
	if g == nil {
		return
	}
	return g.Mask
}

func (g *GptImage25FlareRequest) GetQuality() (value *GptImage25FlareRequestQuality) {
	if g == nil {
		return
	}
	return g.Quality
}

func (g *GptImage25FlareRequest) GetAspectRatio() (value *GptImage25FlareRequestAspectRatio) {
	if g == nil {
		return
	}
	return g.AspectRatio
}

func (g *GptImage25FlareRequest) GetResolution() (value *GptImage25FlareRequestResolution) {
	if g == nil {
		return
	}
	return g.Resolution
}

func (g *GptImage25FlareRequest) SetWebhook(webhook *WebhookTarget) {
	g.Webhook = webhook
}

func (g *GptImage25FlareRequest) SetPrompt(prompt string) {
	g.Prompt = prompt
}

func (g *GptImage25FlareRequest) SetImages(images []*ImageReference) {
	g.Images = images
}

func (g *GptImage25FlareRequest) SetMask(mask *ImageReference) {
	g.Mask = mask
}

func (g *GptImage25FlareRequest) SetQuality(quality *GptImage25FlareRequestQuality) {
	g.Quality = quality
}

func (g *GptImage25FlareRequest) SetAspectRatio(aspectRatio *GptImage25FlareRequestAspectRatio) {
	g.AspectRatio = aspectRatio
}

func (g *GptImage25FlareRequest) SetResolution(resolution *GptImage25FlareRequestResolution) {
	g.Resolution = resolution
}

func (g *GptImage25FlareRequest) ModelId() string {
	return g.modelId
}

func (g *GptImage25FlareRequest) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GptImage25FlareRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GptImage25FlareRequest to string"
	}
	return string(jsonData)
}

func (g *GptImage25FlareRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "gpt-image-2.5-flare" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", g, "gpt-image-2.5-flare", *literals.ModelId)
	}
	g.modelId = "gpt-image-2.5-flare"
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "model_id")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GptImage25FlareRequest) MarshalJSON() ([]byte, error) {
	type embed GptImage25FlareRequest
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*g),
		ModelId: "gpt-image-2.5-flare",
	}
	return json.Marshal(marshaler)
}
