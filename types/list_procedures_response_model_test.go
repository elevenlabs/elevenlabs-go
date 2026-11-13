package types

import (
	"testing"
)

func TestListProceduresResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "procedures": [
    {
      "procedure_id": "value",
      "version_id": "value",
      "name": "value",
      "type": "free_form",
      "trigger": "value",
      "has_draft": true,
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
      "folder_parent_id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(ListProceduresResponseModel))
}
