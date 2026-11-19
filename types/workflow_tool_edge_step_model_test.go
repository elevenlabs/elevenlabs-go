package types

import (
	"testing"
)

func TestWorkflowToolEdgeStepModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "step_latency_secs": 9.99,
  "type": "edge",
  "edge_id": "value",
  "target_node_id": "value"
}`
	assertRoundTrip(t, wire, new(WorkflowToolEdgeStepModel))
}

func TestWorkflowToolEdgeStepModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "step_latency_secs": 9.99,
  "type": "edge",
  "edge_id": "value",
  "target_node_id": "value"
}`
	assertRoundTrip(t, wire, new(WorkflowToolEdgeStepModel))
}
