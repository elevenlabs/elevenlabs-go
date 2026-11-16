package types

import (
	"fmt"
)

type TestSharingMode string

const (
	TestSharingModeAll          TestSharingMode = "all"
	TestSharingModeSharedWithMe TestSharingMode = "shared_with_me"
)

func NewTestSharingModeFromString(s string) (TestSharingMode, error) {
	switch s {
	case "all":
		return TestSharingModeAll, nil
	case "shared_with_me":
		return TestSharingModeSharedWithMe, nil
	}
	var t TestSharingMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TestSharingMode) Ptr() *TestSharingMode {
	return &e
}
