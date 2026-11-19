package types

import (
	"testing"
)

func TestTestRunResultSummaryRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "test_id": "value",
  "test_name": "value",
  "workflow_node_id": "value",
  "buckets": [
    {
      "test_run_ids": [
        "value"
      ],
      "title": "value",
      "reason": "value",
      "status": "pending"
    }
  ]
}`
	assertRoundTrip(t, wire, new(TestRunResultSummary))
}
