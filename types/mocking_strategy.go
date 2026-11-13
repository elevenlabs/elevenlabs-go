package types

import (
	"fmt"
)

type MockingStrategy string

const (
	MockingStrategyAll      MockingStrategy = "all"
	MockingStrategySelected MockingStrategy = "selected"
	MockingStrategyNone     MockingStrategy = "none"
)

func NewMockingStrategyFromString(s string) (MockingStrategy, error) {
	switch s {
	case "all":
		return MockingStrategyAll, nil
	case "selected":
		return MockingStrategySelected, nil
	case "none":
		return MockingStrategyNone, nil
	}
	var t MockingStrategy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MockingStrategy) Ptr() *MockingStrategy {
	return &e
}
