package types

import (
	"fmt"
)

type AgentSortBy string

const (
	AgentSortByName        AgentSortBy = "name"
	AgentSortByCreatedAt   AgentSortBy = "created_at"
	AgentSortByUpdatedAt   AgentSortBy = "updated_at"
	AgentSortByCallCount7D AgentSortBy = "call_count_7d"
)

func NewAgentSortByFromString(s string) (AgentSortBy, error) {
	switch s {
	case "name":
		return AgentSortByName, nil
	case "created_at":
		return AgentSortByCreatedAt, nil
	case "updated_at":
		return AgentSortByUpdatedAt, nil
	case "call_count_7d":
		return AgentSortByCallCount7D, nil
	}
	var t AgentSortBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentSortBy) Ptr() *AgentSortBy {
	return &e
}
