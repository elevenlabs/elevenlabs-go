package types

import (
	"fmt"
)

type MergeProposalCloseReason string

const (
	MergeProposalCloseReasonWithdrawn      MergeProposalCloseReason = "withdrawn"
	MergeProposalCloseReasonRejected       MergeProposalCloseReason = "rejected"
	MergeProposalCloseReasonBranchArchived MergeProposalCloseReason = "branch_archived"
)

func NewMergeProposalCloseReasonFromString(s string) (MergeProposalCloseReason, error) {
	switch s {
	case "withdrawn":
		return MergeProposalCloseReasonWithdrawn, nil
	case "rejected":
		return MergeProposalCloseReasonRejected, nil
	case "branch_archived":
		return MergeProposalCloseReasonBranchArchived, nil
	}
	var t MergeProposalCloseReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MergeProposalCloseReason) Ptr() *MergeProposalCloseReason {
	return &e
}
