package types

import (
	"testing"
)

func TestDeleteVoiceSampleResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(DeleteVoiceSampleResponseModel))
}
