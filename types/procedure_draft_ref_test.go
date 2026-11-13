package types

import (
	"testing"
)

func TestProcedureDraftRefRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "procedure_id": "value",
  "version_id": {}
}`
	assertRoundTrip(t, wire, new(ProcedureDraftRef))
}
