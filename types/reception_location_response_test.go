package types

import (
	"testing"
)

func TestReceptionLocationResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "address": "value",
  "notes": "value",
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
  "customer_facing_name": "value",
  "max_concurrent_appointments": 123,
  "id": "value",
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "is_active": true
}`
	assertRoundTrip(t, wire, new(ReceptionLocationResponse))
}
