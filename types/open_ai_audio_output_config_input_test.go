package types

import (
	"testing"
)

func TestOpenAiAudioOutputConfigInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "format": {
    "type": "audio/pcm"
  },
  "voice": "value"
}`
	assertRoundTrip(t, wire, new(OpenAiAudioOutputConfigInput))
}
