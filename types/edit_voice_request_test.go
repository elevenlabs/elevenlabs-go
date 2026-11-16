package types

import (
	"testing"
)

func TestEditVoiceRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "remove_background_noise": true,
  "description": "value",
  "labels": {
    "key": "value"
  },
  "moderate_metadata": true
}`
	assertRoundTrip(t, wire, new(EditVoiceRequest))
}
