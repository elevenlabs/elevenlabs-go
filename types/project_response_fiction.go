package types

import (
	"fmt"
)

type ProjectResponseFiction string

const (
	ProjectResponseFictionFiction    ProjectResponseFiction = "fiction"
	ProjectResponseFictionNonFiction ProjectResponseFiction = "non-fiction"
)

func NewProjectResponseFictionFromString(s string) (ProjectResponseFiction, error) {
	switch s {
	case "fiction":
		return ProjectResponseFictionFiction, nil
	case "non-fiction":
		return ProjectResponseFictionNonFiction, nil
	}
	var t ProjectResponseFiction
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectResponseFiction) Ptr() *ProjectResponseFiction {
	return &e
}
