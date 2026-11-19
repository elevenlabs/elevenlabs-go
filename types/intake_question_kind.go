package types

import (
	"fmt"
)

// Answer shape for a per-item intake question.
type IntakeQuestionKind string

const (
	IntakeQuestionKindText   IntakeQuestionKind = "text"
	IntakeQuestionKindNumber IntakeQuestionKind = "number"
	IntakeQuestionKindChoice IntakeQuestionKind = "choice"
)

func NewIntakeQuestionKindFromString(s string) (IntakeQuestionKind, error) {
	switch s {
	case "text":
		return IntakeQuestionKindText, nil
	case "number":
		return IntakeQuestionKindNumber, nil
	case "choice":
		return IntakeQuestionKindChoice, nil
	}
	var t IntakeQuestionKind
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e IntakeQuestionKind) Ptr() *IntakeQuestionKind {
	return &e
}
