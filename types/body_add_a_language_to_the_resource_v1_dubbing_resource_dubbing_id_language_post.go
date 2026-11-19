package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyAddALanguageToTheResourceV1DubbingResourceDubbingIdLanguagePost struct {
	// The Target language.
	Language        *string `json:"language" xml:"language" required:"true"`
	extraProperties map[string]interface{}
}

func (b *BodyAddALanguageToTheResourceV1DubbingResourceDubbingIdLanguagePost) GetLanguage() (value *string) {
	if b == nil {
		return
	}
	return b.Language
}

func (b *BodyAddALanguageToTheResourceV1DubbingResourceDubbingIdLanguagePost) SetLanguage(language *string) {
	b.Language = language
}

func (b *BodyAddALanguageToTheResourceV1DubbingResourceDubbingIdLanguagePost) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BodyAddALanguageToTheResourceV1DubbingResourceDubbingIdLanguagePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyAddALanguageToTheResourceV1DubbingResourceDubbingIdLanguagePost to string"
	}
	return string(jsonData)
}

func (b *BodyAddALanguageToTheResourceV1DubbingResourceDubbingIdLanguagePost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BodyAddALanguageToTheResourceV1DubbingResourceDubbingIdLanguagePost) MarshalJSON() ([]byte, error) {
	type embed BodyAddALanguageToTheResourceV1DubbingResourceDubbingIdLanguagePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
