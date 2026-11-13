package types

import (
	"testing"
)

func TestWorkflowToolResponseModelInputStepsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "step_latency_secs": 9.99,
  "type": "edge",
  "edge_id": "value",
  "target_node_id": "value"
}`
	assertRoundTrip(t, wire, new(WorkflowToolResponseModelInputStepsItem))
}
