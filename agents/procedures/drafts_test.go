package procedures

import (
	"testing"
)

func TestUpdateProcedureDraftRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "content": "value",
  "type": "free_form",
  "trigger": "value"
}`
	assertRoundTrip(t, wire, new(UpdateProcedureDraftRequestModel))
}
