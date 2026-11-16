package types

import (
	"fmt"
)

type ProjectResponseSourceType string

const (
	ProjectResponseSourceTypeBlank      ProjectResponseSourceType = "blank"
	ProjectResponseSourceTypeBook       ProjectResponseSourceType = "book"
	ProjectResponseSourceTypeArticle    ProjectResponseSourceType = "article"
	ProjectResponseSourceTypeGenfm      ProjectResponseSourceType = "genfm"
	ProjectResponseSourceTypeVideo      ProjectResponseSourceType = "video"
	ProjectResponseSourceTypeScreenplay ProjectResponseSourceType = "screenplay"
)

func NewProjectResponseSourceTypeFromString(s string) (ProjectResponseSourceType, error) {
	switch s {
	case "blank":
		return ProjectResponseSourceTypeBlank, nil
	case "book":
		return ProjectResponseSourceTypeBook, nil
	case "article":
		return ProjectResponseSourceTypeArticle, nil
	case "genfm":
		return ProjectResponseSourceTypeGenfm, nil
	case "video":
		return ProjectResponseSourceTypeVideo, nil
	case "screenplay":
		return ProjectResponseSourceTypeScreenplay, nil
	}
	var t ProjectResponseSourceType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectResponseSourceType) Ptr() *ProjectResponseSourceType {
	return &e
}
