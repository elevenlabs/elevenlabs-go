package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LanguageInfo struct {
	// The language code (e.g. 'en', 'fr', 'es-ES').
	Code string `json:"code" xml:"code" required:"true"`
	// The human-readable language name (e.g. 'English', 'French', 'Spanish (Spain)').
	Label           string `json:"label" xml:"label" required:"true"`
	extraProperties map[string]interface{}
}

func (l *LanguageInfo) GetCode() (value string) {
	if l == nil {
		return
	}
	return l.Code
}

func (l *LanguageInfo) GetLabel() (value string) {
	if l == nil {
		return
	}
	return l.Label
}

func (l *LanguageInfo) SetCode(code string) {
	l.Code = code
}

func (l *LanguageInfo) SetLabel(label string) {
	l.Label = label
}

func (l *LanguageInfo) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LanguageInfo) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LanguageInfo to string"
	}
	return string(jsonData)
}

func (l *LanguageInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LanguageInfo
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LanguageInfo(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LanguageInfo) MarshalJSON() ([]byte, error) {
	type embed LanguageInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
