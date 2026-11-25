package types

import (
	"testing"
)

func TestSubtitlingEndOfStreamRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "end_of_stream"
}`
	assertRoundTrip(t, wire, new(SubtitlingEndOfStream))
}
