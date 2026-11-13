package types

import (
	"testing"
)

func TestWorkflowPhoneNumberNodeModelOutputPostDialDigitsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "static",
  "value": "0"
}`
	assertRoundTrip(t, wire, new(WorkflowPhoneNumberNodeModelOutputPostDialDigits))
}
