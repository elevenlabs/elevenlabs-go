package types

import (
	"testing"
)

func TestLeadDeliveryViewRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "pending",
  "stage": "queued",
  "sent_at_unix": 123,
  "delivered_at_unix": 123,
  "read_at_unix": 123,
  "replied_at_unix": 123,
  "failed_at_unix": 123,
  "error_message": "value"
}`
	assertRoundTrip(t, wire, new(LeadDeliveryView))
}

func TestLeadDeliveryViewRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "pending",
  "stage": "queued",
  "sent_at_unix": 123,
  "delivered_at_unix": 123,
  "read_at_unix": 123,
  "replied_at_unix": 123,
  "failed_at_unix": 123,
  "error_message": "value"
}`
	assertRoundTrip(t, wire, new(LeadDeliveryView))
}
