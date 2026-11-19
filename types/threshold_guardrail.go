package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ThresholdGuardrail struct {
	IsEnabled       *bool    `json:"is_enabled,omitempty" xml:"is_enabled,omitempty"`
	Threshold       *float64 `json:"threshold,omitempty" xml:"threshold,omitempty" min:"0" max:"1"`
	extraProperties map[string]interface{}
}

func (t *ThresholdGuardrail) GetIsEnabled() (value *bool) {
	if t == nil {
		return
	}
	return t.IsEnabled
}

func (t *ThresholdGuardrail) GetThreshold() (value *float64) {
	if t == nil {
		return
	}
	return t.Threshold
}

func (t *ThresholdGuardrail) SetIsEnabled(isEnabled *bool) {
	t.IsEnabled = isEnabled
}

func (t *ThresholdGuardrail) SetThreshold(threshold *float64) {
	t.Threshold = threshold
}

func (t *ThresholdGuardrail) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t ThresholdGuardrail) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ThresholdGuardrail to string"
	}
	return string(jsonData)
}

func (t *ThresholdGuardrail) UnmarshalJSON(data []byte) error {
	type unmarshaler ThresholdGuardrail
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThresholdGuardrail(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *ThresholdGuardrail) MarshalJSON() ([]byte, error) {
	type embed ThresholdGuardrail
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
