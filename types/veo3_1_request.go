package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the Google Veo 3.1 video model.
type Veo31Request struct {
	// Include to send the generation's result to the workspace's configured flows webhooks once it completes or fails. The webhook payload matches the terminal response of the corresponding GET endpoint.
	Webhook *WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// A text description of the video to generate.
	Prompt string `json:"prompt" xml:"prompt" required:"true"`
	// A text description of what the video should avoid.
	NegativePrompt *string `json:"negative_prompt,omitempty" xml:"negative_prompt,omitempty"`
	// A seed for reproducible generation: the same seed and inputs give similar output across generations. Omit for random.
	Seed *int `json:"seed,omitempty" xml:"seed,omitempty"`
	// Whether the model may rewrite the prompt to improve results.
	EnhancePrompt *bool `json:"enhance_prompt,omitempty" xml:"enhance_prompt,omitempty"`
	// The duration of the output video in seconds.
	DurationSecs *int `json:"duration_secs,omitempty" xml:"duration_secs,omitempty"`
	// The aspect ratio of the output video.
	AspectRatio *Veo31RequestAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// The resolution of the output video.
	Resolution *Veo31RequestResolution `json:"resolution,omitempty" xml:"resolution,omitempty"`
	// Whether to generate audio with the video.
	GenerateAudio *bool `json:"generate_audio,omitempty" xml:"generate_audio,omitempty"`
	// The image to use as the video's first frame.
	StartFrame *ImageReference `json:"start_frame,omitempty" xml:"start_frame,omitempty"`
	// The image to use as the video's last frame.
	EndFrame *ImageReference `json:"end_frame,omitempty" xml:"end_frame,omitempty"`
	// Up to 3 reference images to draw subjects or style from. Cannot be combined with `start_frame`/`end_frame`, and requires the 8-second duration.
	Images []*VeoImageReference `json:"images,omitempty" xml:"images,omitempty" maxItems:"3"`
	// The model to use for the generation.
	modelId         string
	extraProperties map[string]interface{}
}

func (v *Veo31Request) GetWebhook() (value *WebhookTarget) {
	if v == nil {
		return
	}
	return v.Webhook
}

func (v *Veo31Request) GetPrompt() (value string) {
	if v == nil {
		return
	}
	return v.Prompt
}

func (v *Veo31Request) GetNegativePrompt() (value *string) {
	if v == nil {
		return
	}
	return v.NegativePrompt
}

func (v *Veo31Request) GetSeed() (value *int) {
	if v == nil {
		return
	}
	return v.Seed
}

func (v *Veo31Request) GetEnhancePrompt() (value *bool) {
	if v == nil {
		return
	}
	return v.EnhancePrompt
}

func (v *Veo31Request) GetDurationSecs() (value *int) {
	if v == nil {
		return
	}
	return v.DurationSecs
}

func (v *Veo31Request) GetAspectRatio() (value *Veo31RequestAspectRatio) {
	if v == nil {
		return
	}
	return v.AspectRatio
}

func (v *Veo31Request) GetResolution() (value *Veo31RequestResolution) {
	if v == nil {
		return
	}
	return v.Resolution
}

func (v *Veo31Request) GetGenerateAudio() (value *bool) {
	if v == nil {
		return
	}
	return v.GenerateAudio
}

func (v *Veo31Request) GetStartFrame() (value *ImageReference) {
	if v == nil {
		return
	}
	return v.StartFrame
}

func (v *Veo31Request) GetEndFrame() (value *ImageReference) {
	if v == nil {
		return
	}
	return v.EndFrame
}

func (v *Veo31Request) GetImages() (value []*VeoImageReference) {
	if v == nil {
		return
	}
	return v.Images
}

func (v *Veo31Request) SetWebhook(webhook *WebhookTarget) {
	v.Webhook = webhook
}

func (v *Veo31Request) SetPrompt(prompt string) {
	v.Prompt = prompt
}

func (v *Veo31Request) SetNegativePrompt(negativePrompt *string) {
	v.NegativePrompt = negativePrompt
}

func (v *Veo31Request) SetSeed(seed *int) {
	v.Seed = seed
}

func (v *Veo31Request) SetEnhancePrompt(enhancePrompt *bool) {
	v.EnhancePrompt = enhancePrompt
}

func (v *Veo31Request) SetDurationSecs(durationSecs *int) {
	v.DurationSecs = durationSecs
}

func (v *Veo31Request) SetAspectRatio(aspectRatio *Veo31RequestAspectRatio) {
	v.AspectRatio = aspectRatio
}

func (v *Veo31Request) SetResolution(resolution *Veo31RequestResolution) {
	v.Resolution = resolution
}

func (v *Veo31Request) SetGenerateAudio(generateAudio *bool) {
	v.GenerateAudio = generateAudio
}

func (v *Veo31Request) SetStartFrame(startFrame *ImageReference) {
	v.StartFrame = startFrame
}

func (v *Veo31Request) SetEndFrame(endFrame *ImageReference) {
	v.EndFrame = endFrame
}

func (v *Veo31Request) SetImages(images []*VeoImageReference) {
	v.Images = images
}

func (v *Veo31Request) ModelId() string {
	return v.modelId
}

func (v *Veo31Request) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v Veo31Request) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: Veo31Request to string"
	}
	return string(jsonData)
}

func (v *Veo31Request) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "veo-3.1-generate-001" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", v, "veo-3.1-generate-001", *literals.ModelId)
	}
	v.modelId = "veo-3.1-generate-001"
	extraProperties, err := internal.ExtractExtraProperties(data, *v, "model_id")
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *Veo31Request) MarshalJSON() ([]byte, error) {
	type embed Veo31Request
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*v),
		ModelId: "veo-3.1-generate-001",
	}
	return json.Marshal(marshaler)
}
