package types

import (
	"testing"
)

func TestOpenAiAudioOutputConfigOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "format": {
    "type": "audio/pcm"
  },
  "voice": "value"
}`
	assertRoundTrip(t, wire, new(OpenAiAudioOutputConfigOutput))
}
