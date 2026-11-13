package types

import (
	"fmt"
)

type ToolExecutionMode string

const (
	ToolExecutionModeImmediate      ToolExecutionMode = "immediate"
	ToolExecutionModePostToolSpeech ToolExecutionMode = "post_tool_speech"
	ToolExecutionModeAsync          ToolExecutionMode = "async"
)

func NewToolExecutionModeFromString(s string) (ToolExecutionMode, error) {
	switch s {
	case "immediate":
		return ToolExecutionModeImmediate, nil
	case "post_tool_speech":
		return ToolExecutionModePostToolSpeech, nil
	case "async":
		return ToolExecutionModeAsync, nil
	}
	var t ToolExecutionMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ToolExecutionMode) Ptr() *ToolExecutionMode {
	return &e
}
