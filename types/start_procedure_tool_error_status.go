package types

import (
	"fmt"
)

type StartProcedureToolErrorStatus string

const (
	StartProcedureToolErrorStatusNotFound      StartProcedureToolErrorStatus = "not_found"
	StartProcedureToolErrorStatusInvalidName   StartProcedureToolErrorStatus = "invalid_name"
	StartProcedureToolErrorStatusAlreadyActive StartProcedureToolErrorStatus = "already_active"
)

func NewStartProcedureToolErrorStatusFromString(s string) (StartProcedureToolErrorStatus, error) {
	switch s {
	case "not_found":
		return StartProcedureToolErrorStatusNotFound, nil
	case "invalid_name":
		return StartProcedureToolErrorStatusInvalidName, nil
	case "already_active":
		return StartProcedureToolErrorStatusAlreadyActive, nil
	}
	var t StartProcedureToolErrorStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e StartProcedureToolErrorStatus) Ptr() *StartProcedureToolErrorStatus {
	return &e
}
