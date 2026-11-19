package voices

import (
	"testing"
)

func TestBodyCreateVoiceCollectionV1VoicesCollectionsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "icon": "value"
}`
	assertRoundTrip(t, wire, new(BodyCreateVoiceCollectionV1VoicesCollectionsPost))
}

func TestBodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "icon": "value"
}`
	assertRoundTrip(t, wire, new(BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch))
}
