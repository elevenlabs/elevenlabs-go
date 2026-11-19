package types

import (
	"testing"
)

func TestRawLocationInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "address": "value",
  "operating_hours": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(RawLocationInfo))
}
