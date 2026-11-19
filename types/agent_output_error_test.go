package types

import (
	"testing"
)

func TestAgentOutputErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "value",
  "code": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(AgentOutputError))
}
