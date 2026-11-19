package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateStringEnvironmentVariableRequest struct {
	type_ string
	// Unique label for the environment variable.
	Label string `json:"label" xml:"label" required:"true"`
	// Environment-specific values. Must include 'production' key.
	Values          map[string]string `json:"values" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateStringEnvironmentVariableRequest) GetLabel() (value string) {
	if c == nil {
		return
	}
	return c.Label
}

func (c *CreateStringEnvironmentVariableRequest) GetValues() (value map[string]string) {
	if c == nil {
		return
	}
	return c.Values
}

func (c *CreateStringEnvironmentVariableRequest) SetLabel(label string) {
	c.Label = label
}

func (c *CreateStringEnvironmentVariableRequest) SetValues(values map[string]string) {
	c.Values = values
}

func (c *CreateStringEnvironmentVariableRequest) Type() string {
	return c.type_
}

func (c *CreateStringEnvironmentVariableRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateStringEnvironmentVariableRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateStringEnvironmentVariableRequest to string"
	}
	return string(jsonData)
}

func (c *CreateStringEnvironmentVariableRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateStringEnvironmentVariableRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateStringEnvironmentVariableRequest(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "string" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "string", *literals.Type)
	}
	c.type_ = "string"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateStringEnvironmentVariableRequest) MarshalJSON() ([]byte, error) {
	type embed CreateStringEnvironmentVariableRequest
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "string",
	}
	return json.Marshal(marshaler)
}
