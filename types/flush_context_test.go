package types

import (
	"testing"
)

func TestFlushContextRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "context_id": "value",
  "text": "value",
  "flush": true
}`
	assertRoundTrip(t, wire, new(FlushContext))
}
