package types

import (
	"testing"
)

func TestWorkflowPhoneNumberNodeModelOutputCustomSipHeadersItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "static",
  "key": "value",
  "value": "value"
}`
	assertRoundTrip(t, wire, new(WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem))
}
