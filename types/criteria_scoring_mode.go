package types

import (
	"fmt"
)

type CriteriaScoringMode string

const (
	CriteriaScoringModeBinary         CriteriaScoringMode = "binary"
	CriteriaScoringModeNumericUniform CriteriaScoringMode = "numeric_uniform"
)

func NewCriteriaScoringModeFromString(s string) (CriteriaScoringMode, error) {
	switch s {
	case "binary":
		return CriteriaScoringModeBinary, nil
	case "numeric_uniform":
		return CriteriaScoringModeNumericUniform, nil
	}
	var t CriteriaScoringMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CriteriaScoringMode) Ptr() *CriteriaScoringMode {
	return &e
}
