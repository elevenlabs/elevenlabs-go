package types

import (
	"fmt"
)

// Whether a discovered MCP tool has an approval and whether it still matches.
type McpToolApprovalState string

const (
	McpToolApprovalStateUpToDate    McpToolApprovalState = "up_to_date"
	McpToolApprovalStateNeedsReview McpToolApprovalState = "needs_review"
	McpToolApprovalStateNotApproved McpToolApprovalState = "not_approved"
)

func NewMcpToolApprovalStateFromString(s string) (McpToolApprovalState, error) {
	switch s {
	case "up_to_date":
		return McpToolApprovalStateUpToDate, nil
	case "needs_review":
		return McpToolApprovalStateNeedsReview, nil
	case "not_approved":
		return McpToolApprovalStateNotApproved, nil
	}
	var t McpToolApprovalState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e McpToolApprovalState) Ptr() *McpToolApprovalState {
	return &e
}
