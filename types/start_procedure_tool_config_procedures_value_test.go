package types

import (
	"testing"
)

func TestStartProcedureToolConfigProceduresValueRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "procedure_id": "value",
  "version_id": "value"
}`
	assertRoundTrip(t, wire, new(StartProcedureToolConfigProceduresValue))
}
