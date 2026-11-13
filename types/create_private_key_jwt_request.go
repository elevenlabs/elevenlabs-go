package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request model for creating Private Key JWT auth connections - inherits common settings and includes sensitive fields
type CreatePrivateKeyJwtRequest struct {
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
	ExtraHeaders    map[string]string `json:"extra_headers,omitempty"`
	SecretKey       string            `json:"secret_key" xml:"secret_key" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreatePrivateKeyJwtRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreatePrivateKeyJwtRequest) GetAuthType() (value *string) {
	if c == nil {
		return
	}
	return c.AuthType
}

func (c *CreatePrivateKeyJwtRequest) GetProvider() (value string) {
	if c == nil {
		return
	}
	return c.Provider
}

func (c *CreatePrivateKeyJwtRequest) GetAlgorithm() (value *OAuth2JwtResponseAlgorithm) {
	if c == nil {
		return
	}
	return c.Algorithm
}

func (c *CreatePrivateKeyJwtRequest) GetKeyId() (value *string) {
	if c == nil {
		return
	}
	return c.KeyId
}

func (c *CreatePrivateKeyJwtRequest) GetIssuer() (value string) {
	if c == nil {
		return
	}
	return c.Issuer
}

func (c *CreatePrivateKeyJwtRequest) GetAudience() (value string) {
	if c == nil {
		return
	}
	return c.Audience
}

func (c *CreatePrivateKeyJwtRequest) GetSubject() (value string) {
	if c == nil {
		return
	}
	return c.Subject
}

func (c *CreatePrivateKeyJwtRequest) GetExpirationSeconds() (value *int) {
	if c == nil {
		return
	}
	return c.ExpirationSeconds
}

func (c *CreatePrivateKeyJwtRequest) GetExtraParams() (value map[string]string) {
	if c == nil {
		return
	}
	return c.ExtraParams
}

func (c *CreatePrivateKeyJwtRequest) GetExtraHeaders() (value map[string]string) {
	if c == nil {
		return
	}
	return c.ExtraHeaders
}

func (c *CreatePrivateKeyJwtRequest) GetSecretKey() (value string) {
	if c == nil {
		return
	}
	return c.SecretKey
}

func (c *CreatePrivateKeyJwtRequest) SetName(name string) {
	c.Name = name
}

func (c *CreatePrivateKeyJwtRequest) SetAuthType(authType *string) {
	c.AuthType = authType
}

func (c *CreatePrivateKeyJwtRequest) SetProvider(provider string) {
	c.Provider = provider
}

func (c *CreatePrivateKeyJwtRequest) SetAlgorithm(algorithm *OAuth2JwtResponseAlgorithm) {
	c.Algorithm = algorithm
}

func (c *CreatePrivateKeyJwtRequest) SetKeyId(keyId *string) {
	c.KeyId = keyId
}

func (c *CreatePrivateKeyJwtRequest) SetIssuer(issuer string) {
	c.Issuer = issuer
}

func (c *CreatePrivateKeyJwtRequest) SetAudience(audience string) {
	c.Audience = audience
}

func (c *CreatePrivateKeyJwtRequest) SetSubject(subject string) {
	c.Subject = subject
}

func (c *CreatePrivateKeyJwtRequest) SetExpirationSeconds(expirationSeconds *int) {
	c.ExpirationSeconds = expirationSeconds
}

func (c *CreatePrivateKeyJwtRequest) SetExtraParams(extraParams map[string]string) {
	c.ExtraParams = extraParams
}

func (c *CreatePrivateKeyJwtRequest) SetExtraHeaders(extraHeaders map[string]string) {
	c.ExtraHeaders = extraHeaders
}

func (c *CreatePrivateKeyJwtRequest) SetSecretKey(secretKey string) {
	c.SecretKey = secretKey
}

func (c *CreatePrivateKeyJwtRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreatePrivateKeyJwtRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePrivateKeyJwtRequest to string"
	}
	return string(jsonData)
}

func (c *CreatePrivateKeyJwtRequest) UnmarshalJSON(data []byte) error {
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

func (c *CreatePrivateKeyJwtRequest) MarshalJSON() ([]byte, error) {
	type embed CreatePrivateKeyJwtRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
