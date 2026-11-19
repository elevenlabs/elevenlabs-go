package mcpservers

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// Overrides the server's tool_call_sound setting for this tool. A sound name plays that sound; 'off' overrides to no sound (silence); null means do not override (inherit the server default).
type McpToolConfigOverrideCreateRequestModelToolCallSound struct {
	ToolCallSoundType types.ToolCallSoundType
	OffStringLiteral  string
	typ               string
}

func (mcpToolConfigOverrideCreateRequestModelToolCallSound *McpToolConfigOverrideCreateRequestModelToolCallSound) GetToolCallSoundType() types.ToolCallSoundType {
	if mcpToolConfigOverrideCreateRequestModelToolCallSound == nil {
		return ""
	}
	return mcpToolConfigOverrideCreateRequestModelToolCallSound.ToolCallSoundType
}

func (mcpToolConfigOverrideCreateRequestModelToolCallSound *McpToolConfigOverrideCreateRequestModelToolCallSound) GetOffStringLiteral() string {
	if mcpToolConfigOverrideCreateRequestModelToolCallSound == nil {
		return ""
	}
	return mcpToolConfigOverrideCreateRequestModelToolCallSound.OffStringLiteral
}

func (mcpToolConfigOverrideCreateRequestModelToolCallSound *McpToolConfigOverrideCreateRequestModelToolCallSound) UnmarshalJSON(data []byte) error {
	var valueToolCallSoundType types.ToolCallSoundType
	if err := json.Unmarshal(data, &valueToolCallSoundType); err == nil {
		mcpToolConfigOverrideCreateRequestModelToolCallSound.typ = "ToolCallSoundType"
		mcpToolConfigOverrideCreateRequestModelToolCallSound.ToolCallSoundType = valueToolCallSoundType
		return nil
	}
	var valueOffStringLiteral string
	if err := json.Unmarshal(data, &valueOffStringLiteral); err == nil {
		mcpToolConfigOverrideCreateRequestModelToolCallSound.typ = "OffStringLiteral"
		mcpToolConfigOverrideCreateRequestModelToolCallSound.OffStringLiteral = valueOffStringLiteral
		if mcpToolConfigOverrideCreateRequestModelToolCallSound.OffStringLiteral != "off" {
			return fmt.Errorf("unexpected value for literal on type %T; expected off got %v", mcpToolConfigOverrideCreateRequestModelToolCallSound, mcpToolConfigOverrideCreateRequestModelToolCallSound.OffStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, mcpToolConfigOverrideCreateRequestModelToolCallSound)
}

func (mcpToolConfigOverrideCreateRequestModelToolCallSound McpToolConfigOverrideCreateRequestModelToolCallSound) MarshalJSON() ([]byte, error) {
	if mcpToolConfigOverrideCreateRequestModelToolCallSound.typ == "ToolCallSoundType" || mcpToolConfigOverrideCreateRequestModelToolCallSound.ToolCallSoundType != "" {
		return json.Marshal(mcpToolConfigOverrideCreateRequestModelToolCallSound.ToolCallSoundType)
	}
	if mcpToolConfigOverrideCreateRequestModelToolCallSound.typ == "OffStringLiteral" || mcpToolConfigOverrideCreateRequestModelToolCallSound.OffStringLiteral != "" {
		return json.Marshal(mcpToolConfigOverrideCreateRequestModelToolCallSound.OffStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", mcpToolConfigOverrideCreateRequestModelToolCallSound)
}

type McpToolConfigOverrideCreateRequestModelToolCallSoundVisitor interface {
	VisitToolCallSoundType(types.ToolCallSoundType) error
	VisitOffStringLiteral(string) error
}

func (mcpToolConfigOverrideCreateRequestModelToolCallSound *McpToolConfigOverrideCreateRequestModelToolCallSound) Accept(visitor McpToolConfigOverrideCreateRequestModelToolCallSoundVisitor) error {
	if mcpToolConfigOverrideCreateRequestModelToolCallSound.typ == "ToolCallSoundType" || mcpToolConfigOverrideCreateRequestModelToolCallSound.ToolCallSoundType != "" {
		return visitor.VisitToolCallSoundType(mcpToolConfigOverrideCreateRequestModelToolCallSound.ToolCallSoundType)
	}
	if mcpToolConfigOverrideCreateRequestModelToolCallSound.typ == "OffStringLiteral" || mcpToolConfigOverrideCreateRequestModelToolCallSound.OffStringLiteral != "" {
		return visitor.VisitOffStringLiteral(mcpToolConfigOverrideCreateRequestModelToolCallSound.OffStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", mcpToolConfigOverrideCreateRequestModelToolCallSound)
}

func NewMcpToolConfigOverrideCreateRequestModelToolCallSoundWithOffStringLiteral() *McpToolConfigOverrideCreateRequestModelToolCallSound {
	return &McpToolConfigOverrideCreateRequestModelToolCallSound{typ: "OffStringLiteral", OffStringLiteral: "off"}
}
