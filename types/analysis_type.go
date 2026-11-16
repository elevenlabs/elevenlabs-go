package types

import (
	"fmt"
)

type AnalysisType string

const (
	AnalysisTypeEvaluationCriteria AnalysisType = "evaluation_criteria"
	AnalysisTypeDataCollection     AnalysisType = "data_collection"
)

func NewAnalysisTypeFromString(s string) (AnalysisType, error) {
	switch s {
	case "evaluation_criteria":
		return AnalysisTypeEvaluationCriteria, nil
	case "data_collection":
		return AnalysisTypeDataCollection, nil
	}
	var t AnalysisType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AnalysisType) Ptr() *AnalysisType {
	return &e
}
