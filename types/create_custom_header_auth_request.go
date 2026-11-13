package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateCustomHeaderAuthRequest struct {
	Name     string  `json:"name" xml:"name" required:"true"`
	AuthType *string `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	Provider string  `json:"provider" xml:"provider" required:"true"`
	// The name of the header to use for authentication (e.g., 'x-api-key')
	HeaderName      string `json:"header_name" xml:"header_name" required:"true"`
	Token           string `json:"token" xml:"token" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateCustomHeaderAuthRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateCustomHeaderAuthRequest) GetAuthType() (value *string) {
	if c == nil {
		return
	}
	return c.AuthType
}

func (c *CreateCustomHeaderAuthRequest) GetProvider() (value string) {
	if c == nil {
		return
	}
	return c.Provider
}

func (c *CreateCustomHeaderAuthRequest) GetHeaderName() (value string) {
	if c == nil {
		return
	}
	return c.HeaderName
}

func (c *CreateCustomHeaderAuthRequest) GetToken() (value string) {
	if c == nil {
		return
	}
	return c.Token
}

func (c *CreateCustomHeaderAuthRequest) SetName(name string) {
	c.Name = name
}

func (c *CreateCustomHeaderAuthRequest) SetAuthType(authType *string) {
	c.AuthType = authType
}

func (c *CreateCustomHeaderAuthRequest) SetProvider(provider string) {
	c.Provider = provider
}

func (c *CreateCustomHeaderAuthRequest) SetHeaderName(headerName string) {
	c.HeaderName = headerName
}

func (c *CreateCustomHeaderAuthRequest) SetToken(token string) {
	c.Token = token
}

func (c *CreateCustomHeaderAuthRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateCustomHeaderAuthRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateCustomHeaderAuthRequest to string"
	}
	return string(jsonData)
}

func (c *CreateCustomHeaderAuthRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateCustomHeaderAuthRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateCustomHeaderAuthRequest(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateCustomHeaderAuthRequest) MarshalJSON() ([]byte, error) {
	type embed CreateCustomHeaderAuthRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
