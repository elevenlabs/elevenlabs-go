package types

import (
	"fmt"
)

// Version of the turn detection model to use.
type TurnModel string

const (
	TurnModelTurnV2 TurnModel = "turn_v2"
	TurnModelTurnV3 TurnModel = "turn_v3"
)

func NewTurnModelFromString(s string) (TurnModel, error) {
	switch s {
	case "turn_v2":
		return TurnModelTurnV2, nil
	case "turn_v3":
		return TurnModelTurnV3, nil
	}
	var t TurnModel
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TurnModel) Ptr() *TurnModel {
	return &e
}
