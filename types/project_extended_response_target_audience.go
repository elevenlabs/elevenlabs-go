package types

import (
	"fmt"
)

type ProjectExtendedResponseTargetAudience string

const (
	ProjectExtendedResponseTargetAudienceChildren   ProjectExtendedResponseTargetAudience = "children"
	ProjectExtendedResponseTargetAudienceYoungAdult ProjectExtendedResponseTargetAudience = "young adult"
	ProjectExtendedResponseTargetAudienceAdult      ProjectExtendedResponseTargetAudience = "adult"
	ProjectExtendedResponseTargetAudienceAllAges    ProjectExtendedResponseTargetAudience = "all ages"
)

func NewProjectExtendedResponseTargetAudienceFromString(s string) (ProjectExtendedResponseTargetAudience, error) {
	switch s {
	case "children":
		return ProjectExtendedResponseTargetAudienceChildren, nil
	case "young adult":
		return ProjectExtendedResponseTargetAudienceYoungAdult, nil
	case "adult":
		return ProjectExtendedResponseTargetAudienceAdult, nil
	case "all ages":
		return ProjectExtendedResponseTargetAudienceAllAges, nil
	}
	var t ProjectExtendedResponseTargetAudience
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectExtendedResponseTargetAudience) Ptr() *ProjectExtendedResponseTargetAudience {
	return &e
}
