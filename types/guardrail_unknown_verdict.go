package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GuardrailUnknownVerdict struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GuardrailUnknownVerdict) GetType() (value *string) {
	if g == nil {
		return
	}
	return g.Type
}

func (g *GuardrailUnknownVerdict) SetType(type_ *string) {
	g.Type = type_
}

func (g *GuardrailUnknownVerdict) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GuardrailUnknownVerdict) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GuardrailUnknownVerdict to string"
	}
	return string(jsonData)
}

func (g *GuardrailUnknownVerdict) UnmarshalJSON(data []byte) error {
	type unmarshaler GuardrailUnknownVerdict
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuardrailUnknownVerdict(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GuardrailUnknownVerdict) MarshalJSON() ([]byte, error) {
	type embed GuardrailUnknownVerdict
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
