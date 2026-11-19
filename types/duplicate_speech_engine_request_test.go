package types

import (
	"testing"
)

func TestDuplicateSpeechEngineRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value"
}`
	assertRoundTrip(t, wire, new(DuplicateSpeechEngineRequest))
}
