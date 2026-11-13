package agents

import (
	"testing"
)

func TestCreateProcedureRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "content": "value",
  "type": "free_form",
  "trigger": "value",
  "folder_parent_id": "value"
}`
	assertRoundTrip(t, wire, new(CreateProcedureRequestModel))
}
