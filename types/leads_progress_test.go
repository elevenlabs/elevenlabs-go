package types

import (
	"testing"
)

func TestLeadsProgressRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "total": 123,
  "uncalled": 123,
  "leads_needing_attention": 123,
  "active": 123,
  "finished": 123,
  "finished_by_disposition": {
    "key": 123
  }
}`
	assertRoundTrip(t, wire, new(LeadsProgress))
}
