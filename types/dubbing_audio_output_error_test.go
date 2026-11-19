package types

import (
	"testing"
)

func TestDubbingAudioOutputErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "audio_output_error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(DubbingAudioOutputError))
}
