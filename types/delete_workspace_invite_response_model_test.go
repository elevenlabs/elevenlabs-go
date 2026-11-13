package types

import (
	"testing"
)

func TestDeleteWorkspaceInviteResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(DeleteWorkspaceInviteResponseModel))
}
