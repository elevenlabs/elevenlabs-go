package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type PromptInjectionGuardrail struct {
	IsEnabled       *bool `json:"is_enabled,omitempty" xml:"is_enabled,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PromptInjectionGuardrail) GetIsEnabled() (value *bool) {
	if p == nil {
		return
	}
	return p.IsEnabled
}

func (p *PromptInjectionGuardrail) SetIsEnabled(isEnabled *bool) {
	p.IsEnabled = isEnabled
}

func (p *PromptInjectionGuardrail) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PromptInjectionGuardrail) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PromptInjectionGuardrail to string"
	}
	return string(jsonData)
}

func (p *PromptInjectionGuardrail) UnmarshalJSON(data []byte) error {
	type unmarshaler PromptInjectionGuardrail
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PromptInjectionGuardrail(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PromptInjectionGuardrail) MarshalJSON() ([]byte, error) {
	type embed PromptInjectionGuardrail
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
