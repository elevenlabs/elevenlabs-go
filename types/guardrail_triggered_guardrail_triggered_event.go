package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GuardrailTriggeredGuardrailTriggeredEvent struct {
	GuardrailName   string `json:"guardrail_name" xml:"guardrail_name" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GuardrailTriggeredGuardrailTriggeredEvent) GetGuardrailName() (value string) {
	if g == nil {
		return
	}
	return g.GuardrailName
}

func (g *GuardrailTriggeredGuardrailTriggeredEvent) SetGuardrailName(guardrailName string) {
	g.GuardrailName = guardrailName
}

func (g *GuardrailTriggeredGuardrailTriggeredEvent) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GuardrailTriggeredGuardrailTriggeredEvent) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GuardrailTriggeredGuardrailTriggeredEvent to string"
	}
	return string(jsonData)
}

func (g *GuardrailTriggeredGuardrailTriggeredEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GuardrailTriggeredGuardrailTriggeredEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GuardrailTriggeredGuardrailTriggeredEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GuardrailTriggeredGuardrailTriggeredEvent) MarshalJSON() ([]byte, error) {
	type embed GuardrailTriggeredGuardrailTriggeredEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
