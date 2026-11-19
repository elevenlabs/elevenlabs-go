package types

import (
	"encoding/json"
	"fmt"
)

// The URL of the MCP server, if this contains a secret please store as a workspace secret, otherwise store as a plain string. Must use https
type McpServerConfigInputUrl struct {
	String              string
	ConvAiSecretLocator *ConvAiSecretLocator
	typ                 string
}

func (mcpServerConfigInputUrl *McpServerConfigInputUrl) GetString() string {
	if mcpServerConfigInputUrl == nil {
		return ""
	}
	return mcpServerConfigInputUrl.String
}

func (mcpServerConfigInputUrl *McpServerConfigInputUrl) GetConvAiSecretLocator() *ConvAiSecretLocator {
	if mcpServerConfigInputUrl == nil {
		return nil
	}
	if mcpServerConfigInputUrl.ConvAiSecretLocator == nil {
		return nil
	}
	return mcpServerConfigInputUrl.ConvAiSecretLocator
}

func (mcpServerConfigInputUrl *McpServerConfigInputUrl) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		mcpServerConfigInputUrl.typ = "String"
		mcpServerConfigInputUrl.String = valueString
		return nil
	}
	var valueConvAiSecretLocator *ConvAiSecretLocator
	if err := json.Unmarshal(data, &valueConvAiSecretLocator); err == nil {
		mcpServerConfigInputUrl.typ = "ConvAiSecretLocator"
		mcpServerConfigInputUrl.ConvAiSecretLocator = valueConvAiSecretLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, mcpServerConfigInputUrl)
}

func (mcpServerConfigInputUrl McpServerConfigInputUrl) MarshalJSON() ([]byte, error) {
	if mcpServerConfigInputUrl.typ == "String" || mcpServerConfigInputUrl.String != "" {
		return json.Marshal(mcpServerConfigInputUrl.String)
	}
	if mcpServerConfigInputUrl.typ == "ConvAiSecretLocator" || mcpServerConfigInputUrl.ConvAiSecretLocator != nil {
		return json.Marshal(mcpServerConfigInputUrl.ConvAiSecretLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", mcpServerConfigInputUrl)
}

type McpServerConfigInputUrlVisitor interface {
	VisitString(string) error
	VisitConvAiSecretLocator(*ConvAiSecretLocator) error
}

func (mcpServerConfigInputUrl *McpServerConfigInputUrl) Accept(visitor McpServerConfigInputUrlVisitor) error {
	if mcpServerConfigInputUrl.typ == "String" || mcpServerConfigInputUrl.String != "" {
		return visitor.VisitString(mcpServerConfigInputUrl.String)
	}
	if mcpServerConfigInputUrl.typ == "ConvAiSecretLocator" || mcpServerConfigInputUrl.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(mcpServerConfigInputUrl.ConvAiSecretLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", mcpServerConfigInputUrl)
}
