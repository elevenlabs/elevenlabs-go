package types

import (
	"testing"
)

func TestEndProcedureToolConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "system_tool_type": "end_procedure",
  "procedures": {
    "key": {
      "procedure_id": "value",
      "version_id": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(EndProcedureToolConfig))
}
