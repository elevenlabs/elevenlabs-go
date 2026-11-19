package types

import (
	"testing"
)

func TestAsrConversationalConfigOverrideConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "keywords": true
}`
	assertRoundTrip(t, wire, new(AsrConversationalConfigOverrideConfig))
}
