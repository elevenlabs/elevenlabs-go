package types

import (
	"testing"
)

func TestProcedureValidationErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "path": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(ProcedureValidationError))
}
