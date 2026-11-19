package agents

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// Optional workspace auth connection for authentication. Only auth connections that produce an Authorization Bearer token are supported; Basic auth, mTLS, custom header, and URL secret auth connections are not supported.
type McpServerConfigUpdateRequestModelAuthConnection struct {
	AuthConnectionLocator            *types.AuthConnectionLocator
	EnvironmentAuthConnectionLocator *types.EnvironmentAuthConnectionLocator
	typ                              string
}

func (mcpServerConfigUpdateRequestModelAuthConnection *McpServerConfigUpdateRequestModelAuthConnection) GetAuthConnectionLocator() *types.AuthConnectionLocator {
	if mcpServerConfigUpdateRequestModelAuthConnection == nil {
		return nil
	}
	if mcpServerConfigUpdateRequestModelAuthConnection.AuthConnectionLocator == nil {
		return nil
	}
	return mcpServerConfigUpdateRequestModelAuthConnection.AuthConnectionLocator
}

func (mcpServerConfigUpdateRequestModelAuthConnection *McpServerConfigUpdateRequestModelAuthConnection) GetEnvironmentAuthConnectionLocator() *types.EnvironmentAuthConnectionLocator {
	if mcpServerConfigUpdateRequestModelAuthConnection == nil {
		return nil
	}
	if mcpServerConfigUpdateRequestModelAuthConnection.EnvironmentAuthConnectionLocator == nil {
		return nil
	}
	return mcpServerConfigUpdateRequestModelAuthConnection.EnvironmentAuthConnectionLocator
}

func (mcpServerConfigUpdateRequestModelAuthConnection *McpServerConfigUpdateRequestModelAuthConnection) UnmarshalJSON(data []byte) error {
	var valueAuthConnectionLocator *types.AuthConnectionLocator
	if err := json.Unmarshal(data, &valueAuthConnectionLocator); err == nil {
		mcpServerConfigUpdateRequestModelAuthConnection.typ = "AuthConnectionLocator"
		mcpServerConfigUpdateRequestModelAuthConnection.AuthConnectionLocator = valueAuthConnectionLocator
		return nil
	}
	var valueEnvironmentAuthConnectionLocator *types.EnvironmentAuthConnectionLocator
	if err := json.Unmarshal(data, &valueEnvironmentAuthConnectionLocator); err == nil {
		mcpServerConfigUpdateRequestModelAuthConnection.typ = "EnvironmentAuthConnectionLocator"
		mcpServerConfigUpdateRequestModelAuthConnection.EnvironmentAuthConnectionLocator = valueEnvironmentAuthConnectionLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, mcpServerConfigUpdateRequestModelAuthConnection)
}

func (mcpServerConfigUpdateRequestModelAuthConnection McpServerConfigUpdateRequestModelAuthConnection) MarshalJSON() ([]byte, error) {
	if mcpServerConfigUpdateRequestModelAuthConnection.typ == "AuthConnectionLocator" || mcpServerConfigUpdateRequestModelAuthConnection.AuthConnectionLocator != nil {
		return json.Marshal(mcpServerConfigUpdateRequestModelAuthConnection.AuthConnectionLocator)
	}
	if mcpServerConfigUpdateRequestModelAuthConnection.typ == "EnvironmentAuthConnectionLocator" || mcpServerConfigUpdateRequestModelAuthConnection.EnvironmentAuthConnectionLocator != nil {
		return json.Marshal(mcpServerConfigUpdateRequestModelAuthConnection.EnvironmentAuthConnectionLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", mcpServerConfigUpdateRequestModelAuthConnection)
}

type McpServerConfigUpdateRequestModelAuthConnectionVisitor interface {
	VisitAuthConnectionLocator(*types.AuthConnectionLocator) error
	VisitEnvironmentAuthConnectionLocator(*types.EnvironmentAuthConnectionLocator) error
}

func (mcpServerConfigUpdateRequestModelAuthConnection *McpServerConfigUpdateRequestModelAuthConnection) Accept(visitor McpServerConfigUpdateRequestModelAuthConnectionVisitor) error {
	if mcpServerConfigUpdateRequestModelAuthConnection.typ == "AuthConnectionLocator" || mcpServerConfigUpdateRequestModelAuthConnection.AuthConnectionLocator != nil {
		return visitor.VisitAuthConnectionLocator(mcpServerConfigUpdateRequestModelAuthConnection.AuthConnectionLocator)
	}
	if mcpServerConfigUpdateRequestModelAuthConnection.typ == "EnvironmentAuthConnectionLocator" || mcpServerConfigUpdateRequestModelAuthConnection.EnvironmentAuthConnectionLocator != nil {
		return visitor.VisitEnvironmentAuthConnectionLocator(mcpServerConfigUpdateRequestModelAuthConnection.EnvironmentAuthConnectionLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", mcpServerConfigUpdateRequestModelAuthConnection)
}
