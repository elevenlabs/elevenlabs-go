package types

import (
	"testing"
)

func TestRetryTriggerActionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "retry",
  "feedback": "value"
}`
	assertRoundTrip(t, wire, new(RetryTriggerAction))
}
