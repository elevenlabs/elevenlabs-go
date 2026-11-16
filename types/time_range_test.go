package types

import (
	"testing"
)

func TestTimeRangeRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_ms": 123,
  "end_ms": 123
}`
	assertRoundTrip(t, wire, new(TimeRange))
}
