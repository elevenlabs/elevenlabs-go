package types

import (
	"testing"
)

func TestAgentTransferOpPushRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "push",
  "return_node_id": "value"
}`
	assertRoundTrip(t, wire, new(AgentTransferOpPush))
}
