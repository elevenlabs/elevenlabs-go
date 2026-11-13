package types

import (
	"fmt"
)

// The field to use for ordering results from this query.
type ListDubbingRequestOrderBy string

const (
	ListDubbingRequestOrderByCreatedAt ListDubbingRequestOrderBy = "created_at"
	ListDubbingRequestOrderByName      ListDubbingRequestOrderBy = "name"
)

func NewListDubbingRequestOrderByFromString(s string) (ListDubbingRequestOrderBy, error) {
	switch s {
	case "created_at":
		return ListDubbingRequestOrderByCreatedAt, nil
	case "name":
		return ListDubbingRequestOrderByName, nil
	}
	var t ListDubbingRequestOrderBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListDubbingRequestOrderBy) Ptr() *ListDubbingRequestOrderBy {
	return &e
}
