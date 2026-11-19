package types

import (
	"testing"
)

func TestWebhookTargetIdsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "ids",
  "ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(WebhookTargetIds))
}
