package types

import (
	"fmt"
)

type MergeProposalReviewReviewerRole string

const (
	MergeProposalReviewReviewerRoleAdmin     MergeProposalReviewReviewerRole = "admin"
	MergeProposalReviewReviewerRoleEditor    MergeProposalReviewReviewerRole = "editor"
	MergeProposalReviewReviewerRoleCommenter MergeProposalReviewReviewerRole = "commenter"
	MergeProposalReviewReviewerRoleViewer    MergeProposalReviewReviewerRole = "viewer"
)

func NewMergeProposalReviewReviewerRoleFromString(s string) (MergeProposalReviewReviewerRole, error) {
	switch s {
	case "admin":
		return MergeProposalReviewReviewerRoleAdmin, nil
	case "editor":
		return MergeProposalReviewReviewerRoleEditor, nil
	case "commenter":
		return MergeProposalReviewReviewerRoleCommenter, nil
	case "viewer":
		return MergeProposalReviewReviewerRoleViewer, nil
	}
	var t MergeProposalReviewReviewerRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MergeProposalReviewReviewerRole) Ptr() *MergeProposalReviewReviewerRole {
	return &e
}
