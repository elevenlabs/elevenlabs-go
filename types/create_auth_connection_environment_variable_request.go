package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateAuthConnectionEnvironmentVariableRequest struct {
	type_ string
	// Unique label for the environment variable.
	Label string `json:"label" xml:"label" required:"true"`
	// Environment-specific auth connection references. Must include 'production' key.
	Values          map[string]*EnvironmentVariableAuthConnectionValueRequest `json:"values" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateAuthConnectionEnvironmentVariableRequest) GetLabel() (value string) {
	if c == nil {
		return
	}
	return c.Label
}

func (c *CreateAuthConnectionEnvironmentVariableRequest) GetValues() (value map[string]*EnvironmentVariableAuthConnectionValueRequest) {
	if c == nil {
		return
	}
	return c.Values
}

func (c *CreateAuthConnectionEnvironmentVariableRequest) SetLabel(label string) {
	c.Label = label
}

func (c *CreateAuthConnectionEnvironmentVariableRequest) SetValues(values map[string]*EnvironmentVariableAuthConnectionValueRequest) {
	c.Values = values
}

func (c *CreateAuthConnectionEnvironmentVariableRequest) Type() string {
	return c.type_
}

func (c *CreateAuthConnectionEnvironmentVariableRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateAuthConnectionEnvironmentVariableRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateAuthConnectionEnvironmentVariableRequest to string"
	}
	return string(jsonData)
}

func (c *CreateAuthConnectionEnvironmentVariableRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateAuthConnectionEnvironmentVariableRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateAuthConnectionEnvironmentVariableRequest(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "auth_connection" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "auth_connection", *literals.Type)
	}
	c.type_ = "auth_connection"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateAuthConnectionEnvironmentVariableRequest) MarshalJSON() ([]byte, error) {
	type embed CreateAuthConnectionEnvironmentVariableRequest
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "auth_connection",
	}
	return json.Marshal(marshaler)
}
