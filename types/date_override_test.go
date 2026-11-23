package types

import (
	"testing"
)

func TestDateOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_date": "2024-01-01",
  "end_date": "2024-01-01",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(DateOverride))
}
