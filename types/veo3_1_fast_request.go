package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the Google Veo 3.1 Fast video model.
type Veo31FastRequest struct {
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
	AspectRatio *Veo31FastRequestAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// The resolution of the output video.
	Resolution *Veo31FastRequestResolution `json:"resolution,omitempty" xml:"resolution,omitempty"`
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

func (v *Veo31FastRequest) GetWebhook() (value *WebhookTarget) {
	if v == nil {
		return
	}
	return v.Webhook
}

func (v *Veo31FastRequest) GetPrompt() (value string) {
	if v == nil {
		return
	}
	return v.Prompt
}

func (v *Veo31FastRequest) GetNegativePrompt() (value *string) {
	if v == nil {
		return
	}
	return v.NegativePrompt
}

func (v *Veo31FastRequest) GetSeed() (value *int) {
	if v == nil {
		return
	}
	return v.Seed
}

func (v *Veo31FastRequest) GetEnhancePrompt() (value *bool) {
	if v == nil {
		return
	}
	return v.EnhancePrompt
}

func (v *Veo31FastRequest) GetDurationSecs() (value *int) {
	if v == nil {
		return
	}
	return v.DurationSecs
}

func (v *Veo31FastRequest) GetAspectRatio() (value *Veo31FastRequestAspectRatio) {
	if v == nil {
		return
	}
	return v.AspectRatio
}

func (v *Veo31FastRequest) GetResolution() (value *Veo31FastRequestResolution) {
	if v == nil {
		return
	}
	return v.Resolution
}

func (v *Veo31FastRequest) GetGenerateAudio() (value *bool) {
	if v == nil {
		return
	}
	return v.GenerateAudio
}

func (v *Veo31FastRequest) GetStartFrame() (value *ImageReference) {
	if v == nil {
		return
	}
	return v.StartFrame
}

func (v *Veo31FastRequest) GetEndFrame() (value *ImageReference) {
	if v == nil {
		return
	}
	return v.EndFrame
}

func (v *Veo31FastRequest) GetImages() (value []*VeoImageReference) {
	if v == nil {
		return
	}
	return v.Images
}

func (v *Veo31FastRequest) SetWebhook(webhook *WebhookTarget) {
	v.Webhook = webhook
}

func (v *Veo31FastRequest) SetPrompt(prompt string) {
	v.Prompt = prompt
}

func (v *Veo31FastRequest) SetNegativePrompt(negativePrompt *string) {
	v.NegativePrompt = negativePrompt
}

func (v *Veo31FastRequest) SetSeed(seed *int) {
	v.Seed = seed
}

func (v *Veo31FastRequest) SetEnhancePrompt(enhancePrompt *bool) {
	v.EnhancePrompt = enhancePrompt
}

func (v *Veo31FastRequest) SetDurationSecs(durationSecs *int) {
	v.DurationSecs = durationSecs
}

func (v *Veo31FastRequest) SetAspectRatio(aspectRatio *Veo31FastRequestAspectRatio) {
	v.AspectRatio = aspectRatio
}

func (v *Veo31FastRequest) SetResolution(resolution *Veo31FastRequestResolution) {
	v.Resolution = resolution
}

func (v *Veo31FastRequest) SetGenerateAudio(generateAudio *bool) {
	v.GenerateAudio = generateAudio
}

func (v *Veo31FastRequest) SetStartFrame(startFrame *ImageReference) {
	v.StartFrame = startFrame
}

func (v *Veo31FastRequest) SetEndFrame(endFrame *ImageReference) {
	v.EndFrame = endFrame
}

func (v *Veo31FastRequest) SetImages(images []*VeoImageReference) {
	v.Images = images
}

func (v *Veo31FastRequest) ModelId() string {
	return v.modelId
}

func (v *Veo31FastRequest) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v Veo31FastRequest) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: Veo31FastRequest to string"
	}
	return string(jsonData)
}

func (v *Veo31FastRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "veo-3.1-fast-generate-001" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", v, "veo-3.1-fast-generate-001", *literals.ModelId)
	}
	v.modelId = "veo-3.1-fast-generate-001"
	extraProperties, err := internal.ExtractExtraProperties(data, *v, "model_id")
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *Veo31FastRequest) MarshalJSON() ([]byte, error) {
	type embed Veo31FastRequest
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*v),
		ModelId: "veo-3.1-fast-generate-001",
	}
	return json.Marshal(marshaler)
}
