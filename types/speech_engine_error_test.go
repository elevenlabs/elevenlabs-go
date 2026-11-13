package types

import (
	"testing"
)

func TestSpeechEngineErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "error",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(SpeechEngineError))
}
