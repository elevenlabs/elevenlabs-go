package types

import (
	"fmt"
)

type TestType string

const (
	TestTypeLlm        TestType = "llm"
	TestTypeTool       TestType = "tool"
	TestTypeSimulation TestType = "simulation"
	TestTypeFolder     TestType = "folder"
)

func NewTestTypeFromString(s string) (TestType, error) {
	switch s {
	case "llm":
		return TestTypeLlm, nil
	case "tool":
		return TestTypeTool, nil
	case "simulation":
		return TestTypeSimulation, nil
	case "folder":
		return TestTypeFolder, nil
	}
	var t TestType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TestType) Ptr() *TestType {
	return &e
}
