package types

import (
	"fmt"
)

type MetronomeUsageColumnFilterOperation string

const (
	MetronomeUsageColumnFilterOperationEq    MetronomeUsageColumnFilterOperation = "eq"
	MetronomeUsageColumnFilterOperationIn    MetronomeUsageColumnFilterOperation = "in"
	MetronomeUsageColumnFilterOperationNeq   MetronomeUsageColumnFilterOperation = "neq"
	MetronomeUsageColumnFilterOperationNotIn MetronomeUsageColumnFilterOperation = "not_in"
)

func NewMetronomeUsageColumnFilterOperationFromString(s string) (MetronomeUsageColumnFilterOperation, error) {
	switch s {
	case "eq":
		return MetronomeUsageColumnFilterOperationEq, nil
	case "in":
		return MetronomeUsageColumnFilterOperationIn, nil
	case "neq":
		return MetronomeUsageColumnFilterOperationNeq, nil
	case "not_in":
		return MetronomeUsageColumnFilterOperationNotIn, nil
	}
	var t MetronomeUsageColumnFilterOperation
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MetronomeUsageColumnFilterOperation) Ptr() *MetronomeUsageColumnFilterOperation {
	return &e
}
