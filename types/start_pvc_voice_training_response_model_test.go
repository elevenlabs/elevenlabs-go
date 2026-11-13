package types

import (
	"testing"
)

func TestStartPvcVoiceTrainingResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(StartPvcVoiceTrainingResponseModel))
}
