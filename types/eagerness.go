package types

import (
	"fmt"
)

type Eagerness string

const (
	EagernessLow    Eagerness = "low"
	EagernessMedium Eagerness = "medium"
	EagernessHigh   Eagerness = "high"
)

func NewEagernessFromString(s string) (Eagerness, error) {
	switch s {
	case "low":
		return EagernessLow, nil
	case "medium":
		return EagernessMedium, nil
	case "high":
		return EagernessHigh, nil
	}
	var t Eagerness
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Eagerness) Ptr() *Eagerness {
	return &e
}
