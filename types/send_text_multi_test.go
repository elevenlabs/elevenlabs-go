package types

import (
	"testing"
)

func TestSendTextMultiRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "context_id": "value",
  "flush": true
}`
	assertRoundTrip(t, wire, new(SendTextMulti))
}
