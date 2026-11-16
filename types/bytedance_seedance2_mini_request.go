package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the ByteDance Seedance 2.0 Mini video model.
//
// ByteDance models are disabled by default and require explicit approval
// before use. Contact support to request access.
type BytedanceSeedance2MiniRequest struct {
	// Include to send the generation's result to the workspace's configured flows webhooks once it completes or fails. The webhook payload matches the terminal response of the corresponding GET endpoint.
	Webhook *WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// A text description of the video to generate.
	Prompt string `json:"prompt" xml:"prompt" required:"true"`
	// The aspect ratio of the output video. With `auto`, the model picks an aspect ratio based on the inputs.
	AspectRatio *BytedanceSeedance2MiniRequestAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// The duration of the output video in seconds.
	DurationSecs *int `json:"duration_secs,omitempty" xml:"duration_secs,omitempty" min:"4" max:"15"`
	// A seed for reproducible generation: the same seed and inputs give similar output across generations. Omit for random.
	Seed *int `json:"seed,omitempty" xml:"seed,omitempty"`
	// Whether to generate audio with the video.
	GenerateAudio *bool `json:"generate_audio,omitempty" xml:"generate_audio,omitempty"`
	// The image to use as the video's first frame.
	StartFrame *ImageReference `json:"start_frame,omitempty" xml:"start_frame,omitempty"`
	// The image to use as the video's last frame.
	EndFrame *ImageReference `json:"end_frame,omitempty" xml:"end_frame,omitempty"`
	// Up to 9 reference images to draw subjects from. Cannot be combined with `start_frame`/`end_frame`.
	Images []*ImageReference `json:"images,omitempty" xml:"images,omitempty" maxItems:"9"`
	// Up to 3 reference videos to draw subjects or motion from. Cannot be combined with `start_frame`/`end_frame`.
	Videos []*VideoReference `json:"videos,omitempty" xml:"videos,omitempty" maxItems:"3"`
	// Up to 3 reference audios, e.g. for lipsync. Requires at least one of `images` or `videos`, and cannot be combined with `start_frame`/`end_frame`.
	Audios []*AudioReference `json:"audios,omitempty" xml:"audios,omitempty" maxItems:"3"`
	// The model to use for the generation.
	modelId string
	// The resolution of the output video.
	Resolution      *BytedanceSeedance2MiniRequestResolution `json:"resolution,omitempty" xml:"resolution,omitempty"`
	extraProperties map[string]interface{}
}

func (b *BytedanceSeedance2MiniRequest) GetWebhook() (value *WebhookTarget) {
	if b == nil {
		return
	}
	return b.Webhook
}

func (b *BytedanceSeedance2MiniRequest) GetPrompt() (value string) {
	if b == nil {
		return
	}
	return b.Prompt
}

func (b *BytedanceSeedance2MiniRequest) GetAspectRatio() (value *BytedanceSeedance2MiniRequestAspectRatio) {
	if b == nil {
		return
	}
	return b.AspectRatio
}

func (b *BytedanceSeedance2MiniRequest) GetDurationSecs() (value *int) {
	if b == nil {
		return
	}
	return b.DurationSecs
}

func (b *BytedanceSeedance2MiniRequest) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BytedanceSeedance2MiniRequest) GetGenerateAudio() (value *bool) {
	if b == nil {
		return
	}
	return b.GenerateAudio
}

func (b *BytedanceSeedance2MiniRequest) GetStartFrame() (value *ImageReference) {
	if b == nil {
		return
	}
	return b.StartFrame
}

func (b *BytedanceSeedance2MiniRequest) GetEndFrame() (value *ImageReference) {
	if b == nil {
		return
	}
	return b.EndFrame
}

func (b *BytedanceSeedance2MiniRequest) GetImages() (value []*ImageReference) {
	if b == nil {
		return
	}
	return b.Images
}

func (b *BytedanceSeedance2MiniRequest) GetVideos() (value []*VideoReference) {
	if b == nil {
		return
	}
	return b.Videos
}

func (b *BytedanceSeedance2MiniRequest) GetAudios() (value []*AudioReference) {
	if b == nil {
		return
	}
	return b.Audios
}

func (b *BytedanceSeedance2MiniRequest) GetResolution() (value *BytedanceSeedance2MiniRequestResolution) {
	if b == nil {
		return
	}
	return b.Resolution
}

func (b *BytedanceSeedance2MiniRequest) SetWebhook(webhook *WebhookTarget) {
	b.Webhook = webhook
}

func (b *BytedanceSeedance2MiniRequest) SetPrompt(prompt string) {
	b.Prompt = prompt
}

func (b *BytedanceSeedance2MiniRequest) SetAspectRatio(aspectRatio *BytedanceSeedance2MiniRequestAspectRatio) {
	b.AspectRatio = aspectRatio
}

func (b *BytedanceSeedance2MiniRequest) SetDurationSecs(durationSecs *int) {
	b.DurationSecs = durationSecs
}

func (b *BytedanceSeedance2MiniRequest) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BytedanceSeedance2MiniRequest) SetGenerateAudio(generateAudio *bool) {
	b.GenerateAudio = generateAudio
}

func (b *BytedanceSeedance2MiniRequest) SetStartFrame(startFrame *ImageReference) {
	b.StartFrame = startFrame
}

func (b *BytedanceSeedance2MiniRequest) SetEndFrame(endFrame *ImageReference) {
	b.EndFrame = endFrame
}

func (b *BytedanceSeedance2MiniRequest) SetImages(images []*ImageReference) {
	b.Images = images
}

func (b *BytedanceSeedance2MiniRequest) SetVideos(videos []*VideoReference) {
	b.Videos = videos
}

func (b *BytedanceSeedance2MiniRequest) SetAudios(audios []*AudioReference) {
	b.Audios = audios
}

func (b *BytedanceSeedance2MiniRequest) SetResolution(resolution *BytedanceSeedance2MiniRequestResolution) {
	b.Resolution = resolution
}

func (b *BytedanceSeedance2MiniRequest) ModelId() string {
	return b.modelId
}

func (b *BytedanceSeedance2MiniRequest) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BytedanceSeedance2MiniRequest) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BytedanceSeedance2MiniRequest to string"
	}
	return string(jsonData)
}

func (b *BytedanceSeedance2MiniRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "bytedance-seedance-v2-mini" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", b, "bytedance-seedance-v2-mini", *literals.ModelId)
	}
	b.modelId = "bytedance-seedance-v2-mini"
	extraProperties, err := internal.ExtractExtraProperties(data, *b, "model_id")
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BytedanceSeedance2MiniRequest) MarshalJSON() ([]byte, error) {
	type embed BytedanceSeedance2MiniRequest
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*b),
		ModelId: "bytedance-seedance-v2-mini",
	}
	return json.Marshal(marshaler)
}
