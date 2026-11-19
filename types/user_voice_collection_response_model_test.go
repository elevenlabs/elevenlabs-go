package types

import (
	"testing"
)

func TestUserVoiceCollectionResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "collection_id": "value",
  "title": "value",
  "icon": "value",
  "permission_on_resource": "admin"
}`
	assertRoundTrip(t, wire, new(UserVoiceCollectionResponseModel))
}
