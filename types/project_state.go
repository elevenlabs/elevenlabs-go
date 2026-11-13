package types

import (
	"fmt"
)

// The state of the project.
type ProjectState string

const (
	ProjectStateCreating   ProjectState = "creating"
	ProjectStateDefault    ProjectState = "default"
	ProjectStateConverting ProjectState = "converting"
	ProjectStateInQueue    ProjectState = "in_queue"
)

func NewProjectStateFromString(s string) (ProjectState, error) {
	switch s {
	case "creating":
		return ProjectStateCreating, nil
	case "default":
		return ProjectStateDefault, nil
	case "converting":
		return ProjectStateConverting, nil
	case "in_queue":
		return ProjectStateInQueue, nil
	}
	var t ProjectState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectState) Ptr() *ProjectState {
	return &e
}
