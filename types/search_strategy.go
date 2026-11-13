package types

import (
	"fmt"
)

type SearchStrategy string

const (
	SearchStrategyCat      SearchStrategy = "cat"
	SearchStrategyKeyword  SearchStrategy = "keyword"
	SearchStrategySemantic SearchStrategy = "semantic"
	SearchStrategyLs       SearchStrategy = "ls"
)

func NewSearchStrategyFromString(s string) (SearchStrategy, error) {
	switch s {
	case "cat":
		return SearchStrategyCat, nil
	case "keyword":
		return SearchStrategyKeyword, nil
	case "semantic":
		return SearchStrategySemantic, nil
	case "ls":
		return SearchStrategyLs, nil
	}
	var t SearchStrategy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SearchStrategy) Ptr() *SearchStrategy {
	return &e
}
