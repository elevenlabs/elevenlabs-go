package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Used to reference a secret from the agent's secret store.
type ConvAiSecretLocator struct {
	SecretId        string `json:"secret_id" xml:"secret_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConvAiSecretLocator) GetSecretId() (value string) {
	if c == nil {
		return
	}
	return c.SecretId
}

func (c *ConvAiSecretLocator) SetSecretId(secretId string) {
	c.SecretId = secretId
}

func (c *ConvAiSecretLocator) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConvAiSecretLocator) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConvAiSecretLocator to string"
	}
	return string(jsonData)
}

func (c *ConvAiSecretLocator) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConvAiSecretLocator
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConvAiSecretLocator(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConvAiSecretLocator) MarshalJSON() ([]byte, error) {
	type embed ConvAiSecretLocator
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
