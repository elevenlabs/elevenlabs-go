package types

import (
	"testing"
)

func TestReceptionAgentInfoResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "business_name": "value",
  "business_country": "value",
  "business_description": "value",
  "timezone": "value",
  "currency": "USD",
  "operating_hours": {
    "hours": {
      "key": [
        {
          "start_hour": 123,
          "start_minute": 123,
          "end_hour": 24,
          "end_minute": 123
        }
      ]
    },
    "timezone": "value"
  },
  "tool_call_sound": "typing",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "workspace_now": {
    "today": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "hour": 123,
    "minute": 123,
    "fetched_at_ms": 123
  }
}`
	assertRoundTrip(t, wire, new(ReceptionAgentInfoResponse))
}
