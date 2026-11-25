package types

import (
	"testing"
)

func TestContextUsageContextUsageEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "event_id": 123,
  "model": "value",
  "context_tokens": 123,
  "context_limit_tokens": 123
}`
	assertRoundTrip(t, wire, new(ContextUsageContextUsageEvent))
}
