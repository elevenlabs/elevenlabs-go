package reception

import (
	"testing"
)

func TestReceptionUpdateAgentInfoRequestRoundTrip(t *testing.T) {
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
  "tool_call_sound": "typing"
}`
	assertRoundTrip(t, wire, new(ReceptionUpdateAgentInfoRequest))
}
