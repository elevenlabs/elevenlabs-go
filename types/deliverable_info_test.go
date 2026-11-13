package types

import (
	"testing"
)

func TestDeliverableInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "signed_url": "value",
  "content_type": "value",
  "name": "value",
  "version": 123
}`
	assertRoundTrip(t, wire, new(DeliverableInfo))
}
