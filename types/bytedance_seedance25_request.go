package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the ByteDance Seedance 2.5 video model.
//
// Diverges from the Seedance 2.0 public shape: no 4K, durations up to 30s,
// larger reference caps, audio-only input allowed, and no “seed“ (Ark
// tolerates it but does not honour it).
//
// ByteDance models are disabled by default and require explicit approval
// before use. Contact support to request access.
type BytedanceSeedance25Request struct {
	// Include to send the generation's result to the workspace's configured flows webhooks once it completes or fails. The webhook payload matches the terminal response of the corresponding GET endpoint.
	Webhook *WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// The model to use for the generation.
	modelId string
	// A text description of the video to generate.
	Prompt string `json:"prompt" xml:"prompt" required:"true"`
	// The aspect ratio of the output video. With `auto`, the model picks an aspect ratio based on the inputs. First-frame / first-and-last-frame tasks always use `auto`.
	AspectRatio *BytedanceSeedance25RequestAspectRatio `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// The resolution of the output video.
	Resolution *BytedanceSeedance25RequestResolution `json:"resolution,omitempty" xml:"resolution,omitempty"`
	// The duration of the output video in seconds.
	DurationSecs *int `json:"duration_secs,omitempty" xml:"duration_secs,omitempty" min:"4" max:"30"`
	// Whether to generate audio with the video.
	GenerateAudio *bool `json:"generate_audio,omitempty" xml:"generate_audio,omitempty"`
	// The image to use as the video's first frame.
	StartFrame *ImageReference `json:"start_frame,omitempty" xml:"start_frame,omitempty"`
	// The image to use as the video's last frame.
	EndFrame *ImageReference `json:"end_frame,omitempty" xml:"end_frame,omitempty"`
	// Up to 30 reference images to draw subjects from. Cannot be combined with `start_frame`/`end_frame`.
	Images []*ImageReference `json:"images,omitempty" xml:"images,omitempty" maxItems:"30"`
	// Up to 10 reference videos to draw subjects or motion from. Cannot be combined with `start_frame`/`end_frame`.
	Videos []*VideoReference `json:"videos,omitempty" xml:"videos,omitempty" maxItems:"10"`
	// Up to 10 reference audios, e.g. for lipsync. Cannot be combined with `start_frame`/`end_frame`.
	Audios          []*AudioReference `json:"audios,omitempty" xml:"audios,omitempty" maxItems:"10"`
	extraProperties map[string]interface{}
}

func (b *BytedanceSeedance25Request) GetWebhook() (value *WebhookTarget) {
	if b == nil {
		return
	}
	return b.Webhook
}

func (b *BytedanceSeedance25Request) GetPrompt() (value string) {
	if b == nil {
		return
	}
	return b.Prompt
}

func (b *BytedanceSeedance25Request) GetAspectRatio() (value *BytedanceSeedance25RequestAspectRatio) {
	if b == nil {
		return
	}
	return b.AspectRatio
}

func (b *BytedanceSeedance25Request) GetResolution() (value *BytedanceSeedance25RequestResolution) {
	if b == nil {
		return
	}
	return b.Resolution
}

func (b *BytedanceSeedance25Request) GetDurationSecs() (value *int) {
	if b == nil {
		return
	}
	return b.DurationSecs
}

func (b *BytedanceSeedance25Request) GetGenerateAudio() (value *bool) {
	if b == nil {
		return
	}
	return b.GenerateAudio
}

func (b *BytedanceSeedance25Request) GetStartFrame() (value *ImageReference) {
	if b == nil {
		return
	}
	return b.StartFrame
}

func (b *BytedanceSeedance25Request) GetEndFrame() (value *ImageReference) {
	if b == nil {
		return
	}
	return b.EndFrame
}

func (b *BytedanceSeedance25Request) GetImages() (value []*ImageReference) {
	if b == nil {
		return
	}
	return b.Images
}

func (b *BytedanceSeedance25Request) GetVideos() (value []*VideoReference) {
	if b == nil {
		return
	}
	return b.Videos
}

func (b *BytedanceSeedance25Request) GetAudios() (value []*AudioReference) {
	if b == nil {
		return
	}
	return b.Audios
}

func (b *BytedanceSeedance25Request) SetWebhook(webhook *WebhookTarget) {
	b.Webhook = webhook
}

func (b *BytedanceSeedance25Request) SetPrompt(prompt string) {
	b.Prompt = prompt
}

func (b *BytedanceSeedance25Request) SetAspectRatio(aspectRatio *BytedanceSeedance25RequestAspectRatio) {
	b.AspectRatio = aspectRatio
}

func (b *BytedanceSeedance25Request) SetResolution(resolution *BytedanceSeedance25RequestResolution) {
	b.Resolution = resolution
}

func (b *BytedanceSeedance25Request) SetDurationSecs(durationSecs *int) {
	b.DurationSecs = durationSecs
}

func (b *BytedanceSeedance25Request) SetGenerateAudio(generateAudio *bool) {
	b.GenerateAudio = generateAudio
}

func (b *BytedanceSeedance25Request) SetStartFrame(startFrame *ImageReference) {
	b.StartFrame = startFrame
}

func (b *BytedanceSeedance25Request) SetEndFrame(endFrame *ImageReference) {
	b.EndFrame = endFrame
}

func (b *BytedanceSeedance25Request) SetImages(images []*ImageReference) {
	b.Images = images
}

func (b *BytedanceSeedance25Request) SetVideos(videos []*VideoReference) {
	b.Videos = videos
}

func (b *BytedanceSeedance25Request) SetAudios(audios []*AudioReference) {
	b.Audios = audios
}

func (b *BytedanceSeedance25Request) ModelId() string {
	return b.modelId
}

func (b *BytedanceSeedance25Request) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BytedanceSeedance25Request) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BytedanceSeedance25Request to string"
	}
	return string(jsonData)
}

func (b *BytedanceSeedance25Request) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "bytedance-seedance-v2.5" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", b, "bytedance-seedance-v2.5", *literals.ModelId)
	}
	b.modelId = "bytedance-seedance-v2.5"
	extraProperties, err := internal.ExtractExtraProperties(data, *b, "model_id")
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BytedanceSeedance25Request) MarshalJSON() ([]byte, error) {
	type embed BytedanceSeedance25Request
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*b),
		ModelId: "bytedance-seedance-v2.5",
	}
	return json.Marshal(marshaler)
}
