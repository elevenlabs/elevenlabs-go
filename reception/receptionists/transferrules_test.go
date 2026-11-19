package receptionists

import (
	"testing"
)

func TestReceptionCreateTransferRuleRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0"
}`
	assertRoundTrip(t, wire, new(ReceptionCreateTransferRuleRequest))
}

func TestReceptionUpdateTransferRuleRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0"
}`
	assertRoundTrip(t, wire, new(ReceptionUpdateTransferRuleRequest))
}
