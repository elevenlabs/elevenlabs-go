package types

import (
	"fmt"
)

// Voice category used for filtering
type FindSimilarVoicesVoicesRequestCategory string

const (
	FindSimilarVoicesVoicesRequestCategoryProfessional FindSimilarVoicesVoicesRequestCategory = "professional"
	FindSimilarVoicesVoicesRequestCategoryFamous       FindSimilarVoicesVoicesRequestCategory = "famous"
	FindSimilarVoicesVoicesRequestCategoryHighQuality  FindSimilarVoicesVoicesRequestCategory = "high_quality"
)

func NewFindSimilarVoicesVoicesRequestCategoryFromString(s string) (FindSimilarVoicesVoicesRequestCategory, error) {
	switch s {
	case "professional":
		return FindSimilarVoicesVoicesRequestCategoryProfessional, nil
	case "famous":
		return FindSimilarVoicesVoicesRequestCategoryFamous, nil
	case "high_quality":
		return FindSimilarVoicesVoicesRequestCategoryHighQuality, nil
	}
	var t FindSimilarVoicesVoicesRequestCategory
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e FindSimilarVoicesVoicesRequestCategory) Ptr() *FindSimilarVoicesVoicesRequestCategory {
	return &e
}
