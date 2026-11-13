package types

import (
	"testing"
)

func TestOpenerConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "llm": "gpt-4o-mini"
}`
	assertRoundTrip(t, wire, new(OpenerConfig))
}
