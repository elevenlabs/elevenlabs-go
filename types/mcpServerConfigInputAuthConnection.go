package types

import (
	"encoding/json"
	"fmt"
)

// Optional workspace auth connection for authentication. Only auth connections that produce an Authorization Bearer token are supported; Basic auth, mTLS, custom header, and URL secret auth connections are not supported.
type McpServerConfigInputAuthConnection struct {
	AuthConnectionLocator            *AuthConnectionLocator
	EnvironmentAuthConnectionLocator *EnvironmentAuthConnectionLocator
	typ                              string
}

func (mcpServerConfigInputAuthConnection *McpServerConfigInputAuthConnection) GetAuthConnectionLocator() *AuthConnectionLocator {
	if mcpServerConfigInputAuthConnection == nil {
		return nil
	}
	if mcpServerConfigInputAuthConnection.AuthConnectionLocator == nil {
		return nil
	}
	return mcpServerConfigInputAuthConnection.AuthConnectionLocator
}

func (mcpServerConfigInputAuthConnection *McpServerConfigInputAuthConnection) GetEnvironmentAuthConnectionLocator() *EnvironmentAuthConnectionLocator {
	if mcpServerConfigInputAuthConnection == nil {
		return nil
	}
	if mcpServerConfigInputAuthConnection.EnvironmentAuthConnectionLocator == nil {
		return nil
	}
	return mcpServerConfigInputAuthConnection.EnvironmentAuthConnectionLocator
}

func (mcpServerConfigInputAuthConnection *McpServerConfigInputAuthConnection) UnmarshalJSON(data []byte) error {
	var valueAuthConnectionLocator *AuthConnectionLocator
	if err := json.Unmarshal(data, &valueAuthConnectionLocator); err == nil {
		mcpServerConfigInputAuthConnection.typ = "AuthConnectionLocator"
		mcpServerConfigInputAuthConnection.AuthConnectionLocator = valueAuthConnectionLocator
		return nil
	}
	var valueEnvironmentAuthConnectionLocator *EnvironmentAuthConnectionLocator
	if err := json.Unmarshal(data, &valueEnvironmentAuthConnectionLocator); err == nil {
		mcpServerConfigInputAuthConnection.typ = "EnvironmentAuthConnectionLocator"
		mcpServerConfigInputAuthConnection.EnvironmentAuthConnectionLocator = valueEnvironmentAuthConnectionLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, mcpServerConfigInputAuthConnection)
}

func (mcpServerConfigInputAuthConnection McpServerConfigInputAuthConnection) MarshalJSON() ([]byte, error) {
	if mcpServerConfigInputAuthConnection.typ == "AuthConnectionLocator" || mcpServerConfigInputAuthConnection.AuthConnectionLocator != nil {
		return json.Marshal(mcpServerConfigInputAuthConnection.AuthConnectionLocator)
	}
	if mcpServerConfigInputAuthConnection.typ == "EnvironmentAuthConnectionLocator" || mcpServerConfigInputAuthConnection.EnvironmentAuthConnectionLocator != nil {
		return json.Marshal(mcpServerConfigInputAuthConnection.EnvironmentAuthConnectionLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", mcpServerConfigInputAuthConnection)
}

type McpServerConfigInputAuthConnectionVisitor interface {
	VisitAuthConnectionLocator(*AuthConnectionLocator) error
	VisitEnvironmentAuthConnectionLocator(*EnvironmentAuthConnectionLocator) error
}

func (mcpServerConfigInputAuthConnection *McpServerConfigInputAuthConnection) Accept(visitor McpServerConfigInputAuthConnectionVisitor) error {
	if mcpServerConfigInputAuthConnection.typ == "AuthConnectionLocator" || mcpServerConfigInputAuthConnection.AuthConnectionLocator != nil {
		return visitor.VisitAuthConnectionLocator(mcpServerConfigInputAuthConnection.AuthConnectionLocator)
	}
	if mcpServerConfigInputAuthConnection.typ == "EnvironmentAuthConnectionLocator" || mcpServerConfigInputAuthConnection.EnvironmentAuthConnectionLocator != nil {
		return visitor.VisitEnvironmentAuthConnectionLocator(mcpServerConfigInputAuthConnection.EnvironmentAuthConnectionLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", mcpServerConfigInputAuthConnection)
}
