package types

import (
	"testing"
)

func TestAgentMergeProposalResponseOutcomeRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "open"
}`
	assertRoundTrip(t, wire, new(AgentMergeProposalResponseOutcome))
}
