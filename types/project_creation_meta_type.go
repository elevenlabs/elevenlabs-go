package types

import (
	"fmt"
)

type ProjectCreationMetaType string

const (
	ProjectCreationMetaTypeBlank            ProjectCreationMetaType = "blank"
	ProjectCreationMetaTypeGeneratePodcast  ProjectCreationMetaType = "generate_podcast"
	ProjectCreationMetaTypeAutoAssignVoices ProjectCreationMetaType = "auto_assign_voices"
	ProjectCreationMetaTypeDubVideo         ProjectCreationMetaType = "dub_video"
	ProjectCreationMetaTypeImportSpeech     ProjectCreationMetaType = "import_speech"
)

func NewProjectCreationMetaTypeFromString(s string) (ProjectCreationMetaType, error) {
	switch s {
	case "blank":
		return ProjectCreationMetaTypeBlank, nil
	case "generate_podcast":
		return ProjectCreationMetaTypeGeneratePodcast, nil
	case "auto_assign_voices":
		return ProjectCreationMetaTypeAutoAssignVoices, nil
	case "dub_video":
		return ProjectCreationMetaTypeDubVideo, nil
	case "import_speech":
		return ProjectCreationMetaTypeImportSpeech, nil
	}
	var t ProjectCreationMetaType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectCreationMetaType) Ptr() *ProjectCreationMetaType {
	return &e
}
