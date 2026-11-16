package types

import (
	"testing"
)

func TestStudioTextStyleShadowModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "color": "value",
  "opacity": 9.99,
  "blur": 9.99,
  "offset_x": 9.99,
  "offset_y": 9.99
}`
	assertRoundTrip(t, wire, new(StudioTextStyleShadowModel))
}
