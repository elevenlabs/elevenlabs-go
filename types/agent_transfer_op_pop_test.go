package types

import (
	"testing"
)

func TestAgentTransferOpPopRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "pop"
}`
	assertRoundTrip(t, wire, new(AgentTransferOpPop))
}
