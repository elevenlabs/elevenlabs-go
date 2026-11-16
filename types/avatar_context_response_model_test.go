package types

import (
	"testing"
)

func TestAvatarContextResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "avatar_id": "value",
  "avatar_style_id": "value",
  "avatar_name": "value",
  "avatar_style_name": "value"
}`
	assertRoundTrip(t, wire, new(AvatarContextResponseModel))
}
