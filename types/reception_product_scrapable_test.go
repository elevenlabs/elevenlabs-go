package types

import (
	"testing"
)

func TestReceptionProductScrapableRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "pricing_mode": "fixed",
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99
}`
	assertRoundTrip(t, wire, new(ReceptionProductScrapable))
}
