package types

import (
	"testing"
)

func TestReceptionProcedureResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "trigger": "value",
  "body": "value",
  "id": "value",
  "created_at_unix": 123,
  "source_key": "value",
  "missing_tools": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ReceptionProcedureResponse))
}
