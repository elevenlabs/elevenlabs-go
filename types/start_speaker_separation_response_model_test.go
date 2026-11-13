package types

import (
	"testing"
)

func TestStartSpeakerSeparationResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(StartSpeakerSeparationResponseModel))
}
