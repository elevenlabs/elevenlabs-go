package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyCreatePodcastV1StudioPodcastsPost struct {
	// Used for moderation. Your workspace must be allowlisted to use this feature.
	SafetyIdentifier *string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"safety-identifier"`
	// The ID of the model to be used for this Studio project, you can query GET /v1/models to list all available models.
	ModelId string `json:"model_id" xml:"model_id" required:"true"`
	// The type of podcast to generate. Can be 'conversation', an interaction between two voices, or 'bulletin', a monologue.
	Mode *BodyCreatePodcastV1StudioPodcastsPostMode `json:"mode" xml:"mode" required:"true"`
	// The source content for the Podcast.
	Source *BodyCreatePodcastV1StudioPodcastsPostSource `json:"source" xml:"source" required:"true"`
	// Output quality of the generated audio. Must be one of:
	// 'standard' - standard output format, 128kbps with 44.1kHz sample rate.
	// 'high' - high quality output format, 192kbps with 44.1kHz sample rate and major improvements on our side.
	// 'ultra' - ultra quality output format, 192kbps with 44.1kHz sample rate and highest improvements on our side.
	// 'ultra_lossless' - ultra quality output format, 705.6kbps with 44.1kHz sample rate and highest improvements on our side in a fully lossless format.
	// If not provided, defaults to the highest quality preset available on your subscription tier.
	//
	QualityPreset *QualityPresetType `json:"quality_preset,omitempty" xml:"quality_preset,omitempty"`
	// Duration of the generated podcast. Must be one of:
	// short - produces podcasts shorter than 3 minutes.
	// default - produces podcasts roughly between 3-7 minutes.
	// long - produces podcasts longer than 7 minutes.
	//
	DurationScale *BodyCreatePodcastV1StudioPodcastsPostDurationScale `json:"duration_scale,omitempty" xml:"duration_scale,omitempty"`
	// An optional language of the Studio project. Two-letter language code (ISO 639-1).
	Language *string `json:"language,omitempty" xml:"language,omitempty" maxLength:"2" minLength:"2"`
	// The intro text that will always be added to the beginning of the podcast.
	Intro *string `json:"intro,omitempty" xml:"intro,omitempty" maxLength:"1500"`
	// The outro text that will always be added to the end of the podcast.
	Outro *string `json:"outro,omitempty" xml:"outro,omitempty" maxLength:"1500"`
	// Additional instructions prompt for the podcast generation used to adjust the podcast's style and tone.
	InstructionsPrompt *string `json:"instructions_prompt,omitempty" xml:"instructions_prompt,omitempty" maxLength:"3000"`
	// A brief summary or highlights of the Studio project's content, providing key points or themes. This should be between 10 and 70 characters.
	Highlights []string `json:"highlights,omitempty" xml:"highlights,omitempty"`
	//
	// A url that will be called by our service when the Studio project is converted. Request will contain a json blob containing the status of the conversion
	// Messages:
	// 1. When project was converted successfully:
	// {
	// type: "project_conversion_status",
	// event_timestamp: 1234567890,
	// data: {
	// request_id: "1234567890",
	// project_id: "21m00Tcm4TlvDq8ikWAM",
	// conversion_status: "success",
	// project_snapshot_id: "22m00Tcm4TlvDq8ikMAT",
	// error_details: None,
	// }
	// }
	// 2. When project conversion failed:
	// {
	// type: "project_conversion_status",
	// event_timestamp: 1234567890,
	// data: {
	// request_id: "1234567890",
	// project_id: "21m00Tcm4TlvDq8ikWAM",
	// conversion_status: "error",
	// project_snapshot_id: None,
	// error_details: "Error details if conversion failed"
	// }
	// }
	//
	// 3. When chapter was converted successfully:
	// {
	// type: "chapter_conversion_status",
	// event_timestamp: 1234567890,
	// data: {
	// request_id: "1234567890",
	// project_id: "21m00Tcm4TlvDq8ikWAM",
	// chapter_id: "22m00Tcm4TlvDq8ikMAT",
	// conversion_status: "success",
	// chapter_snapshot_id: "23m00Tcm4TlvDq8ikMAV",
	// error_details: None,
	// }
	// }
	// 4. When chapter conversion failed:
	// {
	// type: "chapter_conversion_status",
	// event_timestamp: 1234567890,
	// data: {
	// request_id: "1234567890",
	// project_id: "21m00Tcm4TlvDq8ikWAM",
	// chapter_id: "22m00Tcm4TlvDq8ikMAT",
	// conversion_status: "error",
	// chapter_snapshot_id: None,
	// error_details: "Error details if conversion failed"
	// }
	// }
	//
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty"`
	//
	// This parameter controls text normalization with four modes: 'auto', 'on', 'apply_english' and 'off'.
	// When set to 'auto', the system will automatically decide whether to apply text normalization
	// (e.g., spelling out numbers). With 'on', text normalization will always be applied, while
	// with 'off', it will be skipped. 'apply_english' is the same as 'on' but will assume that text is in English.
	//
	ApplyTextNormalization *BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization `json:"apply_text_normalization,omitempty" xml:"apply_text_normalization,omitempty"`
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetSafetyIdentifier() (value *string) {
	if b == nil {
		return
	}
	return b.SafetyIdentifier
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetModelId() (value string) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetMode() (value *BodyCreatePodcastV1StudioPodcastsPostMode) {
	if b == nil {
		return
	}
	return b.Mode
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetSource() (value *BodyCreatePodcastV1StudioPodcastsPostSource) {
	if b == nil {
		return
	}
	return b.Source
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetQualityPreset() (value *QualityPresetType) {
	if b == nil {
		return
	}
	return b.QualityPreset
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetDurationScale() (value *BodyCreatePodcastV1StudioPodcastsPostDurationScale) {
	if b == nil {
		return
	}
	return b.DurationScale
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetLanguage() (value *string) {
	if b == nil {
		return
	}
	return b.Language
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetIntro() (value *string) {
	if b == nil {
		return
	}
	return b.Intro
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetOutro() (value *string) {
	if b == nil {
		return
	}
	return b.Outro
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetInstructionsPrompt() (value *string) {
	if b == nil {
		return
	}
	return b.InstructionsPrompt
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetHighlights() (value []string) {
	if b == nil {
		return
	}
	return b.Highlights
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetCallbackUrl() (value *string) {
	if b == nil {
		return
	}
	return b.CallbackUrl
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetApplyTextNormalization() (value *BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization) {
	if b == nil {
		return
	}
	return b.ApplyTextNormalization
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetSafetyIdentifier(safetyIdentifier *string) {
	b.SafetyIdentifier = safetyIdentifier
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetModelId(modelId string) {
	b.ModelId = modelId
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetMode(mode *BodyCreatePodcastV1StudioPodcastsPostMode) {
	b.Mode = mode
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetSource(source *BodyCreatePodcastV1StudioPodcastsPostSource) {
	b.Source = source
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetQualityPreset(qualityPreset *QualityPresetType) {
	b.QualityPreset = qualityPreset
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetDurationScale(durationScale *BodyCreatePodcastV1StudioPodcastsPostDurationScale) {
	b.DurationScale = durationScale
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetLanguage(language *string) {
	b.Language = language
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetIntro(intro *string) {
	b.Intro = intro
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetOutro(outro *string) {
	b.Outro = outro
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetInstructionsPrompt(instructionsPrompt *string) {
	b.InstructionsPrompt = instructionsPrompt
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetHighlights(highlights []string) {
	b.Highlights = highlights
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetCallbackUrl(callbackUrl *string) {
	b.CallbackUrl = callbackUrl
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) SetApplyTextNormalization(applyTextNormalization *BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization) {
	b.ApplyTextNormalization = applyTextNormalization
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreatePodcastV1StudioPodcastsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreatePodcastV1StudioPodcastsPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreatePodcastV1StudioPodcastsPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreatePodcastV1StudioPodcastsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
