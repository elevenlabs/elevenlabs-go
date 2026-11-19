package types

import (
	"fmt"
)

// Determines how the tool call sound should be played.
type ToolCallSoundBehavior string

const (
	ToolCallSoundBehaviorAuto   ToolCallSoundBehavior = "auto"
	ToolCallSoundBehaviorAlways ToolCallSoundBehavior = "always"
)

func NewToolCallSoundBehaviorFromString(s string) (ToolCallSoundBehavior, error) {
	switch s {
	case "auto":
		return ToolCallSoundBehaviorAuto, nil
	case "always":
		return ToolCallSoundBehaviorAlways, nil
	}
	var t ToolCallSoundBehavior
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ToolCallSoundBehavior) Ptr() *ToolCallSoundBehavior {
	return &e
}
