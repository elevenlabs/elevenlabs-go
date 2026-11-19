package workspace

import (
	"testing"
)

func TestInviteUserRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "email": "value",
  "workspace_permission": "value",
  "seat_type": "workspace_admin",
  "group_ids": [
    "value"
  ],
  "usage_limit": 123
}`
	assertRoundTrip(t, wire, new(InviteUserRequest))
}

func TestBodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "emails": [
    "value"
  ],
  "seat_type": "workspace_admin",
  "group_ids": [
    "value"
  ],
  "usage_limit": 123
}`
	assertRoundTrip(t, wire, new(BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost))
}

func TestBodyDeleteExistingInvitationV1WorkspaceInvitesDeleteRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "email": "value"
}`
	assertRoundTrip(t, wire, new(BodyDeleteExistingInvitationV1WorkspaceInvitesDelete))
}
