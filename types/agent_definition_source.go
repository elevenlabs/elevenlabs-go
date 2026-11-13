package types

import (
	"fmt"
)

type AgentDefinitionSource string

const (
	AgentDefinitionSourceCli       AgentDefinitionSource = "cli"
	AgentDefinitionSourceUi        AgentDefinitionSource = "ui"
	AgentDefinitionSourceApi       AgentDefinitionSource = "api"
	AgentDefinitionSourceTemplate  AgentDefinitionSource = "template"
	AgentDefinitionSourceArchitect AgentDefinitionSource = "architect"
	AgentDefinitionSourceMcp       AgentDefinitionSource = "mcp"
	AgentDefinitionSourceUnknown   AgentDefinitionSource = "unknown"
)

func NewAgentDefinitionSourceFromString(s string) (AgentDefinitionSource, error) {
	switch s {
	case "cli":
		return AgentDefinitionSourceCli, nil
	case "ui":
		return AgentDefinitionSourceUi, nil
	case "api":
		return AgentDefinitionSourceApi, nil
	case "template":
		return AgentDefinitionSourceTemplate, nil
	case "architect":
		return AgentDefinitionSourceArchitect, nil
	case "mcp":
		return AgentDefinitionSourceMcp, nil
	case "unknown":
		return AgentDefinitionSourceUnknown, nil
	}
	var t AgentDefinitionSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentDefinitionSource) Ptr() *AgentDefinitionSource {
	return &e
}
