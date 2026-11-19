package types

import (
	"testing"
)

func TestTextOutputItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "text",
  "role": "agent",
  "id": "value",
  "text": "value",
  "created_at_unix_ms": 123
}`
	assertRoundTrip(t, wire, new(TextOutputItem))
}
