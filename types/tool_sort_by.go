package types

import (
	"fmt"
)

type ToolSortBy string

const (
	ToolSortByName      ToolSortBy = "name"
	ToolSortByCreatedAt ToolSortBy = "created_at"
)

func NewToolSortByFromString(s string) (ToolSortBy, error) {
	switch s {
	case "name":
		return ToolSortByName, nil
	case "created_at":
		return ToolSortByCreatedAt, nil
	}
	var t ToolSortBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ToolSortBy) Ptr() *ToolSortBy {
	return &e
}
