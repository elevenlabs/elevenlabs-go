package types

import (
	"testing"
)

func TestConversationHistoryTranscriptToolCallCommonModelInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "system",
  "request_id": "value",
  "tool_name": "value",
  "tool_id": "value",
  "params_as_json": "value",
  "tool_has_been_called": true,
  "tool_details": {
    "type": "webhook",
    "method": "value",
    "url": "value",
    "headers": {
      "key": "value"
    },
    "path_params": {
      "key": "value"
    },
    "query_params": {
      "key": "value"
    },
    "body": "value"
  }
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptToolCallCommonModelInput))
}
