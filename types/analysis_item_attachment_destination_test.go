package types

import (
	"testing"
)

func TestAnalysisItemAttachmentDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "branch_id": "value"
}`
	assertRoundTrip(t, wire, new(AnalysisItemAttachmentDestination))
}
