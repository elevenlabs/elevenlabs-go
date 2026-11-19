package types

import (
	"testing"
)

func TestAddWorkspaceInviteResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(AddWorkspaceInviteResponseModel))
}
