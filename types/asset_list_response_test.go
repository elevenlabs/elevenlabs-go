package types

import (
	"testing"
)

func TestAssetListResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "assets": [
    {
      "asset_id": "value",
      "name": "value",
      "mime_type": "value",
      "created_at_unix": 123,
      "content_url": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(AssetListResponse))
}
