package types

import (
	"fmt"
)

type AnalysisScope string

const (
	AnalysisScopeConversation AnalysisScope = "conversation"
	AnalysisScopeAgent        AnalysisScope = "agent"
)

func NewAnalysisScopeFromString(s string) (AnalysisScope, error) {
	switch s {
	case "conversation":
		return AnalysisScopeConversation, nil
	case "agent":
		return AnalysisScopeAgent, nil
	}
	var t AnalysisScope
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AnalysisScope) Ptr() *AnalysisScope {
	return &e
}
