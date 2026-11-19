package receptionists

import (
	"testing"
)

func TestReceptionCreateProcedureRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "trigger": "value",
  "body": "value"
}`
	assertRoundTrip(t, wire, new(ReceptionCreateProcedureRequest))
}

func TestReceptionUpdateProcedureRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "trigger": "value",
  "body": "value"
}`
	assertRoundTrip(t, wire, new(ReceptionUpdateProcedureRequest))
}
