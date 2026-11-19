package types

import (
	"testing"
)

func TestDeleteProjectResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(DeleteProjectResponseModel))
}
