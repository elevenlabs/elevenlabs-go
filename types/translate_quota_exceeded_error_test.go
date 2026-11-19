package types

import (
	"testing"
)

func TestTranslateQuotaExceededErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "quota_exceeded",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(TranslateQuotaExceededError))
}
