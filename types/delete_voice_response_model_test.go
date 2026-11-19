package types

import (
	"testing"
)

func TestDeleteVoiceResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(DeleteVoiceResponseModel))
}
