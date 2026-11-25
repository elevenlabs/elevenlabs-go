package types

import (
	"testing"
)

func TestReceptionTransferRuleRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0",
  "id": "value",
  "created_at_unix": 123
}`
	assertRoundTrip(t, wire, new(ReceptionTransferRule))
}

func TestReceptionTransferRuleRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0",
  "id": "value",
  "created_at_unix": 123
}`
	assertRoundTrip(t, wire, new(ReceptionTransferRule))
}
