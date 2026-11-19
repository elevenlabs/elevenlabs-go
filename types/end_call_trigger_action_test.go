package types

import (
	"testing"
)

func TestEndCallTriggerActionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "end_call"
}`
	assertRoundTrip(t, wire, new(EndCallTriggerAction))
}
