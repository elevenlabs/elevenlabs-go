package types

import (
	"testing"
)

func TestAgentResponseMetadataAgentResponseMetadataEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "metadata": {
    "key": {}
  },
  "event_id": 123
}`
	assertRoundTrip(t, wire, new(AgentResponseMetadataAgentResponseMetadataEvent))
}
