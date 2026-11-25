package types

import (
	"fmt"
)

type ConflictResolutionChoice string

const (
	ConflictResolutionChoiceSource ConflictResolutionChoice = "source"
	ConflictResolutionChoiceTarget ConflictResolutionChoice = "target"
	ConflictResolutionChoiceCustom ConflictResolutionChoice = "custom"
)

func NewConflictResolutionChoiceFromString(s string) (ConflictResolutionChoice, error) {
	switch s {
	case "source":
		return ConflictResolutionChoiceSource, nil
	case "target":
		return ConflictResolutionChoiceTarget, nil
	case "custom":
		return ConflictResolutionChoiceCustom, nil
	}
	var t ConflictResolutionChoice
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConflictResolutionChoice) Ptr() *ConflictResolutionChoice {
	return &e
}
