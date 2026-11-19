package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ClientSecretResponse struct {
	Value           string                     `json:"value" xml:"value" required:"true"`
	ExpiresAt       int                        `json:"expires_at" xml:"expires_at" required:"true"`
	Session         *OpenAiSessionConfigOutput `json:"session,omitempty" xml:"session,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ClientSecretResponse) GetValue() (value string) {
	if c == nil {
		return
	}
	return c.Value
}

func (c *ClientSecretResponse) GetExpiresAt() (value int) {
	if c == nil {
		return
	}
	return c.ExpiresAt
}

func (c *ClientSecretResponse) GetSession() (value *OpenAiSessionConfigOutput) {
	if c == nil {
		return
	}
	return c.Session
}

func (c *ClientSecretResponse) SetValue(value string) {
	c.Value = value
}

func (c *ClientSecretResponse) SetExpiresAt(expiresAt int) {
	c.ExpiresAt = expiresAt
}

func (c *ClientSecretResponse) SetSession(session *OpenAiSessionConfigOutput) {
	c.Session = session
}

func (c *ClientSecretResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClientSecretResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClientSecretResponse to string"
	}
	return string(jsonData)
}

func (c *ClientSecretResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClientSecretResponse) MarshalJSON() ([]byte, error) {
	type embed ClientSecretResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
