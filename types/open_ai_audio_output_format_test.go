package types

import (
	"testing"
)

func TestOpenAiAudioOutputFormatRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "audio/pcm"
}`
	assertRoundTrip(t, wire, new(OpenAiAudioOutputFormat))
}
