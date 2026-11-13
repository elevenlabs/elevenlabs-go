package types

import (
	"testing"
)

func TestVisitedAgentRefRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "branch_id": "value"
}`
	assertRoundTrip(t, wire, new(VisitedAgentRef))
}
