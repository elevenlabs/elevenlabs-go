package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DynamicVariableSchemaOverride struct {
	// The name of the dynamic variable to use
	DynamicVariable string `json:"dynamic_variable" xml:"dynamic_variable" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DynamicVariableSchemaOverride) GetDynamicVariable() (value string) {
	if d == nil {
		return
	}
	return d.DynamicVariable
}

func (d *DynamicVariableSchemaOverride) SetDynamicVariable(dynamicVariable string) {
	d.DynamicVariable = dynamicVariable
}

func (d *DynamicVariableSchemaOverride) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DynamicVariableSchemaOverride) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DynamicVariableSchemaOverride to string"
	}
	return string(jsonData)
}

func (d *DynamicVariableSchemaOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DynamicVariableSchemaOverride
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DynamicVariableSchemaOverride(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "source")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DynamicVariableSchemaOverride) MarshalJSON() ([]byte, error) {
	type embed DynamicVariableSchemaOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
