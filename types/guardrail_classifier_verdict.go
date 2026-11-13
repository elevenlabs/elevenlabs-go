package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GuardrailClassifierVerdict struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Classifier score for the triggered category, from 0 to 1.
	Score float64 `json:"score" xml:"score" required:"true"`
	// Score at or above which the guardrail triggers.
	Threshold       float64 `json:"threshold" xml:"threshold" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GuardrailClassifierVerdict) GetType() (value *string) {
	if g == nil {
		return
	}
	return g.Type
}

func (g *GuardrailClassifierVerdict) GetScore() (value float64) {
	if g == nil {
		return
	}
	return g.Score
}

func (g *GuardrailClassifierVerdict) GetThreshold() (value float64) {
	if g == nil {
		return
	}
	return g.Threshold
}

func (g *GuardrailClassifierVerdict) SetType(type_ *string) {
	g.Type = type_
}

func (g *GuardrailClassifierVerdict) SetScore(score float64) {
	g.Score = score
}

func (g *GuardrailClassifierVerdict) SetThreshold(threshold float64) {
	g.Threshold = threshold
}

func (g *GuardrailClassifierVerdict) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GuardrailClassifierVerdict) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GuardrailClassifierVerdict to string"
	}
	return string(jsonData)
}

func (g *GuardrailClassifierVerdict) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GuardrailClassifierVerdict
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GuardrailClassifierVerdict(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GuardrailClassifierVerdict) MarshalJSON() ([]byte, error) {
	type embed GuardrailClassifierVerdict
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
