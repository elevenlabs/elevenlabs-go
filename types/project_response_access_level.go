package types

import (
	"fmt"
)

// The access level of the project.
type ProjectResponseAccessLevel string

const (
	ProjectResponseAccessLevelAdmin     ProjectResponseAccessLevel = "admin"
	ProjectResponseAccessLevelEditor    ProjectResponseAccessLevel = "editor"
	ProjectResponseAccessLevelCommenter ProjectResponseAccessLevel = "commenter"
	ProjectResponseAccessLevelViewer    ProjectResponseAccessLevel = "viewer"
)

func NewProjectResponseAccessLevelFromString(s string) (ProjectResponseAccessLevel, error) {
	switch s {
	case "admin":
		return ProjectResponseAccessLevelAdmin, nil
	case "editor":
		return ProjectResponseAccessLevelEditor, nil
	case "commenter":
		return ProjectResponseAccessLevelCommenter, nil
	case "viewer":
		return ProjectResponseAccessLevelViewer, nil
	}
	var t ProjectResponseAccessLevel
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectResponseAccessLevel) Ptr() *ProjectResponseAccessLevel {
	return &e
}
