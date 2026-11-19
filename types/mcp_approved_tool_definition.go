package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Snapshot of the MCP tool definition that was approved.
type McpApprovedToolDefinition struct {
	// The MCP server-provided tool description at approval time
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The MCP server-provided JSON input schema at approval time
	InputSchema     map[string]any `json:"input_schema,omitempty"`
	extraProperties map[string]interface{}
}

func (m *McpApprovedToolDefinition) GetDescription() (value *string) {
	if m == nil {
		return
	}
	return m.Description
}

func (m *McpApprovedToolDefinition) GetInputSchema() (value map[string]any) {
	if m == nil {
		return
	}
	return m.InputSchema
}

func (m *McpApprovedToolDefinition) SetDescription(description *string) {
	m.Description = description
}

func (m *McpApprovedToolDefinition) SetInputSchema(inputSchema map[string]any) {
	m.InputSchema = inputSchema
}

func (m *McpApprovedToolDefinition) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpApprovedToolDefinition) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpApprovedToolDefinition to string"
	}
	return string(jsonData)
}

func (m *McpApprovedToolDefinition) UnmarshalJSON(data []byte) error {
	type unmarshaler McpApprovedToolDefinition
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = McpApprovedToolDefinition(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpApprovedToolDefinition) MarshalJSON() ([]byte, error) {
	type embed McpApprovedToolDefinition
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
