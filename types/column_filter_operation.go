package types

import (
	"fmt"
)

type ColumnFilterOperation string

const (
	ColumnFilterOperationIn    ColumnFilterOperation = "in"
	ColumnFilterOperationNotIn ColumnFilterOperation = "not_in"
	ColumnFilterOperationLe    ColumnFilterOperation = "le"
	ColumnFilterOperationGe    ColumnFilterOperation = "ge"
	ColumnFilterOperationLt    ColumnFilterOperation = "lt"
	ColumnFilterOperationGt    ColumnFilterOperation = "gt"
	ColumnFilterOperationEq    ColumnFilterOperation = "eq"
	ColumnFilterOperationNeq   ColumnFilterOperation = "neq"
)

func NewColumnFilterOperationFromString(s string) (ColumnFilterOperation, error) {
	switch s {
	case "in":
		return ColumnFilterOperationIn, nil
	case "not_in":
		return ColumnFilterOperationNotIn, nil
	case "le":
		return ColumnFilterOperationLe, nil
	case "ge":
		return ColumnFilterOperationGe, nil
	case "lt":
		return ColumnFilterOperationLt, nil
	case "gt":
		return ColumnFilterOperationGt, nil
	case "eq":
		return ColumnFilterOperationEq, nil
	case "neq":
		return ColumnFilterOperationNeq, nil
	}
	var t ColumnFilterOperation
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ColumnFilterOperation) Ptr() *ColumnFilterOperation {
	return &e
}
