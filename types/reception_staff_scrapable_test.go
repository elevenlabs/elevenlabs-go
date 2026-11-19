package types

import (
	"testing"
)

func TestReceptionStaffScrapableRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(ReceptionStaffScrapable))
}
