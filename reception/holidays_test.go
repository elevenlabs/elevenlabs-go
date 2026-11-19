package reception

import (
	"testing"
)

func TestReceptionCreateHolidayRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value"
}`
	assertRoundTrip(t, wire, new(ReceptionCreateHolidayRequest))
}

func TestReceptionUpdateHolidayRequestPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value"
}`
	assertRoundTrip(t, wire, new(ReceptionUpdateHolidayRequestPatch))
}
