package types

import (
	"encoding/json"
	"fmt"
)

// The URL of the MCP server, if this contains a secret please store as a workspace secret, otherwise store as a plain string. Must use https
type McpServerConfigOutputUrl struct {
	String              string
	ConvAiSecretLocator *ConvAiSecretLocator
	typ                 string
}

func (m *McpServerConfigOutputUrl) GetString() string {
	if m == nil {
		return ""
	}
	return m.String
}

func (m *McpServerConfigOutputUrl) GetConvAiSecretLocator() *ConvAiSecretLocator {
	if m == nil {
		return nil
	}
	if m.ConvAiSecretLocator == nil {
		return nil
	}
	return m.ConvAiSecretLocator
}

func (m *McpServerConfigOutputUrl) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		m.typ = "String"
		m.String = valueString
		return nil
	}
	var valueConvAiSecretLocator *ConvAiSecretLocator
	if err := json.Unmarshal(data, &valueConvAiSecretLocator); err == nil {
		m.typ = "ConvAiSecretLocator"
		m.ConvAiSecretLocator = valueConvAiSecretLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m McpServerConfigOutputUrl) MarshalJSON() ([]byte, error) {
	if m.typ == "String" || m.String != "" {
		return json.Marshal(m.String)
	}
	if m.typ == "ConvAiSecretLocator" || m.ConvAiSecretLocator != nil {
		return json.Marshal(m.ConvAiSecretLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type McpServerConfigOutputUrlVisitor interface {
	VisitString(string) error
	VisitConvAiSecretLocator(*ConvAiSecretLocator) error
}

func (m *McpServerConfigOutputUrl) Accept(visitor McpServerConfigOutputUrlVisitor) error {
	if m.typ == "String" || m.String != "" {
		return visitor.VisitString(m.String)
	}
	if m.typ == "ConvAiSecretLocator" || m.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(m.ConvAiSecretLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}
