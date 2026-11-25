package types

import (
	"encoding/json"
	"fmt"
)

// Optional mTLS auth connection. When set, its client certificate is used for the TLS handshake with the LLM endpoint and, when auth_connection is OAuth2 client credentials or OAuth2 JWT, with its token endpoint too. Application auth stays on auth_connection or api_key. auth_connection must then be OAuth2 client credentials, OAuth2 JWT or bearer.
type CustomLlmMtlsAuthConnection struct {
	AuthConnectionLocator            *AuthConnectionLocator
	EnvironmentAuthConnectionLocator *EnvironmentAuthConnectionLocator
	typ                              string
}

func (c *CustomLlmMtlsAuthConnection) GetAuthConnectionLocator() *AuthConnectionLocator {
	if c == nil {
		return nil
	}
	if c.AuthConnectionLocator == nil {
		return nil
	}
	return c.AuthConnectionLocator
}

func (c *CustomLlmMtlsAuthConnection) GetEnvironmentAuthConnectionLocator() *EnvironmentAuthConnectionLocator {
	if c == nil {
		return nil
	}
	if c.EnvironmentAuthConnectionLocator == nil {
		return nil
	}
	return c.EnvironmentAuthConnectionLocator
}

func (c *CustomLlmMtlsAuthConnection) UnmarshalJSON(data []byte) error {
	var valueAuthConnectionLocator *AuthConnectionLocator
	if err := json.Unmarshal(data, &valueAuthConnectionLocator); err == nil {
		c.typ = "AuthConnectionLocator"
		c.AuthConnectionLocator = valueAuthConnectionLocator
		return nil
	}
	var valueEnvironmentAuthConnectionLocator *EnvironmentAuthConnectionLocator
	if err := json.Unmarshal(data, &valueEnvironmentAuthConnectionLocator); err == nil {
		c.typ = "EnvironmentAuthConnectionLocator"
		c.EnvironmentAuthConnectionLocator = valueEnvironmentAuthConnectionLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CustomLlmMtlsAuthConnection) MarshalJSON() ([]byte, error) {
	if c.typ == "AuthConnectionLocator" || c.AuthConnectionLocator != nil {
		return json.Marshal(c.AuthConnectionLocator)
	}
	if c.typ == "EnvironmentAuthConnectionLocator" || c.EnvironmentAuthConnectionLocator != nil {
		return json.Marshal(c.EnvironmentAuthConnectionLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CustomLlmMtlsAuthConnectionVisitor interface {
	VisitAuthConnectionLocator(*AuthConnectionLocator) error
	VisitEnvironmentAuthConnectionLocator(*EnvironmentAuthConnectionLocator) error
}

func (c *CustomLlmMtlsAuthConnection) Accept(visitor CustomLlmMtlsAuthConnectionVisitor) error {
	if c.typ == "AuthConnectionLocator" || c.AuthConnectionLocator != nil {
		return visitor.VisitAuthConnectionLocator(c.AuthConnectionLocator)
	}
	if c.typ == "EnvironmentAuthConnectionLocator" || c.EnvironmentAuthConnectionLocator != nil {
		return visitor.VisitEnvironmentAuthConnectionLocator(c.EnvironmentAuthConnectionLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
