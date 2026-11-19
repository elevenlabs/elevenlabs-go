package types

import (
	"testing"
)

func TestWorkflowPhoneNumberNodeModelInputTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "phone",
  "phone_number": "value"
}`
	assertRoundTrip(t, wire, new(WorkflowPhoneNumberNodeModelInputTransferDestination))
}
