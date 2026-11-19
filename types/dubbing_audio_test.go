package types

import (
	"testing"
)

func TestDubbingAudioRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "audio",
  "data": "value",
  "sample_rate": 123
}`
	assertRoundTrip(t, wire, new(DubbingAudio))
}
