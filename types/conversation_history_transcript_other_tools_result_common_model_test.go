package types

import (
	"testing"
)

func TestConversationHistoryTranscriptOtherToolsResultCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "request_id": "value",
  "tool_name": "value",
  "tool_id": "value",
  "result_value": "value",
  "is_error": true,
  "is_blocked": true,
  "tool_has_been_called": true,
  "tool_latency_secs": 9.99,
  "error_type": "value",
  "raw_error_message": "value",
  "dynamic_variable_updates": [
    {
      "variable_name": "value",
      "old_value": "value",
      "new_value": "value",
      "updated_at": 9.99,
      "tool_name": "value",
      "tool_request_id": "value"
    }
  ],
  "type": "client"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptOtherToolsResultCommonModel))
}

func TestConversationHistoryTranscriptOtherToolsResultCommonModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "request_id": "value",
  "tool_name": "value",
  "tool_id": "value",
  "result_value": "value",
  "is_error": true,
  "is_blocked": true,
  "tool_has_been_called": true,
  "tool_latency_secs": 9.99,
  "error_type": "value",
  "raw_error_message": "value",
  "dynamic_variable_updates": [
    {
      "variable_name": "value",
      "old_value": "value",
      "new_value": "value",
      "updated_at": 9.99,
      "tool_name": "value",
      "tool_request_id": "value"
    }
  ],
  "type": "client"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptOtherToolsResultCommonModel))
}

func TestConversationHistoryTranscriptOtherToolsResultCommonModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "request_id": "value",
  "tool_name": "value",
  "tool_id": "value",
  "result_value": "value",
  "is_error": true,
  "is_blocked": true,
  "tool_has_been_called": true,
  "tool_latency_secs": 9.99,
  "error_type": "value",
  "raw_error_message": "value",
  "dynamic_variable_updates": [
    {
      "variable_name": "value",
      "old_value": "value",
      "new_value": "value",
      "updated_at": 9.99,
      "tool_name": "value",
      "tool_request_id": "value"
    }
  ],
  "type": "client"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptOtherToolsResultCommonModel))
}
