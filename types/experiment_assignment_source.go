package types

import (
	"fmt"
)

type ExperimentAssignmentSource string

const (
	ExperimentAssignmentSourceServerBranch   ExperimentAssignmentSource = "server_branch"
	ExperimentAssignmentSourceClientDeclared ExperimentAssignmentSource = "client_declared"
)

func NewExperimentAssignmentSourceFromString(s string) (ExperimentAssignmentSource, error) {
	switch s {
	case "server_branch":
		return ExperimentAssignmentSourceServerBranch, nil
	case "client_declared":
		return ExperimentAssignmentSourceClientDeclared, nil
	}
	var t ExperimentAssignmentSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ExperimentAssignmentSource) Ptr() *ExperimentAssignmentSource {
	return &e
}
