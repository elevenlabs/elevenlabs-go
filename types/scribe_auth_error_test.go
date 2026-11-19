package types

import (
	"testing"
)

func TestScribeAuthErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "auth_error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(ScribeAuthError))
}
