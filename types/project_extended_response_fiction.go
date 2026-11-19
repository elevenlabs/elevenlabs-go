package types

import (
	"fmt"
)

type ProjectExtendedResponseFiction string

const (
	ProjectExtendedResponseFictionFiction    ProjectExtendedResponseFiction = "fiction"
	ProjectExtendedResponseFictionNonFiction ProjectExtendedResponseFiction = "non-fiction"
)

func NewProjectExtendedResponseFictionFromString(s string) (ProjectExtendedResponseFiction, error) {
	switch s {
	case "fiction":
		return ProjectExtendedResponseFictionFiction, nil
	case "non-fiction":
		return ProjectExtendedResponseFictionNonFiction, nil
	}
	var t ProjectExtendedResponseFiction
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectExtendedResponseFiction) Ptr() *ProjectExtendedResponseFiction {
	return &e
}
