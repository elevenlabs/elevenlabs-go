package types

import (
	"testing"
)

func TestCaptionStyleVerticalPlacementModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "align": "top",
  "translate_pct": 9.99
}`
	assertRoundTrip(t, wire, new(CaptionStyleVerticalPlacementModel))
}
