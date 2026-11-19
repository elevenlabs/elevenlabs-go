package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateSecretEnvironmentVariableRequest struct {
	type_ string
	// Unique label for the environment variable.
	Label string `json:"label" xml:"label" required:"true"`
	// Environment-specific secret references. Must include 'production' key.
	Values          map[string]*EnvironmentVariableSecretValueRequest `json:"values" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateSecretEnvironmentVariableRequest) GetLabel() (value string) {
	if c == nil {
		return
	}
	return c.Label
}

func (c *CreateSecretEnvironmentVariableRequest) GetValues() (value map[string]*EnvironmentVariableSecretValueRequest) {
	if c == nil {
		return
	}
	return c.Values
}

func (c *CreateSecretEnvironmentVariableRequest) SetLabel(label string) {
	c.Label = label
}

func (c *CreateSecretEnvironmentVariableRequest) SetValues(values map[string]*EnvironmentVariableSecretValueRequest) {
	c.Values = values
}

func (c *CreateSecretEnvironmentVariableRequest) Type() string {
	return c.type_
}

func (c *CreateSecretEnvironmentVariableRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateSecretEnvironmentVariableRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateSecretEnvironmentVariableRequest to string"
	}
	return string(jsonData)
}

func (c *CreateSecretEnvironmentVariableRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateSecretEnvironmentVariableRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateSecretEnvironmentVariableRequest(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "secret" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "secret", *literals.Type)
	}
	c.type_ = "secret"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateSecretEnvironmentVariableRequest) MarshalJSON() ([]byte, error) {
	type embed CreateSecretEnvironmentVariableRequest
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "secret",
	}
	return json.Marshal(marshaler)
}
