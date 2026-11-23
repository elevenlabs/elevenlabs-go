package types

import (
	"testing"
)

func TestAgentResponseMetadataAgentResponseMetadataEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "metadata": {},
  "event_id": 123
}`
	assertRoundTrip(t, wire, new(AgentResponseMetadataAgentResponseMetadataEvent))
}
