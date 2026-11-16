package types

import (
	"fmt"
)

type TurnMode string

const (
	TurnModeSilence TurnMode = "silence"
	TurnModeTurn    TurnMode = "turn"
)

func NewTurnModeFromString(s string) (TurnMode, error) {
	switch s {
	case "silence":
		return TurnModeSilence, nil
	case "turn":
		return TurnModeTurn, nil
	}
	var t TurnMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TurnMode) Ptr() *TurnMode {
	return &e
}
