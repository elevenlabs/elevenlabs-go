package types

import (
	"testing"
)

func TestAgentWorkspaceOverridesInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_initiation_client_data_webhook": {
    "url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "webhooks": {
    "post_call_webhook_id": "value",
    "events": [
      "transcript"
    ],
    "transcript_format": "json",
    "exclude_transcript": true,
    "send_audio": true
  }
}`
	assertRoundTrip(t, wire, new(AgentWorkspaceOverridesInput))
}
