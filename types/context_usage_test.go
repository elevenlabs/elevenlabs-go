package types

import (
	"testing"
)

func TestContextUsageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "context_usage_event": {
    "event_id": 123,
    "model": "value",
    "context_tokens": 123,
    "context_limit_tokens": 123
  },
  "type": "context_usage"
}`
	assertRoundTrip(t, wire, new(ContextUsage))
}
