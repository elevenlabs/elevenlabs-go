package types

import (
	"fmt"
)

type MockNoMatchBehavior string

const (
	MockNoMatchBehaviorCallRealTool MockNoMatchBehavior = "call_real_tool"
	MockNoMatchBehaviorRaiseError   MockNoMatchBehavior = "raise_error"
)

func NewMockNoMatchBehaviorFromString(s string) (MockNoMatchBehavior, error) {
	switch s {
	case "call_real_tool":
		return MockNoMatchBehaviorCallRealTool, nil
	case "raise_error":
		return MockNoMatchBehaviorRaiseError, nil
	}
	var t MockNoMatchBehavior
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MockNoMatchBehavior) Ptr() *MockNoMatchBehavior {
	return &e
}
