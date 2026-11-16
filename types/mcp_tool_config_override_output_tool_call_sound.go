package types

import (
	"encoding/json"
	"fmt"
)

// Overrides the server's tool_call_sound setting for this tool. A sound name plays that sound; 'off' overrides to no sound (silence); null means do not override (inherit the server default).
type McpToolConfigOverrideOutputToolCallSound struct {
	ToolCallSoundType ToolCallSoundType
	OffStringLiteral  string
	typ               string
}

func (m *McpToolConfigOverrideOutputToolCallSound) GetToolCallSoundType() ToolCallSoundType {
	if m == nil {
		return ""
	}
	return m.ToolCallSoundType
}

func (m *McpToolConfigOverrideOutputToolCallSound) GetOffStringLiteral() string {
	if m == nil {
		return ""
	}
	return m.OffStringLiteral
}

func (m *McpToolConfigOverrideOutputToolCallSound) UnmarshalJSON(data []byte) error {
	var valueToolCallSoundType ToolCallSoundType
	if err := json.Unmarshal(data, &valueToolCallSoundType); err == nil {
		m.typ = "ToolCallSoundType"
		m.ToolCallSoundType = valueToolCallSoundType
		return nil
	}
	var valueOffStringLiteral string
	if err := json.Unmarshal(data, &valueOffStringLiteral); err == nil {
		m.typ = "OffStringLiteral"
		m.OffStringLiteral = valueOffStringLiteral
		if m.OffStringLiteral != "off" {
			return fmt.Errorf("unexpected value for literal on type %T; expected off got %v", m, m.OffStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m McpToolConfigOverrideOutputToolCallSound) MarshalJSON() ([]byte, error) {
	if m.typ == "ToolCallSoundType" || m.ToolCallSoundType != "" {
		return json.Marshal(m.ToolCallSoundType)
	}
	if m.typ == "OffStringLiteral" || m.OffStringLiteral != "" {
		return json.Marshal(m.OffStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type McpToolConfigOverrideOutputToolCallSoundVisitor interface {
	VisitToolCallSoundType(ToolCallSoundType) error
	VisitOffStringLiteral(string) error
}

func (m *McpToolConfigOverrideOutputToolCallSound) Accept(visitor McpToolConfigOverrideOutputToolCallSoundVisitor) error {
	if m.typ == "ToolCallSoundType" || m.ToolCallSoundType != "" {
		return visitor.VisitToolCallSoundType(m.ToolCallSoundType)
	}
	if m.typ == "OffStringLiteral" || m.OffStringLiteral != "" {
		return visitor.VisitOffStringLiteral(m.OffStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}

func NewMcpToolConfigOverrideOutputToolCallSoundWithOffStringLiteral() *McpToolConfigOverrideOutputToolCallSound {
	return &McpToolConfigOverrideOutputToolCallSound{typ: "OffStringLiteral", OffStringLiteral: "off"}
}
