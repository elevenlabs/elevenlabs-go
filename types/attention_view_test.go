package types

import (
	"testing"
)

func TestAttentionViewRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "attention_id": "value",
  "campaign_id": "value",
  "reason": "value",
  "detail": "value",
  "status": "open",
  "retryable": true,
  "resolution": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`
	assertRoundTrip(t, wire, new(AttentionView))
}
