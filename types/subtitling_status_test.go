package types

import (
	"testing"
)

func TestSubtitlingStatusRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "status",
  "status": "value"
}`
	assertRoundTrip(t, wire, new(SubtitlingStatus))
}
