package types

import (
	"testing"
)

func TestImageAvatarRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "image",
  "url": "value"
}`
	assertRoundTrip(t, wire, new(ImageAvatar))
}
