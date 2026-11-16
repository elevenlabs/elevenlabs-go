package types

import (
	"testing"
)

func TestCreateAgentBranchResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "created_branch_id": "value",
  "created_version_id": "value"
}`
	assertRoundTrip(t, wire, new(CreateAgentBranchResponseModel))
}
