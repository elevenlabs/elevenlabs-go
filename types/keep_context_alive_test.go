package types

import (
	"testing"
)

func TestKeepContextAliveRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "",
  "context_id": "value"
}`
	assertRoundTrip(t, wire, new(KeepContextAlive))
}
