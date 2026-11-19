package types

import (
	"testing"
)

func TestWebhookSendResultResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "delivery_id": "value",
  "outcome": "delivered",
  "response_code": 123,
  "latency_ms": 123,
  "response_content_type": "value",
  "response_body_bytes": 123,
  "failure_origin": "value"
}`
	assertRoundTrip(t, wire, new(WebhookSendResultResponseModel))
}
