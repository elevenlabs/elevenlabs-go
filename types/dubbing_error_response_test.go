package types

import (
	"testing"
)

func TestDubbingErrorResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "code": "value",
  "message": "value",
  "retryable": true
}`
	assertRoundTrip(t, wire, new(DubbingErrorResponse))
}

func TestDubbingErrorResponseRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "code": "value",
  "message": "value",
  "retryable": true
}`
	assertRoundTrip(t, wire, new(DubbingErrorResponse))
}
