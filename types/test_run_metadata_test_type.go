package types

import (
	"fmt"
)

type TestRunMetadataTestType string

const (
	TestRunMetadataTestTypeLlm        TestRunMetadataTestType = "llm"
	TestRunMetadataTestTypeToolCall   TestRunMetadataTestType = "tool_call"
	TestRunMetadataTestTypeSimulation TestRunMetadataTestType = "simulation"
)

func NewTestRunMetadataTestTypeFromString(s string) (TestRunMetadataTestType, error) {
	switch s {
	case "llm":
		return TestRunMetadataTestTypeLlm, nil
	case "tool_call":
		return TestRunMetadataTestTypeToolCall, nil
	case "simulation":
		return TestRunMetadataTestTypeSimulation, nil
	}
	var t TestRunMetadataTestType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TestRunMetadataTestType) Ptr() *TestRunMetadataTestType {
	return &e
}
