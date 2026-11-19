package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LanguageResponse struct {
	// The unique identifier of the language.
	LanguageId string `json:"language_id" xml:"language_id" required:"true"`
	// The name of the language.
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (l *LanguageResponse) GetLanguageId() (value string) {
	if l == nil {
		return
	}
	return l.LanguageId
}

func (l *LanguageResponse) GetName() (value string) {
	if l == nil {
		return
	}
	return l.Name
}

func (l *LanguageResponse) SetLanguageId(languageId string) {
	l.LanguageId = languageId
}

func (l *LanguageResponse) SetName(name string) {
	l.Name = name
}

func (l *LanguageResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LanguageResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LanguageResponse to string"
	}
	return string(jsonData)
}

func (l *LanguageResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LanguageResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LanguageResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LanguageResponse) MarshalJSON() ([]byte, error) {
	type embed LanguageResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
