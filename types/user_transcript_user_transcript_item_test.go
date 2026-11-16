package types

import (
	"testing"
)

func TestUserTranscriptUserTranscriptItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "role": "value",
  "content": "value"
}`
	assertRoundTrip(t, wire, new(UserTranscriptUserTranscriptItem))
}
