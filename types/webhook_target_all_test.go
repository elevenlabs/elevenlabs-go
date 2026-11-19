package types

import (
	"testing"
)

func TestWebhookTargetAllRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "all"
}`
	assertRoundTrip(t, wire, new(WebhookTargetAll))
}
