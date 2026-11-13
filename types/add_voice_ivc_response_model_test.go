package types

import (
	"testing"
)

func TestAddVoiceIvcResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_id": "value",
  "requires_verification": true,
  "library_voice_match": {
    "library_voice_id": "value",
    "public_owner_id": "value"
  }
}`
	assertRoundTrip(t, wire, new(AddVoiceIvcResponseModel))
}
