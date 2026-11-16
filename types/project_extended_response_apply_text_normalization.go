package types

import (
	"fmt"
)

// Whether text normalization is applied to the project.
type ProjectExtendedResponseApplyTextNormalization string

const (
	ProjectExtendedResponseApplyTextNormalizationAuto         ProjectExtendedResponseApplyTextNormalization = "auto"
	ProjectExtendedResponseApplyTextNormalizationOn           ProjectExtendedResponseApplyTextNormalization = "on"
	ProjectExtendedResponseApplyTextNormalizationOff          ProjectExtendedResponseApplyTextNormalization = "off"
	ProjectExtendedResponseApplyTextNormalizationApplyEnglish ProjectExtendedResponseApplyTextNormalization = "apply_english"
)

func NewProjectExtendedResponseApplyTextNormalizationFromString(s string) (ProjectExtendedResponseApplyTextNormalization, error) {
	switch s {
	case "auto":
		return ProjectExtendedResponseApplyTextNormalizationAuto, nil
	case "on":
		return ProjectExtendedResponseApplyTextNormalizationOn, nil
	case "off":
		return ProjectExtendedResponseApplyTextNormalizationOff, nil
	case "apply_english":
		return ProjectExtendedResponseApplyTextNormalizationApplyEnglish, nil
	}
	var t ProjectExtendedResponseApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProjectExtendedResponseApplyTextNormalization) Ptr() *ProjectExtendedResponseApplyTextNormalization {
	return &e
}
