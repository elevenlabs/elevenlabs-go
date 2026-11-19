package types

import (
	"fmt"
)

type AnalysisItemKind string

const (
	AnalysisItemKindEvaluation     AnalysisItemKind = "evaluation"
	AnalysisItemKindDataCollection AnalysisItemKind = "data_collection"
)

func NewAnalysisItemKindFromString(s string) (AnalysisItemKind, error) {
	switch s {
	case "evaluation":
		return AnalysisItemKindEvaluation, nil
	case "data_collection":
		return AnalysisItemKindDataCollection, nil
	}
	var t AnalysisItemKind
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AnalysisItemKind) Ptr() *AnalysisItemKind {
	return &e
}
