package types

import (
	"testing"
)

func TestStudioTextStyleOutlineModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "color": "value",
  "opacity": 9.99,
  "width": 9.99
}`
	assertRoundTrip(t, wire, new(StudioTextStyleOutlineModel))
}
