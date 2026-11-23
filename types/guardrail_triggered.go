package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GuardrailTriggered struct {
	GuardrailTriggeredEvent *GuardrailTriggeredGuardrailTriggeredEvent `json:"guardrail_triggered_event" xml:"guardrail_triggered_event" required:"true"`
	Type                    *string                                    `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties         map[string]interface{}
}

func (g *GuardrailTriggered) GetGuardrailTriggeredEvent() (value *GuardrailTriggeredGuardrailTriggeredEvent) {
	if g == nil {
		return
	}
	return g.GuardrailTriggeredEvent
}

func (g *GuardrailTriggered) GetType() (value *string) {
	if g == nil {
		return
	}
	return g.Type
}

func (g *GuardrailTriggered) SetGuardrailTriggeredEvent(guardrailTriggeredEvent *GuardrailTriggeredGuardrailTriggeredEvent) {
	g.GuardrailTriggeredEvent = guardrailTriggeredEvent
}

func (g *GuardrailTriggered) SetType(type_ *string) {
	g.Type = type_
}

func (g *GuardrailTriggered) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GuardrailTriggered) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GuardrailTriggered to string"
	}
	return string(jsonData)
}

func (g *GuardrailTriggered) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GuardrailTriggered
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GuardrailTriggered(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GuardrailTriggered) MarshalJSON() ([]byte, error) {
	type embed GuardrailTriggered
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
