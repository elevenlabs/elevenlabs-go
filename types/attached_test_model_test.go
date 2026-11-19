package types

import (
	"testing"
)

func TestAttachedTestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "test_id": "value",
  "workflow_node_id": "value"
}`
	assertRoundTrip(t, wire, new(AttachedTestModel))
}
