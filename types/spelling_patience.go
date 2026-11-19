package types

import (
	"fmt"
)

// Controls if the agent should be more patient when user is spelling numbers and named entities.
type SpellingPatience string

const (
	SpellingPatienceAuto SpellingPatience = "auto"
	SpellingPatienceOff  SpellingPatience = "off"
)

func NewSpellingPatienceFromString(s string) (SpellingPatience, error) {
	switch s {
	case "auto":
		return SpellingPatienceAuto, nil
	case "off":
		return SpellingPatienceOff, nil
	}
	var t SpellingPatience
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SpellingPatience) Ptr() *SpellingPatience {
	return &e
}
