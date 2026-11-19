package types

import (
	"testing"
)

func TestSayNodePromptMessageOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "prompt",
  "prompt": "value"
}`
	assertRoundTrip(t, wire, new(SayNodePromptMessageOutput))
}
