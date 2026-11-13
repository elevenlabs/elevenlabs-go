package types

import (
	"testing"
)

func TestForcedAlignmentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "characters": [
    {
      "text": "value",
      "start": 9.99,
      "end": 9.99
    }
  ],
  "words": [
    {
      "text": "value",
      "start": 9.99,
      "end": 9.99,
      "loss": 9.99
    }
  ],
  "loss": 9.99
}`
	assertRoundTrip(t, wire, new(ForcedAlignmentResponseModel))
}
