package types

import (
	"fmt"
)

type KnowledgeBaseDependentType string

const (
	KnowledgeBaseDependentTypeDirect     KnowledgeBaseDependentType = "direct"
	KnowledgeBaseDependentTypeTransitive KnowledgeBaseDependentType = "transitive"
	KnowledgeBaseDependentTypeAll        KnowledgeBaseDependentType = "all"
)

func NewKnowledgeBaseDependentTypeFromString(s string) (KnowledgeBaseDependentType, error) {
	switch s {
	case "direct":
		return KnowledgeBaseDependentTypeDirect, nil
	case "transitive":
		return KnowledgeBaseDependentTypeTransitive, nil
	case "all":
		return KnowledgeBaseDependentTypeAll, nil
	}
	var t KnowledgeBaseDependentType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e KnowledgeBaseDependentType) Ptr() *KnowledgeBaseDependentType {
	return &e
}
