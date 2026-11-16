package types

import (
	"fmt"
)

// Whether reaching an End node completes successfully or fails.
type WorkflowEndNodeOutcome string

const (
	WorkflowEndNodeOutcomeSuccess WorkflowEndNodeOutcome = "success"
	WorkflowEndNodeOutcomeFailure WorkflowEndNodeOutcome = "failure"
)

func NewWorkflowEndNodeOutcomeFromString(s string) (WorkflowEndNodeOutcome, error) {
	switch s {
	case "success":
		return WorkflowEndNodeOutcomeSuccess, nil
	case "failure":
		return WorkflowEndNodeOutcomeFailure, nil
	}
	var t WorkflowEndNodeOutcome
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WorkflowEndNodeOutcome) Ptr() *WorkflowEndNodeOutcome {
	return &e
}
