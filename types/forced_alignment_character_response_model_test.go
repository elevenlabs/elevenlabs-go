package types

import (
	"testing"
)

func TestForcedAlignmentCharacterResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "start": 9.99,
  "end": 9.99
}`
	assertRoundTrip(t, wire, new(ForcedAlignmentCharacterResponseModel))
}
