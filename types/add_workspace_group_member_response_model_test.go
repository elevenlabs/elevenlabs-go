package types

import (
	"testing"
)

func TestAddWorkspaceGroupMemberResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(AddWorkspaceGroupMemberResponseModel))
}
