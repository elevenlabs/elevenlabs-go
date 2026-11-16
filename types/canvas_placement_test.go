package types

import (
	"testing"
)

func TestCanvasPlacementRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "x_relative": 9.99,
  "y_relative": 9.99,
  "scale_x": 9.99,
  "scale_y": 9.99,
  "pivot_x": 9.99,
  "pivot_y": 9.99,
  "skew_x": 9.99,
  "skew_y": 9.99,
  "crop_top": 9.99,
  "crop_right": 9.99,
  "crop_bottom": 9.99,
  "crop_left": 9.99,
  "flip_x": true,
  "flip_y": true
}`
	assertRoundTrip(t, wire, new(CanvasPlacement))
}
