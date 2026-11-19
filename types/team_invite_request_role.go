package types

import (
	"fmt"
)

// The role the invitee joins with: 'workspace_admin' or 'workspace_member'.
type TeamInviteRequestRole string

const (
	TeamInviteRequestRoleWorkspaceAdmin  TeamInviteRequestRole = "workspace_admin"
	TeamInviteRequestRoleWorkspaceMember TeamInviteRequestRole = "workspace_member"
)

func NewTeamInviteRequestRoleFromString(s string) (TeamInviteRequestRole, error) {
	switch s {
	case "workspace_admin":
		return TeamInviteRequestRoleWorkspaceAdmin, nil
	case "workspace_member":
		return TeamInviteRequestRoleWorkspaceMember, nil
	}
	var t TeamInviteRequestRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TeamInviteRequestRole) Ptr() *TeamInviteRequestRole {
	return &e
}
