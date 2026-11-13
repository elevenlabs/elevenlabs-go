package types

import (
	"fmt"
)

// Controls how tool errors are processed before being shared with the agent.
type ToolErrorHandlingMode string

const (
	ToolErrorHandlingModeAuto        ToolErrorHandlingMode = "auto"
	ToolErrorHandlingModeSummarized  ToolErrorHandlingMode = "summarized"
	ToolErrorHandlingModePassthrough ToolErrorHandlingMode = "passthrough"
	ToolErrorHandlingModeHide        ToolErrorHandlingMode = "hide"
)

func NewToolErrorHandlingModeFromString(s string) (ToolErrorHandlingMode, error) {
	switch s {
	case "auto":
		return ToolErrorHandlingModeAuto, nil
	case "summarized":
		return ToolErrorHandlingModeSummarized, nil
	case "passthrough":
		return ToolErrorHandlingModePassthrough, nil
	case "hide":
		return ToolErrorHandlingModeHide, nil
	}
	var t ToolErrorHandlingMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ToolErrorHandlingMode) Ptr() *ToolErrorHandlingMode {
	return &e
}
