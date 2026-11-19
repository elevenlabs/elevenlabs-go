package types

import (
	"testing"
)

func TestSpeakerUpdatedResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": 123
}`
	assertRoundTrip(t, wire, new(SpeakerUpdatedResponse))
}
