package types

import (
	"fmt"
)

type ConflictSection string

const (
	ConflictSectionConversationConfig ConflictSection = "conversation_config"
	ConflictSectionPlatformSettings   ConflictSection = "platform_settings"
	ConflictSectionProcedures         ConflictSection = "procedures"
	ConflictSectionWorkflow           ConflictSection = "workflow"
)

func NewConflictSectionFromString(s string) (ConflictSection, error) {
	switch s {
	case "conversation_config":
		return ConflictSectionConversationConfig, nil
	case "platform_settings":
		return ConflictSectionPlatformSettings, nil
	case "procedures":
		return ConflictSectionProcedures, nil
	case "workflow":
		return ConflictSectionWorkflow, nil
	}
	var t ConflictSection
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConflictSection) Ptr() *ConflictSection {
	return &e
}
