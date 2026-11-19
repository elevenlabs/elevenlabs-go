package types

import (
	"testing"
)

func TestStartProcedureToolResultErrorModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "start_procedure_error",
  "status": "not_found",
  "procedure_id": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(StartProcedureToolResultErrorModel))
}

func TestStartProcedureToolResultErrorModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "start_procedure_error",
  "status": "not_found",
  "procedure_id": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(StartProcedureToolResultErrorModel))
}
