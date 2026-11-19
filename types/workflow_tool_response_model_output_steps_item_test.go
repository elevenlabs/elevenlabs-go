package types

import (
	"testing"
)

func TestWorkflowToolResponseModelOutputStepsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "step_latency_secs": 9.99,
  "type": "edge",
  "edge_id": "value",
  "target_node_id": "value"
}`
	assertRoundTrip(t, wire, new(WorkflowToolResponseModelOutputStepsItem))
}
