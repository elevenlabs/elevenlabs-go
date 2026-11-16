package types

import (
	"testing"
)

func TestAssetResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "asset_id": "value",
  "name": "value",
  "mime_type": "value",
  "created_at_unix": 123,
  "content_url": "value"
}`
	assertRoundTrip(t, wire, new(AssetResponse))
}
