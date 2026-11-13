package types

import (
	"fmt"
)

type MergeProposalStatus string

const (
	MergeProposalStatusOpen   MergeProposalStatus = "open"
	MergeProposalStatusMerged MergeProposalStatus = "merged"
	MergeProposalStatusClosed MergeProposalStatus = "closed"
)

func NewMergeProposalStatusFromString(s string) (MergeProposalStatus, error) {
	switch s {
	case "open":
		return MergeProposalStatusOpen, nil
	case "merged":
		return MergeProposalStatusMerged, nil
	case "closed":
		return MergeProposalStatusClosed, nil
	}
	var t MergeProposalStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MergeProposalStatus) Ptr() *MergeProposalStatus {
	return &e
}
