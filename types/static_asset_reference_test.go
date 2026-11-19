package types

import (
	"testing"
)

func TestStaticAssetReferenceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "asset",
  "asset_id": "value"
}`
	assertRoundTrip(t, wire, new(StaticAssetReference))
}

func TestStaticAssetReferenceRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "asset",
  "asset_id": "value"
}`
	assertRoundTrip(t, wire, new(StaticAssetReference))
}
