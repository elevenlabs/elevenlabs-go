package types

import (
	"fmt"
)

type AgentDeploymentSource string

const (
	AgentDeploymentSourceManual      AgentDeploymentSource = "manual"
	AgentDeploymentSourceBranchMerge AgentDeploymentSource = "branch_merge"
	AgentDeploymentSourceDuplicate   AgentDeploymentSource = "duplicate"
)

func NewAgentDeploymentSourceFromString(s string) (AgentDeploymentSource, error) {
	switch s {
	case "manual":
		return AgentDeploymentSourceManual, nil
	case "branch_merge":
		return AgentDeploymentSourceBranchMerge, nil
	case "duplicate":
		return AgentDeploymentSourceDuplicate, nil
	}
	var t AgentDeploymentSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentDeploymentSource) Ptr() *AgentDeploymentSource {
	return &e
}
