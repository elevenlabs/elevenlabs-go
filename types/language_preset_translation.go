package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LanguagePresetTranslation struct {
	SourceHash      string `json:"source_hash" xml:"source_hash" required:"true"`
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (l *LanguagePresetTranslation) GetSourceHash() (value string) {
	if l == nil {
		return
	}
	return l.SourceHash
}

func (l *LanguagePresetTranslation) GetText() (value string) {
	if l == nil {
		return
	}
	return l.Text
}

func (l *LanguagePresetTranslation) SetSourceHash(sourceHash string) {
	l.SourceHash = sourceHash
}

func (l *LanguagePresetTranslation) SetText(text string) {
	l.Text = text
}

func (l *LanguagePresetTranslation) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LanguagePresetTranslation) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LanguagePresetTranslation to string"
	}
	return string(jsonData)
}

func (l *LanguagePresetTranslation) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LanguagePresetTranslation
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LanguagePresetTranslation(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LanguagePresetTranslation) MarshalJSON() ([]byte, error) {
	type embed LanguagePresetTranslation
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
