package types

import (
	"testing"
)

func TestWorkingHoursRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(WorkingHours))
}

func TestWorkingHoursRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(WorkingHours))
}

func TestWorkingHoursRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(WorkingHours))
}

func TestWorkingHoursRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(WorkingHours))
}
