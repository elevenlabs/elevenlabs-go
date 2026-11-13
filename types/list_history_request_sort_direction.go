package types

import (
	"fmt"
)

type ListHistoryRequestSortDirection string

const (
	ListHistoryRequestSortDirectionAsc  ListHistoryRequestSortDirection = "asc"
	ListHistoryRequestSortDirectionDesc ListHistoryRequestSortDirection = "desc"
)

func NewListHistoryRequestSortDirectionFromString(s string) (ListHistoryRequestSortDirection, error) {
	switch s {
	case "asc":
		return ListHistoryRequestSortDirectionAsc, nil
	case "desc":
		return ListHistoryRequestSortDirectionDesc, nil
	}
	var t ListHistoryRequestSortDirection
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListHistoryRequestSortDirection) Ptr() *ListHistoryRequestSortDirection {
	return &e
}
