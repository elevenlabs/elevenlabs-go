package types

import (
	"testing"
)

func TestPendingCancellationResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "cancellation",
  "timestamp_seconds": 123
}`
	assertRoundTrip(t, wire, new(PendingCancellationResponseModel))
}
