package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request model for creating mTLS auth connections.
type CreateMtlsAuthRequest struct {
	Name              string  `json:"name" xml:"name" required:"true"`
	AuthType          *string `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	Provider          string  `json:"provider" xml:"provider" required:"true"`
	ClientCertificate string  `json:"client_certificate" xml:"client_certificate" required:"true"`
	ClientKey         string  `json:"client_key" xml:"client_key" required:"true"`
	CaCertificate     *string `json:"ca_certificate,omitempty" xml:"ca_certificate,omitempty"`
	KeyPassphrase     *string `json:"key_passphrase,omitempty" xml:"key_passphrase,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *CreateMtlsAuthRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateMtlsAuthRequest) GetAuthType() (value *string) {
	if c == nil {
		return
	}
	return c.AuthType
}

func (c *CreateMtlsAuthRequest) GetProvider() (value string) {
	if c == nil {
		return
	}
	return c.Provider
}

func (c *CreateMtlsAuthRequest) GetClientCertificate() (value string) {
	if c == nil {
		return
	}
	return c.ClientCertificate
}

func (c *CreateMtlsAuthRequest) GetClientKey() (value string) {
	if c == nil {
		return
	}
	return c.ClientKey
}

func (c *CreateMtlsAuthRequest) GetCaCertificate() (value *string) {
	if c == nil {
		return
	}
	return c.CaCertificate
}

func (c *CreateMtlsAuthRequest) GetKeyPassphrase() (value *string) {
	if c == nil {
		return
	}
	return c.KeyPassphrase
}

func (c *CreateMtlsAuthRequest) SetName(name string) {
	c.Name = name
}

func (c *CreateMtlsAuthRequest) SetAuthType(authType *string) {
	c.AuthType = authType
}

func (c *CreateMtlsAuthRequest) SetProvider(provider string) {
	c.Provider = provider
}

func (c *CreateMtlsAuthRequest) SetClientCertificate(clientCertificate string) {
	c.ClientCertificate = clientCertificate
}

func (c *CreateMtlsAuthRequest) SetClientKey(clientKey string) {
	c.ClientKey = clientKey
}

func (c *CreateMtlsAuthRequest) SetCaCertificate(caCertificate *string) {
	c.CaCertificate = caCertificate
}

func (c *CreateMtlsAuthRequest) SetKeyPassphrase(keyPassphrase *string) {
	c.KeyPassphrase = keyPassphrase
}

func (c *CreateMtlsAuthRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateMtlsAuthRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateMtlsAuthRequest to string"
	}
	return string(jsonData)
}

func (c *CreateMtlsAuthRequest) UnmarshalJSON(data []byte) error {
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

func (c *CreateMtlsAuthRequest) MarshalJSON() ([]byte, error) {
	type embed CreateMtlsAuthRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
