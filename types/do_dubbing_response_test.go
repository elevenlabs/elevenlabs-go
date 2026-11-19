package types

import (
	"testing"
)

func TestDoDubbingResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "dubbing_id": "value",
  "expected_duration_sec": 9.99
}`
	assertRoundTrip(t, wire, new(DoDubbingResponse))
}
