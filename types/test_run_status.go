package types

import (
	"fmt"
)

type TestRunStatus string

const (
	TestRunStatusPending   TestRunStatus = "pending"
	TestRunStatusPassed    TestRunStatus = "passed"
	TestRunStatusFailed    TestRunStatus = "failed"
	TestRunStatusCancelled TestRunStatus = "cancelled"
)

func NewTestRunStatusFromString(s string) (TestRunStatus, error) {
	switch s {
	case "pending":
		return TestRunStatusPending, nil
	case "passed":
		return TestRunStatusPassed, nil
	case "failed":
		return TestRunStatusFailed, nil
	case "cancelled":
		return TestRunStatusCancelled, nil
	}
	var t TestRunStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TestRunStatus) Ptr() *TestRunStatus {
	return &e
}
