package types

import (
	"fmt"
)

type EndProcedureToolErrorStatus string

const (
	EndProcedureToolErrorStatusNotFound  EndProcedureToolErrorStatus = "not_found"
	EndProcedureToolErrorStatusInvalidId EndProcedureToolErrorStatus = "invalid_id"
)

func NewEndProcedureToolErrorStatusFromString(s string) (EndProcedureToolErrorStatus, error) {
	switch s {
	case "not_found":
		return EndProcedureToolErrorStatusNotFound, nil
	case "invalid_id":
		return EndProcedureToolErrorStatusInvalidId, nil
	}
	var t EndProcedureToolErrorStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EndProcedureToolErrorStatus) Ptr() *EndProcedureToolErrorStatus {
	return &e
}
