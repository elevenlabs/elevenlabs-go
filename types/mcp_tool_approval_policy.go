package types

import (
	"fmt"
)

// Defines the tool-level approval policy.
type McpToolApprovalPolicy string

const (
	McpToolApprovalPolicyAutoApproved     McpToolApprovalPolicy = "auto_approved"
	McpToolApprovalPolicyRequiresApproval McpToolApprovalPolicy = "requires_approval"
)

func NewMcpToolApprovalPolicyFromString(s string) (McpToolApprovalPolicy, error) {
	switch s {
	case "auto_approved":
		return McpToolApprovalPolicyAutoApproved, nil
	case "requires_approval":
		return McpToolApprovalPolicyRequiresApproval, nil
	}
	var t McpToolApprovalPolicy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e McpToolApprovalPolicy) Ptr() *McpToolApprovalPolicy {
	return &e
}
