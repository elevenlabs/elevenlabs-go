package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request model for creating Bearer Auth connections - inherits common settings and includes sensitive fields
type CreateBearerAuthRequest struct {
	Name            string  `json:"name" xml:"name" required:"true"`
	AuthType        *string `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	Provider        string  `json:"provider" xml:"provider" required:"true"`
	Token           string  `json:"token" xml:"token" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateBearerAuthRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateBearerAuthRequest) GetAuthType() (value *string) {
	if c == nil {
		return
	}
	return c.AuthType
}

func (c *CreateBearerAuthRequest) GetProvider() (value string) {
	if c == nil {
		return
	}
	return c.Provider
}

func (c *CreateBearerAuthRequest) GetToken() (value string) {
	if c == nil {
		return
	}
	return c.Token
}

func (c *CreateBearerAuthRequest) SetName(name string) {
	c.Name = name
}

func (c *CreateBearerAuthRequest) SetAuthType(authType *string) {
	c.AuthType = authType
}

func (c *CreateBearerAuthRequest) SetProvider(provider string) {
	c.Provider = provider
}

func (c *CreateBearerAuthRequest) SetToken(token string) {
	c.Token = token
}

func (c *CreateBearerAuthRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateBearerAuthRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateBearerAuthRequest to string"
	}
	return string(jsonData)
}

func (c *CreateBearerAuthRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateBearerAuthRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateBearerAuthRequest(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateBearerAuthRequest) MarshalJSON() ([]byte, error) {
	type embed CreateBearerAuthRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
