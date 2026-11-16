package types

import (
	"fmt"
)

type ToolInterruptionMode string

const (
	ToolInterruptionModeAllow                    ToolInterruptionMode = "allow"
	ToolInterruptionModeDisableDuringTool        ToolInterruptionMode = "disable_during_tool"
	ToolInterruptionModeDisableDuringToolAndTurn ToolInterruptionMode = "disable_during_tool_and_turn"
)

func NewToolInterruptionModeFromString(s string) (ToolInterruptionMode, error) {
	switch s {
	case "allow":
		return ToolInterruptionModeAllow, nil
	case "disable_during_tool":
		return ToolInterruptionModeDisableDuringTool, nil
	case "disable_during_tool_and_turn":
		return ToolInterruptionModeDisableDuringToolAndTurn, nil
	}
	var t ToolInterruptionMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ToolInterruptionMode) Ptr() *ToolInterruptionMode {
	return &e
}
