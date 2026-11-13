package types

import (
	"testing"
)

func TestRequestPvcManualVerificationResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(RequestPvcManualVerificationResponseModel))
}
