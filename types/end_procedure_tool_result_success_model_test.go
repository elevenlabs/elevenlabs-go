package types

import (
	"testing"
)

func TestEndProcedureToolResultSuccessModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "end_procedure_success",
  "status": "success",
  "procedure_id": "value",
  "procedure_name": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(EndProcedureToolResultSuccessModel))
}

func TestEndProcedureToolResultSuccessModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "end_procedure_success",
  "status": "success",
  "procedure_id": "value",
  "procedure_name": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(EndProcedureToolResultSuccessModel))
}
