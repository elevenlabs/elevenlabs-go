package types

import (
	"testing"
)

func TestUtteranceResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start": 9.99,
  "end": 9.99
}`
	assertRoundTrip(t, wire, new(UtteranceResponseModel))
}
