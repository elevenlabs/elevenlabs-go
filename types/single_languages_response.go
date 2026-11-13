package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SingleLanguagesResponse struct {
	// Indicates this response contains single languages (not source-to-destination pairs).
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The list of available languages.
	Languages       []*LanguageInfo `json:"languages" xml:"languages" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SingleLanguagesResponse) GetKind() (value *string) {
	if s == nil {
		return
	}
	return s.Kind
}

func (s *SingleLanguagesResponse) GetLanguages() (value []*LanguageInfo) {
	if s == nil {
		return
	}
	return s.Languages
}

func (s *SingleLanguagesResponse) SetKind(kind *string) {
	s.Kind = kind
}

func (s *SingleLanguagesResponse) SetLanguages(languages []*LanguageInfo) {
	s.Languages = languages
}

func (s *SingleLanguagesResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SingleLanguagesResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SingleLanguagesResponse to string"
	}
	return string(jsonData)
}

func (s *SingleLanguagesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SingleLanguagesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SingleLanguagesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SingleLanguagesResponse) MarshalJSON() ([]byte, error) {
	type embed SingleLanguagesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
