package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Used to reference an environment variable by label.
type ConvAiEnvVarLocator struct {
	EnvVarLabel     string `json:"env_var_label" xml:"env_var_label" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConvAiEnvVarLocator) GetEnvVarLabel() (value string) {
	if c == nil {
		return
	}
	return c.EnvVarLabel
}

func (c *ConvAiEnvVarLocator) SetEnvVarLabel(envVarLabel string) {
	c.EnvVarLabel = envVarLabel
}

func (c *ConvAiEnvVarLocator) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConvAiEnvVarLocator) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConvAiEnvVarLocator to string"
	}
	return string(jsonData)
}

func (c *ConvAiEnvVarLocator) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConvAiEnvVarLocator
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConvAiEnvVarLocator(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConvAiEnvVarLocator) MarshalJSON() ([]byte, error) {
	type embed ConvAiEnvVarLocator
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
