package types

import (
	"testing"
)

func TestWatermarkDetailsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "product": "value",
  "model": "value",
  "generated_at": "value",
  "unique_id": "value"
}`
	assertRoundTrip(t, wire, new(WatermarkDetails))
}
