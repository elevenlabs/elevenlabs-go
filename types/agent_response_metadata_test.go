package types

import (
	"testing"
)

func TestAgentResponseMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_response_metadata_event": {
    "metadata": {},
    "event_id": 123
  },
  "type": "agent_response_metadata"
}`
	assertRoundTrip(t, wire, new(AgentResponseMetadata))
}
