package project

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

type BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost struct {
	// Identifier of the parent dubbing project.
	ProjectId string `json:"-" url:"-"`
	// BCP-47 language tag to dub the project into (for example, `fr` or `es-MX`). Must be one of the [languages the project's dubbing model supports](https://elevenlabs.io/docs/help-center/product/dubbing/which-languages-are-supported-in-dubbing), and a region-qualified tag must be one of the supported dialects.
	TargetLanguage string `json:"target_language" xml:"target_language" required:"true"`
	// Voice settings applied to every speaker in this language. Omit to use the defaults.
	VoiceSettings VoiceSettings `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	// Enterprise only. Optional translations to use instead of machine translation. A map from each source segment's `external_id` (or its `id`, if you supplied none) to the translated text; every source segment must be covered exactly once. At most 20,000 entries, totaling at most 4 MiB of text. See [Bring your own transcript](https://elevenlabs.io/docs/eleven-api/guides/how-to/dubbing/bring-your-own-transcript).
	Translations map[string]*string `json:"translations,omitempty"`
}

func (b *BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) GetProjectId() (value string) {
	if b == nil {
		return
	}
	return b.ProjectId
}

func (b *BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) GetTargetLanguage() (value string) {
	if b == nil {
		return
	}
	return b.TargetLanguage
}

func (b *BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) GetVoiceSettings() (value VoiceSettings) {
	if b == nil {
		return
	}
	return b.VoiceSettings
}

func (b *BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) GetTranslations() (value map[string]*string) {
	if b == nil {
		return
	}
	return b.Translations
}

func (b *BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) SetProjectId(projectId string) {
	b.ProjectId = projectId
}

func (b *BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) SetTargetLanguage(targetLanguage string) {
	b.TargetLanguage = targetLanguage
}

func (b *BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) SetVoiceSettings(voiceSettings VoiceSettings) {
	b.VoiceSettings = voiceSettings
}

func (b *BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) SetTranslations(translations map[string]*string) {
	b.Translations = translations
}

func (b *BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	listLanguageRequestFieldCursor   = big.NewInt(1 << 0)
	listLanguageRequestFieldPageSize = big.NewInt(1 << 1)
	listLanguageRequestFieldStatus   = big.NewInt(1 << 2)
)

type ListLanguageRequest struct {
	Cursor   *string `json:"-" url:"cursor,omitempty"`
	PageSize *int    `json:"-" url:"page_size,omitempty"`
	Status   *string `json:"-" url:"status,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListLanguageRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLanguageRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listLanguageRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLanguageRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listLanguageRequestFieldPageSize)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLanguageRequest) SetStatus(status *string) {
	l.Status = status
	l.require(listLanguageRequestFieldStatus)
}
