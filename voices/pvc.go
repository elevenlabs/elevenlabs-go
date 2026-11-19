package voices

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreatePvcVoiceRequest struct {
	// The name that identifies this voice. This will be displayed in the dropdown of the website.
	Name string `json:"name" xml:"name" required:"true" maxLength:"100"`
	// Language used in the samples.
	Language string `json:"language" xml:"language" required:"true"`
	// Description to use for the created voice.
	Description *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	// Labels for the voice. Keys can be language, accent, gender, or age.
	Labels map[string]*string `json:"labels,omitempty"`
}

func (c *CreatePvcVoiceRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreatePvcVoiceRequest) GetLanguage() (value string) {
	if c == nil {
		return
	}
	return c.Language
}

func (c *CreatePvcVoiceRequest) GetDescription() (value *string) {
	if c == nil {
		return
	}
	return c.Description
}

func (c *CreatePvcVoiceRequest) GetLabels() (value map[string]*string) {
	if c == nil {
		return
	}
	return c.Labels
}

func (c *CreatePvcVoiceRequest) SetName(name string) {
	c.Name = name
}

func (c *CreatePvcVoiceRequest) SetLanguage(language string) {
	c.Language = language
}

func (c *CreatePvcVoiceRequest) SetDescription(description *string) {
	c.Description = description
}

func (c *CreatePvcVoiceRequest) SetLabels(labels map[string]*string) {
	c.Labels = labels
}

func (c *CreatePvcVoiceRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreatePvcVoiceRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePvcVoiceRequest to string"
	}
	return string(jsonData)
}

func (c *CreatePvcVoiceRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

func (c *CreatePvcVoiceRequest) MarshalJSON() ([]byte, error) {
	type embed CreatePvcVoiceRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

type BodyEditPvcVoiceV1VoicesPvcVoiceIdPost struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// The name that identifies this voice. This will be displayed in the dropdown of the website.
	Name *string `json:"name,omitempty" xml:"name,omitempty" maxLength:"100"`
	// Language used in the samples.
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// Description to use for the created voice.
	Description *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	// Labels for the voice. Keys can be language, accent, gender, or age.
	Labels map[string]*string `json:"labels,omitempty"`
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) GetVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.VoiceId
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) GetLanguage() (value *string) {
	if b == nil {
		return
	}
	return b.Language
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) GetDescription() (value *string) {
	if b == nil {
		return
	}
	return b.Description
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) GetLabels() (value map[string]*string) {
	if b == nil {
		return
	}
	return b.Labels
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) SetVoiceId(voiceId string) {
	b.VoiceId = voiceId
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) SetName(name *string) {
	b.Name = name
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) SetLanguage(language *string) {
	b.Language = language
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) SetDescription(description *string) {
	b.Description = description
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) SetLabels(labels map[string]*string) {
	b.Labels = labels
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyEditPvcVoiceV1VoicesPvcVoiceIdPost to string"
	}
	return string(jsonData)
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyEditPvcVoiceV1VoicesPvcVoiceIdPost) MarshalJSON() ([]byte, error) {
	type embed BodyEditPvcVoiceV1VoicesPvcVoiceIdPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// The model ID to use for the conversion.
	ModelId *string `json:"model_id,omitempty" xml:"model_id,omitempty"`
}

func (b *BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost) GetVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.VoiceId
}

func (b *BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost) GetModelId() (value *string) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost) SetVoiceId(voiceId string) {
	b.VoiceId = voiceId
}

func (b *BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost) SetModelId(modelId *string) {
	b.ModelId = modelId
}

func (b *BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost to string"
	}
	return string(jsonData)
}

func (b *BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost) MarshalJSON() ([]byte, error) {
	type embed BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
