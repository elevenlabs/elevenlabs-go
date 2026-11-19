package types

import (
	"fmt"
)

// Sort criteria. Must be one of: created_date, usage_character_count_1y, trending, cloned_by_count.
type GetSharedVoicesRequestSort string

const (
	GetSharedVoicesRequestSortCreatedDate           GetSharedVoicesRequestSort = "created_date"
	GetSharedVoicesRequestSortUsageCharacterCount1Y GetSharedVoicesRequestSort = "usage_character_count_1y"
	GetSharedVoicesRequestSortTrending              GetSharedVoicesRequestSort = "trending"
	GetSharedVoicesRequestSortClonedByCount         GetSharedVoicesRequestSort = "cloned_by_count"
)

func NewGetSharedVoicesRequestSortFromString(s string) (GetSharedVoicesRequestSort, error) {
	switch s {
	case "created_date":
		return GetSharedVoicesRequestSortCreatedDate, nil
	case "usage_character_count_1y":
		return GetSharedVoicesRequestSortUsageCharacterCount1Y, nil
	case "trending":
		return GetSharedVoicesRequestSortTrending, nil
	case "cloned_by_count":
		return GetSharedVoicesRequestSortClonedByCount, nil
	}
	var t GetSharedVoicesRequestSort
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GetSharedVoicesRequestSort) Ptr() *GetSharedVoicesRequestSort {
	return &e
}
