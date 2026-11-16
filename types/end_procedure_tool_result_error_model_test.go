package types

import (
	"testing"
)

func TestEndProcedureToolResultErrorModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "end_procedure_error",
  "status": "not_found",
  "procedure_id": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(EndProcedureToolResultErrorModel))
}

func TestEndProcedureToolResultErrorModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "end_procedure_error",
  "status": "not_found",
  "procedure_id": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(EndProcedureToolResultErrorModel))
}
