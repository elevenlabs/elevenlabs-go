package types

import (
	"testing"
)

func TestUnitTestWorkflowNodeTransitionEvaluationNodeIdRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "node_id",
  "agent_id": "value",
  "target_node_id": "value"
}`
	assertRoundTrip(t, wire, new(UnitTestWorkflowNodeTransitionEvaluationNodeId))
}

func TestUnitTestWorkflowNodeTransitionEvaluationNodeIdRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "node_id",
  "agent_id": "value",
  "target_node_id": "value"
}`
	assertRoundTrip(t, wire, new(UnitTestWorkflowNodeTransitionEvaluationNodeId))
}
