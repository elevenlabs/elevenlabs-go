package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request model for creating OAuth2 JWT auth connections - inherits common settings and includes sensitive fields
type CreateOAuth2JwtRequest struct {
	Name     string  `json:"name" xml:"name" required:"true"`
	AuthType *string `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	Provider string  `json:"provider" xml:"provider" required:"true"`
	// JWT signing algorithm
	Algorithm *OAuth2JwtResponseAlgorithm `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Key ID (kid) for JWT header - useful for key rotation
	KeyId *string `json:"key_id,omitempty" xml:"key_id,omitempty"`
	// JWT issuer (iss claim)
	Issuer string `json:"issuer" xml:"issuer" required:"true"`
	// JWT audience (aud claim)
	Audience string `json:"audience" xml:"audience" required:"true"`
	// JWT subject (sub claim)
	Subject string `json:"subject" xml:"subject" required:"true"`
	// Token expiration time in seconds
	ExpirationSeconds *int `json:"expiration_seconds,omitempty" xml:"expiration_seconds,omitempty" min:"60" max:"86400"`
	// Additional custom claims to include in the JWT
	ExtraParams map[string]string `json:"extra_params,omitempty"`
	// Additional JWT header parameters (e.g. x5t). 'alg' and 'kid' are set from algorithm and key_id.
	ExtraHeaders map[string]string `json:"extra_headers,omitempty"`
	// Token endpoint URL for exchanging JWT for access token
	TokenUrl string `json:"token_url" xml:"token_url" required:"true"`
	// OAuth2 scopes to request when exchanging JWT for access token
	Scopes []string `json:"scopes,omitempty" xml:"scopes,omitempty"`
	// Token field to extract from the token endpoint response.
	TokenResponseField *OAuth2JwtResponseTokenResponseField `json:"token_response_field,omitempty" xml:"token_response_field,omitempty"`
	// How the JWT is used at the token endpoint. 'authorization_grant' sends it as a jwt-bearer grant. 'client_authentication' sends it as a client_assertion on a client_credentials request, with the subject as client_id and scopes in the request body.
	JwtUsage        *OAuth2JwtResponseJwtUsage `json:"jwt_usage,omitempty" xml:"jwt_usage,omitempty"`
	SecretKey       string                     `json:"secret_key" xml:"secret_key" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateOAuth2JwtRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateOAuth2JwtRequest) GetAuthType() (value *string) {
	if c == nil {
		return
	}
	return c.AuthType
}

func (c *CreateOAuth2JwtRequest) GetProvider() (value string) {
	if c == nil {
		return
	}
	return c.Provider
}

func (c *CreateOAuth2JwtRequest) GetAlgorithm() (value *OAuth2JwtResponseAlgorithm) {
	if c == nil {
		return
	}
	return c.Algorithm
}

func (c *CreateOAuth2JwtRequest) GetKeyId() (value *string) {
	if c == nil {
		return
	}
	return c.KeyId
}

func (c *CreateOAuth2JwtRequest) GetIssuer() (value string) {
	if c == nil {
		return
	}
	return c.Issuer
}

func (c *CreateOAuth2JwtRequest) GetAudience() (value string) {
	if c == nil {
		return
	}
	return c.Audience
}

func (c *CreateOAuth2JwtRequest) GetSubject() (value string) {
	if c == nil {
		return
	}
	return c.Subject
}

func (c *CreateOAuth2JwtRequest) GetExpirationSeconds() (value *int) {
	if c == nil {
		return
	}
	return c.ExpirationSeconds
}

func (c *CreateOAuth2JwtRequest) GetExtraParams() (value map[string]string) {
	if c == nil {
		return
	}
	return c.ExtraParams
}

func (c *CreateOAuth2JwtRequest) GetExtraHeaders() (value map[string]string) {
	if c == nil {
		return
	}
	return c.ExtraHeaders
}

func (c *CreateOAuth2JwtRequest) GetTokenUrl() (value string) {
	if c == nil {
		return
	}
	return c.TokenUrl
}

func (c *CreateOAuth2JwtRequest) GetScopes() (value []string) {
	if c == nil {
		return
	}
	return c.Scopes
}

func (c *CreateOAuth2JwtRequest) GetTokenResponseField() (value *OAuth2JwtResponseTokenResponseField) {
	if c == nil {
		return
	}
	return c.TokenResponseField
}

func (c *CreateOAuth2JwtRequest) GetJwtUsage() (value *OAuth2JwtResponseJwtUsage) {
	if c == nil {
		return
	}
	return c.JwtUsage
}

func (c *CreateOAuth2JwtRequest) GetSecretKey() (value string) {
	if c == nil {
		return
	}
	return c.SecretKey
}

func (c *CreateOAuth2JwtRequest) SetName(name string) {
	c.Name = name
}

func (c *CreateOAuth2JwtRequest) SetAuthType(authType *string) {
	c.AuthType = authType
}

func (c *CreateOAuth2JwtRequest) SetProvider(provider string) {
	c.Provider = provider
}

func (c *CreateOAuth2JwtRequest) SetAlgorithm(algorithm *OAuth2JwtResponseAlgorithm) {
	c.Algorithm = algorithm
}

func (c *CreateOAuth2JwtRequest) SetKeyId(keyId *string) {
	c.KeyId = keyId
}

func (c *CreateOAuth2JwtRequest) SetIssuer(issuer string) {
	c.Issuer = issuer
}

func (c *CreateOAuth2JwtRequest) SetAudience(audience string) {
	c.Audience = audience
}

func (c *CreateOAuth2JwtRequest) SetSubject(subject string) {
	c.Subject = subject
}

func (c *CreateOAuth2JwtRequest) SetExpirationSeconds(expirationSeconds *int) {
	c.ExpirationSeconds = expirationSeconds
}

func (c *CreateOAuth2JwtRequest) SetExtraParams(extraParams map[string]string) {
	c.ExtraParams = extraParams
}

func (c *CreateOAuth2JwtRequest) SetExtraHeaders(extraHeaders map[string]string) {
	c.ExtraHeaders = extraHeaders
}

func (c *CreateOAuth2JwtRequest) SetTokenUrl(tokenUrl string) {
	c.TokenUrl = tokenUrl
}

func (c *CreateOAuth2JwtRequest) SetScopes(scopes []string) {
	c.Scopes = scopes
}

func (c *CreateOAuth2JwtRequest) SetTokenResponseField(tokenResponseField *OAuth2JwtResponseTokenResponseField) {
	c.TokenResponseField = tokenResponseField
}

func (c *CreateOAuth2JwtRequest) SetJwtUsage(jwtUsage *OAuth2JwtResponseJwtUsage) {
	c.JwtUsage = jwtUsage
}

func (c *CreateOAuth2JwtRequest) SetSecretKey(secretKey string) {
	c.SecretKey = secretKey
}

func (c *CreateOAuth2JwtRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateOAuth2JwtRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateOAuth2JwtRequest to string"
	}
	return string(jsonData)
}

func (c *CreateOAuth2JwtRequest) UnmarshalJSON(data []byte) error {
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

func (c *CreateOAuth2JwtRequest) MarshalJSON() ([]byte, error) {
	type embed CreateOAuth2JwtRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
