package types

import (
	"fmt"
)

// Predefined tool call sounds; “None“ means no sound.
type ToolCallSoundType string

const (
	ToolCallSoundTypeTyping    ToolCallSoundType = "typing"
	ToolCallSoundTypeElevator1 ToolCallSoundType = "elevator1"
	ToolCallSoundTypeElevator2 ToolCallSoundType = "elevator2"
	ToolCallSoundTypeElevator3 ToolCallSoundType = "elevator3"
	ToolCallSoundTypeElevator4 ToolCallSoundType = "elevator4"
)

func NewToolCallSoundTypeFromString(s string) (ToolCallSoundType, error) {
	switch s {
	case "typing":
		return ToolCallSoundTypeTyping, nil
	case "elevator1":
		return ToolCallSoundTypeElevator1, nil
	case "elevator2":
		return ToolCallSoundTypeElevator2, nil
	case "elevator3":
		return ToolCallSoundTypeElevator3, nil
	case "elevator4":
		return ToolCallSoundTypeElevator4, nil
	}
	var t ToolCallSoundType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ToolCallSoundType) Ptr() *ToolCallSoundType {
	return &e
}
