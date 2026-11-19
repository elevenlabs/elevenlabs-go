package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GuardrailsV1Output struct {
	Version         *string                   `json:"version,omitempty" xml:"version,omitempty"`
	Focus           *FocusGuardrail           `json:"focus,omitempty" xml:"focus,omitempty"`
	PromptInjection *PromptInjectionGuardrail `json:"prompt_injection,omitempty" xml:"prompt_injection,omitempty"`
	Content         *ContentGuardrailOutput   `json:"content,omitempty" xml:"content,omitempty"`
	// Container for custom guardrails, matching ModerationGuardrail pattern
	Custom          *CustomGuardrailOutput `json:"custom,omitempty" xml:"custom,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GuardrailsV1Output) GetVersion() (value *string) {
	if g == nil {
		return
	}
	return g.Version
}

func (g *GuardrailsV1Output) GetFocus() (value *FocusGuardrail) {
	if g == nil {
		return
	}
	return g.Focus
}

func (g *GuardrailsV1Output) GetPromptInjection() (value *PromptInjectionGuardrail) {
	if g == nil {
		return
	}
	return g.PromptInjection
}

func (g *GuardrailsV1Output) GetContent() (value *ContentGuardrailOutput) {
	if g == nil {
		return
	}
	return g.Content
}

func (g *GuardrailsV1Output) GetCustom() (value *CustomGuardrailOutput) {
	if g == nil {
		return
	}
	return g.Custom
}

func (g *GuardrailsV1Output) SetVersion(version *string) {
	g.Version = version
}

func (g *GuardrailsV1Output) SetFocus(focus *FocusGuardrail) {
	g.Focus = focus
}

func (g *GuardrailsV1Output) SetPromptInjection(promptInjection *PromptInjectionGuardrail) {
	g.PromptInjection = promptInjection
}

func (g *GuardrailsV1Output) SetContent(content *ContentGuardrailOutput) {
	g.Content = content
}

func (g *GuardrailsV1Output) SetCustom(custom *CustomGuardrailOutput) {
	g.Custom = custom
}

func (g *GuardrailsV1Output) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GuardrailsV1Output) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GuardrailsV1Output to string"
	}
	return string(jsonData)
}

func (g *GuardrailsV1Output) UnmarshalJSON(data []byte) error {
	type unmarshaler GuardrailsV1Output
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuardrailsV1Output(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GuardrailsV1Output) MarshalJSON() ([]byte, error) {
	type embed GuardrailsV1Output
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
