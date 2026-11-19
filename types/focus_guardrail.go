package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type FocusGuardrail struct {
	IsEnabled       *bool `json:"is_enabled,omitempty" xml:"is_enabled,omitempty"`
	extraProperties map[string]interface{}
}

func (f *FocusGuardrail) GetIsEnabled() (value *bool) {
	if f == nil {
		return
	}
	return f.IsEnabled
}

func (f *FocusGuardrail) SetIsEnabled(isEnabled *bool) {
	f.IsEnabled = isEnabled
}

func (f *FocusGuardrail) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FocusGuardrail) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FocusGuardrail to string"
	}
	return string(jsonData)
}

func (f *FocusGuardrail) UnmarshalJSON(data []byte) error {
	type unmarshaler FocusGuardrail
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FocusGuardrail(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FocusGuardrail) MarshalJSON() ([]byte, error) {
	type embed FocusGuardrail
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
