package types

import (
	"testing"
)

func TestProcedureDraftResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "procedure_id": "value",
  "name": "value",
  "content": "value",
  "type": "free_form",
  "trigger": "value",
  "folder_parent_id": "value"
}`
	assertRoundTrip(t, wire, new(ProcedureDraftResponseModel))
}
