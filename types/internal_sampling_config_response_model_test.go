package types

import (
	"testing"
)

func TestInternalSamplingConfigResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "focus": 0.5,
  "similarity": 0.5,
  "stability": 0.5
}`
	assertRoundTrip(t, wire, new(InternalSamplingConfigResponseModel))
}
