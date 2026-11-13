package types

import (
	"testing"
)

func TestCaptionStyleHorizontalPlacementModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "align": "left",
  "translate_pct": 9.99
}`
	assertRoundTrip(t, wire, new(CaptionStyleHorizontalPlacementModel))
}
