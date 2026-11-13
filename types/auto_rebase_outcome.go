package types

import (
	"fmt"
)

type AutoRebaseOutcome string

const (
	AutoRebaseOutcomeSucceeded        AutoRebaseOutcome = "succeeded"
	AutoRebaseOutcomeConflict         AutoRebaseOutcome = "conflict"
	AutoRebaseOutcomeNoCommonAncestor AutoRebaseOutcome = "no_common_ancestor"
	AutoRebaseOutcomeRejected         AutoRebaseOutcome = "rejected"
	AutoRebaseOutcomeError            AutoRebaseOutcome = "error"
)

func NewAutoRebaseOutcomeFromString(s string) (AutoRebaseOutcome, error) {
	switch s {
	case "succeeded":
		return AutoRebaseOutcomeSucceeded, nil
	case "conflict":
		return AutoRebaseOutcomeConflict, nil
	case "no_common_ancestor":
		return AutoRebaseOutcomeNoCommonAncestor, nil
	case "rejected":
		return AutoRebaseOutcomeRejected, nil
	case "error":
		return AutoRebaseOutcomeError, nil
	}
	var t AutoRebaseOutcome
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AutoRebaseOutcome) Ptr() *AutoRebaseOutcome {
	return &e
}
