package types

import (
	"testing"
)

func TestOrderDeliverablesResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "deliverables": [
    {
      "signed_url": "value",
      "content_type": "value",
      "name": "value",
      "version": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(OrderDeliverablesResponse))
}
