package types

import (
	"testing"
)

func TestAddVoiceToUserVoiceCollectionResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "collection_id": "value",
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(AddVoiceToUserVoiceCollectionResponseModel))
}
