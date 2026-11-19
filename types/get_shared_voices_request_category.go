package types

import (
	"fmt"
)

// Voice category used for filtering
type GetSharedVoicesRequestCategory string

const (
	GetSharedVoicesRequestCategoryProfessional GetSharedVoicesRequestCategory = "professional"
	GetSharedVoicesRequestCategoryFamous       GetSharedVoicesRequestCategory = "famous"
	GetSharedVoicesRequestCategoryHighQuality  GetSharedVoicesRequestCategory = "high_quality"
)

func NewGetSharedVoicesRequestCategoryFromString(s string) (GetSharedVoicesRequestCategory, error) {
	switch s {
	case "professional":
		return GetSharedVoicesRequestCategoryProfessional, nil
	case "famous":
		return GetSharedVoicesRequestCategoryFamous, nil
	case "high_quality":
		return GetSharedVoicesRequestCategoryHighQuality, nil
	}
	var t GetSharedVoicesRequestCategory
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GetSharedVoicesRequestCategory) Ptr() *GetSharedVoicesRequestCategory {
	return &e
}
