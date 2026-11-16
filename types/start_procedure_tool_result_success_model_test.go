package types

import (
	"testing"
)

func TestStartProcedureToolResultSuccessModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "start_procedure_success",
  "status": "success",
  "procedure_id": "value",
  "procedure_name": "value",
  "procedure_entry_workflow_node": "value",
  "procedure_return_workflow_node": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(StartProcedureToolResultSuccessModel))
}

func TestStartProcedureToolResultSuccessModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "start_procedure_success",
  "status": "success",
  "procedure_id": "value",
  "procedure_name": "value",
  "procedure_entry_workflow_node": "value",
  "procedure_return_workflow_node": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(StartProcedureToolResultSuccessModel))
}
