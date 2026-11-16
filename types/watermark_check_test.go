package types

import (
	"testing"
)

func TestWatermarkCheckRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_detected": true,
  "details": {
    "product": "value",
    "model": "value",
    "generated_at": "value",
    "unique_id": "value"
  }
}`
	assertRoundTrip(t, wire, new(WatermarkCheck))
}
