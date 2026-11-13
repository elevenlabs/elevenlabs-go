package types

import (
	"testing"
)

func TestRecipientDeliveryCountsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "sent": 123,
  "delivered": 123,
  "read": 123,
  "replied": 123,
  "failed": 123
}`
	assertRoundTrip(t, wire, new(RecipientDeliveryCounts))
}
