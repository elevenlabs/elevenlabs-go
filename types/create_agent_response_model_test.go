package types

import (
	"testing"
)

func TestCreateAgentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value"
}`
	assertRoundTrip(t, wire, new(CreateAgentResponseModel))
}
