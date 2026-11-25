package types

import (
	"testing"
)

func TestReceptionProcedureRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "trigger": "value",
  "body": "value",
  "id": "value",
  "created_at_unix": 123,
  "source_key": "value"
}`
	assertRoundTrip(t, wire, new(ReceptionProcedure))
}

func TestReceptionProcedureRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "trigger": "value",
  "body": "value",
  "id": "value",
  "created_at_unix": 123,
  "source_key": "value"
}`
	assertRoundTrip(t, wire, new(ReceptionProcedure))
}
