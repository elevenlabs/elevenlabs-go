package types

import (
	"fmt"
)

type ToolExecutionTaskSupport string

const (
	ToolExecutionTaskSupportForbidden ToolExecutionTaskSupport = "forbidden"
	ToolExecutionTaskSupportOptional  ToolExecutionTaskSupport = "optional"
	ToolExecutionTaskSupportRequired  ToolExecutionTaskSupport = "required"
)

func NewToolExecutionTaskSupportFromString(s string) (ToolExecutionTaskSupport, error) {
	switch s {
	case "forbidden":
		return ToolExecutionTaskSupportForbidden, nil
	case "optional":
		return ToolExecutionTaskSupportOptional, nil
	case "required":
		return ToolExecutionTaskSupportRequired, nil
	}
	var t ToolExecutionTaskSupport
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ToolExecutionTaskSupport) Ptr() *ToolExecutionTaskSupport {
	return &e
}
