package types

import (
	"testing"
)

func TestLibraryVoiceMatchResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "library_voice_id": "value",
  "public_owner_id": "value"
}`
	assertRoundTrip(t, wire, new(LibraryVoiceMatchResponseModel))
}
