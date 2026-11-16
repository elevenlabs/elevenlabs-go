package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LanguagePairInfo struct {
	SourceLanguage *LanguageInfo `json:"source_language" xml:"source_language" required:"true"`
	// The available destination languages for this source language.
	DestinationLanguages []*LanguageInfo `json:"destination_languages" xml:"destination_languages" required:"true"`
	extraProperties      map[string]interface{}
}

func (l *LanguagePairInfo) GetSourceLanguage() (value *LanguageInfo) {
	if l == nil {
		return
	}
	return l.SourceLanguage
}

func (l *LanguagePairInfo) GetDestinationLanguages() (value []*LanguageInfo) {
	if l == nil {
		return
	}
	return l.DestinationLanguages
}

func (l *LanguagePairInfo) SetSourceLanguage(sourceLanguage *LanguageInfo) {
	l.SourceLanguage = sourceLanguage
}

func (l *LanguagePairInfo) SetDestinationLanguages(destinationLanguages []*LanguageInfo) {
	l.DestinationLanguages = destinationLanguages
}

func (l *LanguagePairInfo) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LanguagePairInfo) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LanguagePairInfo to string"
	}
	return string(jsonData)
}

func (l *LanguagePairInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LanguagePairInfo
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LanguagePairInfo(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LanguagePairInfo) MarshalJSON() ([]byte, error) {
	type embed LanguagePairInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
