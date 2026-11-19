package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PostDialDigitsDynamicVariable struct {
	// The dynamic variable name to resolve
	Value           string `json:"value" xml:"value" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (p *PostDialDigitsDynamicVariable) GetValue() (value string) {
	if p == nil {
		return
	}
	return p.Value
}

func (p *PostDialDigitsDynamicVariable) SetValue(value string) {
	p.Value = value
}

func (p *PostDialDigitsDynamicVariable) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PostDialDigitsDynamicVariable) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PostDialDigitsDynamicVariable to string"
	}
	return string(jsonData)
}

func (p *PostDialDigitsDynamicVariable) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PostDialDigitsDynamicVariable
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PostDialDigitsDynamicVariable(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PostDialDigitsDynamicVariable) MarshalJSON() ([]byte, error) {
	type embed PostDialDigitsDynamicVariable
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
