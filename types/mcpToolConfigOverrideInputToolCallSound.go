package types

import (
	"encoding/json"
	"fmt"
)

// Overrides the server's tool_call_sound setting for this tool. A sound name plays that sound; 'off' overrides to no sound (silence); null means do not override (inherit the server default).
type McpToolConfigOverrideInputToolCallSound struct {
	ToolCallSoundType ToolCallSoundType
	OffStringLiteral  string
	typ               string
}

func (mcpToolConfigOverrideInputToolCallSound *McpToolConfigOverrideInputToolCallSound) GetToolCallSoundType() ToolCallSoundType {
	if mcpToolConfigOverrideInputToolCallSound == nil {
		return ""
	}
	return mcpToolConfigOverrideInputToolCallSound.ToolCallSoundType
}

func (mcpToolConfigOverrideInputToolCallSound *McpToolConfigOverrideInputToolCallSound) GetOffStringLiteral() string {
	if mcpToolConfigOverrideInputToolCallSound == nil {
		return ""
	}
	return mcpToolConfigOverrideInputToolCallSound.OffStringLiteral
}

func (mcpToolConfigOverrideInputToolCallSound *McpToolConfigOverrideInputToolCallSound) UnmarshalJSON(data []byte) error {
	var valueToolCallSoundType ToolCallSoundType
	if err := json.Unmarshal(data, &valueToolCallSoundType); err == nil {
		mcpToolConfigOverrideInputToolCallSound.typ = "ToolCallSoundType"
		mcpToolConfigOverrideInputToolCallSound.ToolCallSoundType = valueToolCallSoundType
		return nil
	}
	var valueOffStringLiteral string
	if err := json.Unmarshal(data, &valueOffStringLiteral); err == nil {
		mcpToolConfigOverrideInputToolCallSound.typ = "OffStringLiteral"
		mcpToolConfigOverrideInputToolCallSound.OffStringLiteral = valueOffStringLiteral
		if mcpToolConfigOverrideInputToolCallSound.OffStringLiteral != "off" {
			return fmt.Errorf("unexpected value for literal on type %T; expected off got %v", mcpToolConfigOverrideInputToolCallSound, mcpToolConfigOverrideInputToolCallSound.OffStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, mcpToolConfigOverrideInputToolCallSound)
}

func (mcpToolConfigOverrideInputToolCallSound McpToolConfigOverrideInputToolCallSound) MarshalJSON() ([]byte, error) {
	if mcpToolConfigOverrideInputToolCallSound.typ == "ToolCallSoundType" || mcpToolConfigOverrideInputToolCallSound.ToolCallSoundType != "" {
		return json.Marshal(mcpToolConfigOverrideInputToolCallSound.ToolCallSoundType)
	}
	if mcpToolConfigOverrideInputToolCallSound.typ == "OffStringLiteral" || mcpToolConfigOverrideInputToolCallSound.OffStringLiteral != "" {
		return json.Marshal(mcpToolConfigOverrideInputToolCallSound.OffStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", mcpToolConfigOverrideInputToolCallSound)
}

type McpToolConfigOverrideInputToolCallSoundVisitor interface {
	VisitToolCallSoundType(ToolCallSoundType) error
	VisitOffStringLiteral(string) error
}

func (mcpToolConfigOverrideInputToolCallSound *McpToolConfigOverrideInputToolCallSound) Accept(visitor McpToolConfigOverrideInputToolCallSoundVisitor) error {
	if mcpToolConfigOverrideInputToolCallSound.typ == "ToolCallSoundType" || mcpToolConfigOverrideInputToolCallSound.ToolCallSoundType != "" {
		return visitor.VisitToolCallSoundType(mcpToolConfigOverrideInputToolCallSound.ToolCallSoundType)
	}
	if mcpToolConfigOverrideInputToolCallSound.typ == "OffStringLiteral" || mcpToolConfigOverrideInputToolCallSound.OffStringLiteral != "" {
		return visitor.VisitOffStringLiteral(mcpToolConfigOverrideInputToolCallSound.OffStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", mcpToolConfigOverrideInputToolCallSound)
}

func NewMcpToolConfigOverrideInputToolCallSoundWithOffStringLiteral() *McpToolConfigOverrideInputToolCallSound {
	return &McpToolConfigOverrideInputToolCallSound{typ: "OffStringLiteral", OffStringLiteral: "off"}
}
