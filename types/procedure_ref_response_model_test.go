package types

import (
	"testing"
)

func TestProcedureRefResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "procedure_id": "value",
  "version_id": "value",
  "name": "value",
  "type": "free_form",
  "trigger": "value",
  "referenced_tool_ids": [
    "value"
  ],
  "referenced_kb_ids": [
    "value"
  ],
  "referenced_procedure_ids": [
    "value"
  ],
  "referenced_dynamic_variables": [
    "value"
  ],
  "written_dynamic_variables": [
    "value"
  ],
  "folder_parent_id": "value"
}`
	assertRoundTrip(t, wire, new(ProcedureRefResponseModel))
}

func TestProcedureRefResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "procedure_id": "value",
  "version_id": "value",
  "name": "value",
  "type": "free_form",
  "trigger": "value",
  "referenced_tool_ids": [
    "value"
  ],
  "referenced_kb_ids": [
    "value"
  ],
  "referenced_procedure_ids": [
    "value"
  ],
  "referenced_dynamic_variables": [
    "value"
  ],
  "written_dynamic_variables": [
    "value"
  ],
  "folder_parent_id": "value"
}`
	assertRoundTrip(t, wire, new(ProcedureRefResponseModel))
}
