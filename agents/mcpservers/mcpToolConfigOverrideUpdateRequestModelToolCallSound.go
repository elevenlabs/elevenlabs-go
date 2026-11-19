package mcpservers

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// Overrides the server's tool_call_sound setting for this tool. A sound name plays that sound; 'off' overrides to no sound (silence); null means do not override (inherit the server default).
type McpToolConfigOverrideUpdateRequestModelToolCallSound struct {
	ToolCallSoundType types.ToolCallSoundType
	OffStringLiteral  string
	typ               string
}

func (mcpToolConfigOverrideUpdateRequestModelToolCallSound *McpToolConfigOverrideUpdateRequestModelToolCallSound) GetToolCallSoundType() types.ToolCallSoundType {
	if mcpToolConfigOverrideUpdateRequestModelToolCallSound == nil {
		return ""
	}
	return mcpToolConfigOverrideUpdateRequestModelToolCallSound.ToolCallSoundType
}

func (mcpToolConfigOverrideUpdateRequestModelToolCallSound *McpToolConfigOverrideUpdateRequestModelToolCallSound) GetOffStringLiteral() string {
	if mcpToolConfigOverrideUpdateRequestModelToolCallSound == nil {
		return ""
	}
	return mcpToolConfigOverrideUpdateRequestModelToolCallSound.OffStringLiteral
}

func (mcpToolConfigOverrideUpdateRequestModelToolCallSound *McpToolConfigOverrideUpdateRequestModelToolCallSound) UnmarshalJSON(data []byte) error {
	var valueToolCallSoundType types.ToolCallSoundType
	if err := json.Unmarshal(data, &valueToolCallSoundType); err == nil {
		mcpToolConfigOverrideUpdateRequestModelToolCallSound.typ = "ToolCallSoundType"
		mcpToolConfigOverrideUpdateRequestModelToolCallSound.ToolCallSoundType = valueToolCallSoundType
		return nil
	}
	var valueOffStringLiteral string
	if err := json.Unmarshal(data, &valueOffStringLiteral); err == nil {
		mcpToolConfigOverrideUpdateRequestModelToolCallSound.typ = "OffStringLiteral"
		mcpToolConfigOverrideUpdateRequestModelToolCallSound.OffStringLiteral = valueOffStringLiteral
		if mcpToolConfigOverrideUpdateRequestModelToolCallSound.OffStringLiteral != "off" {
			return fmt.Errorf("unexpected value for literal on type %T; expected off got %v", mcpToolConfigOverrideUpdateRequestModelToolCallSound, mcpToolConfigOverrideUpdateRequestModelToolCallSound.OffStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, mcpToolConfigOverrideUpdateRequestModelToolCallSound)
}

func (mcpToolConfigOverrideUpdateRequestModelToolCallSound McpToolConfigOverrideUpdateRequestModelToolCallSound) MarshalJSON() ([]byte, error) {
	if mcpToolConfigOverrideUpdateRequestModelToolCallSound.typ == "ToolCallSoundType" || mcpToolConfigOverrideUpdateRequestModelToolCallSound.ToolCallSoundType != "" {
		return json.Marshal(mcpToolConfigOverrideUpdateRequestModelToolCallSound.ToolCallSoundType)
	}
	if mcpToolConfigOverrideUpdateRequestModelToolCallSound.typ == "OffStringLiteral" || mcpToolConfigOverrideUpdateRequestModelToolCallSound.OffStringLiteral != "" {
		return json.Marshal(mcpToolConfigOverrideUpdateRequestModelToolCallSound.OffStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", mcpToolConfigOverrideUpdateRequestModelToolCallSound)
}

type McpToolConfigOverrideUpdateRequestModelToolCallSoundVisitor interface {
	VisitToolCallSoundType(types.ToolCallSoundType) error
	VisitOffStringLiteral(string) error
}

func (mcpToolConfigOverrideUpdateRequestModelToolCallSound *McpToolConfigOverrideUpdateRequestModelToolCallSound) Accept(visitor McpToolConfigOverrideUpdateRequestModelToolCallSoundVisitor) error {
	if mcpToolConfigOverrideUpdateRequestModelToolCallSound.typ == "ToolCallSoundType" || mcpToolConfigOverrideUpdateRequestModelToolCallSound.ToolCallSoundType != "" {
		return visitor.VisitToolCallSoundType(mcpToolConfigOverrideUpdateRequestModelToolCallSound.ToolCallSoundType)
	}
	if mcpToolConfigOverrideUpdateRequestModelToolCallSound.typ == "OffStringLiteral" || mcpToolConfigOverrideUpdateRequestModelToolCallSound.OffStringLiteral != "" {
		return visitor.VisitOffStringLiteral(mcpToolConfigOverrideUpdateRequestModelToolCallSound.OffStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", mcpToolConfigOverrideUpdateRequestModelToolCallSound)
}

func NewMcpToolConfigOverrideUpdateRequestModelToolCallSoundWithOffStringLiteral() *McpToolConfigOverrideUpdateRequestModelToolCallSound {
	return &McpToolConfigOverrideUpdateRequestModelToolCallSound{typ: "OffStringLiteral", OffStringLiteral: "off"}
}
