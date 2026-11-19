package types

import (
	"testing"
)

func TestAgentResponseClientEventDataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_response": "value",
  "event_id": 123,
  "response_id": "value",
  "in_response_to_ids": [
    "value"
  ],
  "response_type": "text",
  "attachments": [
    {
      "url": "value",
      "name": "value",
      "mime_type": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AgentResponseClientEventData))
}
