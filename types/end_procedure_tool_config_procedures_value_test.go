package types

import (
	"testing"
)

func TestEndProcedureToolConfigProceduresValueRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "procedure_id": "value",
  "version_id": "value"
}`
	assertRoundTrip(t, wire, new(EndProcedureToolConfigProceduresValue))
}
