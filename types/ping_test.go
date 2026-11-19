package types

import (
	"testing"
)

func TestPingRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "ping"
}`
	assertRoundTrip(t, wire, new(Ping))
}
