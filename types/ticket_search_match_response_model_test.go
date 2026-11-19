package types

import (
	"testing"
)

func TestTicketSearchMatchResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "field": "title",
  "snippet": "value",
  "highlight_start": 123,
  "highlight_end": 123,
  "turn_index": 123
}`
	assertRoundTrip(t, wire, new(TicketSearchMatchResponseModel))
}
