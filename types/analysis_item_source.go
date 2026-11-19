package types

import (
	"fmt"
)

type AnalysisItemSource string

const (
	AnalysisItemSourceUser   AnalysisItemSource = "user"
	AnalysisItemSourceSystem AnalysisItemSource = "system"
)

func NewAnalysisItemSourceFromString(s string) (AnalysisItemSource, error) {
	switch s {
	case "user":
		return AnalysisItemSourceUser, nil
	case "system":
		return AnalysisItemSourceSystem, nil
	}
	var t AnalysisItemSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AnalysisItemSource) Ptr() *AnalysisItemSource {
	return &e
}
