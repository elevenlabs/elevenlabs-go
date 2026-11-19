package types

import (
	"fmt"
)

// The status of the project creation action.
type ProjectCreationMetaResponseModelStatus string

const (
	ProjectCreationMetaResponseModelStatusDraft    ProjectCreationMetaResponseModelStatus = "draft"
	ProjectCreationMetaResponseModelStatusPending  ProjectCreationMetaResponseModelStatus = "pending"
	ProjectCreationMetaResponseModelStatusCreating ProjectCreationMetaResponseModelStatus = "creating"
	ProjectCreationMetaResponseModelStatusFinished ProjectCreationMetaResponseModelStatus = "finished"
	ProjectCreationMetaResponseModelStatusFailed   ProjectCreationMetaResponseModelStatus = "failed"
)

func NewProjectCreationMetaResponseModelStatusFromString(s string) (ProjectCreationMetaResponseModelStatus, error) {
	switch s {
	case "draft":
		return ProjectCreationMetaResponseModelStatusDraft, nil
	case "pending":
		return ProjectCreationMetaResponseModelStatusPending, nil
	case "creating":
		return ProjectCreationMetaResponseModelStatusCreating, nil
	case "finished":
		return ProjectCreationMetaResponseModelStatusFinished, nil
	case "failed":
		return ProjectCreationMetaResponseModelStatusFailed, nil
	}
	var t ProjectCreationMetaResponseModelStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectCreationMetaResponseModelStatus) Ptr() *ProjectCreationMetaResponseModelStatus {
	return &e
}
