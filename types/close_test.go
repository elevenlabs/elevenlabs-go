package types

import (
	"testing"
)

func TestCloseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "close"
}`
	assertRoundTrip(t, wire, new(Close))
}
