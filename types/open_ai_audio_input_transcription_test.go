package types

import (
	"testing"
)

func TestOpenAiAudioInputTranscriptionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "language": "value"
}`
	assertRoundTrip(t, wire, new(OpenAiAudioInputTranscription))
}
