package types

import (
	"testing"
)

func TestInitRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "init",
  "conversation_id": "value"
}`
	assertRoundTrip(t, wire, new(Init))
}
