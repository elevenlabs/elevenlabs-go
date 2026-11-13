package types

import (
	"testing"
)

func TestAgentBranchBasicInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(AgentBranchBasicInfo))
}
