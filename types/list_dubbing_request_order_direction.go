package types

import (
	"fmt"
)

// The order direction to use for results from this query.
type ListDubbingRequestOrderDirection string

const (
	ListDubbingRequestOrderDirectionDescending ListDubbingRequestOrderDirection = "DESCENDING"
	ListDubbingRequestOrderDirectionAscending  ListDubbingRequestOrderDirection = "ASCENDING"
)

func NewListDubbingRequestOrderDirectionFromString(s string) (ListDubbingRequestOrderDirection, error) {
	switch s {
	case "DESCENDING":
		return ListDubbingRequestOrderDirectionDescending, nil
	case "ASCENDING":
		return ListDubbingRequestOrderDirectionAscending, nil
	}
	var t ListDubbingRequestOrderDirection
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListDubbingRequestOrderDirection) Ptr() *ListDubbingRequestOrderDirection {
	return &e
}
