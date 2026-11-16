package types

import (
	"fmt"
)

type AnalysisItemAttachmentAction string

const (
	AnalysisItemAttachmentActionAttach AnalysisItemAttachmentAction = "attach"
	AnalysisItemAttachmentActionDetach AnalysisItemAttachmentAction = "detach"
)

func NewAnalysisItemAttachmentActionFromString(s string) (AnalysisItemAttachmentAction, error) {
	switch s {
	case "attach":
		return AnalysisItemAttachmentActionAttach, nil
	case "detach":
		return AnalysisItemAttachmentActionDetach, nil
	}
	var t AnalysisItemAttachmentAction
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AnalysisItemAttachmentAction) Ptr() *AnalysisItemAttachmentAction {
	return &e
}
