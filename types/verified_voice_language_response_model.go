package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VerifiedVoiceLanguageResponseModel struct {
	// The language of the voice.
	Language string `json:"language" xml:"language" required:"true"`
	// The voice's model ID.
	ModelId string `json:"model_id" xml:"model_id" required:"true"`
	// The voice's accent, if applicable.
	Accent *string `json:"accent,omitempty" xml:"accent,omitempty"`
	// The voice's locale, if applicable.
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// The voice's preview URL, if applicable.
	PreviewUrl      *string `json:"preview_url,omitempty" xml:"preview_url,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VerifiedVoiceLanguageResponseModel) GetLanguage() (value string) {
	if v == nil {
		return
	}
	return v.Language
}

func (v *VerifiedVoiceLanguageResponseModel) GetModelId() (value string) {
	if v == nil {
		return
	}
	return v.ModelId
}

func (v *VerifiedVoiceLanguageResponseModel) GetAccent() (value *string) {
	if v == nil {
		return
	}
	return v.Accent
}

func (v *VerifiedVoiceLanguageResponseModel) GetLocale() (value *string) {
	if v == nil {
		return
	}
	return v.Locale
}

func (v *VerifiedVoiceLanguageResponseModel) GetPreviewUrl() (value *string) {
	if v == nil {
		return
	}
	return v.PreviewUrl
}

func (v *VerifiedVoiceLanguageResponseModel) SetLanguage(language string) {
	v.Language = language
}

func (v *VerifiedVoiceLanguageResponseModel) SetModelId(modelId string) {
	v.ModelId = modelId
}

func (v *VerifiedVoiceLanguageResponseModel) SetAccent(accent *string) {
	v.Accent = accent
}

func (v *VerifiedVoiceLanguageResponseModel) SetLocale(locale *string) {
	v.Locale = locale
}

func (v *VerifiedVoiceLanguageResponseModel) SetPreviewUrl(previewUrl *string) {
	v.PreviewUrl = previewUrl
}

func (v *VerifiedVoiceLanguageResponseModel) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VerifiedVoiceLanguageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VerifiedVoiceLanguageResponseModel to string"
	}
	return string(jsonData)
}

func (v *VerifiedVoiceLanguageResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VerifiedVoiceLanguageResponseModel) MarshalJSON() ([]byte, error) {
	type embed VerifiedVoiceLanguageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
