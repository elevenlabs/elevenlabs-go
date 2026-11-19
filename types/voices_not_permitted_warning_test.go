package types

import (
	"testing"
)

func TestVoicesNotPermittedWarningRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "voices_not_permitted",
  "speaker_ids": [
    "value"
  ],
  "message": "value"
}`
	assertRoundTrip(t, wire, new(VoicesNotPermittedWarning))
}

func TestVoicesNotPermittedWarningRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "voices_not_permitted",
  "speaker_ids": [
    "value"
  ],
  "message": "value"
}`
	assertRoundTrip(t, wire, new(VoicesNotPermittedWarning))
}
