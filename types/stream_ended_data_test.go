package types

import (
	"testing"
)

func TestStreamEndedDataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "in_progress",
  "error": {
    "type": "value",
    "code": "value",
    "message": "value"
  },
  "required_actions": [
    {
      "type": "client_tool_call",
      "tool_call_id": "value",
      "tool_name": "value",
      "parameters": {
        "key": {}
      }
    }
  ]
}`
	assertRoundTrip(t, wire, new(StreamEndedData))
}
