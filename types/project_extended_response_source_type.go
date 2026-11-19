package types

import (
	"fmt"
)

type ProjectExtendedResponseSourceType string

const (
	ProjectExtendedResponseSourceTypeBlank      ProjectExtendedResponseSourceType = "blank"
	ProjectExtendedResponseSourceTypeBook       ProjectExtendedResponseSourceType = "book"
	ProjectExtendedResponseSourceTypeArticle    ProjectExtendedResponseSourceType = "article"
	ProjectExtendedResponseSourceTypeGenfm      ProjectExtendedResponseSourceType = "genfm"
	ProjectExtendedResponseSourceTypeVideo      ProjectExtendedResponseSourceType = "video"
	ProjectExtendedResponseSourceTypeScreenplay ProjectExtendedResponseSourceType = "screenplay"
)

func NewProjectExtendedResponseSourceTypeFromString(s string) (ProjectExtendedResponseSourceType, error) {
	switch s {
	case "blank":
		return ProjectExtendedResponseSourceTypeBlank, nil
	case "book":
		return ProjectExtendedResponseSourceTypeBook, nil
	case "article":
		return ProjectExtendedResponseSourceTypeArticle, nil
	case "genfm":
		return ProjectExtendedResponseSourceTypeGenfm, nil
	case "video":
		return ProjectExtendedResponseSourceTypeVideo, nil
	case "screenplay":
		return ProjectExtendedResponseSourceTypeScreenplay, nil
	}
	var t ProjectExtendedResponseSourceType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectExtendedResponseSourceType) Ptr() *ProjectExtendedResponseSourceType {
	return &e
}
