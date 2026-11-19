package types

import (
	"testing"
)

func TestForcedAlignmentWordResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "start": 9.99,
  "end": 9.99,
  "loss": 9.99
}`
	assertRoundTrip(t, wire, new(ForcedAlignmentWordResponseModel))
}
