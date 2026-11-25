package types

import (
	"testing"
)

func TestContextualUpdateRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "contextual_update",
  "text": "value",
  "context_id": "value"
}`
	assertRoundTrip(t, wire, new(ContextualUpdate))
}
