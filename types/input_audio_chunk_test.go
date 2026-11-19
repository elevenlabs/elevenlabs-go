package types

import (
	"testing"
)

func TestInputAudioChunkRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "value",
  "audio_base_64": "value",
  "commit": true,
  "previous_text": "value",
  "keyterm_library": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(InputAudioChunk))
}
