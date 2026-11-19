package types

import (
	"testing"
)

func TestTranslateQueueOverflowErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "queue_overflow",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(TranslateQueueOverflowError))
}
