package types

import (
	"testing"
)

func TestWebhookTargetRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "all"
}`
	assertRoundTrip(t, wire, new(WebhookTarget))
}

func TestWebhookTargetRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "all"
}`
	assertRoundTrip(t, wire, new(WebhookTarget))
}

func TestWebhookTargetRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "all"
}`
	assertRoundTrip(t, wire, new(WebhookTarget))
}

func TestWebhookTargetRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "all"
}`
	assertRoundTrip(t, wire, new(WebhookTarget))
}
