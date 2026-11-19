package types

import (
	"fmt"
)

type AnalysisPropertyType string

const (
	AnalysisPropertyTypeBoolean AnalysisPropertyType = "boolean"
	AnalysisPropertyTypeString  AnalysisPropertyType = "string"
	AnalysisPropertyTypeInteger AnalysisPropertyType = "integer"
	AnalysisPropertyTypeNumber  AnalysisPropertyType = "number"
)

func NewAnalysisPropertyTypeFromString(s string) (AnalysisPropertyType, error) {
	switch s {
	case "boolean":
		return AnalysisPropertyTypeBoolean, nil
	case "string":
		return AnalysisPropertyTypeString, nil
	case "integer":
		return AnalysisPropertyTypeInteger, nil
	case "number":
		return AnalysisPropertyTypeNumber, nil
	}
	var t AnalysisPropertyType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AnalysisPropertyType) Ptr() *AnalysisPropertyType {
	return &e
}
