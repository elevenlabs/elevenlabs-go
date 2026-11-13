package types

import (
	"testing"
)

func TestStartProcedureToolConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "system_tool_type": "start_procedure",
  "procedures": {
    "key": {
      "procedure_id": "value",
      "version_id": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(StartProcedureToolConfig))
}
