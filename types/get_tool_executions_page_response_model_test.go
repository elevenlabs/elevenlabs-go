package types

import (
	"testing"
)

func TestGetToolExecutionsPageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "executions": [
    {
      "tool_id": "value",
      "tool_request_id": "value",
      "conversation_id": "value",
      "agent_id": "value",
      "branch_id": "value",
      "timestamp": 9.99,
      "latency_secs": 9.99,
      "is_error": true,
      "request_payload": "value",
      "response_payload": "value",
      "error_message": "value",
      "error_type": "value",
      "id": "value",
      "tool_call_details": {
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
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(GetToolExecutionsPageResponseModel))
}
