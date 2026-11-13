package types

import (
	"fmt"
)

type ProjectResponseTargetAudience string

const (
	ProjectResponseTargetAudienceChildren   ProjectResponseTargetAudience = "children"
	ProjectResponseTargetAudienceYoungAdult ProjectResponseTargetAudience = "young adult"
	ProjectResponseTargetAudienceAdult      ProjectResponseTargetAudience = "adult"
	ProjectResponseTargetAudienceAllAges    ProjectResponseTargetAudience = "all ages"
)

func NewProjectResponseTargetAudienceFromString(s string) (ProjectResponseTargetAudience, error) {
	switch s {
	case "children":
		return ProjectResponseTargetAudienceChildren, nil
	case "young adult":
		return ProjectResponseTargetAudienceYoungAdult, nil
	case "adult":
		return ProjectResponseTargetAudienceAdult, nil
	case "all ages":
		return ProjectResponseTargetAudienceAllAges, nil
	}
	var t ProjectResponseTargetAudience
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectResponseTargetAudience) Ptr() *ProjectResponseTargetAudience {
	return &e
}
