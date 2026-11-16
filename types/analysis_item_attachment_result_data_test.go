package types

import (
	"testing"
)

func TestAnalysisItemAttachmentResultDataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "branch_id": "value",
  "action": "attach",
  "changed": true
}`
	assertRoundTrip(t, wire, new(AnalysisItemAttachmentResultData))
}
