package types

import (
	"testing"
)

func TestPongRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "pong"
}`
	assertRoundTrip(t, wire, new(Pong))
}
