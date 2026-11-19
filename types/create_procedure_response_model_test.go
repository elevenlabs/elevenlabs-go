package types

import (
	"testing"
)

func TestCreateProcedureResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "procedure_id": "value"
}`
	assertRoundTrip(t, wire, new(CreateProcedureResponseModel))
}
