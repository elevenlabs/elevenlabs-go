package types

import (
	"testing"
)

func TestGetUserVoiceCollectionsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "collections": [
    {
      "collection_id": "value",
      "title": "value",
      "icon": "value",
      "permission_on_resource": "admin"
    }
  ]
}`
	assertRoundTrip(t, wire, new(GetUserVoiceCollectionsResponseModel))
}
