package types

import (
	"testing"
)

func TestSuggestedAudioTagRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tag": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(SuggestedAudioTag))
}

func TestSuggestedAudioTagRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tag": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(SuggestedAudioTag))
}
