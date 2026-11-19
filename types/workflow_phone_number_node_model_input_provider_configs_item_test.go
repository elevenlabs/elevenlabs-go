package types

import (
	"testing"
)

func TestWorkflowPhoneNumberNodeModelInputProviderConfigsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "twilio",
  "config": {
    "transfer_destination": {
      "destination": "value",
      "type": "phone"
    },
    "post_dial_digits": {
      "type": "static",
      "value": "0"
    },
    "type": "twilio_warm_conference_transfer"
  }
}`
	assertRoundTrip(t, wire, new(WorkflowPhoneNumberNodeModelInputProviderConfigsItem))
}
