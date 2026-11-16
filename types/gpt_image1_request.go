package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the OpenAI GPT Image 1 model.
type GptImage1Request struct {
	// Include to send the generation's result to the workspace's configured flows webhooks once it completes or fails. The webhook payload matches the terminal response of the corresponding GET endpoint.
	Webhook *WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// A text description of the image to generate.
	Prompt string `json:"prompt" xml:"prompt" required:"true"`
	// Up to 5 reference images to edit or draw from.
	Images []*ImageReference `json:"images,omitempty" xml:"images,omitempty" maxItems:"5"`
	// An image whose fully transparent areas mark where the first reference image may be edited; requires `images`.
	Mask *ImageReference `json:"mask,omitempty" xml:"mask,omitempty"`
	// The quality of the output image.
	Quality *GptImage1RequestQuality `json:"quality,omitempty" xml:"quality,omitempty"`
	// The background of the output image. With `auto`, the model picks the background that suits the image.
	Background *GptImage1RequestBackground `json:"background,omitempty" xml:"background,omitempty"`
	// The aspect ratio of the output image.
	AspectRatio *GptImage1RequestAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// The model to use for the generation.
	modelId         string
	extraProperties map[string]interface{}
}

func (g *GptImage1Request) GetWebhook() (value *WebhookTarget) {
	if g == nil {
		return
	}
	return g.Webhook
}

func (g *GptImage1Request) GetPrompt() (value string) {
	if g == nil {
		return
	}
	return g.Prompt
}

func (g *GptImage1Request) GetImages() (value []*ImageReference) {
	if g == nil {
		return
	}
	return g.Images
}

func (g *GptImage1Request) GetMask() (value *ImageReference) {
	if g == nil {
		return
	}
	return g.Mask
}

func (g *GptImage1Request) GetQuality() (value *GptImage1RequestQuality) {
	if g == nil {
		return
	}
	return g.Quality
}

func (g *GptImage1Request) GetBackground() (value *GptImage1RequestBackground) {
	if g == nil {
		return
	}
	return g.Background
}

func (g *GptImage1Request) GetAspectRatio() (value *GptImage1RequestAspectRatio) {
	if g == nil {
		return
	}
	return g.AspectRatio
}

func (g *GptImage1Request) SetWebhook(webhook *WebhookTarget) {
	g.Webhook = webhook
}

func (g *GptImage1Request) SetPrompt(prompt string) {
	g.Prompt = prompt
}

func (g *GptImage1Request) SetImages(images []*ImageReference) {
	g.Images = images
}

func (g *GptImage1Request) SetMask(mask *ImageReference) {
	g.Mask = mask
}

func (g *GptImage1Request) SetQuality(quality *GptImage1RequestQuality) {
	g.Quality = quality
}

func (g *GptImage1Request) SetBackground(background *GptImage1RequestBackground) {
	g.Background = background
}

func (g *GptImage1Request) SetAspectRatio(aspectRatio *GptImage1RequestAspectRatio) {
	g.AspectRatio = aspectRatio
}

func (g *GptImage1Request) ModelId() string {
	return g.modelId
}

func (g *GptImage1Request) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GptImage1Request) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GptImage1Request to string"
	}
	return string(jsonData)
}

func (g *GptImage1Request) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "gpt-image-1" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", g, "gpt-image-1", *literals.ModelId)
	}
	g.modelId = "gpt-image-1"
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "model_id")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GptImage1Request) MarshalJSON() ([]byte, error) {
	type embed GptImage1Request
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*g),
		ModelId: "gpt-image-1",
	}
	return json.Marshal(marshaler)
}
