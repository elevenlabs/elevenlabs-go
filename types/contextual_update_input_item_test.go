package types

import (
	"testing"
)

func TestContextualUpdateInputItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "contextual_update",
  "text": "value",
  "context_id": "value",
  "placement": "in_place"
}`
	assertRoundTrip(t, wire, new(ContextualUpdateInputItem))
}
