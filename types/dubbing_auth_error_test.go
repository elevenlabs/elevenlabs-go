package types

import (
	"testing"
)

func TestDubbingAuthErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "auth_error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(DubbingAuthError))
}
