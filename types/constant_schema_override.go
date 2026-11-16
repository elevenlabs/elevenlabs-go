package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConstantSchemaOverride struct {
	// The constant value to use
	ConstantValue   *ConstantSchemaOverrideConstantValue `json:"constant_value" xml:"constant_value" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConstantSchemaOverride) GetConstantValue() (value *ConstantSchemaOverrideConstantValue) {
	if c == nil {
		return
	}
	return c.ConstantValue
}

func (c *ConstantSchemaOverride) SetConstantValue(constantValue *ConstantSchemaOverrideConstantValue) {
	c.ConstantValue = constantValue
}

func (c *ConstantSchemaOverride) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConstantSchemaOverride) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConstantSchemaOverride to string"
	}
	return string(jsonData)
}

func (c *ConstantSchemaOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "source")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConstantSchemaOverride) MarshalJSON() ([]byte, error) {
	type embed ConstantSchemaOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
