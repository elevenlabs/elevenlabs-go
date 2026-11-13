package types

import (
	"testing"
)

func TestMediaGenerationFailedResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "failed",
  "failure_reason": "timeout",
  "error_message": "value"
}`
	assertRoundTrip(t, wire, new(MediaGenerationFailedResponse))
}
