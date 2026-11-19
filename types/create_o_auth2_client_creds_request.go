package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request model for creating auth connections - inherits common settings and includes sensitive fields
type CreateOAuth2ClientCredsRequest struct {
	Name        string            `json:"name" xml:"name" required:"true"`
	AuthType    *string           `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	Provider    string            `json:"provider" xml:"provider" required:"true"`
	ClientId    string            `json:"client_id" xml:"client_id" required:"true"`
	TokenUrl    string            `json:"token_url" xml:"token_url" required:"true"`
	Scopes      []string          `json:"scopes,omitempty" xml:"scopes,omitempty"`
	ExtraParams map[string]string `json:"extra_params,omitempty"`
	// If True, send client credentials in Authorization header as Basic Auth instead of request body
	BasicAuthInHeader *bool             `json:"basic_auth_in_header,omitempty" xml:"basic_auth_in_header,omitempty"`
	ClientSecret      string            `json:"client_secret" xml:"client_secret" required:"true"`
	CustomHeaders     map[string]string `json:"custom_headers,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *CreateOAuth2ClientCredsRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateOAuth2ClientCredsRequest) GetAuthType() (value *string) {
	if c == nil {
		return
	}
	return c.AuthType
}

func (c *CreateOAuth2ClientCredsRequest) GetProvider() (value string) {
	if c == nil {
		return
	}
	return c.Provider
}

func (c *CreateOAuth2ClientCredsRequest) GetClientId() (value string) {
	if c == nil {
		return
	}
	return c.ClientId
}

func (c *CreateOAuth2ClientCredsRequest) GetTokenUrl() (value string) {
	if c == nil {
		return
	}
	return c.TokenUrl
}

func (c *CreateOAuth2ClientCredsRequest) GetScopes() (value []string) {
	if c == nil {
		return
	}
	return c.Scopes
}

func (c *CreateOAuth2ClientCredsRequest) GetExtraParams() (value map[string]string) {
	if c == nil {
		return
	}
	return c.ExtraParams
}

func (c *CreateOAuth2ClientCredsRequest) GetBasicAuthInHeader() (value *bool) {
	if c == nil {
		return
	}
	return c.BasicAuthInHeader
}

func (c *CreateOAuth2ClientCredsRequest) GetClientSecret() (value string) {
	if c == nil {
		return
	}
	return c.ClientSecret
}

func (c *CreateOAuth2ClientCredsRequest) GetCustomHeaders() (value map[string]string) {
	if c == nil {
		return
	}
	return c.CustomHeaders
}

func (c *CreateOAuth2ClientCredsRequest) SetName(name string) {
	c.Name = name
}

func (c *CreateOAuth2ClientCredsRequest) SetAuthType(authType *string) {
	c.AuthType = authType
}

func (c *CreateOAuth2ClientCredsRequest) SetProvider(provider string) {
	c.Provider = provider
}

func (c *CreateOAuth2ClientCredsRequest) SetClientId(clientId string) {
	c.ClientId = clientId
}

func (c *CreateOAuth2ClientCredsRequest) SetTokenUrl(tokenUrl string) {
	c.TokenUrl = tokenUrl
}

func (c *CreateOAuth2ClientCredsRequest) SetScopes(scopes []string) {
	c.Scopes = scopes
}

func (c *CreateOAuth2ClientCredsRequest) SetExtraParams(extraParams map[string]string) {
	c.ExtraParams = extraParams
}

func (c *CreateOAuth2ClientCredsRequest) SetBasicAuthInHeader(basicAuthInHeader *bool) {
	c.BasicAuthInHeader = basicAuthInHeader
}

func (c *CreateOAuth2ClientCredsRequest) SetClientSecret(clientSecret string) {
	c.ClientSecret = clientSecret
}

func (c *CreateOAuth2ClientCredsRequest) SetCustomHeaders(customHeaders map[string]string) {
	c.CustomHeaders = customHeaders
}

func (c *CreateOAuth2ClientCredsRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateOAuth2ClientCredsRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateOAuth2ClientCredsRequest to string"
	}
	return string(jsonData)
}

func (c *CreateOAuth2ClientCredsRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateOAuth2ClientCredsRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateOAuth2ClientCredsRequest(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateOAuth2ClientCredsRequest) MarshalJSON() ([]byte, error) {
	type embed CreateOAuth2ClientCredsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
