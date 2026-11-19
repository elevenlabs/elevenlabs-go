package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// User-specific secret model that are not shared with other users in a workspace.
type ConvAiUserSecretDbModel struct {
	Name            string `json:"name" xml:"name" required:"true"`
	EncryptedValue  string `json:"encrypted_value" xml:"encrypted_value" required:"true"`
	Nonce           string `json:"nonce" xml:"nonce" required:"true"`
	Id              string `json:"id" xml:"id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConvAiUserSecretDbModel) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *ConvAiUserSecretDbModel) GetEncryptedValue() (value string) {
	if c == nil {
		return
	}
	return c.EncryptedValue
}

func (c *ConvAiUserSecretDbModel) GetNonce() (value string) {
	if c == nil {
		return
	}
	return c.Nonce
}

func (c *ConvAiUserSecretDbModel) GetId() (value string) {
	if c == nil {
		return
	}
	return c.Id
}

func (c *ConvAiUserSecretDbModel) SetName(name string) {
	c.Name = name
}

func (c *ConvAiUserSecretDbModel) SetEncryptedValue(encryptedValue string) {
	c.EncryptedValue = encryptedValue
}

func (c *ConvAiUserSecretDbModel) SetNonce(nonce string) {
	c.Nonce = nonce
}

func (c *ConvAiUserSecretDbModel) SetId(id string) {
	c.Id = id
}

func (c *ConvAiUserSecretDbModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConvAiUserSecretDbModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConvAiUserSecretDbModel to string"
	}
	return string(jsonData)
}

func (c *ConvAiUserSecretDbModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConvAiUserSecretDbModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConvAiUserSecretDbModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConvAiUserSecretDbModel) MarshalJSON() ([]byte, error) {
	type embed ConvAiUserSecretDbModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
