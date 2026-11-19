package types

import (
	"testing"
)

func TestUpdateWorkspaceMemberResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(UpdateWorkspaceMemberResponseModel))
}
