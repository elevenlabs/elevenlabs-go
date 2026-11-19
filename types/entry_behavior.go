package types

import (
	"fmt"
)

type EntryBehavior string

const (
	EntryBehaviorGenerateImmediately EntryBehavior = "generate_immediately"
	EntryBehaviorWaitForUser         EntryBehavior = "wait_for_user"
	EntryBehaviorAuto                EntryBehavior = "auto"
)

func NewEntryBehaviorFromString(s string) (EntryBehavior, error) {
	switch s {
	case "generate_immediately":
		return EntryBehaviorGenerateImmediately, nil
	case "wait_for_user":
		return EntryBehaviorWaitForUser, nil
	case "auto":
		return EntryBehaviorAuto, nil
	}
	var t EntryBehavior
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EntryBehavior) Ptr() *EntryBehavior {
	return &e
}
