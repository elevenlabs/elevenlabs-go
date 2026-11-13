package types

import (
	"fmt"
)

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

func NewSortDirectionFromString(s string) (SortDirection, error) {
	switch s {
	case "asc":
		return SortDirectionAsc, nil
	case "desc":
		return SortDirectionDesc, nil
	}
	var t SortDirection
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SortDirection) Ptr() *SortDirection {
	return &e
}
