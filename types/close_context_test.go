package types

import (
	"testing"
)

func TestCloseContextRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "context_id": "value",
  "close_context": true
}`
	assertRoundTrip(t, wire, new(CloseContext))
}
