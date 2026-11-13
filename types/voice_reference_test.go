package types

import (
	"testing"
)

func TestVoiceReferenceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "voice",
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(VoiceReference))
}
