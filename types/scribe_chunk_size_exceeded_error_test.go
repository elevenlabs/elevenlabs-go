package types

import (
	"testing"
)

func TestScribeChunkSizeExceededErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "chunk_size_exceeded",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(ScribeChunkSizeExceededError))
}
