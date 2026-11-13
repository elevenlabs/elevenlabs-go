package agents

import (
	"testing"
)

func TestPatchConvAiSettingsRequestRoundTrip(t *testing.T) {
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
  },
  "can_use_mcp_servers": true,
  "rag_retention_period_days": 30,
  "conversation_embedding_retention_days": 123,
  "default_livekit_stack": "standard"
}`
	assertRoundTrip(t, wire, new(PatchConvAiSettingsRequest))
}
