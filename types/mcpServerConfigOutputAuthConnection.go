package types

import (
	"encoding/json"
	"fmt"
)

// Optional workspace auth connection for authentication. Only auth connections that produce an Authorization Bearer token are supported; Basic auth, mTLS, custom header, and URL secret auth connections are not supported.
type McpServerConfigOutputAuthConnection struct {
	AuthConnectionLocator            *AuthConnectionLocator
	EnvironmentAuthConnectionLocator *EnvironmentAuthConnectionLocator
	typ                              string
}

func (mcpServerConfigOutputAuthConnection *McpServerConfigOutputAuthConnection) GetAuthConnectionLocator() *AuthConnectionLocator {
	if mcpServerConfigOutputAuthConnection == nil {
		return nil
	}
	if mcpServerConfigOutputAuthConnection.AuthConnectionLocator == nil {
		return nil
	}
	return mcpServerConfigOutputAuthConnection.AuthConnectionLocator
}

func (mcpServerConfigOutputAuthConnection *McpServerConfigOutputAuthConnection) GetEnvironmentAuthConnectionLocator() *EnvironmentAuthConnectionLocator {
	if mcpServerConfigOutputAuthConnection == nil {
		return nil
	}
	if mcpServerConfigOutputAuthConnection.EnvironmentAuthConnectionLocator == nil {
		return nil
	}
	return mcpServerConfigOutputAuthConnection.EnvironmentAuthConnectionLocator
}

func (mcpServerConfigOutputAuthConnection *McpServerConfigOutputAuthConnection) UnmarshalJSON(data []byte) error {
	var valueAuthConnectionLocator *AuthConnectionLocator
	if err := json.Unmarshal(data, &valueAuthConnectionLocator); err == nil {
		mcpServerConfigOutputAuthConnection.typ = "AuthConnectionLocator"
		mcpServerConfigOutputAuthConnection.AuthConnectionLocator = valueAuthConnectionLocator
		return nil
	}
	var valueEnvironmentAuthConnectionLocator *EnvironmentAuthConnectionLocator
	if err := json.Unmarshal(data, &valueEnvironmentAuthConnectionLocator); err == nil {
		mcpServerConfigOutputAuthConnection.typ = "EnvironmentAuthConnectionLocator"
		mcpServerConfigOutputAuthConnection.EnvironmentAuthConnectionLocator = valueEnvironmentAuthConnectionLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, mcpServerConfigOutputAuthConnection)
}

func (mcpServerConfigOutputAuthConnection McpServerConfigOutputAuthConnection) MarshalJSON() ([]byte, error) {
	if mcpServerConfigOutputAuthConnection.typ == "AuthConnectionLocator" || mcpServerConfigOutputAuthConnection.AuthConnectionLocator != nil {
		return json.Marshal(mcpServerConfigOutputAuthConnection.AuthConnectionLocator)
	}
	if mcpServerConfigOutputAuthConnection.typ == "EnvironmentAuthConnectionLocator" || mcpServerConfigOutputAuthConnection.EnvironmentAuthConnectionLocator != nil {
		return json.Marshal(mcpServerConfigOutputAuthConnection.EnvironmentAuthConnectionLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", mcpServerConfigOutputAuthConnection)
}

type McpServerConfigOutputAuthConnectionVisitor interface {
	VisitAuthConnectionLocator(*AuthConnectionLocator) error
	VisitEnvironmentAuthConnectionLocator(*EnvironmentAuthConnectionLocator) error
}

func (mcpServerConfigOutputAuthConnection *McpServerConfigOutputAuthConnection) Accept(visitor McpServerConfigOutputAuthConnectionVisitor) error {
	if mcpServerConfigOutputAuthConnection.typ == "AuthConnectionLocator" || mcpServerConfigOutputAuthConnection.AuthConnectionLocator != nil {
		return visitor.VisitAuthConnectionLocator(mcpServerConfigOutputAuthConnection.AuthConnectionLocator)
	}
	if mcpServerConfigOutputAuthConnection.typ == "EnvironmentAuthConnectionLocator" || mcpServerConfigOutputAuthConnection.EnvironmentAuthConnectionLocator != nil {
		return visitor.VisitEnvironmentAuthConnectionLocator(mcpServerConfigOutputAuthConnection.EnvironmentAuthConnectionLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", mcpServerConfigOutputAuthConnection)
}
