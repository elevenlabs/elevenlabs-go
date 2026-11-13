package types

import (
	"testing"
)

func TestBatchSuccessResponseModelAnalysisItemAttachmentResultDataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "success",
  "data": {
    "agent_id": "value",
    "branch_id": "value",
    "action": "attach",
    "changed": true
  }
}`
	assertRoundTrip(t, wire, new(BatchSuccessResponseModelAnalysisItemAttachmentResultData))
}
