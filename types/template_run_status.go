package types

import (
	"fmt"
)

// Where a run, or one of its outputs, is in its lifecycle. The run's is rolled
// up from its outputs' (“aggregate_output_statuses“). A class rather than a
// “Literal“ so it publishes as one named enum instead of being inlined into
// every kind.
type TemplateRunStatus string

const (
	TemplateRunStatusPending    TemplateRunStatus = "pending"
	TemplateRunStatusGenerating TemplateRunStatus = "generating"
	TemplateRunStatusCompleted  TemplateRunStatus = "completed"
	TemplateRunStatusFailed     TemplateRunStatus = "failed"
)

func NewTemplateRunStatusFromString(s string) (TemplateRunStatus, error) {
	switch s {
	case "pending":
		return TemplateRunStatusPending, nil
	case "generating":
		return TemplateRunStatusGenerating, nil
	case "completed":
		return TemplateRunStatusCompleted, nil
	case "failed":
		return TemplateRunStatusFailed, nil
	}
	var t TemplateRunStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TemplateRunStatus) Ptr() *TemplateRunStatus {
	return &e
}
