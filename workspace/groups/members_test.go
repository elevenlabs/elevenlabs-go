package groups

import (
	"testing"
)

func TestBodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "email": "value"
}`
	assertRoundTrip(t, wire, new(BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost))
}

func TestAddMemberToGroupRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "email": "value"
}`
	assertRoundTrip(t, wire, new(AddMemberToGroupRequest))
}
