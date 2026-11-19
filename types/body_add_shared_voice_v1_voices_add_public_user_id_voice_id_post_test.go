package types

import (
	"testing"
)

func TestBodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "new_name": "value",
  "bookmarked": true
}`
	assertRoundTrip(t, wire, new(BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost))
}
