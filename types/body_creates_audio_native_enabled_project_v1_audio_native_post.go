package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyCreatesAudioNativeEnabledProjectV1AudioNativePost struct {
	// Project name.
	Name string `json:"name" xml:"name" required:"true"`
	// (Deprecated) Image URL used in the player. If not provided, default image set in the Player settings is used.
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// Author used in the player and inserted at the start of the uploaded article. If not provided, the default author set in the Player settings is used.
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// Title used in the player and inserted at the top of the uploaded article. If not provided, the default title set in the Player settings is used.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// (Deprecated) Whether to use small player or not. If not provided, default value set in the Player settings is used.
	Small *bool `json:"small,omitempty" xml:"small,omitempty"`
	// Text color used in the player. If not provided, default text color set in the Player settings is used.
	TextColor *string `json:"text_color,omitempty" xml:"text_color,omitempty"`
	// Background color used in the player. If not provided, default background color set in the Player settings is used.
	BackgroundColor *string `json:"background_color,omitempty" xml:"background_color,omitempty"`
	// (Deprecated) Specifies for how many minutes to persist the session across page reloads. If not provided, default sessionization set in the Player settings is used.
	Sessionization *int `json:"sessionization,omitempty" xml:"sessionization,omitempty"`
	// Voice ID used to voice the content. If not provided, default voice ID set in the Player settings is used.
	VoiceId *string `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	// TTS Model ID used in the player. If not provided, default model ID set in the Player settings is used.
	ModelId *string `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Either txt or HTML input file containing the article content. HTML should be formatted as follows '&lt;html&gt;&lt;body&gt;&lt;div&gt;&lt;p&gt;Your content&lt;/p&gt;&lt;h3&gt;More of your content&lt;/h3&gt;&lt;p&gt;Some more of your content&lt;/p&gt;&lt;/div&gt;&lt;/body&gt;&lt;/html&gt;'
	File io.Reader `json:"file,omitempty" xml:"file,omitempty"`
	// Whether to auto convert the project to audio or not.
	AutoConvert *bool `json:"auto_convert,omitempty" xml:"auto_convert,omitempty"`
	//
	// This parameter controls text normalization with four modes: 'auto', 'on', 'apply_english' and 'off'.
	// When set to 'auto', the system will automatically decide whether to apply text normalization
	// (e.g., spelling out numbers). With 'on', text normalization will always be applied, while
	// with 'off', it will be skipped. 'apply_english' is the same as 'on' but will assume that text is in English.
	//
	ApplyTextNormalization *CreateAudioNativeRequestApplyTextNormalization `json:"apply_text_normalization,omitempty" xml:"apply_text_normalization,omitempty"`
	// A list of pronunciation dictionary locators (pronunciation_dictionary_id, version_id) encoded as a list of JSON strings for pronunciation dictionaries to be applied to the text. A list of json encoded strings is required as adding projects may occur through formData as opposed to jsonBody. To specify multiple dictionaries use multiple --form lines in your curl, such as --form 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"Vmd4Zor6fplcA7WrINey\",\"version_id\":\"hRPaxjlTdR7wFMhV4w0b\"}"' --form 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"JzWtcGQMJ6bnlWwyMo7e\",\"version_id\":\"lbmwxiLu4q6txYxgdZqn\"}"'.
	PronunciationDictionaryLocators []string `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetImage() (value *string) {
	if b == nil {
		return
	}
	return b.Image
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetAuthor() (value *string) {
	if b == nil {
		return
	}
	return b.Author
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetTitle() (value *string) {
	if b == nil {
		return
	}
	return b.Title
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetSmall() (value *bool) {
	if b == nil {
		return
	}
	return b.Small
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetTextColor() (value *string) {
	if b == nil {
		return
	}
	return b.TextColor
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetBackgroundColor() (value *string) {
	if b == nil {
		return
	}
	return b.BackgroundColor
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetSessionization() (value *int) {
	if b == nil {
		return
	}
	return b.Sessionization
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetVoiceId() (value *string) {
	if b == nil {
		return
	}
	return b.VoiceId
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetModelId() (value *string) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetAutoConvert() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoConvert
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetApplyTextNormalization() (value *CreateAudioNativeRequestApplyTextNormalization) {
	if b == nil {
		return
	}
	return b.ApplyTextNormalization
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetPronunciationDictionaryLocators() (value []string) {
	if b == nil {
		return
	}
	return b.PronunciationDictionaryLocators
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetName(name string) {
	b.Name = name
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetImage(image *string) {
	b.Image = image
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetAuthor(author *string) {
	b.Author = author
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetTitle(title *string) {
	b.Title = title
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetSmall(small *bool) {
	b.Small = small
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetTextColor(textColor *string) {
	b.TextColor = textColor
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetBackgroundColor(backgroundColor *string) {
	b.BackgroundColor = backgroundColor
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetSessionization(sessionization *int) {
	b.Sessionization = sessionization
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetVoiceId(voiceId *string) {
	b.VoiceId = voiceId
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetModelId(modelId *string) {
	b.ModelId = modelId
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetFile(file io.Reader) {
	b.File = file
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetAutoConvert(autoConvert *bool) {
	b.AutoConvert = autoConvert
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetApplyTextNormalization(applyTextNormalization *CreateAudioNativeRequestApplyTextNormalization) {
	b.ApplyTextNormalization = applyTextNormalization
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []string) {
	b.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreatesAudioNativeEnabledProjectV1AudioNativePost to string"
	}
	return string(jsonData)
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreatesAudioNativeEnabledProjectV1AudioNativePost) MarshalJSON() ([]byte, error) {
	type embed BodyCreatesAudioNativeEnabledProjectV1AudioNativePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
