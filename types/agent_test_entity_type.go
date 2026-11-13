package types

import (
	"fmt"
)

type AgentTestEntityType string

const (
	AgentTestEntityTypeTest   AgentTestEntityType = "test"
	AgentTestEntityTypeFolder AgentTestEntityType = "folder"
)

func NewAgentTestEntityTypeFromString(s string) (AgentTestEntityType, error) {
	switch s {
	case "test":
		return AgentTestEntityTypeTest, nil
	case "folder":
		return AgentTestEntityTypeFolder, nil
	}
	var t AgentTestEntityType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentTestEntityType) Ptr() *AgentTestEntityType {
	return &e
}
