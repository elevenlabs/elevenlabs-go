package types

import (
	"fmt"
)

type ProjectExtendedResponseAspectRatio string

const (
	ProjectExtendedResponseAspectRatioOne69  ProjectExtendedResponseAspectRatio = "16:9"
	ProjectExtendedResponseAspectRatioNine16 ProjectExtendedResponseAspectRatio = "9:16"
	ProjectExtendedResponseAspectRatioFour5  ProjectExtendedResponseAspectRatio = "4:5"
	ProjectExtendedResponseAspectRatioOne1   ProjectExtendedResponseAspectRatio = "1:1"
)

func NewProjectExtendedResponseAspectRatioFromString(s string) (ProjectExtendedResponseAspectRatio, error) {
	switch s {
	case "16:9":
		return ProjectExtendedResponseAspectRatioOne69, nil
	case "9:16":
		return ProjectExtendedResponseAspectRatioNine16, nil
	case "4:5":
		return ProjectExtendedResponseAspectRatioFour5, nil
	case "1:1":
		return ProjectExtendedResponseAspectRatioOne1, nil
	}
	var t ProjectExtendedResponseAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectExtendedResponseAspectRatio) Ptr() *ProjectExtendedResponseAspectRatio {
	return &e
}
