package types

import (
	"fmt"
)

type ProjectResponseAspectRatio string

const (
	ProjectResponseAspectRatioOne69  ProjectResponseAspectRatio = "16:9"
	ProjectResponseAspectRatioNine16 ProjectResponseAspectRatio = "9:16"
	ProjectResponseAspectRatioFour5  ProjectResponseAspectRatio = "4:5"
	ProjectResponseAspectRatioOne1   ProjectResponseAspectRatio = "1:1"
)

func NewProjectResponseAspectRatioFromString(s string) (ProjectResponseAspectRatio, error) {
	switch s {
	case "16:9":
		return ProjectResponseAspectRatioOne69, nil
	case "9:16":
		return ProjectResponseAspectRatioNine16, nil
	case "4:5":
		return ProjectResponseAspectRatioFour5, nil
	case "1:1":
		return ProjectResponseAspectRatioOne1, nil
	}
	var t ProjectResponseAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectResponseAspectRatio) Ptr() *ProjectResponseAspectRatio {
	return &e
}
