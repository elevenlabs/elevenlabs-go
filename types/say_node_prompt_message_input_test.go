package types

import (
	"testing"
)

func TestSayNodePromptMessageInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "prompt",
  "prompt": "value"
}`
	assertRoundTrip(t, wire, new(SayNodePromptMessageInput))
}
