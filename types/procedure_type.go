package types

import (
	"fmt"
)

type ProcedureType string

const (
	ProcedureTypeFreeForm      ProcedureType = "free_form"
	ProcedureTypeDeterministic ProcedureType = "deterministic"
	ProcedureTypeFolder        ProcedureType = "folder"
)

func NewProcedureTypeFromString(s string) (ProcedureType, error) {
	switch s {
	case "free_form":
		return ProcedureTypeFreeForm, nil
	case "deterministic":
		return ProcedureTypeDeterministic, nil
	case "folder":
		return ProcedureTypeFolder, nil
	}
	var t ProcedureType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProcedureType) Ptr() *ProcedureType {
	return &e
}
