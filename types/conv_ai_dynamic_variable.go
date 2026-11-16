package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Used to reference a dynamic variable.
type ConvAiDynamicVariable struct {
	VariableName    string `json:"variable_name" xml:"variable_name" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConvAiDynamicVariable) GetVariableName() (value string) {
	if c == nil {
		return
	}
	return c.VariableName
}

func (c *ConvAiDynamicVariable) SetVariableName(variableName string) {
	c.VariableName = variableName
}

func (c *ConvAiDynamicVariable) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConvAiDynamicVariable) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConvAiDynamicVariable to string"
	}
	return string(jsonData)
}

func (c *ConvAiDynamicVariable) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConvAiDynamicVariable
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConvAiDynamicVariable(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConvAiDynamicVariable) MarshalJSON() ([]byte, error) {
	type embed ConvAiDynamicVariable
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
