package types

import (
	"fmt"
)

// How quickly the receptionist decides the caller has finished speaking.
type ReceptionTurnEagerness string

const (
	ReceptionTurnEagernessPatient ReceptionTurnEagerness = "patient"
	ReceptionTurnEagernessNormal  ReceptionTurnEagerness = "normal"
	ReceptionTurnEagernessEager   ReceptionTurnEagerness = "eager"
)

func NewReceptionTurnEagernessFromString(s string) (ReceptionTurnEagerness, error) {
	switch s {
	case "patient":
		return ReceptionTurnEagernessPatient, nil
	case "normal":
		return ReceptionTurnEagernessNormal, nil
	case "eager":
		return ReceptionTurnEagernessEager, nil
	}
	var t ReceptionTurnEagerness
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionTurnEagerness) Ptr() *ReceptionTurnEagerness {
	return &e
}
