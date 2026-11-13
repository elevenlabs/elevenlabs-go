package types

import (
	"encoding/json"
	"fmt"
)

// Optional workspace auth connection for authentication. Only auth connections that produce an Authorization Bearer token are supported; Basic auth, mTLS, custom header, and URL secret auth connections are not supported.
type CustomLlmAuthConnection struct {
	AuthConnectionLocator            *AuthConnectionLocator
	EnvironmentAuthConnectionLocator *EnvironmentAuthConnectionLocator
	typ                              string
}

func (c *CustomLlmAuthConnection) GetAuthConnectionLocator() *AuthConnectionLocator {
	if c == nil {
		return nil
	}
	if c.AuthConnectionLocator == nil {
		return nil
	}
	return c.AuthConnectionLocator
}

func (c *CustomLlmAuthConnection) GetEnvironmentAuthConnectionLocator() *EnvironmentAuthConnectionLocator {
	if c == nil {
		return nil
	}
	if c.EnvironmentAuthConnectionLocator == nil {
		return nil
	}
	return c.EnvironmentAuthConnectionLocator
}

func (c *CustomLlmAuthConnection) UnmarshalJSON(data []byte) error {
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

func (c CustomLlmAuthConnection) MarshalJSON() ([]byte, error) {
	if c.typ == "AuthConnectionLocator" || c.AuthConnectionLocator != nil {
		return json.Marshal(c.AuthConnectionLocator)
	}
	if c.typ == "EnvironmentAuthConnectionLocator" || c.EnvironmentAuthConnectionLocator != nil {
		return json.Marshal(c.EnvironmentAuthConnectionLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CustomLlmAuthConnectionVisitor interface {
	VisitAuthConnectionLocator(*AuthConnectionLocator) error
	VisitEnvironmentAuthConnectionLocator(*EnvironmentAuthConnectionLocator) error
}

func (c *CustomLlmAuthConnection) Accept(visitor CustomLlmAuthConnectionVisitor) error {
	if c.typ == "AuthConnectionLocator" || c.AuthConnectionLocator != nil {
		return visitor.VisitAuthConnectionLocator(c.AuthConnectionLocator)
	}
	if c.typ == "EnvironmentAuthConnectionLocator" || c.EnvironmentAuthConnectionLocator != nil {
		return visitor.VisitEnvironmentAuthConnectionLocator(c.EnvironmentAuthConnectionLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
