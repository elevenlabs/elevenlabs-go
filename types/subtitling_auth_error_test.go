package types

import (
	"testing"
)

func TestSubtitlingAuthErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "auth_error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(SubtitlingAuthError))
}
