package types

import (
	"testing"
)

func TestAgentTransferOpRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "replace"
}`
	assertRoundTrip(t, wire, new(AgentTransferOp))
}
