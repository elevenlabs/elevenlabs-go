package types

import (
	"testing"
)

func TestSubtitlingQuotaExceededErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "quota_exceeded",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(SubtitlingQuotaExceededError))
}
