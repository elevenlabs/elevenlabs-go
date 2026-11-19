package types

import (
	"testing"
)

func TestReceptionStaffResourceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "is_active": true,
  "email": "value",
  "phone": "value",
  "sms_notifications_opted_in": true,
  "photo_url": "value",
  "location_id": "value",
  "working_hours": {
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
  "display_on_booking_page": true,
  "customer_facing_name": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`
	assertRoundTrip(t, wire, new(ReceptionStaffResource))
}
