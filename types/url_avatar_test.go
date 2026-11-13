package types

import (
	"testing"
)

func TestUrlAvatarRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "url",
  "custom_url": "value"
}`
	assertRoundTrip(t, wire, new(UrlAvatar))
}
