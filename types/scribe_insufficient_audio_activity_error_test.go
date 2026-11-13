package types

import (
	"testing"
)

func TestScribeInsufficientAudioActivityErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "insufficient_audio_activity",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(ScribeInsufficientAudioActivityError))
}
