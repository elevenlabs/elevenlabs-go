package types

import (
	"testing"
)

func TestOrbAvatarRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "orb",
  "color_1": "value",
  "color_2": "value"
}`
	assertRoundTrip(t, wire, new(OrbAvatar))
}
