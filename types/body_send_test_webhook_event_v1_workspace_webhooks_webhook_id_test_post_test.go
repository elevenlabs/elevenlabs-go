package types

import (
	"testing"
)

func TestBodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook_type": "project_conversion_status"
}`
	assertRoundTrip(t, wire, new(BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost))
}
