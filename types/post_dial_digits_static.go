package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PostDialDigitsStatic struct {
	// DTMF digits to send after call connects (e.g., 'ww1234' for extension)
	Value           string `json:"value" xml:"value" required:"true" minLength:"1" pattern:"^[0-9*#wW]*$"`
	extraProperties map[string]interface{}
}

func (p *PostDialDigitsStatic) GetValue() (value string) {
	if p == nil {
		return
	}
	return p.Value
}

func (p *PostDialDigitsStatic) SetValue(value string) {
	p.Value = value
}

func (p *PostDialDigitsStatic) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PostDialDigitsStatic) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PostDialDigitsStatic to string"
	}
	return string(jsonData)
}

func (p *PostDialDigitsStatic) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PostDialDigitsStatic
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PostDialDigitsStatic(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PostDialDigitsStatic) MarshalJSON() ([]byte, error) {
	type embed PostDialDigitsStatic
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
