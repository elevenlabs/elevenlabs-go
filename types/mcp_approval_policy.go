package types

import (
	"fmt"
)

// Defines the MCP server-level approval policy for tool execution.
type McpApprovalPolicy string

const (
	McpApprovalPolicyAutoApproveAll         McpApprovalPolicy = "auto_approve_all"
	McpApprovalPolicyRequireApprovalAll     McpApprovalPolicy = "require_approval_all"
	McpApprovalPolicyRequireApprovalPerTool McpApprovalPolicy = "require_approval_per_tool"
)

func NewMcpApprovalPolicyFromString(s string) (McpApprovalPolicy, error) {
	switch s {
	case "auto_approve_all":
		return McpApprovalPolicyAutoApproveAll, nil
	case "require_approval_all":
		return McpApprovalPolicyRequireApprovalAll, nil
	case "require_approval_per_tool":
		return McpApprovalPolicyRequireApprovalPerTool, nil
	}
	var t McpApprovalPolicy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e McpApprovalPolicy) Ptr() *McpApprovalPolicy {
	return &e
}
