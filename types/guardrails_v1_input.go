package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GuardrailsV1Input struct {
	Version         *string                   `json:"version,omitempty" xml:"version,omitempty"`
	Focus           *FocusGuardrail           `json:"focus,omitempty" xml:"focus,omitempty"`
	PromptInjection *PromptInjectionGuardrail `json:"prompt_injection,omitempty" xml:"prompt_injection,omitempty"`
	Content         *ContentGuardrailInput    `json:"content,omitempty" xml:"content,omitempty"`
	// Container for custom guardrails, matching ModerationGuardrail pattern
	Custom          *CustomGuardrailInput `json:"custom,omitempty" xml:"custom,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GuardrailsV1Input) GetVersion() (value *string) {
	if g == nil {
		return
	}
	return g.Version
}

func (g *GuardrailsV1Input) GetFocus() (value *FocusGuardrail) {
	if g == nil {
		return
	}
	return g.Focus
}

func (g *GuardrailsV1Input) GetPromptInjection() (value *PromptInjectionGuardrail) {
	if g == nil {
		return
	}
	return g.PromptInjection
}

func (g *GuardrailsV1Input) GetContent() (value *ContentGuardrailInput) {
	if g == nil {
		return
	}
	return g.Content
}

func (g *GuardrailsV1Input) GetCustom() (value *CustomGuardrailInput) {
	if g == nil {
		return
	}
	return g.Custom
}

func (g *GuardrailsV1Input) SetVersion(version *string) {
	g.Version = version
}

func (g *GuardrailsV1Input) SetFocus(focus *FocusGuardrail) {
	g.Focus = focus
}

func (g *GuardrailsV1Input) SetPromptInjection(promptInjection *PromptInjectionGuardrail) {
	g.PromptInjection = promptInjection
}

func (g *GuardrailsV1Input) SetContent(content *ContentGuardrailInput) {
	g.Content = content
}

func (g *GuardrailsV1Input) SetCustom(custom *CustomGuardrailInput) {
	g.Custom = custom
}

func (g *GuardrailsV1Input) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GuardrailsV1Input) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GuardrailsV1Input to string"
	}
	return string(jsonData)
}

func (g *GuardrailsV1Input) UnmarshalJSON(data []byte) error {
	type unmarshaler GuardrailsV1Input
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuardrailsV1Input(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GuardrailsV1Input) MarshalJSON() ([]byte, error) {
	type embed GuardrailsV1Input
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
