package types

import (
	"testing"
)

func TestProcedureAtVersionResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "procedure_id": "value",
  "version_id": "value",
  "name": "value",
  "content": "value",
  "type": "free_form",
  "trigger": "value",
  "folder_parent_id": "value"
}`
	assertRoundTrip(t, wire, new(ProcedureAtVersionResponseModel))
}
