package types

import (
	"testing"
)

func TestAgentVersionParentsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "in_branch_parent_id": "value",
  "out_of_branch_parent_id": "value",
  "merged_into_branch_id": "value",
  "merged_from_branch_id": "value",
  "merged_from_version_id": "value",
  "rebased_from_version_id": "value"
}`
	assertRoundTrip(t, wire, new(AgentVersionParents))
}
