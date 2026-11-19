package types

import (
	"testing"
)

func TestTimeRangeModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_hour": 123,
  "start_minute": 123,
  "end_hour": 24,
  "end_minute": 123
}`
	assertRoundTrip(t, wire, new(TimeRangeModel))
}

func TestTimeRangeModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_hour": 123,
  "start_minute": 123,
  "end_hour": 24,
  "end_minute": 123
}`
	assertRoundTrip(t, wire, new(TimeRangeModel))
}
