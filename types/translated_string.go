package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslatedString struct {
	Value           string `json:"value" xml:"value" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (t *TranslatedString) GetValue() (value string) {
	if t == nil {
		return
	}
	return t.Value
}

func (t *TranslatedString) SetValue(value string) {
	t.Value = value
}

func (t *TranslatedString) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslatedString) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslatedString to string"
	}
	return string(jsonData)
}

func (t *TranslatedString) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslatedString
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslatedString(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslatedString) MarshalJSON() ([]byte, error) {
	type embed TranslatedString
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
