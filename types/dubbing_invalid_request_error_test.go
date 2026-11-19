package types

import (
	"testing"
)

func TestDubbingInvalidRequestErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "invalid_request",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(DubbingInvalidRequestError))
}
