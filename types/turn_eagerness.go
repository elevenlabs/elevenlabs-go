package types

import (
	"fmt"
)

// Agent's eagerness to respond. Higher values make agent wait for higher turn probability.
type TurnEagerness string

const (
	TurnEagernessPatient TurnEagerness = "patient"
	TurnEagernessNormal  TurnEagerness = "normal"
	TurnEagernessEager   TurnEagerness = "eager"
)

func NewTurnEagernessFromString(s string) (TurnEagerness, error) {
	switch s {
	case "patient":
		return TurnEagernessPatient, nil
	case "normal":
		return TurnEagernessNormal, nil
	case "eager":
		return TurnEagernessEager, nil
	}
	var t TurnEagerness
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TurnEagerness) Ptr() *TurnEagerness {
	return &e
}
