package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConvAiWorkspaceStoredSecretConfig struct {
	type_           string
	SecretId        string                          `json:"secret_id" xml:"secret_id" required:"true"`
	Name            string                          `json:"name" xml:"name" required:"true"`
	UsedBy          *ConvAiStoredSecretDependencies `json:"used_by" xml:"used_by" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConvAiWorkspaceStoredSecretConfig) GetSecretId() (value string) {
	if c == nil {
		return
	}
	return c.SecretId
}

func (c *ConvAiWorkspaceStoredSecretConfig) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *ConvAiWorkspaceStoredSecretConfig) GetUsedBy() (value *ConvAiStoredSecretDependencies) {
	if c == nil {
		return
	}
	return c.UsedBy
}

func (c *ConvAiWorkspaceStoredSecretConfig) SetSecretId(secretId string) {
	c.SecretId = secretId
}

func (c *ConvAiWorkspaceStoredSecretConfig) SetName(name string) {
	c.Name = name
}

func (c *ConvAiWorkspaceStoredSecretConfig) SetUsedBy(usedBy *ConvAiStoredSecretDependencies) {
	c.UsedBy = usedBy
}

func (c *ConvAiWorkspaceStoredSecretConfig) Type() string {
	return c.type_
}

func (c *ConvAiWorkspaceStoredSecretConfig) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConvAiWorkspaceStoredSecretConfig) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConvAiWorkspaceStoredSecretConfig to string"
	}
	return string(jsonData)
}

func (c *ConvAiWorkspaceStoredSecretConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConvAiWorkspaceStoredSecretConfig
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConvAiWorkspaceStoredSecretConfig(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "stored" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "stored", *literals.Type)
	}
	c.type_ = "stored"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConvAiWorkspaceStoredSecretConfig) MarshalJSON() ([]byte, error) {
	type embed ConvAiWorkspaceStoredSecretConfig
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "stored",
	}
	return json.Marshal(marshaler)
}
