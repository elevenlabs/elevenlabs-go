package types

import (
	"testing"
)

func TestDeleteUserVoiceCollectionResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "collection_id": "value"
}`
	assertRoundTrip(t, wire, new(DeleteUserVoiceCollectionResponseModel))
}
