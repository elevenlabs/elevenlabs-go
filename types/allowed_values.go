package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AllowedValues struct {
	// Name of a dynamic variable that must resolve to a JSON array of permitted values, e.g. ["ws_alpha", "ws_beta"]. System variables work only if they resolve to a list.
	DynamicVariable string `json:"dynamic_variable" xml:"dynamic_variable" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (a *AllowedValues) GetDynamicVariable() (value string) {
	if a == nil {
		return
	}
	return a.DynamicVariable
}

func (a *AllowedValues) SetDynamicVariable(dynamicVariable string) {
	a.DynamicVariable = dynamicVariable
}

func (a *AllowedValues) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AllowedValues) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AllowedValues to string"
	}
	return string(jsonData)
}

func (a *AllowedValues) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AllowedValues
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AllowedValues(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AllowedValues) MarshalJSON() ([]byte, error) {
	type embed AllowedValues
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
