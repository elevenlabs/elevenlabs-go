package types

import (
	"testing"
)

func TestWorkspaceGroupByNameResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "id": "value",
  "members_emails": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(WorkspaceGroupByNameResponseModel))
}
