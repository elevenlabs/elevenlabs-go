package types

import (
	"fmt"
)

type MergeProposalReviewState string

const (
	MergeProposalReviewStateApproved         MergeProposalReviewState = "approved"
	MergeProposalReviewStateChangesRequested MergeProposalReviewState = "changes_requested"
)

func NewMergeProposalReviewStateFromString(s string) (MergeProposalReviewState, error) {
	switch s {
	case "approved":
		return MergeProposalReviewStateApproved, nil
	case "changes_requested":
		return MergeProposalReviewStateChangesRequested, nil
	}
	var t MergeProposalReviewState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MergeProposalReviewState) Ptr() *MergeProposalReviewState {
	return &e
}
