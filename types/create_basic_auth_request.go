package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request model for creating Basic Auth connections - inherits common settings and includes sensitive fields
type CreateBasicAuthRequest struct {
	Name            string  `json:"name" xml:"name" required:"true"`
	AuthType        *string `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	Provider        string  `json:"provider" xml:"provider" required:"true"`
	Username        string  `json:"username" xml:"username" required:"true"`
	Password        string  `json:"password" xml:"password" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateBasicAuthRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateBasicAuthRequest) GetAuthType() (value *string) {
	if c == nil {
		return
	}
	return c.AuthType
}

func (c *CreateBasicAuthRequest) GetProvider() (value string) {
	if c == nil {
		return
	}
	return c.Provider
}

func (c *CreateBasicAuthRequest) GetUsername() (value string) {
	if c == nil {
		return
	}
	return c.Username
}

func (c *CreateBasicAuthRequest) GetPassword() (value string) {
	if c == nil {
		return
	}
	return c.Password
}

func (c *CreateBasicAuthRequest) SetName(name string) {
	c.Name = name
}

func (c *CreateBasicAuthRequest) SetAuthType(authType *string) {
	c.AuthType = authType
}

func (c *CreateBasicAuthRequest) SetProvider(provider string) {
	c.Provider = provider
}

func (c *CreateBasicAuthRequest) SetUsername(username string) {
	c.Username = username
}

func (c *CreateBasicAuthRequest) SetPassword(password string) {
	c.Password = password
}

func (c *CreateBasicAuthRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateBasicAuthRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateBasicAuthRequest to string"
	}
	return string(jsonData)
}

func (c *CreateBasicAuthRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateBasicAuthRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateBasicAuthRequest(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateBasicAuthRequest) MarshalJSON() ([]byte, error) {
	type embed CreateBasicAuthRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
