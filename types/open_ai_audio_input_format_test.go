package types

import (
	"testing"
)

func TestOpenAiAudioInputFormatRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "audio/pcm",
  "rate": 123
}`
	assertRoundTrip(t, wire, new(OpenAiAudioInputFormat))
}
