package types

import (
	"testing"
)

func TestDeleteWorkspaceGroupMemberResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(DeleteWorkspaceGroupMemberResponseModel))
}
