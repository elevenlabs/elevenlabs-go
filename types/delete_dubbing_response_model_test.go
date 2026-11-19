package types

import (
	"testing"
)

func TestDeleteDubbingResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(DeleteDubbingResponseModel))
}
