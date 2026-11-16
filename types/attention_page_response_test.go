package types

import (
	"testing"
)

func TestAttentionPageResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "attention": [
    {
      "attention_id": "value",
      "campaign_id": "value",
      "reason": "value",
      "detail": "value",
      "status": "open",
      "retryable": true,
      "resolution": "value",
      "created_at_unix": 123,
      "updated_at_unix": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(AttentionPageResponse))
}
