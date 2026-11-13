package types

import (
	"testing"
)

func TestWidgetConfigAvatarRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "orb",
  "color_1": "value",
  "color_2": "value"
}`
	assertRoundTrip(t, wire, new(WidgetConfigAvatar))
}

func TestWidgetConfigAvatarRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "orb",
  "color_1": "value",
  "color_2": "value"
}`
	assertRoundTrip(t, wire, new(WidgetConfigAvatar))
}
