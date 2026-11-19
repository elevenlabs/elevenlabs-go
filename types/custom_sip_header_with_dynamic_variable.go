package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Custom SIP header for phone transfers with a dynamic variable reference.
// The value is a variable name that will be resolved at runtime.
// Value is not validated here since it will be substituted with actual value later.
type CustomSipHeaderWithDynamicVariable struct {
	// The dynamic variable name to resolve
	Value string `json:"value" xml:"value" required:"true" minLength:"1"`
	type_ string
	// The SIP header name (e.g., 'X-Customer-ID')
	Key             string `json:"key" xml:"key" required:"true" maxLength:"64" minLength:"1" pattern:"^[A-Za-z0-9]+(-[A-Za-z0-9]+)*$"`
	extraProperties map[string]interface{}
}

func (c *CustomSipHeaderWithDynamicVariable) GetValue() (value string) {
	if c == nil {
		return
	}
	return c.Value
}

func (c *CustomSipHeaderWithDynamicVariable) GetKey() (value string) {
	if c == nil {
		return
	}
	return c.Key
}

func (c *CustomSipHeaderWithDynamicVariable) SetValue(value string) {
	c.Value = value
}

func (c *CustomSipHeaderWithDynamicVariable) SetKey(key string) {
	c.Key = key
}

func (c *CustomSipHeaderWithDynamicVariable) Type() string {
	return c.type_
}

func (c *CustomSipHeaderWithDynamicVariable) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CustomSipHeaderWithDynamicVariable) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CustomSipHeaderWithDynamicVariable to string"
	}
	return string(jsonData)
}

func (c *CustomSipHeaderWithDynamicVariable) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CustomSipHeaderWithDynamicVariable
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CustomSipHeaderWithDynamicVariable(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "dynamic" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "dynamic", *literals.Type)
	}
	c.type_ = "dynamic"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CustomSipHeaderWithDynamicVariable) MarshalJSON() ([]byte, error) {
	type embed CustomSipHeaderWithDynamicVariable
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "dynamic",
	}
	return json.Marshal(marshaler)
}
