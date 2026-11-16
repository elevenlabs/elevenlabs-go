package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GuardrailsResult struct {
	Triggered       *bool `json:"triggered,omitempty" xml:"triggered,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GuardrailsResult) GetTriggered() (value *bool) {
	if g == nil {
		return
	}
	return g.Triggered
}

func (g *GuardrailsResult) SetTriggered(triggered *bool) {
	g.Triggered = triggered
}

func (g *GuardrailsResult) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GuardrailsResult) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GuardrailsResult to string"
	}
	return string(jsonData)
}

func (g *GuardrailsResult) UnmarshalJSON(data []byte) error {
	type unmarshaler GuardrailsResult
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuardrailsResult(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GuardrailsResult) MarshalJSON() ([]byte, error) {
	type embed GuardrailsResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
