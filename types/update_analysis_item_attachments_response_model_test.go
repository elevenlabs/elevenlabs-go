package types

import (
	"testing"
)

func TestUpdateAnalysisItemAttachmentsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "results": [
    {
      "status": "success",
      "data": {
        "agent_id": "value",
        "branch_id": "value",
        "action": "attach",
        "changed": true
      }
    }
  ]
}`
	assertRoundTrip(t, wire, new(UpdateAnalysisItemAttachmentsResponseModel))
}
