package types

import (
	"testing"
)

func TestDubbingLanguageOutputsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "lossless_audio": "value",
  "subtitles": "value"
}`
	assertRoundTrip(t, wire, new(DubbingLanguageOutputs))
}
