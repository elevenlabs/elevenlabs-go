package types

import (
	"testing"
)

func TestEditVoiceResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(EditVoiceResponseModel))
}
