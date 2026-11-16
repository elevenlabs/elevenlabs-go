package types

import (
	"testing"
)

func TestAsrConversationalConfigOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "keywords": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(AsrConversationalConfigOverride))
}

func TestAsrConversationalConfigOverrideRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "keywords": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(AsrConversationalConfigOverride))
}
