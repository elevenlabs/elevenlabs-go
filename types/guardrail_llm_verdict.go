package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GuardrailLlmVerdict struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The LLM's explanation for its decision.
	Rationale       string `json:"rationale" xml:"rationale" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GuardrailLlmVerdict) GetType() (value *string) {
	if g == nil {
		return
	}
	return g.Type
}

func (g *GuardrailLlmVerdict) GetRationale() (value string) {
	if g == nil {
		return
	}
	return g.Rationale
}

func (g *GuardrailLlmVerdict) SetType(type_ *string) {
	g.Type = type_
}

func (g *GuardrailLlmVerdict) SetRationale(rationale string) {
	g.Rationale = rationale
}

func (g *GuardrailLlmVerdict) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GuardrailLlmVerdict) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GuardrailLlmVerdict to string"
	}
	return string(jsonData)
}

func (g *GuardrailLlmVerdict) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GuardrailLlmVerdict
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GuardrailLlmVerdict(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GuardrailLlmVerdict) MarshalJSON() ([]byte, error) {
	type embed GuardrailLlmVerdict
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
